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

type ForkParameters struct {
	Altair             Fork   `json:"altair"`
	Bellatrix          Fork   `json:"bellatrix"`
	Capella            Fork   `json:"capella"`
	Deneb              Fork   `json:"deneb"`
	Electra            Fork   `json:"electra"`
	GenesisForkVersion string `json:"genesis_fork_version"`
	GenesisSlot        uint64 `json:"genesis_slot"`
}

type Fork struct {
	Epoch   uint64 `json:"epoch"`
	Version string `json:"version"`
}

func (cs *EthereumClientState) ComputeSyncCommitteePeriodAtSlot(slot uint64) uint64 {
	epoch := slot / cs.SlotsPerEpoch
	return epoch / cs.EpochsPerSyncCommitteePeriod
}

type SyncCommittee struct {
	Pubkeys         []string `json:"pubkeys"`
	AggregatePubkey string   `json:"aggregate_pubkey"`
}

type SummarizedSyncCommittee struct {
	PubkeysHash     string `json:"pubkeys_hash"`
	AggregatePubkey string `json:"aggregate_pubkey"`
}

type LightClientHeader struct {
	Beacon          BeaconBlockHeader      `json:"beacon"`
	Execution       ExecutionPayloadHeader `json:"execution"`
	ExecutionBranch []string               `json:"execution_branch"`
}

type BeaconBlockHeader struct {
	Slot          string `json:"slot"`
	ProposerIndex string `json:"proposer_index"`
	ParentRoot    string `json:"parent_root"`
	StateRoot     string `json:"state_root"`
	BodyRoot      string `json:"body_root"`
}

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

type SyncAggregate struct {
	SyncCommitteeBits      string `json:"sync_committee_bits"`
	SyncCommitteeSignature string `json:"sync_committee_signature"`
}

type LightClientUpdate struct {
	AttestedHeader          LightClientHeader `json:"attested_header"`
	NextSyncCommittee       *SyncCommittee    `json:"next_sync_committee"`
	NextSyncCommitteeBranch []string          `json:"next_sync_committee_branch"`
	FinalizedHeader         LightClientHeader `json:"finalized_header"`
	FinalityBranch          []string          `json:"finality_branch"`
	SyncAggregate           SyncAggregate     `json:"sync_aggregate"`
	SignatureSlot           string            `json:"signature_slot"`
}

type LightClientFinalityUpdate struct {
	AttestedHeader  LightClientHeader `json:"attested_header"`
	FinalizedHeader LightClientHeader `json:"finalized_header"`
	FinalityBranch  []string          `json:"finality_branch"`
	SyncAggregate   SyncAggregate     `json:"sync_aggregate"`
	SignatureSlot   string            `json:"signature_slot"`
}

type ActiveSyncCommittee struct {
	Current *SyncCommittee `json:"Current,omitempty"`
	Next    *SyncCommittee `json:"Next,omitempty"`
}

type EthereumHeader struct {
	ActiveSyncCommittee ActiveSyncCommittee `json:"active_sync_committee"`
	ConsensusUpdate     LightClientUpdate   `json:"consensus_update"`
	TrustedSlot         uint64              `json:"trusted_slot"`
}

type FinalityUpdateResponse struct {
	Version string                    `json:"version"`
	Data    LightClientFinalityUpdate `json:"data"`
}

type LightClientUpdateResponse struct {
	Data LightClientUpdate `json:"data"`
}

type BootstrapResponse struct {
	Data struct {
		Header               LightClientHeader `json:"header"`
		CurrentSyncCommittee SyncCommittee     `json:"current_sync_committee"`
	} `json:"data"`
}

type BeaconBlockRootResponse struct {
	Data struct {
		Root string `json:"root"`
	} `json:"data"`
}

func httpGet[T any](url string) (T, error) {
	var result T

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return result, err
	}
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}

	if resp.StatusCode != 200 {
		return result, fmt.Errorf("request failed with status %d: %s", resp.StatusCode, body)
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return result, err
	}

	return result, nil
}

func GetFinalityUpdate(beaconAPIURL string) (*LightClientFinalityUpdate, error) {
	url := fmt.Sprintf("%s/eth/v1/beacon/light_client/finality_update", beaconAPIURL)
	response, err := httpGet[FinalityUpdateResponse](url)
	if err != nil {
		return nil, fmt.Errorf("failed to get finality update: %w", err)
	}
	return &response.Data, nil
}

func GetLightClientUpdates(beaconAPIURL string, startPeriod, count uint64) ([]LightClientUpdate, error) {
	url := fmt.Sprintf("%s/eth/v1/beacon/light_client/updates?start_period=%d&count=%d", beaconAPIURL, startPeriod, count)
	responses, err := httpGet[[]LightClientUpdateResponse](url)
	if err != nil {
		return nil, fmt.Errorf("failed to get light client updates: %w", err)
	}

	updates := make([]LightClientUpdate, len(responses))
	for i, r := range responses {
		updates[i] = r.Data
	}
	return updates, nil
}

func GetBeaconBlockRoot(beaconAPIURL, blockID string) (string, error) {
	url := fmt.Sprintf("%s/eth/v1/beacon/blocks/%s/root", beaconAPIURL, blockID)
	response, err := httpGet[BeaconBlockRootResponse](url)
	if err != nil {
		return "", fmt.Errorf("failed to get beacon block root: %w", err)
	}
	return response.Data.Root, nil
}

func GetLightClientBootstrap(beaconAPIURL, blockRoot string) (*BootstrapResponse, error) {
	url := fmt.Sprintf("%s/eth/v1/beacon/light_client/bootstrap/%s", beaconAPIURL, blockRoot)
	response, err := httpGet[BootstrapResponse](url)
	if err != nil {
		return nil, fmt.Errorf("failed to get light client bootstrap: %w", err)
	}
	return &response, nil
}

func GetEthereumClientState(cosmosClient *rpchttp.HTTP, clientID string) (*EthereumClientState, error) {
	queryReq := &clienttypes.QueryClientStateRequest{
		ClientId: clientID,
	}

	reqBytes, err := proto.Marshal(queryReq)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal query request: %w", err)
	}

	result, err := cosmosClient.ABCIQuery(context.Background(), "/ibc.core.client.v1.Query/ClientState", reqBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to query client state: %w", err)
	}

	if result.Response.Code != 0 {
		return nil, fmt.Errorf("query failed with code %d: %s", result.Response.Code, result.Response.Log)
	}

	var queryResp clienttypes.QueryClientStateResponse
	if err := proto.Unmarshal(result.Response.Value, &queryResp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal query response: %w", err)
	}

	var wasmClientState ibcwasmtypes.ClientState
	if err := proto.Unmarshal(queryResp.ClientState.Value, &wasmClientState); err != nil {
		return nil, fmt.Errorf("failed to unmarshal wasm client state: %w", err)
	}

	var ethClientState EthereumClientState
	if err := json.Unmarshal(wasmClientState.Data, &ethClientState); err != nil {
		return nil, fmt.Errorf("failed to unmarshal ethereum client state: %w", err)
	}

	return &ethClientState, nil
}
