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
	// Get the beacon API URL and Ethereum client ID from context
	beaconAPIURL := ctx.BeaconAPIURL()
	if beaconAPIURL == "" {
		return fmt.Errorf("beacon API URL is not configured")
	}

	ethClientID := ctx.EthClientID()
	if ethClientID == "" {
		return fmt.Errorf("Ethereum client ID is not configured")
	}

	// Query the Ethereum client state from Cosmos to get the trusted slot
	ethClientState, err := operatorclient.GetEthereumClientState(ctx.CosmosClient(), ethClientID)
	if err != nil {
		return fmt.Errorf("failed to get ethereum client state: %w", err)
	}
	trustedSlot := ethClientState.LatestSlot

	// Get the latest finality update from the beacon API
	finalityUpdate, err := operatorclient.GetFinalityUpdate(beaconAPIURL)
	if err != nil {
		return fmt.Errorf("failed to get finality update: %w", err)
	}

	// Parse the finalized slot to check if update is needed
	finalizedSlot, err := parseSlot(finalityUpdate.FinalizedHeader.Beacon.Slot)
	if err != nil {
		return fmt.Errorf("failed to parse finalized slot: %w", err)
	}

	// Check if update is needed
	if finalizedSlot <= trustedSlot {
		ctx.Logger.Printf("No update needed: finalized slot %d <= trusted slot %d", finalizedSlot, trustedSlot)
		return nil
	}

	// Calculate sync committee periods
	trustedPeriod := ethClientState.ComputeSyncCommitteePeriodAtSlot(trustedSlot)
	targetPeriod := ethClientState.ComputeSyncCommitteePeriodAtSlot(finalizedSlot)

	ctx.Logger.Printf("Updating Ethereum client: trusted slot %d (period %d) -> finalized slot %d (period %d)",
		trustedSlot, trustedPeriod, finalizedSlot, targetPeriod)

	// Check if we're crossing sync committee periods
	if targetPeriod > trustedPeriod {
		// Crossing sync committee periods - fetch light client updates with NextSyncCommittee
		return w.updateEthClientWithPeriodCrossing(ctx, beaconAPIURL, ethClientID, ethClientState, trustedSlot, trustedPeriod, targetPeriod, finalityUpdate, finalizedSlot)
	}

	// Same period - use finality update directly
	return w.updateEthClientSamePeriod(ctx, beaconAPIURL, ethClientID, trustedSlot, finalityUpdate, finalizedSlot)
}

// updateEthClientSamePeriod handles updates within the same sync committee period
func (w *Worker) updateEthClientSamePeriod(ctx Context, beaconAPIURL, ethClientID string, trustedSlot uint64, finalityUpdate *operatorclient.LightClientFinalityUpdate, finalizedSlot uint64) error {
	// Get the block root for the attested slot to fetch the sync committee
	attestedSlot := finalityUpdate.AttestedHeader.Beacon.Slot
	blockRoot, err := operatorclient.GetBeaconBlockRoot(beaconAPIURL, attestedSlot)
	if err != nil {
		return fmt.Errorf("failed to get beacon block root: %w", err)
	}

	// Get bootstrap to get the current sync committee
	bootstrap, err := operatorclient.GetLightClientBootstrap(beaconAPIURL, blockRoot)
	if err != nil {
		return fmt.Errorf("failed to get bootstrap: %w", err)
	}

	syncCommittee := bootstrap.Data.CurrentSyncCommittee

	// Convert finality update to light client update (no next sync committee for same period)
	consensusUpdate := operatorclient.LightClientUpdate{
		AttestedHeader:          finalityUpdate.AttestedHeader,
		NextSyncCommittee:       nil,
		NextSyncCommitteeBranch: nil,
		FinalizedHeader:         finalityUpdate.FinalizedHeader,
		FinalityBranch:          finalityUpdate.FinalityBranch,
		SyncAggregate:           finalityUpdate.SyncAggregate,
		SignatureSlot:           finalityUpdate.SignatureSlot,
	}

	// Build the header with Current sync committee
	header := operatorclient.EthereumHeader{
		ActiveSyncCommittee: operatorclient.ActiveSyncCommittee{
			Current: &syncCommittee,
		},
		ConsensusUpdate: consensusUpdate,
		TrustedSlot:     trustedSlot,
	}

	// Build and send the MsgUpdateClient
	msg, err := buildMsgUpdateClient(ethClientID, header)
	if err != nil {
		return fmt.Errorf("failed to build MsgUpdateClient: %w", err)
	}

	if err := w.txHandler.SendCosmosTx(ctx, msg); err != nil {
		return fmt.Errorf("failed to send update client transaction: %w", err)
	}

	ctx.Logger.Printf("Successfully updated Ethereum client (same period) from slot %d to slot %d", trustedSlot, finalizedSlot)
	return nil
}

