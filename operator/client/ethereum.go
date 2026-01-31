package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
	"github.com/cosmos/gogoproto/proto"
	ibcwasmtypes "github.com/cosmos/ibc-go/modules/light-clients/08-wasm/v10/types"
	clienttypes "github.com/cosmos/ibc-go/v10/modules/core/02-client/types"
)

// EthereumClientState represents the Ethereum light client state stored on Cosmos
type EthereumClientState struct {
	ChainID                      uint64         `json:"chain_id"`
	EpochsPerSyncCommitteePeriod uint64         `json:"epochs_per_sync_committee_period"`
	ForkParameters               ForkParameters `json:"fork_parameters"`
	GenesisSlot                  uint64         `json:"genesis_slot"`
	GenesisTime                  uint64         `json:"genesis_time"`
	GenesisValidatorsRoot        string         `json:"genesis_validators_root"`
	IbcCommitmentSlot            string         `json:"ibc_commitment_slot"`
	IbcContractAddress           string         `json:"ibc_contract_address"`
	IsFrozen                     bool           `json:"is_frozen"`
	LatestExecutionBlockNumber   uint64         `json:"latest_execution_block_number"`
	LatestSlot                   uint64         `json:"latest_slot"`
	MinSyncCommitteeParticipants uint64         `json:"min_sync_committee_participants"`
	SecondsPerSlot               uint64         `json:"seconds_per_slot"`
	SlotsPerEpoch                uint64         `json:"slots_per_epoch"`
	SyncCommitteeSize            uint64         `json:"sync_committee_size"`
}

// ForkParameters represents the fork parameters for Ethereum
type ForkParameters struct {
	Altair             Fork   `json:"altair"`
	Bellatrix          Fork   `json:"bellatrix"`
	Capella            Fork   `json:"capella"`
	Deneb              Fork   `json:"deneb"`
	Electra            Fork   `json:"electra"`
	GenesisForkVersion string `json:"genesis_fork_version"`
	GenesisSlot        uint64 `json:"genesis_slot"`
}

// Fork represents a single fork
type Fork struct {
	Epoch   uint64 `json:"epoch"`
	Version string `json:"version"`
}

// ComputeSyncCommitteePeriodAtSlot computes the sync committee period for a given slot
func (cs *EthereumClientState) ComputeSyncCommitteePeriodAtSlot(slot uint64) uint64 {
	epoch := slot / cs.SlotsPerEpoch
	return epoch / cs.EpochsPerSyncCommitteePeriod
}

// SyncCommittee represents the sync committee data
type SyncCommittee struct {
	Pubkeys         []string `json:"pubkeys"`
	AggregatePubkey string   `json:"aggregate_pubkey"`
}

// SummarizedSyncCommittee represents the summarized sync committee data
type SummarizedSyncCommittee struct {
	PubkeysHash     string `json:"pubkeys_hash"`
	AggregatePubkey string `json:"aggregate_pubkey"`
}

// LightClientHeader represents a light client header
type LightClientHeader struct {
	Beacon          BeaconBlockHeader      `json:"beacon"`
	Execution       ExecutionPayloadHeader `json:"execution"`
	ExecutionBranch []string               `json:"execution_branch"`
}

// BeaconBlockHeader represents a beacon block header
type BeaconBlockHeader struct {
	Slot          string `json:"slot"`
	ProposerIndex string `json:"proposer_index"`
	ParentRoot    string `json:"parent_root"`
	StateRoot     string `json:"state_root"`
	BodyRoot      string `json:"body_root"`
}

// ExecutionPayloadHeader represents the execution payload header
type ExecutionPayloadHeader struct {
	ParentHash       string `json:"parent_hash"`
	FeeRecipient     string `json:"fee_recipient"`
	StateRoot        string `json:"state_root"`
	ReceiptsRoot     string `json:"receipts_root"`
	LogsBloom        string `json:"logs_bloom"`
	PrevRandao       string `json:"prev_randao"`
	BlockNumber      string `json:"block_number"`
	GasLimit         string `json:"gas_limit"`
	GasUsed          string `json:"gas_used"`
	Timestamp        string `json:"timestamp"`
	ExtraData        string `json:"extra_data"`
	BaseFeePerGas    string `json:"base_fee_per_gas"`
	BlockHash        string `json:"block_hash"`
	TransactionsRoot string `json:"transactions_root"`
	WithdrawalsRoot  string `json:"withdrawals_root"`
	BlobGasUsed      string `json:"blob_gas_used"`
	ExcessBlobGas    string `json:"excess_blob_gas"`
}

// SyncAggregate represents the sync aggregate
type SyncAggregate struct {
	SyncCommitteeBits      string `json:"sync_committee_bits"`
	SyncCommitteeSignature string `json:"sync_committee_signature"`
}

// LightClientUpdate represents a light client update from the beacon API
type LightClientUpdate struct {
	AttestedHeader          LightClientHeader `json:"attested_header"`
	NextSyncCommittee       *SyncCommittee    `json:"next_sync_committee"`
	NextSyncCommitteeBranch []string          `json:"next_sync_committee_branch"`
	FinalizedHeader         LightClientHeader `json:"finalized_header"`
	FinalityBranch          []string          `json:"finality_branch"`
	SyncAggregate           SyncAggregate     `json:"sync_aggregate"`
	SignatureSlot           string            `json:"signature_slot"`
}

// LightClientFinalityUpdate represents a finality update from the beacon API
type LightClientFinalityUpdate struct {
	AttestedHeader  LightClientHeader `json:"attested_header"`
	FinalizedHeader LightClientHeader `json:"finalized_header"`
	FinalityBranch  []string          `json:"finality_branch"`
	SyncAggregate   SyncAggregate     `json:"sync_aggregate"`
	SignatureSlot   string            `json:"signature_slot"`
}

