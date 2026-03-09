package services

import (
	"fmt"
	"log"
	"sync"
	"time"

	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Timestamp struct {
	mtx                sync.Mutex
	LatestUpdateTime   time.Time
	LatestUpdateHeight uint64
}

type Context struct {
	Logger *log.Logger
	Config Config

	latestEthTimestamp    *Timestamp
	latestCosmosTimestamp *Timestamp

	cosmosClient *rpchttp.HTTP
	ethClient    *ethclient.Client
	beaconAPIURL string

	// Ethereum light client configuration
	ethClientID string
}

func NewCtx(cosmosClient *rpchttp.HTTP, ethClient *ethclient.Client) Context {
	return Context{
		Logger:       log.Default(),
		cosmosClient: cosmosClient,
		ethClient:    ethClient,
		beaconAPIURL: "",
		ethClientID:  "",
		latestEthTimestamp: &Timestamp{
			LatestUpdateTime:   time.Now(),
			LatestUpdateHeight: 0,
		},
		latestCosmosTimestamp: &Timestamp{
			LatestUpdateTime:   time.Now(),
			LatestUpdateHeight: 0,
		},
	}
}

func NewCtxWithBeacon(cosmosClient *rpchttp.HTTP, ethClient *ethclient.Client, beaconAPIURL string, ethClientID string) Context {
	return Context{
		Logger:       log.Default(),
		cosmosClient: cosmosClient,
		ethClient:    ethClient,
		beaconAPIURL: beaconAPIURL,
		ethClientID:  ethClientID,
		latestEthTimestamp: &Timestamp{
			LatestUpdateTime:   time.Now(),
			LatestUpdateHeight: 0,
		},
		latestCosmosTimestamp: &Timestamp{
			LatestUpdateTime:   time.Now(),
			LatestUpdateHeight: 0,
		},
	}
}

func (c *Context) EthClient() *ethclient.Client {
	return c.ethClient
}

func (c *Context) CosmosClient() *rpchttp.HTTP {
	return c.cosmosClient
}

func (c *Context) BeaconAPIURL() string {
	return c.beaconAPIURL
}

func (c *Context) EthClientID() string {
	return c.ethClientID
}

func (c *Context) StopClient() {
	err := c.cosmosClient.Stop()
	if err != nil {
		panic(fmt.Errorf("failed to terminate cosmos client: %v", err))
	}
	c.ethClient.Close()
}

func (c *Context) LatestCosmosTimestamp() *Timestamp {
	return c.latestCosmosTimestamp
}

func (c *Context) LatestEthTimestamp() *Timestamp {
	return c.latestEthTimestamp
}