// updateEthClientWithPeriodCrossing handles updates that cross sync committee periods
// It sends multiple updates to advance from trustedPeriod to targetPeriod, then a final finality update
func (w *Worker) updateEthClientWithPeriodCrossing(ctx Context, beaconAPIURL, ethClientID string, ethClientState *operatorclient.EthereumClientState, trustedSlot, trustedPeriod, targetPeriod uint64, finalityUpdate *operatorclient.LightClientFinalityUpdate, finalizedSlot uint64) error {
	// Fetch light client updates for the period range
	// We need updates from trustedPeriod to targetPeriod (inclusive)
	count := targetPeriod - trustedPeriod + 1
	lightClientUpdates, err := operatorclient.GetLightClientUpdates(beaconAPIURL, trustedPeriod, count)
	if err != nil {
		return fmt.Errorf("failed to get light client updates: %w", err)
	}

	if len(lightClientUpdates) == 0 {
		return fmt.Errorf("no light client updates available for period range %d to %d", trustedPeriod, targetPeriod)
	}

	ctx.Logger.Printf("Fetched %d light client updates for period crossing", len(lightClientUpdates))

	// Process light client updates for period crossings
	// Following the Rust relayer pattern: process each period crossing update
	latestTrustedSlot := trustedSlot
	latestPeriod := trustedPeriod

	for _, update := range lightClientUpdates {
		updateFinalizedSlot, err := parseSlot(update.FinalizedHeader.Beacon.Slot)
		if err != nil {
			return fmt.Errorf("failed to parse update finalized slot: %w", err)
		}

		// Skip updates that don't advance beyond the trusted slot
		if updateFinalizedSlot <= latestTrustedSlot {
			ctx.Logger.Printf("Skipping update for slot %d (not advancing beyond %d)", updateFinalizedSlot, latestTrustedSlot)
			continue
		}

		updatePeriod := ethClientState.ComputeSyncCommitteePeriodAtSlot(updateFinalizedSlot)

		// Skip updates that don't advance the period
		if updatePeriod == latestPeriod {
			ctx.Logger.Printf("Skipping update for slot %d (same period %d)", updateFinalizedSlot, updatePeriod)
			continue
		}

		// This update crosses a period boundary
		// Get the sync committee for the attested slot (this becomes the "next" sync committee)
		attestedSlot := update.AttestedHeader.Beacon.Slot
		blockRoot, err := operatorclient.GetBeaconBlockRoot(beaconAPIURL, attestedSlot)
		if err != nil {
			return fmt.Errorf("failed to get beacon block root for period crossing: %w", err)
		}

		bootstrap, err := operatorclient.GetLightClientBootstrap(beaconAPIURL, blockRoot)
		if err != nil {
			return fmt.Errorf("failed to get bootstrap for period crossing: %w", err)
		}

		syncCommittee := bootstrap.Data.CurrentSyncCommittee

		// Build header with Next sync committee (for period crossing)
		// The update from GetLightClientUpdates includes NextSyncCommittee and NextSyncCommitteeBranch
		header := operatorclient.EthereumHeader{
			ActiveSyncCommittee: operatorclient.ActiveSyncCommittee{
				Next: &syncCommittee,
			},
			ConsensusUpdate: update, // This update has NextSyncCommittee and NextSyncCommitteeBranch populated
			TrustedSlot:     latestTrustedSlot,
		}

		// Build and send the MsgUpdateClient
		msg, err := buildMsgUpdateClient(ethClientID, header)
		if err != nil {
			return fmt.Errorf("failed to build MsgUpdateClient for period crossing: %w", err)
		}

		if err := w.txHandler.SendCosmosTx(ctx, msg); err != nil {
			return fmt.Errorf("failed to send update client transaction for period crossing: %w", err)
		}

		ctx.Logger.Printf("Successfully updated Ethereum client (period crossing) from slot %d (period %d) to slot %d (period %d)",
			latestTrustedSlot, latestPeriod, updateFinalizedSlot, updatePeriod)

		// Update tracking variables
		latestPeriod = updatePeriod
		latestTrustedSlot = updateFinalizedSlot
	}

	// After processing period crossings, check if we need to send a final finality update
	// This is needed if the finality update's slot is newer than the last period crossing update
	if finalizedSlot > latestTrustedSlot {
		ctx.Logger.Printf("Sending final finality update from slot %d to slot %d", latestTrustedSlot, finalizedSlot)
		return w.updateEthClientSamePeriod(ctx, beaconAPIURL, ethClientID, latestTrustedSlot, finalityUpdate, finalizedSlot)
	}

	return nil
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
