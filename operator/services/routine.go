package services

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	updateclient "operator/bindings/UpdateClient"
	operatorclient "operator/client"
	"time"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	ibcwasmtypes "github.com/cosmos/ibc-go/modules/light-clients/08-wasm/v10/types"
	clienttypes "github.com/cosmos/ibc-go/v10/modules/core/02-client/types"
)

type Worker struct {
	txHandler TransactionHandler
}

func NewWorker(txHandler TransactionHandler) *Worker {
	return &Worker{
		txHandler,
	}
}

func (w *Worker) UpdateCosmosClient(ctx Context, proofType string, trustedBlock int64, trustLevel string) (*operatorclient.LightBlock, error) {
	status, err := ctx.CosmosClient().Status(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get status: %w", err)
	}

	if trustedBlock == 0 {
		trustedBlock = status.SyncInfo.LatestBlockHeight
	}

	trustedLightBlock, err := operatorclient.GetLightBlock(ctx.CosmosClient(), trustedBlock)
	if err != nil {
		return nil, fmt.Errorf("failed to get trusted light block: %w", err)
	}

	latestLightBlock, err := operatorclient.GetLightBlock(ctx.CosmosClient(), status.SyncInfo.LatestBlockHeight)
	if err != nil {
		return nil, fmt.Errorf("failed to get latest light block: %w", err)
	}

	unbondingPeriod, err := operatorclient.GetUnbondingTime(ctx.CosmosClient())
	if err != nil {
		return nil, fmt.Errorf("failed to get unbonding time: %w", err)
	}

	trustingPeriod := uint32(unbondingPeriod * 2 / 3)

	if trustingPeriod > uint32(unbondingPeriod) {
		return nil, fmt.Errorf("trusting period %d cannot be greater than unbonding period %d", trustingPeriod, uint32(unbondingPeriod))
	}

	chainId := trustedLightBlock.SignedHeader.Header.ChainID
	revision := clienttypes.ParseChainID(chainId)

	trustThreshold, err := operatorclient.ParseTrustThreshold(trustLevel)
	if err != nil {
		return nil, fmt.Errorf("failed to parse trust level: %w", err)
	}

	var zkAlgorithm operatorclient.SupportedZkAlgorithm
	switch proofType {
	case "groth16":
		zkAlgorithm = operatorclient.Groth16
	case "plonk":
		zkAlgorithm = operatorclient.Plonk
	default:
		return nil, fmt.Errorf("unsupported proof type: %s, supported types are: groth16, plonk", proofType)
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

	return latestLightBlock, w.txHandler.SendTx(ctx, msg)
}

func (w *Worker) UpdateEthClient(ctx Context) error {
	beaconAPIURL := ctx.BeaconAPIURL()
	if beaconAPIURL == "" {
		return fmt.Errorf("beacon API URL is not configured")
	}

	ethClientID := ctx.EthClientID()
	if ethClientID == "" {
		return fmt.Errorf("ethereum client ID is not configured")
	}

	ethClientState, err := operatorclient.GetEthereumClientState(ctx.CosmosClient(), ethClientID)
	if err != nil {
		return fmt.Errorf("failed to get ethereum client state: %w", err)
	}
	trustedSlot := ethClientState.LatestSlot

	finalityUpdate, err := operatorclient.GetFinalityUpdate(beaconAPIURL)
	if err != nil {
		return fmt.Errorf("failed to get finality update: %w", err)
	}

	finalizedSlot, err := parseSlot(finalityUpdate.FinalizedHeader.Beacon.Slot)
	if err != nil {
		return fmt.Errorf("failed to parse finalized slot: %w", err)
	}

	if finalizedSlot <= trustedSlot {
		return nil
	}

	trustedPeriod := ethClientState.ComputeSyncCommitteePeriodAtSlot(trustedSlot)
	targetPeriod := ethClientState.ComputeSyncCommitteePeriodAtSlot(finalizedSlot)

	if targetPeriod > trustedPeriod {
		return w.updateEthClientWithPeriodCrossing(ctx, beaconAPIURL, ethClientID, ethClientState, trustedSlot, trustedPeriod, targetPeriod, finalityUpdate, finalizedSlot)
	}

	return w.updateEthClientSamePeriod(ctx, beaconAPIURL, ethClientID, trustedSlot, finalityUpdate, finalizedSlot)
}

func (w *Worker) updateEthClientSamePeriod(ctx Context, beaconAPIURL, ethClientID string, trustedSlot uint64, finalityUpdate *operatorclient.LightClientFinalityUpdate, _ uint64) error {
	attestedSlot := finalityUpdate.AttestedHeader.Beacon.Slot
	blockRoot, err := operatorclient.GetBeaconBlockRoot(beaconAPIURL, attestedSlot)
	if err != nil {
		return fmt.Errorf("failed to get beacon block root: %w", err)
	}

	bootstrap, err := operatorclient.GetLightClientBootstrap(beaconAPIURL, blockRoot)
	if err != nil {
		return fmt.Errorf("failed to get light client bootstrap: %w", err)
	}

	syncCommittee := bootstrap.Data.CurrentSyncCommittee

	consensusUpdate := operatorclient.LightClientUpdate{
		AttestedHeader:          finalityUpdate.AttestedHeader,
		NextSyncCommittee:       nil,
		NextSyncCommitteeBranch: nil,
		FinalizedHeader:         finalityUpdate.FinalizedHeader,
		FinalityBranch:          finalityUpdate.FinalityBranch,
		SyncAggregate:           finalityUpdate.SyncAggregate,
		SignatureSlot:           finalityUpdate.SignatureSlot,
	}

	header := operatorclient.EthereumHeader{
		ActiveSyncCommittee: operatorclient.ActiveSyncCommittee{
			Current: &syncCommittee,
		},
		ConsensusUpdate: consensusUpdate,
		TrustedSlot:     trustedSlot,
	}

	msg, err := buildMsgUpdateClient(ethClientID, header)
	if err != nil {
		return fmt.Errorf("failed to build update client message: %w", err)
	}

	return w.txHandler.SendCosmosTx(ctx, msg)
}

func (w *Worker) updateEthClientWithPeriodCrossing(ctx Context, beaconAPIURL, ethClientID string, ethClientState *operatorclient.EthereumClientState, trustedSlot, trustedPeriod, targetPeriod uint64, finalityUpdate *operatorclient.LightClientFinalityUpdate, finalizedSlot uint64) error {
	count := targetPeriod - trustedPeriod + 1
	lightClientUpdates, err := operatorclient.GetLightClientUpdates(beaconAPIURL, trustedPeriod, count)
	if err != nil {
		return fmt.Errorf("failed to get light client updates: %w", err)
	}

	if len(lightClientUpdates) == 0 {
		return fmt.Errorf("no light client updates available for period range %d to %d", trustedPeriod, targetPeriod)
	}

	var msgs []any
	latestTrustedSlot := trustedSlot
	latestPeriod := trustedPeriod

	for _, update := range lightClientUpdates {
		updateFinalizedSlot, err := parseSlot(update.FinalizedHeader.Beacon.Slot)
		if err != nil {
			return fmt.Errorf("failed to parse update finalized slot: %w", err)
		}

		if updateFinalizedSlot <= latestTrustedSlot {
			continue
		}

		updatePeriod := ethClientState.ComputeSyncCommitteePeriodAtSlot(updateFinalizedSlot)
		if updatePeriod == latestPeriod {
			continue
		}

		blockRoot, err := operatorclient.GetBeaconBlockRoot(beaconAPIURL, fmt.Sprintf("%d", updateFinalizedSlot))
		if err != nil {
			return fmt.Errorf("failed to get beacon block root: %w", err)
		}

		bootstrap, err := operatorclient.GetLightClientBootstrap(beaconAPIURL, blockRoot)
		if err != nil {
			return fmt.Errorf("failed to get light client bootstrap: %w", err)
		}

		syncCommittee := bootstrap.Data.CurrentSyncCommittee

		header := operatorclient.EthereumHeader{
			ActiveSyncCommittee: operatorclient.ActiveSyncCommittee{
				Next: &syncCommittee,
			},
			ConsensusUpdate: update,
			TrustedSlot:     latestTrustedSlot,
		}

		msg, err := buildMsgUpdateClient(ethClientID, header)
		if err != nil {
			return fmt.Errorf("failed to build update client message: %w", err)
		}

		msgs = append(msgs, msg)
		latestPeriod = updatePeriod
		latestTrustedSlot = updateFinalizedSlot
	}

	if finalizedSlot > latestTrustedSlot {
		attestedSlot := finalityUpdate.AttestedHeader.Beacon.Slot
		blockRoot, err := operatorclient.GetBeaconBlockRoot(beaconAPIURL, attestedSlot)
		if err != nil {
			return fmt.Errorf("failed to get beacon block root: %w", err)
		}

		bootstrap, err := operatorclient.GetLightClientBootstrap(beaconAPIURL, blockRoot)
		if err != nil {
			return fmt.Errorf("failed to get light client bootstrap: %w", err)
		}

		syncCommittee := bootstrap.Data.CurrentSyncCommittee

		consensusUpdate := operatorclient.LightClientUpdate{
			AttestedHeader:          finalityUpdate.AttestedHeader,
			NextSyncCommittee:       nil,
			NextSyncCommitteeBranch: nil,
			FinalizedHeader:         finalityUpdate.FinalizedHeader,
			FinalityBranch:          finalityUpdate.FinalityBranch,
			SyncAggregate:           finalityUpdate.SyncAggregate,
			SignatureSlot:           finalityUpdate.SignatureSlot,
		}

		header := operatorclient.EthereumHeader{
			ActiveSyncCommittee: operatorclient.ActiveSyncCommittee{
				Current: &syncCommittee,
			},
			ConsensusUpdate: consensusUpdate,
			TrustedSlot:     latestTrustedSlot,
		}

		msg, err := buildMsgUpdateClient(ethClientID, header)
		if err != nil {
			return fmt.Errorf("failed to build update client message: %w", err)
		}

		msgs = append(msgs, msg)
	}

	if len(msgs) == 0 {
		return nil
	}

	return w.txHandler.SendCosmosTxBatch(ctx, msgs)
}

// parseSlot parses a slot string to uint64
func parseSlot(slotStr string) (uint64, error) {
	var slot uint64
	_, err := fmt.Sscanf(slotStr, "%d", &slot)
	return slot, err
}

func bytesToBytes32(data []byte) [32]byte {
	var result [32]byte
	copy(result[:], data)
	return result
}

// buildMsgUpdateClient builds a MsgUpdateClient for the Ethereum light client
func buildMsgUpdateClient(clientID string, header operatorclient.EthereumHeader) (*clienttypes.MsgUpdateClient, error) {
	// Serialize the header to JSON
	headerBytes, err := json.Marshal(header)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal header: %w", err)
	}

	// Create the wasm ClientMessage
	clientMessage := &ibcwasmtypes.ClientMessage{
		Data: headerBytes,
	}

	// Pack into Any type
	clientMessageAny, err := codectypes.NewAnyWithValue(clientMessage)
	if err != nil {
		return nil, fmt.Errorf("failed to create Any for client message: %w", err)
	}

	// Build the MsgUpdateClient
	msg := &clienttypes.MsgUpdateClient{
		ClientId:      clientID,
		ClientMessage: clientMessageAny,
		Signer:        "", // Will be set by the transaction handler
	}

	return msg, nil
}
