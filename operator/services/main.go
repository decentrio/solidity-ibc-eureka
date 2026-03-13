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
	ibcexported "github.com/cosmos/ibc-go/v10/modules/core/exported"
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
	CreateCosmosClientContract(ctx Context, clientState, consensusHash []byte) error
	CreateEthClient(ctx Context, clientState ibcexported.ClientState, consensusState ibcexported.ConsensusState) error
	SendEthTx(ctx Context, msg any) error
	SendRecvPacketTx(ctx Context, msg contractICS26Router.IICS26RouterMsgsMsgRecvPacket) error
	SendCosmosTx(ctx Context, msg any) error
	SendCosmosTxBatch(ctx Context, msgs []any) error
}

type EventListener interface {
	SubscribeCosmos(ctx Context, batchBuilder *BatchBuilder)
	SubscribeEth(ctx Context, batchBuilder *BatchBuilder)
}

type Services struct {
	// event listener
	listener EventListener
	worker   *Worker

	ethConfig    Config
	cosmosConfig Config

	BatchPackets chan BatchPackets
	BatchBuilder *BatchBuilder

	txHandler TransactionHandler
}

func New(rpcEndpoint string, eventListener EventListener, txHandler TransactionHandler, ethConfig, cosmosConfig Config) *Services {
	return &Services{
		listener:     eventListener,
		ethConfig:    ethConfig,
		cosmosConfig: cosmosConfig,
		worker: NewWorker(txHandler, nil),
		BatchPackets: make(chan BatchPackets),
		BatchBuilder: NewBatchBuidler(),
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

	// Set contract addresses from environment
	ics26Router := os.Getenv("ICS26_ROUTER")
	wrapVerifier := os.Getenv("WRAP_VERIFIER")
	membership := os.Getenv("MEMBERSHIP")
	misbehaviour := os.Getenv("MISBEHAVIOUR")
	updateClientAddr := os.Getenv("UPDATE_CLIENT")
	roleManager := os.Getenv("ROLE_MANAGER")

	if ics26Router != "" && wrapVerifier != "" && membership != "" && misbehaviour != "" && updateClientAddr != "" && roleManager != "" {
		ctx.SetAddresses(ics26Router, wrapVerifier, membership, misbehaviour, updateClientAddr, roleManager)
	}

	// listen to new tx events on Eth
	// add it to handler queue
	go func() {
		s.listener.SubscribeCosmos(ctx, s.BatchBuilder)
	}()

	// listen to new tx events on Cosmos
	// add it to handler queue
	go func() {
		s.listener.SubscribeEth(ctx, s.BatchBuilder)
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
				ctx.latestEthTimestamp.mtx.Lock()
				ctx.latestEthTimestamp.LatestUpdateTime = time.Now()
				ctx.latestEthTimestamp.LatestUpdateHeight = uint64(latestBlock.BlockHeight)
				ctx.latestEthTimestamp.mtx.Unlock()
			}

			// update client on Cosmos side routinely
			if ctx.latestCosmosTimestamp.LatestUpdateTime.Add(s.cosmosConfig.IntervalParams.blockTime).After(time.Now()) {
				s.worker.UpdateEthClient(ctx)

				// update latest update time
				ctx.latestCosmosTimestamp.mtx.Lock()
				ctx.latestCosmosTimestamp.LatestUpdateTime = time.Now()
				ctx.latestCosmosTimestamp.mtx.Unlock()

			}
		}
	}()

	// check for batch builder
	go func() {
		for {
			// check for batch every 10 seconds
			time.Sleep(time.Second * 10)
			s.BatchBuilder.CheckBatch(ctx.Config.BatchConfig, s.BatchPackets)
		}
	}()

	// handle packets
	for {
		batch, ok := <-s.BatchPackets
		if !ok {
			fmt.Println("Channel closed, exiting loop")
			break
		}

		// Separate packets by direction
		var cosmosToEthPackets []Packet
		var ethToCosmosPackets []Packet
		for _, pkt := range batch.Packets {
			if pkt.FromEth {
				ethToCosmosPackets = append(ethToCosmosPackets, pkt)
			} else {
				cosmosToEthPackets = append(cosmosToEthPackets, pkt)
			}
		}

		// Handle Cosmos→Eth packets
		if len(cosmosToEthPackets) > 0 {
			s.handleCosmosToEthPackets(ctx, cosmosToEthPackets)
		}

		// Handle Eth→Cosmos packets
		if len(ethToCosmosPackets) > 0 {
			s.handleEthToCosmosPackets(ctx, ethToCosmosPackets)
		}
	}

	defer ctx.StopClient()
}

