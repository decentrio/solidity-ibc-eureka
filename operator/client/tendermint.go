package client

import (
	"context"
	"fmt"
	"math"
	"math/big"
	"os"
	"slices"
	"strconv"
	"strings"

	tendermintContract "operator/bindings/SP1ICS07Tendermint"
	updateClientContract "operator/bindings/UpdateClient"

	"github.com/cometbft/cometbft/p2p"
	rpcclient "github.com/cometbft/cometbft/rpc/client"
	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
	commettypes "github.com/cometbft/cometbft/types"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	clienttypes "github.com/cosmos/ibc-go/v10/modules/core/02-client/types"
	commitmenttypes "github.com/cosmos/ibc-go/v10/modules/core/23-commitment/types"
	ics23 "github.com/cosmos/ics23/go"
)

const (
	MAX_TOTAL_VOTING_POWER = math.MaxInt64 / 8
	ProofType_EXIST        = 0
	ProofType_NON_EXIST    = 1
)

type LightBlock struct {
	SignedHeader commettypes.SignedHeader
	ValSet       commettypes.ValidatorSet
	NextValSet   commettypes.ValidatorSet
	PeerId       p2p.ID
	BlockHeight  int64
}

type SP1ICS07TendermintGenesis struct {
	TrustedClientState    updateClientContract.IICS07TendermintMsgsClientState
	TrustedConsensusState updateClientContract.IICS07TendermintMsgsConsensusState
}

type SupportedZkAlgorithm uint8

const (
	Groth16 SupportedZkAlgorithm = iota
	Plonk
)

// String returns the string representation of the algorithm
func (s SupportedZkAlgorithm) String() string {
	switch s {
	case Groth16:
		return "Groth16"
	case Plonk:
		return "Plonk"
	default:
		return "Unknown"
	}
}

func GetGenesis(trustedBlock int64, trustingPeriod uint32, trustLevel string, proofType string) (*SP1ICS07TendermintGenesis, error) {
	// Read RPC endpoint from environment variable
	rpcEndpoint := os.Getenv("TENDERMINT_RPC_URL")
	if rpcEndpoint == "" {
		return nil, fmt.Errorf("TENDERMINT_RPC_URL environment variable is required in .env file")
	}
	client, err := rpchttp.New(rpcEndpoint, "/websocket")
	if err != nil {
		return nil, fmt.Errorf("failed to create RPC client: %w", err)
	}

	status, err := client.Status(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get status: %w", err)
	}

	if trustedBlock == 0 {
		// get latest block
		trustedBlock = status.SyncInfo.LatestBlockHeight
	}

	trustedLightBlock, err := GetLightBlock(client, trustedBlock)
	if err != nil {
		return nil, fmt.Errorf("failed to get light block: %w", err)
	}
	_ = trustedLightBlock

	unbondingPeriod, err := GetUnbondingTime(client)
	if err != nil {
		return nil, fmt.Errorf("failed to get unbonding time: %w", err)
	}

	if trustingPeriod == 0 {
		trustingPeriod = uint32(unbondingPeriod * 2 / 3)
	}

	if trustingPeriod > uint32(unbondingPeriod) {
		return nil, fmt.Errorf("trusting period %d cannot be greater than unbonding period %d", trustingPeriod, uint32(unbondingPeriod))
	}

	chainId := trustedLightBlock.SignedHeader.Header.ChainID
	revision := clienttypes.ParseChainID(chainId)

	trustThreshold, err := ParseTrustThreshold(trustLevel)
	if err != nil {
		return nil, fmt.Errorf("failed to parse trust level: %w", err)
	}

	var zkAlgorithm SupportedZkAlgorithm
	switch proofType {
	case "groth16":
		zkAlgorithm = Groth16
	case "plonk":
		zkAlgorithm = Plonk
	default:
		return nil, fmt.Errorf("unsupported proof type: %s, supported types are: groth16, plonk", proofType)
	}

	clientState := updateClientContract.IICS07TendermintMsgsClientState{
		ChainId:    chainId,
		TrustLevel: trustThreshold,
		LatestHeight: updateClientContract.IICS02ClientMsgsHeight{
			RevisionNumber: revision,
			RevisionHeight: uint64(trustedLightBlock.SignedHeader.Header.Height),
		},
		IsFrozen:        false,
		ZkAlgorithm:     uint8(zkAlgorithm),
		TrustingPeriod:  trustingPeriod,
		UnbondingPeriod: uint32(unbondingPeriod),
	}

	consensusState := updateClientContract.IICS07TendermintMsgsConsensusState{
		Timestamp:          big.NewInt(trustedLightBlock.SignedHeader.Header.Time.UnixMilli()),
		Root:               bytesToBytes32(trustedLightBlock.SignedHeader.Header.AppHash),
		NextValidatorsHash: bytesToBytes32(trustedLightBlock.SignedHeader.NextValidatorsHash),
	}

	genesis := SP1ICS07TendermintGenesis{
		TrustedClientState:    clientState,
		TrustedConsensusState: consensusState,
	}
	return &genesis, nil
}

