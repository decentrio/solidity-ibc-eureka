package services

import (
	"fmt"
	"os"
	"time"

	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
	"github.com/ethereum/go-ethereum/ethclient"
)

type TransactionHandler interface {
	SendTx(ctx Context, msg any) error
	SendCosmosTx(ctx Context, msg any) error
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

	// pending txs
	txs chan any
}

func New(rpcEndpoint string, eventListener EventListener, txHandler TransactionHandler, ethConfig, cosmosConfig Config) *Services {
	return &Services{
		listener:     eventListener,
		ethConfig:    ethConfig,
		cosmosConfig: cosmosConfig,
		worker: &Worker{
			txHandler: txHandler,
		},
		txs: make(chan any),
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
			if ctx.latestEthTimestamp.latestUpdateTime.Add(s.ethConfig.IntervalParams.blockTime).After(time.Now()) {
				s.worker.UpdateCosmosClient(ctx, "groth16", 1, "2/3")

				// update latest update time
				ctx.latestEthTimestamp.latestUpdateTime = time.Now()
			}

			// update client on Cosmos side routinely
			if ctx.latestCosmosTimestamp.latestUpdateTime.Add(s.cosmosConfig.IntervalParams.blockTime).After(time.Now()) {
				s.worker.UpdateEthClient(ctx)

				// update latest update time
				ctx.latestCosmosTimestamp.latestUpdateTime = time.Now()
			}
		}
	}()

	// handle txs
	for {
		_, ok := <-s.txs
		if !ok {
			fmt.Println("Channel closed, exiting loop")
			break // Exit the loop when the channel is closed
		}

		// handle tx here

	}

	defer ctx.StopClient()
}
