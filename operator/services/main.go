package services

import (
	"fmt"
	"os"
	"time"

	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
	"github.com/ethereum/go-ethereum/ethclient"
)

type TransactionHandler interface {
	CreateCosmosClientContract(ctx Context, clientState, consensusHash []byte) error
	SendEthTx(ctx Context, msg any) error
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

	// packets chan
	// TODO: add retry queue and tracks timeout for packets
	// TODO: interface instead of any here
	packets chan any
}

func New(rpcEndpoint string, eventListener EventListener, txHandler TransactionHandler, ethConfig, cosmosConfig Config) *Services {
	return &Services{
		listener:     eventListener,
		ethConfig:    ethConfig,
		cosmosConfig: cosmosConfig,
		worker: &Worker{
			txHandler,
		},
		packets: make(chan any),
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

	wrapVerifier := os.Getenv("WRAP_VERIFIER")
	if wrapVerifier == "" {
		panic(fmt.Errorf("WRAP_VERIFIER environment variable is required in .env file"))
	}

	membership := os.Getenv("MEMBERSHIP")
	if membership == "" {
		panic(fmt.Errorf("MEMBERSHIP environment variable is required in .env file"))
	}
	misbehaviour := os.Getenv("MISBEHAVIOUR")
	if misbehaviour == "" {
		panic(fmt.Errorf("MISBEHAVIOUR environment variable is required in .env file"))
	}
	updateClient := os.Getenv("UPDATE_CLIENT")
	if updateClient == "" {
		panic(fmt.Errorf("UPDATE_CLIENT environment variable is required in .env file"))
	}
	roleManager := os.Getenv("ROLE_MANAGER")
	if roleManager == "" {
		panic(fmt.Errorf("ROLE_MANAGER environment variable is required in .env file"))
	}

	ctx := NewCtx(cosmosClient, ethClient)
	ctx.SetAddresses(wrapVerifier, membership, misbehaviour, updateClient, roleManager)

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
				s.worker.UpdateCosmosClient(ctx, "groth16", 1, "2/3")

				// update latest update time
				ctx.latestEthTimestamp.LatestUpdateTime = time.Now()
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
		_, ok := <-s.packets
		if !ok {
			fmt.Println("Channel closed, exiting loop")
			break // Exit the loop when the channel is closed
		}

		// handle tx here

	}

	defer ctx.StopClient()
}
