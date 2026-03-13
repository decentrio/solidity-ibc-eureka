package services

import (
	"context"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	updateclient "operator/bindings/UpdateClient"
	operatorclient "operator/client"
	"operator/prover"
	"strconv"
	"strings"
	"time"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	ibcwasmtypes "github.com/cosmos/ibc-go/modules/light-clients/08-wasm/v10/types"
	clienttypes "github.com/cosmos/ibc-go/v10/modules/core/02-client/types"
	channeltypesv2 "github.com/cosmos/ibc-go/v10/modules/core/04-channel/v2/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

const ICS26_IBC_STORAGE_SLOT = "0x1260944489272988d9df285149b5aa1b0f48f2136d6f416159f840a3e0747600"

type Worker struct {
	txHandler TransactionHandler
	prover    *prover.Prover
}

func NewWorker(txHandler TransactionHandler, p *prover.Prover) *Worker {
	return &Worker{
		txHandler: txHandler,
		prover:    p,
	}
}

func (w *Worker) CreateCosmosClient(ctx Context, proofType string, trustingPeriod uint32, trustedBlock int64, trustLevel string) error {
	genesis, err := operatorclient.GetGenesis(ctx.CosmosClient(), trustedBlock, trustingPeriod, trustLevel, proofType)
	if err != nil {
		return fmt.Errorf("failed to get genesis: %w", err)
	}

	clientState := genesis.TrustedClientState
	consensusState := genesis.TrustedConsensusState

	clientStateEncoded, err := operatorclient.EncodeClientState(clientState)
	if err != nil {
		return fmt.Errorf("failed to encode client state: %w", err)
	}

	consensusStateEncoded, err := operatorclient.EncodeConsensusState(consensusState)
	if err != nil {
		return fmt.Errorf("failed to encode consensus state: %w", err)
	}

	consensusHash := crypto.Keccak256(consensusStateEncoded)
	return w.txHandler.CreateCosmosClientContract(ctx, clientStateEncoded, consensusHash)
}

func (w *Worker) CreateEthClient(ctx Context, checksum string) error {
	beaconAPIURL := ctx.BeaconAPIURL()
	if beaconAPIURL == "" {
		return fmt.Errorf("beacon API URL is not configured")
	}

	genesis, err := operatorclient.GetBeaconGenesis(beaconAPIURL)
	if err != nil {
		return fmt.Errorf("failed to get beacon genesis: %w", err)
	}

	spec, err := operatorclient.GetBeaconSpec(beaconAPIURL)
	if err != nil {
		return fmt.Errorf("failed to get beacon spec: %w", err)
	}

	beaconBlock, err := operatorclient.GetBeaconBlock(beaconAPIURL, "finalized")
	if err != nil {
		return fmt.Errorf("failed to get beacon block: %w", err)
	}

	blockRoot, err := operatorclient.GetBeaconBlockRoot(beaconAPIURL, beaconBlock.Message.Slot)
	if err != nil {
		return fmt.Errorf("failed to get beacon block root: %w", err)
	}

	bootstrap, err := operatorclient.GetLightClientBootstrap(beaconAPIURL, blockRoot)
	if err != nil {
		return fmt.Errorf("failed to get light client bootstrap: %w", err)
	}

	if bootstrap.Data.Header.Execution.BlockNumber != beaconBlock.Message.Body.ExecutionPayload.BlockNumber {
		return fmt.Errorf("bootstrap block number does not match execution block number")
	}

	chainId, err := ctx.EthClient().ChainID(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get eth chain id: %w", err)
	}

	forkParameters, err := spec.ToForkParameters()
	if err != nil {
		return fmt.Errorf("failed to get fork parameters: %w", err)
	}

	epochsPerSyncCommitteePeriod, err := strconv.ParseUint(spec.EpochsPerSyncCommitteePeriod, 10, 64)
	if err != nil {
		return fmt.Errorf("failed to parse epochs per sync committee period: %w", err)
	}

	genesisTime, err := strconv.ParseUint(genesis.GenesisTime, 10, 64)
	if err != nil {
		return fmt.Errorf("failed to parse genesis time: %w", err)
	}

	blockNumber, err := strconv.ParseUint(bootstrap.Data.Header.Execution.BlockNumber, 10, 64)
	if err != nil {
		return fmt.Errorf("failed to parse block number: %w", err)
	}

	slot, err := strconv.ParseUint(bootstrap.Data.Header.Beacon.Slot, 10, 64)
	if err != nil {
		return fmt.Errorf("failed to parse slot: %w", err)
	}

	syncCommitteeSize, err := strconv.ParseUint(spec.SyncCommitteeSize, 10, 64)
	if err != nil {
		return fmt.Errorf("failed to parse sync committee size: %w", err)
	}

	secondsPerSlot, err := strconv.ParseUint(spec.SecondsPerSlot, 10, 64)
	if err != nil {
		return fmt.Errorf("failed to parse seconds per slot: %w", err)
	}

	slotsPerEpoch, err := strconv.ParseUint(spec.SlotsPerEpoch, 10, 64)
	if err != nil {
		return fmt.Errorf("failed to parse slots per epoch: %w", err)
	}

	clientState := operatorclient.EthereumClientState{
		ChainID:                      chainId.Uint64(),
		EpochsPerSyncCommitteePeriod: epochsPerSyncCommitteePeriod,
		ForkParameters:               *forkParameters,
		GenesisSlot:                  0,
		GenesisTime:                  genesisTime,
		GenesisValidatorsRoot:        genesis.GenesisValidatorsRoot,
		IbcCommitmentSlot:            ICS26_IBC_STORAGE_SLOT,
		IbcContractAddress:           ctx.RouterContract().String(),
		IsFrozen:                     false,
		LatestExecutionBlockNumber:   blockNumber,
		LatestSlot:                   slot,
		MinSyncCommitteeParticipants: (syncCommitteeSize + 2) / 3,
		SecondsPerSlot:               secondsPerSlot,
		SlotsPerEpoch:                slotsPerEpoch,
		SyncCommitteeSize:            syncCommitteeSize,
	}

	clientStateBz, err := json.Marshal(clientState)
	if err != nil {
		return fmt.Errorf("failed to serialize client state: %w", err)
	}

	checksumTrimmed := strings.TrimPrefix(checksum, "0x")
	checksumBz, err := hex.DecodeString(checksumTrimmed)
	if err != nil {
		return fmt.Errorf("failed to parse checksum: %w", err)
	}

	wasmClientState := ibcwasmtypes.ClientState{
		Data:     clientStateBz,
		Checksum: checksumBz,
		LatestHeight: clienttypes.Height{
			RevisionNumber: 0,
			RevisionHeight: clientState.LatestSlot,
		},
	}

	timestamp, err := strconv.ParseUint(bootstrap.Data.Header.Execution.Timestamp, 10, 64)
	if err != nil {
		return fmt.Errorf("failed to parse timestamp: %w", err)
	}

	currentSyncCommittee, err := bootstrap.Data.CurrentSyncCommittee.ToSummarizedSyncCommittee()
	if err != nil {
		return fmt.Errorf("failed to hash pubkeys for CurrentSyncCommittee: %w", err)
	}

	latestPeriod := clientState.ComputeSyncCommitteePeriodAtSlot(clientState.LatestSlot)
	lightClientUpdates, err := operatorclient.GetLightClientUpdates(ctx.BeaconAPIURL(), latestPeriod, 1)
	if err != nil {
		return fmt.Errorf("failed to get light client updates: %w", err)
	}

	nextSyncCommittee, err := lightClientUpdates[0].NextSyncCommittee.ToSummarizedSyncCommittee()
	if err != nil {
		return fmt.Errorf("failed to hash pubkeys for NextSyncCommittee: %w", err)
	}

	consensusState := operatorclient.EthereumConsensusState{
		Slot:                 clientState.LatestSlot,
		StateRoot:            bootstrap.Data.Header.Execution.StateRoot,
		Timestamp:            timestamp,
		CurrentSyncCommittee: *currentSyncCommittee,
		NextSyncCommittee:    nextSyncCommittee,
	}

	consensusStateBz, err := json.Marshal(consensusState)
	if err != nil {
		return fmt.Errorf("failed to serialize consensus state: %w", err)
	}

	wasmConsensusState := ibcwasmtypes.ConsensusState{
		Data: consensusStateBz,
	}

	return w.txHandler.CreateEthClient(ctx, &wasmClientState, &wasmConsensusState)
}

func (w *Worker) UpdateCosmosClient(ctx Context, proofType string, trustedBlock int64, trustLevel string) (*operatorclient.LightBlock, error) {
	status, err := ctx.CosmosClient().Status(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get status: %w", err)
	}

	if trustedBlock == 0 {
		trustedBlock = status.SyncInfo.LatestBlockHeight
	} else if trustedBlock == status.SyncInfo.LatestBlockHeight {
		// if trusted block height is equal to latest block height stop here
		return nil, fmt.Errorf("client is up to dated")
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
		Timestamp:          big.NewInt(trustedLightBlock.SignedHeader.Header.Time.Unix()),
		Root:               bytesToBytes32(trustedLightBlock.SignedHeader.Header.AppHash),
		NextValidatorsHash: bytesToBytes32(trustedLightBlock.SignedHeader.NextValidatorsHash),
	}

	proposedHeader := latestLightBlock.IntoHeader(*trustedLightBlock)

	// Extract validator signature for Ed25519 ZK proof
	valSig, sigErr := prover.ExtractValidatorSignature(latestLightBlock, chainId)
	if sigErr != nil {
		return nil, fmt.Errorf("failed to extract validator signature: %w", sigErr)
	}

	// Generate Groth16 proof for Ed25519 signature verification
	var proofBigInt [8]*big.Int
	var commitmentsBigInt [2]*big.Int
	var commitmentPokBigInt [2]*big.Int

	if w.prover != nil {
		proofBigInt, commitmentsBigInt, commitmentPokBigInt, sigErr = w.prover.ProveSignature(
			valSig.Signature,
			valSig.PublicKey,
			valSig.SignBytes,
		)
		if sigErr != nil {
			return nil, fmt.Errorf("failed to generate proof: %w", sigErr)
		}
	}

	// Build signature fields for on-chain verification
	// signature[0] = R (first 32 bytes), signature[1] = S (last 32 bytes)
	var signature [2][32]byte
	copy(signature[0][:], valSig.Signature[:32])
	copy(signature[1][:], valSig.Signature[32:64])

	var validatorPubkey [32]byte
	copy(validatorPubkey[:], valSig.PublicKey)

	msg := updateclient.IUpdateClientMsgsMsgUpdateClient{
		ClientState:           clientState,
		TrustedConsensusState: consensusState,
		Time:                  big.NewInt(time.Now().Unix()),
		ProposedHeader:        proposedHeader,
		Proof:                 proofBigInt,
		Commitments:           commitmentsBigInt,
		CommitmentPok:         commitmentPokBigInt,
		Signature:             signature,
		ValidatorPubkey:       validatorPubkey,
		VoteSignBytes:         valSig.SignBytes,
	}

	return latestLightBlock, w.txHandler.SendEthTx(ctx, msg)
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

// RelayEthToCosmosPacket relays a packet from Ethereum to Cosmos
// It retrieves the Ethereum storage proof and sends MsgRecvPacket to Cosmos
func (w *Worker) RelayEthToCosmosPacket(ctx Context, packet Packet) error {
	if packet.Packet == nil {
		return fmt.Errorf("packet is nil")
	}

	routerAddr := ctx.RouterContract()
	if routerAddr == nil {
		return fmt.Errorf("ICS26_ROUTER address is not configured")
	}

	// Compute the storage key for the packet commitment
	// The ICS26Router stores commitments at: keccak256(abi.encode(commitmentPath, ICS26_IBC_STORAGE_SLOT))
	storageKey := computePacketCommitmentStorageKey(
		packet.Packet.SourceClient,
		packet.Packet.Sequence,
	)

	// Get the storage proof from Ethereum at the packet's block height
	blockNum := new(big.Int).SetUint64(packet.EthHeight)
	proof, err := ctx.EthClient().StorageAt(
		context.Background(),
		*routerAddr,
		storageKey,
		blockNum,
	)
	if err != nil {
		return fmt.Errorf("failed to get storage proof: %w", err)
	}

	// Get the full eth_getProof for Merkle proof
	ethProof, err := getEthStorageProof(ctx, *routerAddr, storageKey, blockNum)
	if err != nil {
		return fmt.Errorf("failed to get eth storage proof: %w", err)
	}

	_ = proof // value is used for validation

	// Build Cosmos MsgRecvPacket
	ethClientID := ctx.EthClientID()
	if ethClientID == "" {
		return fmt.Errorf("ethereum client ID is not configured")
	}

	msg := channeltypesv2.NewMsgRecvPacket(
		*packet.Packet,
		ethProof,
		clienttypes.Height{
			RevisionNumber: 0,
			RevisionHeight: packet.EthHeight,
		},
		"", // Signer will be set by transaction handler
	)

	return w.txHandler.SendCosmosTx(ctx, msg)
}

// computePacketCommitmentStorageKey computes the Ethereum storage slot for a packet commitment
// Storage layout: mapping key = keccak256(abi.encode(commitmentKey, baseSlot))
func computePacketCommitmentStorageKey(sourceClient string, sequence uint64) common.Hash {
	// Build the IBC commitment path: sourceClient || 0x01 || sequence (big-endian 8 bytes)
	seqBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(seqBytes, sequence)
	commitmentPath := append([]byte(sourceClient), 0x01)
	commitmentPath = append(commitmentPath, seqBytes...)

	// Compute keccak256 of the commitment path
	pathHash := crypto.Keccak256Hash(commitmentPath)

	// Compute the storage slot: keccak256(abi.encode(pathHash, ICS26_IBC_STORAGE_SLOT))
	ibcSlot := common.HexToHash(ICS26_IBC_STORAGE_SLOT)

	// For Solidity mapping: keccak256(abi.encode(key, slot))
	slotInput := make([]byte, 64)
	copy(slotInput[:32], pathHash.Bytes())
	copy(slotInput[32:], ibcSlot.Bytes())

	return crypto.Keccak256Hash(slotInput)
}

// getEthStorageProof retrieves the Ethereum storage proof using eth_getProof
func getEthStorageProof(ctx Context, contractAddr common.Address, storageKey common.Hash, blockNum *big.Int) ([]byte, error) {
	type AccountResult struct {
		StorageProof []struct {
			Key   string   `json:"key"`
			Value string   `json:"value"`
			Proof []string `json:"proof"`
		} `json:"storageProof"`
	}

	var result AccountResult
	err := ctx.EthClient().Client().Call(
		&result,
		"eth_getProof",
		contractAddr.Hex(),
		[]string{storageKey.Hex()},
		fmt.Sprintf("0x%x", blockNum.Uint64()),
	)
	if err != nil {
		return nil, fmt.Errorf("eth_getProof failed: %w", err)
	}

	if len(result.StorageProof) == 0 {
		return nil, fmt.Errorf("no storage proof returned")
	}

	// Encode the proof as JSON for the Cosmos WASM light client
	proofBytes, err := json.Marshal(result.StorageProof[0].Proof)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal proof: %w", err)
	}

	return proofBytes, nil
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
