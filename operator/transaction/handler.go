package transaction

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"operator/keys"
	"os"
	"strings"

	tendermintContract "operator/bindings/SP1ICS07Tendermint"
	updateclient "operator/bindings/UpdateClient"
	services "operator/services"

	"github.com/cosmos/gogoproto/proto"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type Handler struct {
}

func (h *Handler) SendEthTx(ctx services.Context, msg any) error {
	privKey := os.Getenv("PRIVATE_KEY")
	if privKey == "" {
		return fmt.Errorf("PRIVATE_KEY environment variable is required in .env file")
	}
	privateKey, err := keys.RestoreKey(privKey)
	if err != nil {
		return fmt.Errorf("failed to restore private key: %w", err)
	}

	chainIdEth := os.Getenv("CHAIN_ID")
	if chainIdEth == "" {
		return fmt.Errorf("CHAIN_ID environment variable is required in .env file")
	}

	publicKey, err := keys.PublicKey(privateKey)
	if err != nil {
		log.Fatal(err)
	}

	fromAddress := crypto.PubkeyToAddress(*publicKey)
	nonce, err := ctx.EthClient().PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	gasPrice, err := ctx.EthClient().SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	chainIdInt := big.NewInt(0)
	chainIdInt, ok := chainIdInt.SetString(chainIdEth, 10)
	if !ok {
		return fmt.Errorf("invalid chain id: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainIdInt)
	if err != nil {
		return fmt.Errorf("failed to create auth transactor: %w", err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	hexAddress := os.Getenv("CONTRACT_ADDRESS")
	if hexAddress == "" {
		return fmt.Errorf("CONTRACT_ADDRESS environment variable is required in .env file")
	}

	tendermintAddr := common.HexToAddress(hexAddress)
	ics07Tendermint, err := tendermintContract.NewContractSP1ICS07Tendermint(
		tendermintAddr,
		ctx.EthClient(),
	)
	if err != nil {
		return fmt.Errorf("failed to create ICS07 Tendermint contract: %w", err)
	}

	switch msg := msg.(type) {
	case updateclient.IUpdateClientMsgsMsgUpdateClient:
		parsedABI, err := abi.JSON(strings.NewReader("../../abi/SP1ICS07Tendermint.json"))
		if err != nil {
			panic(err)
		}

		data, err := parsedABI.Pack(
			"transfer",
			msg,
		)
		if err != nil {
			panic(err)
		}

		_, err = ics07Tendermint.UpdateClient(auth, data)
		if err != nil {
			return fmt.Errorf("failed to verify membership: %w", err)
		}
	case tendermintContract.ILightClientMsgsMsgVerifyMembership:
		_, err := ics07Tendermint.VerifyMembership(auth, msg)
		if err != nil {
			return fmt.Errorf("failed to verify membership: %w", err)
		}
	case tendermintContract.ILightClientMsgsMsgVerifyNonMembership:
		_, err := ics07Tendermint.VerifyNonMembership(auth, msg)
		if err != nil {
			return fmt.Errorf("failed to verify membership: %w", err)
		}
	default:
		return fmt.Errorf("unsupported message type")
	}

	return nil
}

func (h *Handler) SendCosmosTx(ctx services.Context, msg proto.Message) error {
	return nil
}
