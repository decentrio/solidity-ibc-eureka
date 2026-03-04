package services

import (
	"fmt"
	"log"
	"math/big"
	"operator/utils"
	"os"
	"strings"
	"time"

	contractICS26Router "operator/bindings/ICS26Router"
	tendermintContract "operator/bindings/SP1ICS07Tendermint"

	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/ethclient"
)

// read abi json file once in runtime
var tendermintAbiJson []byte
var initErr error

func init() {
	tendermintAbiJson, initErr = os.ReadFile("../../abi/SP1ICS07Tendermint.json")
	if initErr != nil {
		log.Fatal(initErr)
	}
}

type TransactionHandler interface {
	SendTx(ctx Context, msg any) error
	SendCosmosTx(ctx Context, msg any) error
	SendCosmosTxBatch(ctx Context, msgs []any) error
}

type EventListener interface {
	SubscribeCosmos(ctx Context)
	SubscribeEth(ctx Context)
}

type Services struct {
	// event listener
	listener EventListener
	worker   *Worker

	ethConfig    Config
	cosmosConfig Config

	txHandler TransactionHandler
}

func New(rpcEndpoint string, eventListener EventListener, txHandler TransactionHandler, ethConfig, cosmosConfig Config) *Services {
	return &Services{
		listener:     eventListener,
		ethConfig:    ethConfig,
		cosmosConfig: cosmosConfig,
		worker: &Worker{
			txHandler: txHandler,
		},
	}
}

func (s *Services) StartLoop() {
	rpcEndpoint := os.Getenv("TENDERMINT_RPC_URL")
	if rpcEndpoint == "" {
		panic(fmt.Errorf("TENDERMINT_RPC_URL environment variable is required in .env file"))
	}
	cosmosClient, err := rpchttp.New(rpcEndpoint, "/websocket")
	if err != nil {
		panic(fmt.Errorf("failed to create RPC client: %w", err))
	}

	ethRpcEndpoint := os.Getenv("ETH_RPC_URL")
	if ethRpcEndpoint == "" {
		panic(fmt.Errorf("ETH_RPC_URL environment variable is required in .env file"))
	}
	ethClient, err := ethclient.Dial(ethRpcEndpoint)
	if err != nil {
		panic(fmt.Errorf("failed to connect to client: %s: ", err.Error()))
	}

	ctx := NewCtx(cosmosClient, ethClient)

	// listen to new tx events on Eth
	// add it to handler queue
	go func() {
		s.listener.SubscribeCosmos(ctx)
	}()

	// listen to new tx events on Cosmos
	// add it to handler queue
	go func() {
		s.listener.SubscribeEth(ctx)
	}()

	// routinely run update client
	go func() {
		for {
			// update client on Eth side routinely
			if ctx.latestEthTimestamp.LatestUpdateTime.Add(s.ethConfig.IntervalParams.blockTime).After(time.Now()) {
				latestBlock, err := s.worker.UpdateCosmosClient(ctx, "groth16", int64(ctx.latestEthTimestamp.LatestUpdateHeight), "2/3")
				if err != nil {
					ctx.Logger.Println(fmt.Errorf("Failed to update cosmos light client: %s", err.Error()))
				}

				// update latest update time
				ctx.latestEthTimestamp.LatestUpdateTime = time.Now()
				ctx.latestEthTimestamp.LatestUpdateHeight = uint64(latestBlock.BlockHeight)
			}

			// update client on Cosmos side routinely
			if ctx.latestCosmosTimestamp.LatestUpdateTime.Add(s.cosmosConfig.IntervalParams.blockTime).After(time.Now()) {
				s.worker.UpdateEthClient(ctx)

				// update latest update time
				ctx.latestCosmosTimestamp.LatestUpdateTime = time.Now()
			}
		}
	}()

	// handle packets
	for {
		batch, ok := <-ctx.BatchPackets
		if !ok {
			fmt.Println("Channel closed, exiting loop")
			break // Exit the loop when the channel is closed
		}

		// update client
		latestLightBlock, err := s.worker.UpdateCosmosClient(ctx, "groth16", int64(ctx.latestEthTimestamp.LatestUpdateHeight), "2/3")
		if err != nil {
			ctx.Logger.Println(fmt.Errorf("Failed to update cosmos light client: %s", err.Error()))
		}

		// update latest update time
		ctx.latestEthTimestamp.LatestUpdateTime = time.Now()
		// update latest trusted block height
		ctx.latestEthTimestamp.LatestUpdateHeight = uint64(latestLightBlock.BlockHeight)

		// handle packets in batch
		for _, packet := range batch.Packets {

			ibcPath := utils.IbcCommitmentPath(*packet.Packet)

			// target height are the latest block height
			value, merkleProof, err := utils.ProvePath(ctx.CosmosClient(), ibcPath, uint64(latestLightBlock.BlockHeight))

			membershipMsg := tendermintContract.ILightClientMsgsMsgVerifyMembership{
				Height: tendermintContract.IICS02ClientMsgsHeight{
					RevisionHeight: uint64(latestLightBlock.BlockHeight),
					RevisionNumber: 0,
				},
				KvPairs: []tendermintContract.IMembershipMsgsKVPair{
					{
						Path:  ibcPath,
						Value: utils.BytesToBytes32(value),
					},
				},
				MerkleProofs: []tendermintContract.IMembershipMsgsMerkleProof{
					*merkleProof,
				},
				// current appHash
				AppHash: utils.BytesToBytes32(latestLightBlock.SignedHeader.AppHash),
				// trusted consensus from revision height block
				TrustedConsensusState: tendermintContract.IICS07TendermintMsgsConsensusState{
					Timestamp:          big.NewInt(latestLightBlock.SignedHeader.Header.Time.Unix()),
					Root:               utils.BytesToBytes32(latestLightBlock.SignedHeader.Header.ConsensusHash),
					NextValidatorsHash: utils.BytesToBytes32(latestLightBlock.SignedHeader.Header.NextValidatorsHash),
				},
				MembershipType: 1,
			}

			parsedABI, err := abi.JSON(strings.NewReader(string(tendermintAbiJson)))
			if err != nil {
				ctx.Logger.Println(fmt.Errorf("Failed to read abi json file: %s", err.Error()))
			}

			calldata, err := parsedABI.Pack("verifyMembership", membershipMsg)
			if err != nil {
				ctx.Logger.Println(fmt.Errorf("Failed to abi encode verify msg: %s", err.Error()))
			}

			payloads := make([]contractICS26Router.IICS26RouterMsgsPayload, len(packet.Packet.Payloads))
			for _, p := range packet.Packet.Payloads {
				payloads = append(payloads, contractICS26Router.IICS26RouterMsgsPayload{
					SourcePort: p.SourcePort,
					DestPort:   p.DestinationPort,
					Version:    p.Version,
					Encoding:   p.Encoding,
					Value:      p.Value,
				})
			}

			msgRecvPacket := contractICS26Router.IICS26RouterMsgsMsgRecvPacket{
				Packet: contractICS26Router.IICS26RouterMsgsPacket{
					Sequence:         packet.Packet.Sequence,
					SourceClient:     packet.Packet.SourceClient,
					DestClient:       packet.Packet.DestinationClient,
					TimeoutTimestamp: packet.Packet.TimeoutTimestamp,
					Payloads:         payloads,
				},
				MembershipMsg: calldata,
			}

			s.txHandler.SendTx(ctx, msgRecvPacket)
		}
	}

	defer ctx.StopClient()
}
