package client

import (
	"context"
	"fmt"
	"math"
	"slices"

	"github.com/cometbft/cometbft/p2p"
	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
	commettypes "github.com/cometbft/cometbft/types"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

const MAX_TOTAL_VOTING_POWER = math.MaxInt64 / 8

type LightBlock struct {
	SignedHeader commettypes.SignedHeader
	ValSet       commettypes.ValidatorSet
	NextValSet   commettypes.ValidatorSet
	PeerId       p2p.ID
	BlockHeight  int64
}

func getUnbondingTime(client *rpchttp.HTTP) (float64, error) {

	// Query path for staking parameters
	queryPath := "/cosmos.staking.v1beta1.Query/Params"

	// Make ABCI query
	result, err := client.ABCIQuery(context.Background(), queryPath, nil)
	if err != nil {
		return 0, fmt.Errorf("ABCI query failed: %w", err)
	}

	if result.Response.Code != 0 {
		return 0, fmt.Errorf("query failed with code %d: %s", result.Response.Code, result.Response.Log)
	}

	// Create codec for decoding
	interfaceRegistry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(interfaceRegistry)

	// Decode the response
	var params stakingtypes.QueryParamsResponse
	if err := cdc.Unmarshal(result.Response.Value, &params); err != nil {
		return 0, fmt.Errorf("failed to unmarshal params: %w", err)
	}
	return params.Params.UnbondingTime.Seconds(), nil
}

func getLightBlock(client *rpchttp.HTTP, height int64) (*LightBlock, error) {
	status, err := client.Status(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get status: %w", err)
	}

	peerId := status.NodeInfo.ID()
	commitResp, err := client.Commit(context.Background(), &height)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch trusted commit: %w", err)
	}

	signedHeader := commitResp.SignedHeader
	proposerAddr := signedHeader.Header.ProposerAddress
	var proposer *commettypes.Validator
	validatorResp, err := client.Validators(context.Background(), &height, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch validators: %w", err)
	}

	for _, resp := range validatorResp.Validators {
		if resp != nil {
			if slices.Equal(resp.Address.Bytes(), proposerAddr.Bytes()) {
				proposer = resp
				break
			}
		}
	}

	valSet := commettypes.NewValidatorSet(validatorResp.Validators)
	valSet.Proposer = proposer

	nextHeight := height + 1
	nextValidatorResp, err := client.Validators(context.Background(), &nextHeight, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch next validators: %w", err)
	}

	for _, resp := range nextValidatorResp.Validators {
		if resp != nil {
			if slices.Equal(resp.Address.Bytes(), proposerAddr.Bytes()) {
				proposer = resp
				break
			}
		}
	}

	nextValSet := commettypes.NewValidatorSet(nextValidatorResp.Validators)
	nextValSet.Proposer = proposer

	return &LightBlock{
		SignedHeader: signedHeader,
		ValSet:       *valSet,
		NextValSet:   *nextValSet,
		PeerId:       peerId,
		BlockHeight:  height,
	}, nil

}
