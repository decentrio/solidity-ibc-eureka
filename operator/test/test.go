package main

import (
	"encoding/json"
	"fmt"
	"log"
	operatorclient "operator/client"
	"operator/services"
	"operator/transaction"
	"os"

	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
	"github.com/ethereum/go-ethereum/ethclient"
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

// const COMETBFT_SEND_PACKET_EVENT = "tm.event = 'Tx' AND message.action = '/ibc.applications.transfer.v1.MsgTransfer'"
// const EVENT_SEND_PACKET_FIELD = "send_packet.encoded_packet_hex"
// const EVENT_TX_HASH_FIELD = "tx.hash"

// func subscribeCosmos(logger *log.Logger, client *rpchttp.HTTP) {
// 	sub, err := client.WSEvents.Subscribe(context.Background(), "", COMETBFT_SEND_PACKET_EVENT)
// 	if err != nil {
// 		logger.Println(err.Error())
// 	}

// 	for {
// 		select {
// 		case e := <-sub:
// 			fmt.Println("events: ", e.Events)
// 			// handle event
// 			sendPacketEvent := e.Events[EVENT_SEND_PACKET_FIELD]
// 			if sendPacketEvent == nil {
// 				fmt.Println("sendPacketEvent is empty")
// 				continue
// 			}
// 			// packetHex := e.Events[channeltypesv2.AttributeKeyEncodedPacketHex]
// 			// fmt.Println("packetHex: ", packetHex)

// 			txHashStr := e.Events[EVENT_TX_HASH_FIELD]
// 			if txHashStr == nil {
// 				fmt.Println("txHashStr is empty")
// 				continue
// 			}
// 			fmt.Println("txHashStr: ", txHashStr)
// 			// txHash, err := hex.DecodeString(txHashStr[0])
// 			// if err != nil {
// 			// 	fmt.Println(fmt.Errorf("Failed to decode tx hash: %s", err.Error()))
// 			// 	continue
// 			// }

// 			packetEncodedStr := sendPacketEvent[0]
// 			packetBytes, err := hex.DecodeString(packetEncodedStr)
// 			if err != nil {
// 				fmt.Println(fmt.Errorf("Failed to decode packet hex: %s", err.Error()))
// 				continue
// 			}

// 			var packet channeltypesv2.Packet
// 			err = proto.Unmarshal(packetBytes, &packet)
// 			if err != nil {
// 				fmt.Println(fmt.Errorf("Failed to unmarshal packet: %s", err.Error()))
// 				continue
// 			}

// 			txHash, err := hex.DecodeString(txHashStr[0])
// 			time.Sleep(time.Second)
// 			txResp, err := client.Tx(context.Background(), txHash, true)
// 			if err != nil {
// 				fmt.Println(fmt.Errorf("Failed to fetch tx from tx hash: %s", err.Error()))
// 				continue
// 			}
// 			fmt.Println("txResp: ", txResp)

// 			revisionHeight := int64(txResp.Height)

// 			sequenceBytes := make([]byte, 8)
// 			binary.BigEndian.PutUint64(sequenceBytes, packet.Sequence)
// 			path := []byte(packet.SourceClient)
// 			path = append(path, []byte{1}...)
// 			path = append(path, sequenceBytes...)
// 			// target height are the latest block height
// 			_, merkleProof, err := utils.ProvePath(client, [][]byte{[]byte("ibc"), path}, uint64(revisionHeight))

// 			fmt.Println(merkleProof)
// 			fmt.Println(err)
// 		}
// 	}
// }

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

	worker := services.NewWorker(&transaction.Handler{}, nil)

	ctx := services.NewCtxWithBeacon(cosmosClient, ethClient, cfg.EthToCosmosConfig.BeaconUrl, "")
	ctx.SetAddresses(cfg.CosmosToEthConfig.ICS26Address, cfg.CosmosToEthConfig.WrapperVerifier, cfg.CosmosToEthConfig.Membership, cfg.CosmosToEthConfig.Misbehaviour, cfg.CosmosToEthConfig.UpdateClient, "0x8943545177806ED17B9F23F0a21ee5948eCaa776")

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
}
