package main

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	contractICS26Router "operator/bindings/ICS26Router"
	tendermintContract "operator/bindings/SP1ICS07Tendermint"
	operatorclient "operator/client"
	"operator/prover"
	"operator/services"
	"operator/subscriber"
	"operator/transaction"
	"operator/utils"
	"os"
	"time"

	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
	"github.com/cosmos/gogoproto/proto"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	channeltypesv2 "github.com/cosmos/ibc-go/v10/modules/core/04-channel/v2/types"
	"github.com/joho/godotenv"
)

type CosmosToEthConfig struct {
	TmRpcUrl        string `json:"tm_rpc_url"`
	ICS26Address    string `json:"ics26_address"`
	EthRpcUrl       string `json:"eth_rpc_url"`
	WrapperVerifier string `json:"wrapper_verifier"`
	Membership      string `json:"membership"`
	Misbehaviour    string `json:"misbehaviour"`
	UpdateClient    string `json:"update_client"`
}

type EthToCosmosConfig struct {
	TmRpcUrl      string `json:"tm_rpc_url"`
	ICS26Address  string `json:"ics26_address"`
	EthRpcUrl     string `json:"eth_rpc_url"`
	BeaconUrl     string `json:"eth_beacon_api_url"`
	SignerAddress string `json:"signer_address"`
}

type Module struct {
	Name     string          `json:"name"`
	SrcChain string          `json:"src_chain"`
	DstChain string          `json:"dst_chain"`
	Config   json.RawMessage `json:"config"`
}

type ServerConfig struct {
	LogLevel string `json:"log_level"`
	Address  string `json:"address"`
	Port     uint64 `json:"port"`
}

type AppJsonConfig struct {
	SeverConfig ServerConfig `json:"server"`
	Modules     []Module     `json:"modules"`
}

type AppConfig struct {
	SeverConfig       ServerConfig
	EthToCosmosConfig EthToCosmosConfig
	CosmosToEthConfig CosmosToEthConfig
}

func loadConfig(configPath string) (*AppConfig, error) {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var appConfig AppJsonConfig
	if err := json.Unmarshal(data, &appConfig); err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}

	var c2eCfg CosmosToEthConfig
	var e2cfg EthToCosmosConfig
	for _, m := range appConfig.Modules {

		if m.Name == "cosmos_to_eth" {
			fmt.Println("m.Config: ", string(m.Config))
			if err := json.Unmarshal(m.Config, &c2eCfg); err != nil {
				return nil, fmt.Errorf("failed to parse cosmos_to_eth config: %w", err)
			}
		}
		if m.Name == "eth_to_cosmos" {
			if err := json.Unmarshal(m.Config, &e2cfg); err != nil {
				return nil, fmt.Errorf("failed to parse eth_to_cosmos config: %w", err)
			}
		}
	}

	return &AppConfig{
		SeverConfig:       appConfig.SeverConfig,
		CosmosToEthConfig: c2eCfg,
		EthToCosmosConfig: e2cfg,
	}, nil
}

// type EurekaEvent struct {
// 	eventType string
// 	packet    channeltypesv2.Packet
// 	ack       *channeltypesv2.Acknowledgement
// }