// handleCosmosToEthPackets handles packets from Cosmos to Ethereum
func (s *Services) handleCosmosToEthPackets(ctx Context, packets []Packet) {
	// Update Cosmos light client on Ethereum
	latestLightBlock, err := s.worker.UpdateCosmosClient(ctx, "groth16", int64(ctx.latestEthTimestamp.LatestUpdateHeight), "2/3")
	if err != nil {
		ctx.Logger.Println(fmt.Errorf("Failed to update cosmos light client: %s", err.Error()))
		return
	}

	ctx.latestEthTimestamp.mtx.Lock()
	ctx.latestEthTimestamp.LatestUpdateTime = time.Now()
	ctx.latestEthTimestamp.LatestUpdateHeight = uint64(latestLightBlock.BlockHeight)
	ctx.latestEthTimestamp.mtx.Unlock()

	parsedABI, err := abi.JSON(strings.NewReader(string(tendermintAbiJson)))
	if err != nil {
		ctx.Logger.Println(fmt.Errorf("Failed to parse ABI: %s", err.Error()))
		return
	}

	verifyMethod, exists := parsedABI.Methods["verifyMembership"]
	if !exists {
		ctx.Logger.Println(fmt.Errorf("verifyMembership method not found in ABI"))
		return
	}

	for _, packet := range packets {
		ibcPath := utils.IbcCommitmentPath(*packet.Packet)

		value, merkleProof, err := utils.ProvePath(ctx.CosmosClient(), ibcPath, uint64(latestLightBlock.BlockHeight))
		if err != nil {
			ctx.Logger.Println(fmt.Errorf("Failed to prove path: %s", err.Error()))
			continue
		}

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
			AppHash: utils.BytesToBytes32(latestLightBlock.SignedHeader.AppHash),
			TrustedConsensusState: tendermintContract.IICS07TendermintMsgsConsensusState{
				Timestamp:          big.NewInt(latestLightBlock.SignedHeader.Header.Time.Unix()),
				Root:               utils.BytesToBytes32(latestLightBlock.SignedHeader.Header.AppHash),
				NextValidatorsHash: utils.BytesToBytes32(latestLightBlock.SignedHeader.Header.NextValidatorsHash),
			},
			MembershipType: 0,
		}

		encodedMsg, err := verifyMethod.Inputs.Pack(membershipMsg)
		if err != nil {
			ctx.Logger.Println(fmt.Errorf("Failed to ABI encode membership msg: %s", err.Error()))
			continue
		}

		var payloads []contractICS26Router.IICS26RouterMsgsPayload
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
			MembershipMsg: encodedMsg,
		}

		if err := s.txHandler.SendRecvPacketTx(ctx, msgRecvPacket); err != nil {
			ctx.Logger.Println(fmt.Errorf("Failed to send recv packet to Eth: %s", err.Error()))
		}
	}
}

// handleEthToCosmosPackets handles packets from Ethereum to Cosmos
func (s *Services) handleEthToCosmosPackets(ctx Context, packets []Packet) {
	// Update Ethereum light client on Cosmos
	if err := s.worker.UpdateEthClient(ctx); err != nil {
		ctx.Logger.Println(fmt.Errorf("Failed to update eth light client: %s", err.Error()))
		return
	}

	ctx.latestCosmosTimestamp.mtx.Lock()
	ctx.latestCosmosTimestamp.LatestUpdateTime = time.Now()
	ctx.latestCosmosTimestamp.mtx.Unlock()

	for _, packet := range packets {
		if err := s.worker.RelayEthToCosmosPacket(ctx, packet); err != nil {
			ctx.Logger.Println(fmt.Errorf("Failed to relay Eth→Cosmos packet: %s", err.Error()))
		}
	}
}