func GetUnbondingTime(client *rpchttp.HTTP) (float64, error) {

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

func GetLightBlock(client *rpchttp.HTTP, height int64) (*LightBlock, error) {
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

func ProvePath(client *rpchttp.HTTP, height int64, path [][]byte) ([]byte, *commitmenttypes.MerkleProof, error) {
	queryPath := fmt.Sprintf("store/%s/key", string(path[0]))
	request := slices.Concat(path[1:]...)

	// Make ABCI query
	result, err := client.ABCIQueryWithOptions(context.Background(), queryPath, request, rpcclient.ABCIQueryOptions{
		// Proof height should be the block before the target block.
		Height: height - 1,
		Prove:  true,
	})
	if err != nil {
		return nil, nil, fmt.Errorf("ABCI query failed: %w", err)
	}

	if result.Response.Code != 0 {
		return nil, nil, fmt.Errorf("query failed with code %d: %s", result.Response.Code, result.Response.Log)
	}

	if result.Response.Height != height-1 {
		return nil, nil, fmt.Errorf("proof height mismatch")
	}

	if !slices.Equal(result.Response.Key, path[1]) {
		return nil, nil, fmt.Errorf("key mismatch")
	}

	proof, err := commitmenttypes.ConvertProofs(result.Response.ProofOps)
	if err != nil {
		return nil, nil, fmt.Errorf("proof could not be retrieved: %w", err)
	}
	if len(proof.Proofs) == 0 {
		return nil, nil, fmt.Errorf("proof is empty")
	}
	return result.Response.Value, &proof, nil
}

// parseTrustThreshold parses a trust threshold fraction string like "2/3"
func ParseTrustThreshold(value string) (updateClientContract.IICS07TendermintMsgsTrustThreshold, error) {
	parts := strings.Split(value, "/")
	if len(parts) != 2 {
		return updateClientContract.IICS07TendermintMsgsTrustThreshold{}, fmt.Errorf("invalid trust threshold format: %s (expected format: 'numerator/denominator')", value)
	}

	numerator, err := strconv.ParseUint(parts[0], 10, 64)
	if err != nil {
		return updateClientContract.IICS07TendermintMsgsTrustThreshold{}, fmt.Errorf("invalid numerator: %s", parts[0])
	}

	denominator, err := strconv.ParseUint(parts[1], 10, 64)
	if err != nil {
		return updateClientContract.IICS07TendermintMsgsTrustThreshold{}, fmt.Errorf("invalid denominator: %s", parts[1])
	}

	if denominator == 0 {
		return updateClientContract.IICS07TendermintMsgsTrustThreshold{}, fmt.Errorf("denominator cannot be zero")
	}

	return updateClientContract.IICS07TendermintMsgsTrustThreshold{
		Numerator:   uint8(numerator),
		Denominator: uint8(denominator),
	}, nil
}

func ParseCommitmentProof(proof *ics23.CommitmentProof) (*tendermintContract.IMembershipMsgsCommitmentProof, error) {
	if proof == nil {
		return nil, fmt.Errorf("proof is nil")
	}

	var parsedProof *tendermintContract.IMembershipMsgsCommitmentProof
	switch p := proof.Proof.(type) {
	case *ics23.CommitmentProof_Exist:
		parsedProof = &tendermintContract.IMembershipMsgsCommitmentProof{
			ProofType: ProofType_EXIST,
			ExistenceProof: tendermintContract.IMembershipMsgsExistenceProof{
				Key:   p.Exist.Key,
				Value: bytesToBytes32(p.Exist.Value),
				Leaf:  ParseLeafOp(p.Exist.Leaf),
				Path:  []tendermintContract.IMembershipMsgsInnerOp{},
			},
			NonExistenceProof: tendermintContract.IMembershipMsgsNonExistenceProof{},
		}

		for _, innerOp := range p.Exist.Path {
			parsedProof.ExistenceProof.Path = append(parsedProof.ExistenceProof.Path, ParseInnerOp(innerOp))
		}
	case *ics23.CommitmentProof_Nonexist:
		if p.Nonexist.Left == nil || p.Nonexist.Right == nil {
			return nil, fmt.Errorf("non-existence proof must at least left or right existence proofs")
		}
		parsedProof = &tendermintContract.IMembershipMsgsCommitmentProof{
			ProofType:      ProofType_NON_EXIST,
			ExistenceProof: tendermintContract.IMembershipMsgsExistenceProof{},
			NonExistenceProof: tendermintContract.IMembershipMsgsNonExistenceProof{
				Key:      p.Nonexist.Key,
				HasLeft:  false,
				Left:     tendermintContract.IMembershipMsgsExistenceProof{},
				HasRight: false,
				Right:    tendermintContract.IMembershipMsgsExistenceProof{},
			},
		}

		if p.Nonexist.Left != nil {
			parsedProof.NonExistenceProof.HasLeft = true
			parsedProof.NonExistenceProof.Left = tendermintContract.IMembershipMsgsExistenceProof{
				Key:   p.Nonexist.Left.Key,
				Value: bytesToBytes32(p.Nonexist.Left.Value),
				Leaf:  ParseLeafOp(p.Nonexist.Left.Leaf),
				Path:  []tendermintContract.IMembershipMsgsInnerOp{},
			}
			for _, innerOp := range p.Nonexist.Left.Path {
				parsedProof.NonExistenceProof.Left.Path = append(parsedProof.NonExistenceProof.Left.Path, ParseInnerOp(innerOp))
			}
		}

		if p.Nonexist.Right != nil {
			parsedProof.NonExistenceProof.HasRight = true
			parsedProof.NonExistenceProof.Right = tendermintContract.IMembershipMsgsExistenceProof{
				Key:   p.Nonexist.Right.Key,
				Value: bytesToBytes32(p.Nonexist.Right.Value),
				Leaf:  ParseLeafOp(p.Nonexist.Right.Leaf),
				Path:  []tendermintContract.IMembershipMsgsInnerOp{},
			}
			for _, innerOp := range p.Nonexist.Right.Path {
				parsedProof.NonExistenceProof.Right.Path = append(parsedProof.NonExistenceProof.Right.Path, ParseInnerOp(innerOp))
			}
		}

	case *ics23.CommitmentProof_Batch:
		if len(p.Batch.GetEntries()) == 0 || p.Batch.GetEntries()[0] == nil {
			return nil, fmt.Errorf("batch proof has empty entry")
		}

		if e := p.Batch.GetEntries()[0].GetExist(); e != nil {
			parsedProof = &tendermintContract.IMembershipMsgsCommitmentProof{
				ProofType: ProofType_EXIST,
				ExistenceProof: tendermintContract.IMembershipMsgsExistenceProof{
					Key:   e.Key,
					Value: bytesToBytes32(e.Value),
					Leaf:  ParseLeafOp(e.Leaf),
					Path:  []tendermintContract.IMembershipMsgsInnerOp{},
				},
				NonExistenceProof: tendermintContract.IMembershipMsgsNonExistenceProof{},
			}

			for _, innerOp := range e.Path {
				parsedProof.ExistenceProof.Path = append(parsedProof.ExistenceProof.Path, ParseInnerOp(innerOp))
			}
		}

		if n := p.Batch.GetEntries()[0].GetNonexist(); n != nil {
			parsedProof = &tendermintContract.IMembershipMsgsCommitmentProof{
				ProofType:      ProofType_NON_EXIST,
				ExistenceProof: tendermintContract.IMembershipMsgsExistenceProof{},
				NonExistenceProof: tendermintContract.IMembershipMsgsNonExistenceProof{
					Key:      n.Key,
					HasLeft:  false,
					Left:     tendermintContract.IMembershipMsgsExistenceProof{},
					HasRight: false,
					Right:    tendermintContract.IMembershipMsgsExistenceProof{},
				},
			}

			if n.Left != nil {
				parsedProof.NonExistenceProof.HasLeft = true
				parsedProof.NonExistenceProof.Left = tendermintContract.IMembershipMsgsExistenceProof{
					Key:   n.Left.Key,
					Value: bytesToBytes32(n.Left.Value),
					Leaf:  ParseLeafOp(n.Left.Leaf),
					Path:  []tendermintContract.IMembershipMsgsInnerOp{},
				}
				for _, innerOp := range n.Left.Path {
					parsedProof.NonExistenceProof.Left.Path = append(parsedProof.NonExistenceProof.Left.Path, ParseInnerOp(innerOp))
				}
			}

			if n.Right != nil {
				parsedProof.NonExistenceProof.HasRight = true
				parsedProof.NonExistenceProof.Right = tendermintContract.IMembershipMsgsExistenceProof{
					Key:   n.Right.Key,
					Value: bytesToBytes32(n.Right.Value),
					Leaf:  ParseLeafOp(n.Right.Leaf),
					Path:  []tendermintContract.IMembershipMsgsInnerOp{},
				}
				for _, innerOp := range n.Right.Path {
					parsedProof.NonExistenceProof.Right.Path = append(parsedProof.NonExistenceProof.Right.Path, ParseInnerOp(innerOp))
				}
			}
		}
	case *ics23.CommitmentProof_Compressed:
		decompressedProof := ics23.Decompress(proof)
		return ParseCommitmentProof(decompressedProof)
	default:
		return nil, fmt.Errorf("unrecognized proof type")
	}

	return parsedProof, nil
}

func ParseLeafOp(leafOp *ics23.LeafOp) tendermintContract.IMembershipMsgsLeafOp {
	if leafOp == nil {
		return tendermintContract.IMembershipMsgsLeafOp{}
	}

	return tendermintContract.IMembershipMsgsLeafOp{
		HashOp:       uint8(leafOp.Hash),
		PrehashKey:   uint8(leafOp.PrehashKey),
		PrehashValue: uint8(leafOp.PrehashValue),
		Prefix:       leafOp.Prefix,
	}
}

func ParseInnerOp(innerOp *ics23.InnerOp) tendermintContract.IMembershipMsgsInnerOp {
	if innerOp == nil {
		return tendermintContract.IMembershipMsgsInnerOp{}
	}

	return tendermintContract.IMembershipMsgsInnerOp{
		HashOp: uint8(innerOp.Hash),
		Prefix: innerOp.Prefix,
		Suffix: innerOp.Suffix,
	}
}

func bytesToBytes32(data []byte) [32]byte {
	var result [32]byte
	copy(result[:], data)
	return result
}