func init() {
	// tendermintAbiJson, initErr = os.ReadFile("../../abi/SP1ICS07Tendermint.json")
	// if initErr != nil {
	// 	log.Fatal(initErr)
	// }
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

type Listener struct{}

func (l *Listener) SubscribeCosmos(ctx services.Context, worker *services.Worker) {
	sub, err := ctx.CosmosClient().WSEvents.Subscribe(context.Background(), "", subscriber.COMETBFT_SEND_PACKET_EVENT)
	if err != nil {
		ctx.Logger.Println(err.Error())
	}
	for {
		select {
		case e := <-sub:
			// handle event
			sendPacketEvent := e.Events[subscriber.EVENT_SEND_PACKET_FIELD]
			if sendPacketEvent == nil {
				continue
			}

			packetEncodedStr := sendPacketEvent[0]
			packetBytes, err := hex.DecodeString(packetEncodedStr)
			if err != nil {
				// TODO handle log here
				ctx.Logger.Println(fmt.Errorf("Failed to decode packet hex: %s", err.Error()))
				continue
			}

			var packet channeltypesv2.Packet
			err = proto.Unmarshal(packetBytes, &packet)
			if err != nil {
				// TODO handle log here
				ctx.Logger.Println(fmt.Errorf("Failed to unmarshal packet: %s", err.Error()))
				continue
			}

			latestEthTimestamp := ctx.LatestCosmosTimestamp()
			latestLightBlock, err := worker.UpdateCosmosClient(ctx, "groth16", int64(latestEthTimestamp.LatestUpdateHeight), "2/3")
			if err != nil {
				ctx.Logger.Println(fmt.Errorf("Failed to update cosmos light client: %s", err.Error()))
				continue
			}
			latestEthTimestamp.LatestUpdateTime = time.Now()
			latestEthTimestamp.LatestUpdateHeight = uint64(latestLightBlock.BlockHeight)

			ibcPath := utils.IbcCommitmentPath(packet)
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

			tendermintAbiJson, err := tendermintContract.ContractSP1ICS07TendermintMetaData.GetAbi()
			if err != nil {
				ctx.Logger.Println(fmt.Errorf("Failed to abi encode verify msg: %s", err.Error()))
			}
			calldata, err := tendermintAbiJson.Pack("verifyMembership", membershipMsg)
			if err != nil {
				ctx.Logger.Println(fmt.Errorf("Failed to abi encode verify msg: %s", err.Error()))
			}

			payloads := make([]contractICS26Router.IICS26RouterMsgsPayload, len(packet.Payloads))
			for _, p := range packet.Payloads {
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
					Sequence:         packet.Sequence,
					SourceClient:     packet.SourceClient,
					DestClient:       packet.DestinationClient,
					TimeoutTimestamp: packet.TimeoutTimestamp,
					Payloads:         payloads,
				},
				MembershipMsg: calldata,
			}

			worker.TxHandler.SendEthTx(ctx, msgRecvPacket)
		}
	}
}
func main() {
	cfg, err := loadConfig("./config.example.json")
	if err != nil {
		panic(fmt.Errorf("failed to load config: %w", err).Error())
	}
	fmt.Println("cfg: ", cfg)

	ethRpcEndpoint := cfg.EthToCosmosConfig.EthRpcUrl
	ethClient, err := ethclient.Dial(ethRpcEndpoint)
	if err != nil {
		panic(fmt.Errorf("failed to connect to client: %s: ", err.Error()))
	}

	cosmosRpcEndpoint := cfg.CosmosToEthConfig.TmRpcUrl
	cosmosClient, err := rpchttp.New(cosmosRpcEndpoint, "/websocket")
	if err != nil {
		panic(fmt.Errorf("failed to create RPC client: %w", err))
	}

	prover, err := prover.NewProver("./prover/data/r1cs.bin", "./prover/data/pk.bin")
	if err != nil {
		panic(fmt.Errorf("failed reading prover key: %w", err))
	}
	worker := services.NewWorker(&transaction.Handler{}, prover)

	ctx := services.NewCtxWithBeacon(cosmosClient, ethClient, cfg.EthToCosmosConfig.BeaconUrl, "")
	ctx.SetAddresses(cfg.CosmosToEthConfig.ICS26Address, cfg.CosmosToEthConfig.WrapperVerifier, cfg.CosmosToEthConfig.Membership, cfg.CosmosToEthConfig.Misbehaviour, cfg.CosmosToEthConfig.UpdateClient, "0x8943545177806ED17B9F23F0a21ee5948eCaa776")

	ics07 := common.HexToAddress("0x6fDA176cb71b4f2b85c17E398b58803797f721e4")
	ctx.SetClient(ics07)
	err = ctx.CosmosClient().Start()
	if err != nil {
		panic(err)
	}
	defer ctx.CosmosClient().Stop()
	listener := Listener{}

	unbondingPeriod, err := operatorclient.GetUnbondingTime(cosmosClient)
	if err != nil {
		panic(fmt.Errorf("failed to fetch unbonding time client: %w", err))
	}
	trustingPeriod := 2 * uint32(unbondingPeriod) / 3
	err = worker.CreateCosmosClient(ctx, "groth16", trustingPeriod, 0, "1/3")
	if err != nil {
		panic(fmt.Errorf("create client err: %w", err))
	}
	err = worker.CreateEthClient(ctx, "0xc6d93045091f05f6c056ca8fa583126902967b4b829085042529d279c188391c")
	if err != nil {
		panic(fmt.Errorf("create client err: %w", err))
	}

	listener.SubscribeCosmos(ctx, worker)
}