// ActiveSyncCommittee represents the active sync committee (either current or next)
type ActiveSyncCommittee struct {
	Current *SyncCommittee `json:"Current,omitempty"`
	Next    *SyncCommittee `json:"Next,omitempty"`
}

// EthereumHeader represents the header for updating the Ethereum light client
type EthereumHeader struct {
	ActiveSyncCommittee ActiveSyncCommittee `json:"active_sync_committee"`
	ConsensusUpdate     LightClientUpdate   `json:"consensus_update"`
	TrustedSlot         uint64              `json:"trusted_slot"`
}

// FinalityUpdateResponse represents the response from the finality update endpoint
type FinalityUpdateResponse struct {
	Version string                    `json:"version"`
	Data    LightClientFinalityUpdate `json:"data"`
}

// LightClientUpdateResponse represents a single light client update response
type LightClientUpdateResponse struct {
	Data LightClientUpdate `json:"data"`
}

// BootstrapResponse represents the response from the bootstrap endpoint
type BootstrapResponse struct {
	Data struct {
		Header               LightClientHeader `json:"header"`
		CurrentSyncCommittee SyncCommittee     `json:"current_sync_committee"`
	} `json:"data"`
}

// BeaconBlockRootResponse represents the response for beacon block root
type BeaconBlockRootResponse struct {
	Data struct {
		Root string `json:"root"`
	} `json:"data"`
}

// GetFinalityUpdate fetches the latest finality update from the beacon API
func GetFinalityUpdate(beaconAPIURL string) (*LightClientFinalityUpdate, error) {
	url := fmt.Sprintf("%s/eth/v1/beacon/light_client/finality_update", beaconAPIURL)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch finality update: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("finality update request failed with status %d: %s", resp.StatusCode, string(body))
	}

	var response FinalityUpdateResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal finality update: %w", err)
	}

	return &response.Data, nil
}

// GetLightClientUpdates fetches light client updates from the beacon API
func GetLightClientUpdates(beaconAPIURL string, startPeriod, count uint64) ([]LightClientUpdate, error) {
	url := fmt.Sprintf("%s/eth/v1/beacon/light_client/updates?start_period=%d&count=%d", beaconAPIURL, startPeriod, count)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch light client updates: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("light client updates request failed with status %d: %s", resp.StatusCode, string(body))
	}

	var responses []LightClientUpdateResponse
	if err := json.Unmarshal(body, &responses); err != nil {
		return nil, fmt.Errorf("failed to unmarshal light client updates: %w", err)
	}

	updates := make([]LightClientUpdate, len(responses))
	for i, r := range responses {
		updates[i] = r.Data
	}
	return updates, nil
}

// GetBeaconBlockRoot fetches the beacon block root for a given block ID
func GetBeaconBlockRoot(beaconAPIURL, blockID string) (string, error) {
	url := fmt.Sprintf("%s/eth/v1/beacon/blocks/%s/root", beaconAPIURL, blockID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to fetch beacon block root: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("beacon block root request failed with status %d: %s", resp.StatusCode, string(body))
	}

	var response BeaconBlockRootResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return "", fmt.Errorf("failed to unmarshal beacon block root: %w", err)
	}

	return response.Data.Root, nil
}

// GetLightClientBootstrap fetches the light client bootstrap for a given block root
func GetLightClientBootstrap(beaconAPIURL, blockRoot string) (*BootstrapResponse, error) {
	url := fmt.Sprintf("%s/eth/v1/beacon/light_client/bootstrap/%s", beaconAPIURL, blockRoot)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch bootstrap: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("bootstrap request failed with status %d: %s", resp.StatusCode, string(body))
	}

	var response BootstrapResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal bootstrap: %w", err)
	}

	return &response, nil
}

// GetEthereumClientState queries the Ethereum client state from Cosmos via ABCI
// This is a simplified version that uses direct proto marshaling without heavy interface registry
func GetEthereumClientState(cosmosClient *rpchttp.HTTP, clientID string) (*EthereumClientState, error) {
	// Build the query request using direct proto marshaling
	queryReq := &clienttypes.QueryClientStateRequest{
		ClientId: clientID,
	}

	reqBytes, err := proto.Marshal(queryReq)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal query request: %w", err)
	}

	// Query path for IBC client state
	queryPath := "/ibc.core.client.v1.Query/ClientState"

	// Make ABCI query
	result, err := cosmosClient.ABCIQuery(context.Background(), queryPath, reqBytes)
	if err != nil {
		return nil, fmt.Errorf("ABCI query failed: %w", err)
	}

	if result.Response.Code != 0 {
		return nil, fmt.Errorf("query failed with code %d: %s", result.Response.Code, result.Response.Log)
	}

	// Decode the response using direct proto unmarshaling
	var queryResp clienttypes.QueryClientStateResponse
	if err := proto.Unmarshal(result.Response.Value, &queryResp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal query response: %w", err)
	}

	// Decode the wasm client state directly from the Any value
	var wasmClientState ibcwasmtypes.ClientState
	if err := proto.Unmarshal(queryResp.ClientState.Value, &wasmClientState); err != nil {
		return nil, fmt.Errorf("failed to unmarshal wasm client state: %w", err)
	}

	// Decode the Ethereum client state from wasm data
	var ethClientState EthereumClientState
	if err := json.Unmarshal(wasmClientState.Data, &ethClientState); err != nil {
		return nil, fmt.Errorf("failed to unmarshal ethereum client state: %w", err)
	}

	return &ethClientState, nil
}
