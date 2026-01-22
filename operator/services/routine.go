package services

import (
	"context"
	"fmt"
	"math/big"
	updateclient "operator/bindings/UpdateClient"
	operatorclient "operator/client"
	"time"

	clienttypes "github.com/cosmos/ibc-go/v10/modules/core/02-client/types"
)

type Worker struct {
	txHandler TransactionHandler
}

func NewWorkerWithConfig(txHandler TransactionHandler) *Worker {
	return &Worker{
		txHandler,
	}
}

func NewWorker(txHandler TransactionHandler) *Worker {
	return &Worker{
		txHandler,
	}
}

func (w *Worker) UpdateCosmosClient(ctx Context, proofType string, trustedBlock int64, trustLevel string) error {
	status, err := ctx.CosmosClient().Status(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get status: %w", err)
	}

	if trustedBlock == 0 {
		trustedBlock = status.SyncInfo.LatestBlockHeight
	}

	trustedLightBlock, err := operatorclient.GetLightBlock(ctx.CosmosClient(), trustedBlock)
	if err != nil {
		return fmt.Errorf("failed to get trusted light block: %w", err)
	}

	latestLightBlock, err := operatorclient.GetLightBlock(ctx.CosmosClient(), status.SyncInfo.LatestBlockHeight)
	if err != nil {
		return fmt.Errorf("failed to get latest light block: %w", err)
	}

	unbondingPeriod, err := operatorclient.GetUnbondingTime(ctx.CosmosClient())
	if err != nil {
		return fmt.Errorf("failed to get unbonding time: %w", err)
	}

	trustingPeriod := uint32(unbondingPeriod * 2 / 3)

	if trustingPeriod > uint32(unbondingPeriod) {
		return fmt.Errorf("trusting period %d cannot be greater than unbonding period %d", trustingPeriod, uint32(unbondingPeriod))
	}

	chainId := trustedLightBlock.SignedHeader.Header.ChainID
	revision := clienttypes.ParseChainID(chainId)

	trustThreshold, err := operatorclient.ParseTrustThreshold(trustLevel)
	if err != nil {
		return fmt.Errorf("failed to parse trust level: %w", err)
	}

	var zkAlgorithm operatorclient.SupportedZkAlgorithm
	switch proofType {
	case "groth16":
		zkAlgorithm = operatorclient.Groth16
	case "plonk":
		zkAlgorithm = operatorclient.Plonk
	default:
		return fmt.Errorf("unsupported proof type: %s, supported types are: groth16, plonk", proofType)
	}

	clientState := updateclient.IICS07TendermintMsgsClientState{
		ChainId:    chainId,
		TrustLevel: trustThreshold,
		LatestHeight: updateclient.IICS02ClientMsgsHeight{
			RevisionNumber: revision,
			RevisionHeight: uint64(trustedLightBlock.SignedHeader.Header.Height),
		},
		IsFrozen:        false,
		ZkAlgorithm:     uint8(zkAlgorithm),
		TrustingPeriod:  trustingPeriod,
		UnbondingPeriod: uint32(unbondingPeriod),
	}

	consensusState := updateclient.IICS07TendermintMsgsConsensusState{
		Timestamp:          big.NewInt(trustedLightBlock.SignedHeader.Header.Time.UnixMilli()),
		Root:               bytesToBytes32(trustedLightBlock.SignedHeader.Header.AppHash),
		NextValidatorsHash: bytesToBytes32(trustedLightBlock.SignedHeader.NextValidatorsHash),
	}

	proposedHeader := latestLightBlock.IntoHeader(*trustedLightBlock)

	// TODO: generate proof

	msg := updateclient.IUpdateClientMsgsMsgUpdateClient{
		ClientState:           clientState,
		TrustedConsensusState: consensusState,
		Time:                  big.NewInt(time.Now().Unix()),
		ProposedHeader:        proposedHeader,
	}

	return w.txHandler.SendTx(ctx, msg)
}

func (w *Worker) UpdateEthClient(ctx Context) error {
	return nil
}

func bytesToBytes32(data []byte) [32]byte {
	var result [32]byte
	copy(result[:], data)
	return result
}
