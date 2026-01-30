// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractSP1ICS07Tendermint

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// IICS02ClientMsgsHeight is an auto generated low-level Go binding around an user-defined struct.
type IICS02ClientMsgsHeight struct {
	RevisionNumber uint64
	RevisionHeight uint64
}

// IICS07TendermintMsgsBlockCommit is an auto generated low-level Go binding around an user-defined struct.
type IICS07TendermintMsgsBlockCommit struct {
	Height     uint64
	Round      uint32
	BlockId    IICS07TendermintMsgsBlockId
	CommitSigs []IICS07TendermintMsgsCommitSig
}

// IICS07TendermintMsgsBlockHeader is an auto generated low-level Go binding around an user-defined struct.
type IICS07TendermintMsgsBlockHeader struct {
	Version            IICS07TendermintMsgsVersion
	ChainId            string
	Height             uint64
	Time               *big.Int
	HasLastBlockId     bool
	LastBlockId        IICS07TendermintMsgsBlockId
	HasLastCommitHash  bool
	LastCommitHash     [32]byte
	HasDataHash        bool
	DataHash           [32]byte
	ValidatorsHash     [32]byte
	NextValidatorsHash [32]byte
	ConsensusHash      [32]byte
	AppHash            [32]byte
	HasLastResultsHash bool
	LastResultsHash    [32]byte
	HasEvidenceHash    bool
	EvidenceHash       [32]byte
	ProposerAddress    []byte
}

// IICS07TendermintMsgsBlockId is an auto generated low-level Go binding around an user-defined struct.
type IICS07TendermintMsgsBlockId struct {
	HashData      [32]byte
	PartSetHeader IICS07TendermintMsgsPartSetHeader
}

// IICS07TendermintMsgsClientState is an auto generated low-level Go binding around an user-defined struct.
type IICS07TendermintMsgsClientState struct {
	ChainId         string
	TrustLevel      IICS07TendermintMsgsTrustThreshold
	LatestHeight    IICS02ClientMsgsHeight
	TrustingPeriod  uint32
	UnbondingPeriod uint32
	IsFrozen        bool
	ZkAlgorithm     uint8
}

// IICS07TendermintMsgsCommitSig is an auto generated low-level Go binding around an user-defined struct.
type IICS07TendermintMsgsCommitSig struct {
	Flag uint8
	Data IICS07TendermintMsgsCommitSigData
}

// IICS07TendermintMsgsCommitSigData is an auto generated low-level Go binding around an user-defined struct.
type IICS07TendermintMsgsCommitSigData struct {
	ValidatorAddress []byte
	Timestamp        *big.Int
	HasSignature     bool
	Signature        []byte
}

// IICS07TendermintMsgsConsensusState is an auto generated low-level Go binding around an user-defined struct.
type IICS07TendermintMsgsConsensusState struct {
	Timestamp          *big.Int
	Root               [32]byte
	NextValidatorsHash [32]byte
}

// IICS07TendermintMsgsHeader is an auto generated low-level Go binding around an user-defined struct.
type IICS07TendermintMsgsHeader struct {
	SignedHeader            IICS07TendermintMsgsSignedHeader
	ValidatorSet            IICS07TendermintMsgsValidatorSet
	TrustedHeight           IICS02ClientMsgsHeight
	TrustedNextValidatorSet IICS07TendermintMsgsValidatorSet
}

// IICS07TendermintMsgsPartSetHeader is an auto generated low-level Go binding around an user-defined struct.
type IICS07TendermintMsgsPartSetHeader struct {
	Total    uint32
	HashData [32]byte
}

// IICS07TendermintMsgsSignedHeader is an auto generated low-level Go binding around an user-defined struct.
type IICS07TendermintMsgsSignedHeader struct {
	Header IICS07TendermintMsgsBlockHeader
	Commit IICS07TendermintMsgsBlockCommit
}

// IICS07TendermintMsgsTrustThreshold is an auto generated low-level Go binding around an user-defined struct.
type IICS07TendermintMsgsTrustThreshold struct {
	Numerator   uint8
	Denominator uint8
}

// IICS07TendermintMsgsValidatorInfo is an auto generated low-level Go binding around an user-defined struct.
type IICS07TendermintMsgsValidatorInfo struct {
	ValAddress       []byte
	PubKey           [32]byte
	VotingPower      uint64
	ProposerPriority int64
}

// IICS07TendermintMsgsValidatorSet is an auto generated low-level Go binding around an user-defined struct.
type IICS07TendermintMsgsValidatorSet struct {
	Validators       []IICS07TendermintMsgsValidatorInfo
	HasProposer      bool
	Proposer         IICS07TendermintMsgsValidatorInfo
	TotalVotingPower uint64
}

// IICS07TendermintMsgsVersion is an auto generated low-level Go binding around an user-defined struct.
type IICS07TendermintMsgsVersion struct {
	BlockVersion uint64
	AppVersion   uint64
}

// ILightClientMsgsMsgVerifyMembership is an auto generated low-level Go binding around an user-defined struct.
type ILightClientMsgsMsgVerifyMembership struct {
	Height                IICS02ClientMsgsHeight
	KvPairs               []IMembershipMsgsKVPair
	MerkleProofs          []IMembershipMsgsMerkleProof
	AppHash               [32]byte
	TrustedConsensusState IICS07TendermintMsgsConsensusState
	MembershipType        uint8
}

// ILightClientMsgsMsgVerifyNonMembership is an auto generated low-level Go binding around an user-defined struct.
type ILightClientMsgsMsgVerifyNonMembership struct {
	Height                IICS02ClientMsgsHeight
	KvPairs               []IMembershipMsgsKVPair
	MerkleProofs          []IMembershipMsgsMerkleProof
	AppHash               [32]byte
	TrustedConsensusState IICS07TendermintMsgsConsensusState
	MembershipType        uint8
}

// IMembershipMsgsCommitmentProof is an auto generated low-level Go binding around an user-defined struct.
type IMembershipMsgsCommitmentProof struct {
	ProofType         uint8
	ExistenceProof    IMembershipMsgsExistenceProof
	NonExistenceProof IMembershipMsgsNonExistenceProof
}

// IMembershipMsgsExistenceProof is an auto generated low-level Go binding around an user-defined struct.
type IMembershipMsgsExistenceProof struct {
	Key   []byte
	Value [32]byte
	Leaf  IMembershipMsgsLeafOp
	Path  []IMembershipMsgsInnerOp
}

// IMembershipMsgsInnerOp is an auto generated low-level Go binding around an user-defined struct.
type IMembershipMsgsInnerOp struct {
	HashOp uint8
	Prefix []byte
	Suffix []byte
}

// IMembershipMsgsKVPair is an auto generated low-level Go binding around an user-defined struct.
type IMembershipMsgsKVPair struct {
	Path  [][]byte
	Value [32]byte
}

// IMembershipMsgsLeafOp is an auto generated low-level Go binding around an user-defined struct.
type IMembershipMsgsLeafOp struct {
	HashOp       uint8
	PrehashKey   uint8
	PrehashValue uint8
	Prefix       []byte
}

// IMembershipMsgsMerkleProof is an auto generated low-level Go binding around an user-defined struct.
type IMembershipMsgsMerkleProof struct {
	Proofs []IMembershipMsgsCommitmentProof
}

// IMembershipMsgsNonExistenceProof is an auto generated low-level Go binding around an user-defined struct.
type IMembershipMsgsNonExistenceProof struct {
	Key      []byte
	HasLeft  bool
	Left     IMembershipMsgsExistenceProof
	HasRight bool
	Right    IMembershipMsgsExistenceProof
}

// IUpdateClientMsgsMsgUpdateClient is an auto generated low-level Go binding around an user-defined struct.
type IUpdateClientMsgsMsgUpdateClient struct {
	ClientState           IICS07TendermintMsgsClientState
	TrustedConsensusState IICS07TendermintMsgsConsensusState
	ProposedHeader        IICS07TendermintMsgsHeader
	Time                  *big.Int
	Proof                 [8]*big.Int
	Commitments           [2]*big.Int
	CommitmentPok         [2]*big.Int
}

// ContractSP1ICS07TendermintMetaData contains all meta data concerning the ContractSP1ICS07Tendermint contract.
var ContractSP1ICS07TendermintMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"verifier\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"membership_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"misbehaviour_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"updateClient_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_clientState\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_consensusState\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"roleManager\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ALLOWED_SP1_CLOCK_DRIFT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MEMBERSHIP\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIMembership\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MISBEHAVIOUR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIMisbehaviour\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PROOF_SUBMITTER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UPDATE_CLIENT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIUpdateClient\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"VERIFIER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIVerifier\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"clientState\",\"inputs\":[],\"outputs\":[{\"name\":\"chainId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"trustLevel\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.TrustThreshold\",\"components\":[{\"name\":\"numerator\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"denominator\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"name\":\"latestHeight\",\"type\":\"tuple\",\"internalType\":\"structIICS02ClientMsgs.Height\",\"components\":[{\"name\":\"revisionNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revisionHeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"trustingPeriod\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"unbondingPeriod\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isFrozen\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"zkAlgorithm\",\"type\":\"uint8\",\"internalType\":\"enumIICS07TendermintMsgs.SupportedZkAlgorithm\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getClientState\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getConsensusStateHash\",\"inputs\":[{\"name\":\"revisionHeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"misbehaviour\",\"inputs\":[{\"name\":\"misbehaviourMsg\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"multicall\",\"inputs\":[{\"name\":\"data\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"outputs\":[{\"name\":\"results\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateClient\",\"inputs\":[{\"name\":\"updateClientMsg\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumILightClientMsgs.UpdateResult\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateClientMsg\",\"inputs\":[{\"name\":\"msg_\",\"type\":\"tuple\",\"internalType\":\"structIUpdateClientMsgs.MsgUpdateClient\",\"components\":[{\"name\":\"clientState\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.ClientState\",\"components\":[{\"name\":\"chainId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"trustLevel\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.TrustThreshold\",\"components\":[{\"name\":\"numerator\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"denominator\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"name\":\"latestHeight\",\"type\":\"tuple\",\"internalType\":\"structIICS02ClientMsgs.Height\",\"components\":[{\"name\":\"revisionNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revisionHeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"trustingPeriod\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"unbondingPeriod\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isFrozen\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"zkAlgorithm\",\"type\":\"uint8\",\"internalType\":\"enumIICS07TendermintMsgs.SupportedZkAlgorithm\"}]},{\"name\":\"trustedConsensusState\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.ConsensusState\",\"components\":[{\"name\":\"timestamp\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"nextValidatorsHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"proposedHeader\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.Header\",\"components\":[{\"name\":\"signedHeader\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.SignedHeader\",\"components\":[{\"name\":\"header\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.BlockHeader\",\"components\":[{\"name\":\"version\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.Version\",\"components\":[{\"name\":\"blockVersion\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"appVersion\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"chainId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"height\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"time\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"hasLastBlockId\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"lastBlockId\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.BlockId\",\"components\":[{\"name\":\"hashData\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"partSetHeader\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.PartSetHeader\",\"components\":[{\"name\":\"total\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashData\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]},{\"name\":\"hasLastCommitHash\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"lastCommitHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"hasDataHash\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"dataHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"validatorsHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"nextValidatorsHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"consensusHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"appHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"hasLastResultsHash\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"lastResultsHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"hasEvidenceHash\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"evidenceHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"proposerAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"commit\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.BlockCommit\",\"components\":[{\"name\":\"height\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"round\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"blockId\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.BlockId\",\"components\":[{\"name\":\"hashData\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"partSetHeader\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.PartSetHeader\",\"components\":[{\"name\":\"total\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashData\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]},{\"name\":\"commitSigs\",\"type\":\"tuple[]\",\"internalType\":\"structIICS07TendermintMsgs.CommitSig[]\",\"components\":[{\"name\":\"flag\",\"type\":\"uint8\",\"internalType\":\"enumIICS07TendermintMsgs.CommitSigFlag\"},{\"name\":\"data\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.CommitSigData\",\"components\":[{\"name\":\"validatorAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"timestamp\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"hasSignature\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]}]}]},{\"name\":\"validatorSet\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.ValidatorSet\",\"components\":[{\"name\":\"validators\",\"type\":\"tuple[]\",\"internalType\":\"structIICS07TendermintMsgs.ValidatorInfo[]\",\"components\":[{\"name\":\"valAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"pubKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"votingPower\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"proposerPriority\",\"type\":\"int64\",\"internalType\":\"int64\"}]},{\"name\":\"hasProposer\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"proposer\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.ValidatorInfo\",\"components\":[{\"name\":\"valAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"pubKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"votingPower\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"proposerPriority\",\"type\":\"int64\",\"internalType\":\"int64\"}]},{\"name\":\"totalVotingPower\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"trustedHeight\",\"type\":\"tuple\",\"internalType\":\"structIICS02ClientMsgs.Height\",\"components\":[{\"name\":\"revisionNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revisionHeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"trustedNextValidatorSet\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.ValidatorSet\",\"components\":[{\"name\":\"validators\",\"type\":\"tuple[]\",\"internalType\":\"structIICS07TendermintMsgs.ValidatorInfo[]\",\"components\":[{\"name\":\"valAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"pubKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"votingPower\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"proposerPriority\",\"type\":\"int64\",\"internalType\":\"int64\"}]},{\"name\":\"hasProposer\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"proposer\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.ValidatorInfo\",\"components\":[{\"name\":\"valAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"pubKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"votingPower\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"proposerPriority\",\"type\":\"int64\",\"internalType\":\"int64\"}]},{\"name\":\"totalVotingPower\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}]},{\"name\":\"time\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"proof\",\"type\":\"uint256[8]\",\"internalType\":\"uint256[8]\"},{\"name\":\"commitments\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"commitmentPok\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumILightClientMsgs.UpdateResult\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeClient\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyMembership\",\"inputs\":[{\"name\":\"msg_\",\"type\":\"tuple\",\"internalType\":\"structILightClientMsgs.MsgVerifyMembership\",\"components\":[{\"name\":\"height\",\"type\":\"tuple\",\"internalType\":\"structIICS02ClientMsgs.Height\",\"components\":[{\"name\":\"revisionNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revisionHeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"kvPairs\",\"type\":\"tuple[]\",\"internalType\":\"structIMembershipMsgs.KVPair[]\",\"components\":[{\"name\":\"path\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"value\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"merkleProofs\",\"type\":\"tuple[]\",\"internalType\":\"structIMembershipMsgs.MerkleProof[]\",\"components\":[{\"name\":\"proofs\",\"type\":\"tuple[]\",\"internalType\":\"structIMembershipMsgs.CommitmentProof[]\",\"components\":[{\"name\":\"proofType\",\"type\":\"uint8\",\"internalType\":\"enumIMembershipMsgs.ProofType\"},{\"name\":\"existenceProof\",\"type\":\"tuple\",\"internalType\":\"structIMembershipMsgs.ExistenceProof\",\"components\":[{\"name\":\"key\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"leaf\",\"type\":\"tuple\",\"internalType\":\"structIMembershipMsgs.LeafOp\",\"components\":[{\"name\":\"hashOp\",\"type\":\"uint8\",\"internalType\":\"enumIMembershipMsgs.HashOp\"},{\"name\":\"prehashKey\",\"type\":\"uint8\",\"internalType\":\"enumIMembershipMsgs.HashOp\"},{\"name\":\"prehashValue\",\"type\":\"uint8\",\"internalType\":\"enumIMembershipMsgs.HashOp\"},{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"path\",\"type\":\"tuple[]\",\"internalType\":\"structIMembershipMsgs.InnerOp[]\",\"components\":[{\"name\":\"hashOp\",\"type\":\"uint8\",\"internalType\":\"enumIMembershipMsgs.HashOp\"},{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"suffix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]},{\"name\":\"nonExistenceProof\",\"type\":\"tuple\",\"internalType\":\"structIMembershipMsgs.NonExistenceProof\",\"components\":[{\"name\":\"key\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"hasLeft\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"left\",\"type\":\"tuple\",\"internalType\":\"structIMembershipMsgs.ExistenceProof\",\"components\":[{\"name\":\"key\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"leaf\",\"type\":\"tuple\",\"internalType\":\"structIMembershipMsgs.LeafOp\",\"components\":[{\"name\":\"hashOp\",\"type\":\"uint8\",\"internalType\":\"enumIMembershipMsgs.HashOp\"},{\"name\":\"prehashKey\",\"type\":\"uint8\",\"internalType\":\"enumIMembershipMsgs.HashOp\"},{\"name\":\"prehashValue\",\"type\":\"uint8\",\"internalType\":\"enumIMembershipMsgs.HashOp\"},{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"path\",\"type\":\"tuple[]\",\"internalType\":\"structIMembershipMsgs.InnerOp[]\",\"components\":[{\"name\":\"hashOp\",\"type\":\"uint8\",\"internalType\":\"enumIMembershipMsgs.HashOp\"},{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"suffix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]},{\"name\":\"hasRight\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"right\",\"type\":\"tuple\",\"internalType\":\"structIMembershipMsgs.ExistenceProof\",\"components\":[{\"name\":\"key\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"leaf\",\"type\":\"tuple\",\"internalType\":\"structIMembershipMsgs.LeafOp\",\"components\":[{\"name\":\"hashOp\",\"type\":\"uint8\",\"internalType\":\"enumIMembershipMsgs.HashOp\"},{\"name\":\"prehashKey\",\"type\":\"uint8\",\"internalType\":\"enumIMembershipMsgs.HashOp\"},{\"name\":\"prehashValue\",\"type\":\"uint8\",\"internalType\":\"enumIMembershipMsgs.HashOp\"},{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"path\",\"type\":\"tuple[]\",\"internalType\":\"structIMembershipMsgs.InnerOp[]\",\"components\":[{\"name\":\"hashOp\",\"type\":\"uint8\",\"internalType\":\"enumIMembershipMsgs.HashOp\"},{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"suffix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]}]}]}]},{\"name\":\"appHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"trustedConsensusState\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.ConsensusState\",\"components\":[{\"name\":\"timestamp\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"nextValidatorsHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"membershipType\",\"type\":\"uint8\",\"internalType\":\"enumIMembershipMsgs.MembershipType\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifyNonMembership\",\"inputs\":[{\"name\":\"msg_\",\"type\":\"tuple\",\"internalType\":\"structILightClientMsgs.MsgVerifyNonMembership\",\"components\":[{\"name\":\"height\",\"type\":\"tuple\",\"internalType\":\"structIICS02ClientMsgs.Height\",\"components\":[{\"name\":\"revisionNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revisionHeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"kvPairs\",\"type\":\"tuple[]\",\"internalType\":\"structIMembershipMsgs.KVPair[]\",\"components\":[{\"name\":\"path\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"value\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"merkleProofs\",\"type\":\"tuple[]\",\"internalType\":\"structIMembershipMsgs.MerkleProof[]\",\"components\":[{\"name\":\"proofs\",\"type\":\"tuple[]\",\"internalType\":\"structIMembershipMsgs.CommitmentProof[]\",\"components\":[{\"name\":\"proofType\",\"type\":\"uint8\",\"internalType\":\"enumIMembershipMsgs.ProofType\"},{\"name\":\"existenceProof\",\"type\":\"tuple\",\"internalType\":\"structIMembershipMsgs.ExistenceProof\",\"components\":[{\"name\":\"key\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"leaf\",\"type\":\"tuple\",\"internalType\":\"structIMembershipMsgs.LeafOp\",\"components\":[{\"name\":\"hashOp\",\"type\":\"uint8\",\"internalType\":\"enumIMembershipMsgs.HashOp\"},{\"name\":\"prehashKey\",\"type\":\"uint8\",\"internalType\":\"enumIMembershipMsgs.HashOp\"},{\"name\":\"prehashValue\",\"type\":\"uint8\",\"internalType\":\"enumIMembershipMsgs.HashOp\"},{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"path\",\"type\":\"tuple[]\",\"internalType\":\"structIMembershipMsgs.InnerOp[]\",\"components\":[{\"name\":\"hashOp\",\"type\":\"uint8\",\"internalType\":\"enumIMembershipMsgs.HashOp\"},{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"suffix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]},{\"name\":\"nonExistenceProof\",\"type\":\"tuple\",\"internalType\":\"structIMembershipMsgs.NonExistenceProof\",\"components\":[{\"name\":\"key\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"hasLeft\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"left\",\"type\":\"tuple\",\"internalType\":\"structIMembershipMsgs.ExistenceProof\",\"components\":[{\"name\":\"key\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"leaf\",\"type\":\"tuple\",\"internalType\":\"structIMembershipMsgs.LeafOp\",\"components\":[{\"name\":\"hashOp\",\"type\":\"uint8\",\"internalType\":\"enumIMembershipMsgs.HashOp\"},{\"name\":\"prehashKey\",\"type\":\"uint8\",\"internalType\":\"enumIMembershipMsgs.HashOp\"},{\"name\":\"prehashValue\",\"type\":\"uint8\",\"internalType\":\"enumIMembershipMsgs.HashOp\"},{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"path\",\"type\":\"tuple[]\",\"internalType\":\"structIMembershipMsgs.InnerOp[]\",\"components\":[{\"name\":\"hashOp\",\"type\":\"uint8\",\"internalType\":\"enumIMembershipMsgs.HashOp\"},{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"suffix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]},{\"name\":\"hasRight\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"right\",\"type\":\"tuple\",\"internalType\":\"structIMembershipMsgs.ExistenceProof\",\"components\":[{\"name\":\"key\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"leaf\",\"type\":\"tuple\",\"internalType\":\"structIMembershipMsgs.LeafOp\",\"components\":[{\"name\":\"hashOp\",\"type\":\"uint8\",\"internalType\":\"enumIMembershipMsgs.HashOp\"},{\"name\":\"prehashKey\",\"type\":\"uint8\",\"internalType\":\"enumIMembershipMsgs.HashOp\"},{\"name\":\"prehashValue\",\"type\":\"uint8\",\"internalType\":\"enumIMembershipMsgs.HashOp\"},{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"path\",\"type\":\"tuple[]\",\"internalType\":\"structIMembershipMsgs.InnerOp[]\",\"components\":[{\"name\":\"hashOp\",\"type\":\"uint8\",\"internalType\":\"enumIMembershipMsgs.HashOp\"},{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"suffix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]}]}]}]},{\"name\":\"appHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"trustedConsensusState\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.ConsensusState\",\"components\":[{\"name\":\"timestamp\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"nextValidatorsHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"membershipType\",\"type\":\"uint8\",\"internalType\":\"enumIMembershipMsgs.MembershipType\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotHandleMisbehavior\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChainIdMismatch\",\"inputs\":[{\"name\":\"expected\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"actual\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"ClientStateMismatch\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"actual\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"ConsensusStateHashMismatch\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"actual\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ConsensusStateNotFound\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ConsensusStateRootMismatch\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"actual\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"EmptyValue\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedToVerifyHeader\",\"inputs\":[{\"name\":\"description\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"FeatureNotSupported\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FrozenClientState\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientMisbehaviourHeaderHeight\",\"inputs\":[{\"name\":\"height1\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"height2\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InsufficientTrustingPeriod\",\"inputs\":[{\"name\":\"durationSinceConsensusState\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"trustingPeriod\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"type\":\"error\",\"name\":\"InvalidConsensusStateTimestamp\",\"inputs\":[{\"name\":\"timestamp\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"type\":\"error\",\"name\":\"InvalidHeaderHeight\",\"inputs\":[{\"name\":\"height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidMembershipProof\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"KeyValuePairNotInCache\",\"inputs\":[{\"name\":\"path\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"value\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"LengthIsOutOfRange\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"min\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"max\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MembershipProofKeyNotFound\",\"inputs\":[{\"name\":\"path\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}]},{\"type\":\"error\",\"name\":\"MembershipProofValueMismatch\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"actual\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"MismatchedRevisionHeights\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"actual\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"MismatchedValidatorHashes\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"actual\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ProofHeightMismatch\",\"inputs\":[{\"name\":\"expectedRevisionNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"expectedRevisionHeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"actualRevisionNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"actualRevisionHeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ProofIsInTheFuture\",\"inputs\":[{\"name\":\"now\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proofTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ProofIsTooOld\",\"inputs\":[{\"name\":\"now\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proofTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"TrustThresholdMismatch\",\"inputs\":[{\"name\":\"expectedNumerator\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expectedDenominator\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"actualNumerator\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"actualDenominator\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"TrustingPeriodMismatch\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"actual\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"TrustingPeriodTooLong\",\"inputs\":[{\"name\":\"trustingPeriod\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"unbondingPeriod\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"UnbondingPeriodMismatch\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"actual\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"UnknownMembershipType\",\"inputs\":[{\"name\":\"membershipType\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"UnknownZkAlgorithm\",\"inputs\":[{\"name\":\"algorithm\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"VerificationKeyMismatch\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"actual\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]",
	Bin: "0x6101006040523461055257614a658038038061001a81610575565b928339810160e082820312610552576100328261059a565b61003e6020840161059a565b61004a6040850161059a565b916100576060860161059a565b60808601519094906001600160401b0381116105525786019080601f83011215610552578151610089926020016105ae565b9461009b60c060a0830151920161059a565b9580518101906020820190602081840312610552576020810151906001600160401b038211610552570191829003601f1981019061012013610552576040519160e083016001600160401b0381118482101761053e5760405260208401516001600160401b038111610552576020908501019080601f83011215610552578151610127926020016105ae565b82526040811261055257604061013b610556565b916101478286016105ee565b8352610155606086016105ee565b602084015260208401928352603f19011261055257610172610556565b61017e608085016105fc565b815261018c60a085016105fc565b6020820152604083019081526101a460c08501610610565b90606084019182526101b860e08601610610565b92608085019384526101008601519586151587036105525760a0860196875261012001519460028610156105525760c08101958652518051906001600160401b03821161053e57600154600181811c91168015610534575b602082101461052057601f81116104bd575b50602090601f83116001146104505763ffffffff95949392915f9183610445575b50508160011b915f199060031b1c1916176001555b5160ff81511661ff00602060025493015160081b169161ffff191617176002555160018060401b0381511660035491602068010000000000000000600160801b0391015160401b169160018060801b031916171760035551169267ffffffff00000000600454925160201b169051151560401b92519160028310156104315769ff00000000000000000068ff00000000000000009360481b169469ff000000000000000000199160018060481b0319161716179116171760045560018060401b0360035460401c165f52600560205260405f205560018060a01b031660805260018060a01b031660a05260018060a01b031660c05260018060a01b031660e05260045463ffffffff8116610708810163ffffffff811161041d5763ffffffff809360201c16928391161161040857826001600160a01b0381166103ef575061039e610717565b505b60405161422b908161079a823960805181610f82015260a051818181610c460152613360015260c0518181816105220152611080015260e051818181610c9601528181611f3b01526130900152f35b806103fc61040292610621565b50610697565b506103a0565b6333fae18560e21b5f5260045260245260445ffd5b634e487b7160e01b5f52601160045260245ffd5b634e487b7160e01b5f52602160045260245ffd5b015190505f80610243565b90601f1983169160015f52815f20925f5b8181106104a5575091600193918563ffffffff99989796941061048d575b505050811b01600155610258565b01515f1960f88460031b161c191690555f808061047f565b92936020600181928786015181550195019301610461565b60015f527fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf6601f840160051c81019160208510610516575b601f0160051c01905b81811061050b5750610222565b5f81556001016104fe565b90915081906104f5565b634e487b7160e01b5f52602260045260245ffd5b90607f1690610210565b634e487b7160e01b5f52604160045260245ffd5b5f80fd5b60408051919082016001600160401b0381118382101761053e57604052565b6040519190601f01601f191682016001600160401b0381118382101761053e57604052565b51906001600160a01b038216820361055257565b9192916001600160401b03821161053e576105d2601f8301601f1916602001610575565b938285528282011161055257815f926020928387015e84010152565b519060ff8216820361055257565b51906001600160401b038216820361055257565b519063ffffffff8216820361055257565b6001600160a01b0381165f9081525f516020614a455f395f51905f52602052604090205460ff16610692576001600160a01b03165f8181525f516020614a455f395f51905f5260205260408120805460ff191660011790553391905f5160206149c55f395f51905f528180a4600190565b505f90565b6001600160a01b0381165f9081525f5160206149e55f395f51905f52602052604090205460ff16610692576001600160a01b03165f8181525f5160206149e55f395f51905f5260205260408120805460ff191660011790553391905f516020614a255f395f51905f52905f5160206149c55f395f51905f529080a4600190565b5f80525f5160206149e55f395f51905f526020525f516020614a055f395f51905f525460ff16610795575f8080525f5160206149e55f395f51905f526020525f516020614a055f395f51905f52805460ff1916600117905533905f516020614a255f395f51905f525f5160206149c55f395f51905f528280a4600190565b5f9056fe6080806040526004361015610012575f80fd5b5f3560e01c90816301ffc9a7146110a45750806302cf295214611054578063038d5cb314610dfd57806304354dac14610fa657806308c84e7014610f565780630bece35614610ebc5780630c6faf5914610dfd57806323842fb814610dcd578063248a9ca314610d9b5780632c3ee47414610d7f5780632f2ff15d14610d5057806336568abe14610cf45780635972185a14610cba57806387d4332f14610c6a57806389df51f114610c1a5780638a8e4c5d14610b7657806391d1485414610b2d578063a217fddf14610b13578063ac9650d8146108e3578063bd3ce6b0146107d9578063d547741f146107a3578063ddba6537146101ed5763ef913a4b14610119575f80fd5b346101e9575f6003193601126101e9576101e560405160208082015261012060408201526101d18161014e61016082016112bc565b60ff600254818116606085015260081c16608083015267ffffffffffffffff60035481811660a085015260401c1660c083015260ff60045463ffffffff811660e085015263ffffffff8160201c16610100850152818160401c16151561012085015260481c166101bd8161140c565b61014083015203601f1981018352826113e9565b604051918291602083526020830190611246565b0390f35b5f80fd5b346101e9576101fb366111ad565b6004549060ff8260401c1661077b575f80527f4c48d1d2b0f4f7485fc28e3db22341d96a20aa29e6efa8149da9751603abd4e06020527f6e0c24a6e293ff9b755263dbaa15ba3796b0b8d3fe17cfb4ddf8143b268eac475460ff161561076e575b8201916020818403126101e95780359067ffffffffffffffff82116101e95701610120818403126101e95760405160a0810181811067ffffffffffffffff82111761074157604052813567ffffffffffffffff81116101e957846102c191840161226d565b8152602082013567ffffffffffffffff81116101e9578201916060838603126101e957604051926102f184611395565b803567ffffffffffffffff81116101e95781016040818803126101e9576040519061031b82611379565b803567ffffffffffffffff81116101e9578161033e8a60209361034695016121f8565b8452016111fe565b60208201528452602081013567ffffffffffffffff81116101e9578661036d918301612537565b602085015260408101359067ffffffffffffffff82116101e95761039391879101612537565b6040840152602082019283526103c06103af866040840161233f565b956040840196875260a0830161233f565b9160806104426103da6101006060850195878752016117de565b95828401958787526fffffffffffffffffffffffffffffffff85519251986105028c51936104d46104a46040519d8e998a997fa6fe8f56000000000000000000000000000000000000000000000000000000008b5261012060048c01526101248b0190612957565b6003198a82030160248b0152604061049383516060845267ffffffffffffffff6020610479835186606089015260a0880190611246565b920151168f85015260208501518482036020860152612ae5565b920151906040818403910152612ae5565b86516fffffffffffffffffffffffffffffffff166044890152602087015160648901526040909601516084880152565b80516fffffffffffffffffffffffffffffffff1660a4870152602081015160c48701526040015160e4860152565b16610104830152038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa938415610736575f94610674575b5067ffffffffffffffff60208061066d9561061a7fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff9998966105bb680100000000000000009c6fffffffffffffffffffffffffffffffff61061299519151935195511690613f03565b6040516105f18582018093604080916fffffffffffffffffffffffffffffffff8151168452602081015160208501520151910152565b606081526106006080826113e9565b51902061061286858a510151166130f3565b808214613a37565b6040516106508382018093604080916fffffffffffffffffffffffffffffffff8151168452602081015160208501520151910152565b6060815261065f6080826113e9565b5190209401510151166130f3565b1617600455005b959493509060803d60801161072f575b61068e81886113e9565b8601906080878303126101e95760208061066d9561061a67ffffffffffffffff946105bb7fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff9b6fffffffffffffffffffffffffffffffff680100000000000000009e6107176106129b604080519361070585611379565b61070f83826114df565b8552016114df565b888201529c9d50509c50509695505095505050610552565b503d610684565b6040513d5f823e3d90fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b61077661317b565b61025c565b7f928b1233000000000000000000000000000000000000000000000000000000005f5260045ffd5b346101e9576107d76107b436611213565b906107d26107cd825f525f602052600160405f20015490565b613203565b613c81565b005b346101e9575f6003193601126101e957610870604051610803816107fc816112bc565b03826113e9565b60405161080f81611379565b60ff600254818116835260081c16602082015260405161082e81611379565b67ffffffffffffffff600354818116835260401c16602082015260ff6004546108aa828260481c169361088a6040519889986101208a526101208a0190611246565b96602089019060ff60208092828151168552015116910152565b606087019067ffffffffffffffff60208092828151168552015116910152565b63ffffffff811660a086015263ffffffff8160201c1660c086015260401c16151560e08401526108d98161140c565b6101008301520390f35b346101e95760206003193601126101e95760043567ffffffffffffffff81116101e957366023820112156101e95780600401359067ffffffffffffffff82116101e9573660248360051b830101116101e9579060206040519061094681836113e9565b5f825280820193601f19820136863761095e846123ca565b9061096c60405192836113e9565b848252601f1961097b866123ca565b015f5b818110610b045750505f907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffbd81360301915b86811015610a855760248160051b83010135838112156101e95782019060248201359167ffffffffffffffff83116101e9576044019180360383136101e9575f806001948a610a2e610a61958f8d906040519483869484860198893784019083820190898252519283915e010185815203601f1981018352826113e9565b5190305af43d15610a7d573d90610a448261146a565b91610a5260405193846113e9565b82523d5f8a84013e5b30614185565b610a6b828761313a565b52610a76818661313a565b50016109b0565b606090610a5b565b5050506040519082820192808352815180945260408301938160408260051b8601019301915f955b828710610aba5785850386f35b909192938280610af4837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08a600196030186528851611246565b9601920196019592919092610aad565b6060848201860152840161097e565b346101e9575f6003193601126101e95760206040515f8152f35b346101e957610b3b36611213565b905f525f60205273ffffffffffffffffffffffffffffffffffffffff60405f2091165f52602052602060ff60405f2054166040519015158152f35b346101e957610b84366111ad565b505060ff60045460401c1661077b575f80527f4c48d1d2b0f4f7485fc28e3db22341d96a20aa29e6efa8149da9751603abd4e06020527f6e0c24a6e293ff9b755263dbaa15ba3796b0b8d3fe17cfb4ddf8143b268eac475460ff1615610c0d575b7fda81d7c2000000000000000000000000000000000000000000000000000000005f5260045ffd5b610c1561317b565b610be5565b346101e9575f6003193601126101e957602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346101e9575f6003193601126101e957602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346101e9575f6003193601126101e95760206040517fbd893629a699470e4ec82a5715bb4981fdaacc5d0a728bf5f55b801d8f4ef10b8152f35b346101e957610d0236611213565b3373ffffffffffffffffffffffffffffffffffffffff821603610d28576107d791613c81565b7f6697b232000000000000000000000000000000000000000000000000000000005f5260045ffd5b346101e9576107d7610d6136611213565b90610d7a6107cd825f525f602052600160405f20015490565b613baf565b346101e9575f6003193601126101e95760206040516107088152f35b346101e95760206003193601126101e9576020610dc56004355f525f602052600160405f20015490565b604051908152f35b346101e95760206003193601126101e95760043567ffffffffffffffff811681036101e957610dc56020916130f3565b346101e957610e0b36611142565b60ff60045460401c1661077b575f80527f4c48d1d2b0f4f7485fc28e3db22341d96a20aa29e6efa8149da9751603abd4e06020527f6e0c24a6e293ff9b755263dbaa15ba3796b0b8d3fe17cfb4ddf8143b268eac475460ff1615610eaf575b610e776040820182611416565b90610e856060840184611416565b9190936101008101359260028410156101e957602095610dc59560a0840194608085013594613269565b610eb761317b565b610e6a565b346101e957610eca366111ad565b9060ff60045460401c1661077b575f80527f4c48d1d2b0f4f7485fc28e3db22341d96a20aa29e6efa8149da9751603abd4e060209081527f6e0c24a6e293ff9b755263dbaa15ba3796b0b8d3fe17cfb4ddf8143b268eac47549092610f3892909160ff1615610f4957612e69565b60405190610f4581611176565b8152f35b610f5161317b565b612e69565b346101e9575f6003193601126101e957602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346101e95760206003193601126101e95760043567ffffffffffffffff81116101e95761024060031982360301126101e95760ff60045460401c1661077b575f80527f4c48d1d2b0f4f7485fc28e3db22341d96a20aa29e6efa8149da9751603abd4e060209081527f6e0c24a6e293ff9b755263dbaa15ba3796b0b8d3fe17cfb4ddf8143b268eac47549091610f389160ff1615611047575b600401611a47565b61104f61317b565b61103f565b346101e9575f6003193601126101e957602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346101e95760206003193601126101e957600435907fffffffff0000000000000000000000000000000000000000000000000000000082168092036101e957817f7965db0b0000000000000000000000000000000000000000000000000000000060209314908115611118575b5015158152f35b7f01ffc9a70000000000000000000000000000000000000000000000000000000091501483611111565b60206003198201126101e9576004359067ffffffffffffffff82116101e95760031982610120920301126101e95760040190565b6003111561118057565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b9060206003198301126101e95760043567ffffffffffffffff81116101e957826023820112156101e95780600401359267ffffffffffffffff84116101e957602484830101116101e9576024019190565b359067ffffffffffffffff821682036101e957565b60031960409101126101e9576004359060243573ffffffffffffffffffffffffffffffffffffffff811681036101e95790565b90601f19601f602080948051918291828752018686015e5f8582860101520116010190565b90600182811c921680156112b2575b602083101461128557565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b91607f169161127a565b6001545f92916112cb8261126b565b808252916001811690811561133f57506001146112e6575050565b60015f9081529293509091907fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf65b838310611325575060209250010190565b600181602092949394548385870101520191019190611314565b60209495507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091509291921683830152151560051b010190565b6040810190811067ffffffffffffffff82111761074157604052565b6060810190811067ffffffffffffffff82111761074157604052565b60e0810190811067ffffffffffffffff82111761074157604052565b6080810190811067ffffffffffffffff82111761074157604052565b90601f601f19910116810190811067ffffffffffffffff82111761074157604052565b6002111561118057565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156101e9570180359067ffffffffffffffff82116101e957602001918160051b360383136101e957565b67ffffffffffffffff811161074157601f01601f191660200190565b9291926114928261146a565b916114a060405193846113e9565b8294818452818301116101e9578281602093845f96015e010152565b519060ff821682036101e957565b519067ffffffffffffffff821682036101e957565b91908260409103126101e9576040516114f781611379565b6020611510818395611508816114ca565b8552016114ca565b910152565b519063ffffffff821682036101e957565b51906fffffffffffffffffffffffffffffffff821682036101e957565b91908260609103126101e95760405161155b81611395565b604080829461156981611526565b8452602081015160208501520151910152565b6020818303126101e95780519067ffffffffffffffff82116101e95701610180818303126101e9576040519160c0830183811067ffffffffffffffff82111761074157604052815167ffffffffffffffff81116101e9578201918282039261012084126101e957604051936115f0856113b1565b815167ffffffffffffffff81116101e957820184601f820112156101e95760409161162486836020601f1995519101611486565b875201126101e95760405161163881611379565b611644602083016114bc565b8152611652604083016114bc565b6020820152602085015261166983606083016114df565b604085015261167a60a08201611515565b606085015261168b60c08201611515565b608085015260e08101519081151582036101e9576101009160a086015201519060028210156101e957836101409260c061171196015285526116d08360208301611543565b60208601526116e28360808301611543565b60408601526116f360e08201611526565b60608601526117068361010083016114df565b6080860152016114df565b60a082015290565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1823603018112156101e957016020813591019167ffffffffffffffff82116101e95781360383136101e957565b601f8260209493601f1993818652868601375f8582860101520116010190565b359060ff821682036101e957565b67ffffffffffffffff6117ba60208093836117b1826111fe565b168652016111fe565b16910152565b359063ffffffff821682036101e957565b359081151582036101e957565b35906fffffffffffffffffffffffffffffffff821682036101e957565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff61823603018112156101e9570190565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc1823603018112156101e9570190565b604080918035845263ffffffff611878602083016117c0565b1660208501520135910152565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1823603018112156101e957016020813591019167ffffffffffffffff82116101e9578160051b360383136101e957565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81823603018112156101e9570190565b35908160070b82036101e957565b9060606119668161193a61192c8680611719565b608087526080870191611769565b946020810135602086015267ffffffffffffffff61195a604083016111fe565b1660408601520161190a565b60070b91015290565b906080810161197e8380611885565b809192608085525260a083019060a08160051b8501019280925f915b8383106119f457505050505067ffffffffffffffff6119ed60606119e681946119c5602089016117d1565b151560208801526119d960408901896118d8565b8782036040890152611918565b95016111fe565b1691015290565b9091929394602080611a38837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff608b60019603018752611a338a876118d8565b611918565b9701930193019193929061199a565b604051907f58d248c00000000000000000000000000000000000000000000000000000000082526020600483015280357ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee1823603018112156101e95781016102406024840152610100611ad1611abd8380611719565b610120610264880152610384870191611769565b9160ff611ae060208301611789565b1661028486015260ff611af560408301611789565b166102a4860152611b0d6102c4860160608301611797565b63ffffffff611b1e60a083016117c0565b1661030486015263ffffffff611b3660c083016117c0565b16610324860152611b4960e082016117d1565b1515610344860152013560028110156101e957611b658161140c565b6103648401526fffffffffffffffffffffffffffffffff611b88602084016117de565b1660448401526040820135606484015260608201356084840152611baf60808301836117fb565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc8482030160a4850152611be4828061182d565b60a08252803590803603917ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffd4183018112156101e95781611db09101604060a0860152611c3360e0860182611797565b611d7f611c5a611c466040840184611719565b6102c06101208a01526103a0890191611769565b9167ffffffffffffffff611c70606083016111fe565b166101408801526fffffffffffffffffffffffffffffffff611c94608083016117de565b16610160880152611ca760a082016117d1565b1515610180880152611cc06101a0880160c0830161185f565b611ccd61012082016117d1565b1515610200880152610140810135610220880152611cee61016082016117d1565b15156102408801526101808101356102608801526101a08101356102808801526101c08101356102a08801526101e08101356102c08801526102008101356102e0880152611d3f61022082016117d1565b1515610300880152610240810135610320880152611d6061026082016117d1565b15156103408801526102808101356103608801526102a0810190611719565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff2087840301610380880152611769565b917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff41602083013591018112156101e957908694939291017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff608383030160c0840152611e5f60c083019167ffffffffffffffff611e2b826111fe565b16845263ffffffff611e3f602083016117c0565b166020850152611e55604085016040830161185f565b60a0810190611885565b80919260c060a08601525260e083019060e08160051b85010193835f915b8383106120e357505050505050610200611edf845f9794611ed1611eb8604096611eab60208c9b01866118d8565b848203602086015261196f565b92611ec7878401888301611797565b60808101906118d8565b90608081840391015261196f565b926fffffffffffffffffffffffffffffffff611efd60a083016117de565b1660c486015261010060c0820160e4870137826101c082016101e487013701610224840137038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115610736575f916120c1575b50611f7581613a6e565b611f7e81613aee565b90611f8882611176565b816120625767ffffffffffffffff6020604060a0840193845183810151600354918683861c1687831611612019575b5050500151604051611ff28382018093604080916fffffffffffffffffffffffffffffffff8151168452602081015160208501520151910152565b606081526120016080826113e9565b51902092510151165f52600560205260405f20555b90565b6fffffffffffffffff0000000000000000877fffffffffffffffffffffffffffffffff0000000000000000000000000000000092511692861b16921617176003555f8080611fb7565b5061206c81611176565b600181036120aa57680100000000000000007fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff600454161760045590565b6120b381611176565b600281036120165750600290565b6120dd91503d805f833e6120d581836113e9565b81019061157c565b5f611f6b565b919395909294969798507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff2082820301865261211e878461182d565b90813560038110156101e9576121e460209261214c600195846121418796611176565b8352848101906118d8565b906040848201526121b46121746121638480611719565b6080604086015260c0850191611769565b926fffffffffffffffffffffffffffffffff6121918783016117de565b1660608401526121a3604082016117d1565b151560808401526060810190611719565b9160a07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc082860301910152611769565b98019601930190918a989796959492611e7d565b81601f820112156101e9576020813591016122128261146a565b9261222060405194856113e9565b828452828201116101e957815f92602092838601378301015290565b91908260409103126101e95760405161225481611379565b6020611510818395612265816111fe565b8552016111fe565b808203929161012084126101e95760405191612288836113b1565b8294813567ffffffffffffffff81116101e9576040916122ad85601f199386016121f8565b865201126101e9576122f5610100926040516122c881611379565b6122d460208501611789565b81526122e260408501611789565b602082015260208601526060830161223c565b604084015261230660a082016117c0565b606084015261231760c082016117c0565b608084015261232860e082016117d1565b60a084015201359060028210156101e95760c00152565b91908260609103126101e95760405161235781611395565b6040808294612365816117de565b8452602081013560208501520135910152565b8092910391606083126101e95760405161239181611379565b6040601f1982958435845201126101e95760209060408051936123b385611379565b6123be8482016117c0565b85520135828401520152565b67ffffffffffffffff81116107415760051b60200190565b9190916080818403126101e957604051906123fc826113cd565b819381359067ffffffffffffffff82116101e9578261242460609492611510948694016121f8565b85526020810135602086015261243c604082016111fe565b60408601520161190a565b9190916080818403126101e95760405190612461826113cd565b8193813567ffffffffffffffff81116101e957820181601f820112156101e957803561248c816123ca565b9161249a60405193846113e9565b81835260208084019260051b820101918483116101e95760208201905b838210612509575050505083526124d0602083016117d1565b602084015260408201359067ffffffffffffffff82116101e957826124fe60609492611510948694016123e2565b6040860152016111fe565b813567ffffffffffffffff81116101e95760209161252c888480948801016123e2565b8152019101906124b7565b919060a0838203126101e95760405190612550826113cd565b8193803567ffffffffffffffff81116101e95781016040818403126101e9576040519061257c82611379565b803567ffffffffffffffff81116101e95781016102c0818603126101e95760405190610260820182811067ffffffffffffffff821117610741576040526125c3868261223c565b8252604081013567ffffffffffffffff81116101e957866125e59183016121f8565b60208301526125f6606082016111fe565b6040830152612607608082016117de565b606083015261261860a082016117d1565b608083015261262a8660c08301612378565b60a083015261263c61012082016117d1565b60c083015261014081013560e083015261265961016082016117d1565b6101008301526101808101356101208301526101a08101356101408301526101c08101356101608301526101e08101356101808301526102008101356101a08301526126a861022082016117d1565b6101c08301526102408101356101e08301526126c761026082016117d1565b6102008301526102808101356102208301526102a08101359067ffffffffffffffff82116101e9576126fb918791016121f8565b610240820152825260208101359067ffffffffffffffff82116101e9570160c0818503126101e95760405190612730826113cd565b612739816111fe565b8252612747602082016117c0565b60208301526127598560408301612378565b604083015260a08101359067ffffffffffffffff82116101e9570184601f820112156101e95780359061278b826123ca565b9161279960405193846113e9565b80835260208084019160051b830101918783116101e95760208101915b8383106128265750505050606082015260208201528352602081013567ffffffffffffffff81116101e957826127ed918301612447565b60208401526127ff826040830161223c565b604084015260808101359167ffffffffffffffff83116101e9576060926115109201612447565b823567ffffffffffffffff81116101e9578201906040601f19838c0301126101e9576040519161285583611379565b602081013560038110156101e9578352604081013567ffffffffffffffff81116101e9576020910101906080828c03126101e95760405192612896846113cd565b823567ffffffffffffffff81116101e9578c6128b39185016121f8565b84526128c1602084016117de565b60208501526128d2604084016117d1565b604085015260608301359367ffffffffffffffff85116101e9576128fb8d6020968796016121f8565b6060820152838201528152019201916127b6565b9080601f830112156101e95760408051929061292b90846113e9565b8290604081019283116101e957905b8282106129475750505090565b813581526020918201910161293a565b9061010060c061297284516101208552610120850190611246565b602080860151805160ff90811687840152910151166040850152936129b56040820151606086019067ffffffffffffffff60208092828151168552015116910152565b63ffffffff60608201511660a085015263ffffffff6080820151168285015260a0810151151560e08501520151916129ec8361140c565b015290565b90606080612a088451608085526080850190611246565b936020810151602085015267ffffffffffffffff6040820151166040850152015160070b91015290565b906080810182519060808352815180915260a0830190602060a08260051b8601019301915f905b828210612a9c575050505067ffffffffffffffff6060612a928193602087015115156020870152604087015186820360408801526129f1565b9401511691015290565b90919293602080612ad7837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff608a6001960301865288516129f1565b960192019201909291612a59565b91909180519260a081526020612c6e8551604060a0850152612b2160e08501825167ffffffffffffffff60208092828151168552015116910152565b610240612b3f848301516102c06101208801526103a0870190611246565b604083015167ffffffffffffffff1661014087015260608301516fffffffffffffffffffffffffffffffff166101608701526080830151151561018087015260a083015180516101a0880152602090810151805163ffffffff166101c089015201516101e08701529160c0810151151561020087015260e08101516102208701526101008101511515828701526101208101516102608701526101408101516102808701526101608101516102a08701526101808101516102c08701526101a08101516102e08701526101c081015115156103008701526101e0810151610320870152610200810151151561034087015261022081015161036087015201517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff2085830301610380860152611246565b940151937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff608282030160c0830152606060c082019567ffffffffffffffff815116835263ffffffff6020820151166020840152612ced6040820151604085019060208060409280518552015163ffffffff815116828501520151910152565b01519460c060a0830152855180915260e0820190602060e08260051b8501019701915f905b828210612d735750505050506060612d3a612016949560208501518482036020860152612a32565b92612d636040820151604085019067ffffffffffffffff60208092828151168552015116910152565b0151906080818403910152612a32565b9091929397602080612e32837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff208d600196030186528289518051612db681611176565b83520151906040848201526060612dd983516080604085015260c0840190611246565b926fffffffffffffffffffffffffffffffff86820151168284015260408101511515608084015201519060a07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc082850301910152611246565b9a96019493919091019101612d12565b905f905b60028210612e5357505050565b6020806001928551815201930191019091612e46565b908101906020818303126101e95780359067ffffffffffffffff82116101e95701610240818303126101e95760405191612ea2836113b1565b813567ffffffffffffffff81116101e95781612ebf91840161226d565b8352612ece816020840161233f565b9160208401928352608081013567ffffffffffffffff81116101e95782612ef6918301612537565b9160408501928352612f0a60a083016117de565b93606086019485528160df840112156101e95760405191612f2d610100846113e9565b6101c08401838282116101e95760c08601905b8282106130e35750508795949392612f8b83610200612f7d6fffffffffffffffffffffffffffffffff9661303a966080612fda9f01998a5261290f565b9760a08b019889520161290f565b9560c0880196875261300a6040519a8b997f58d248c0000000000000000000000000000000000000000000000000000000008b52602060048c01525161024060248c01526102648b0190612957565b925180516fffffffffffffffffffffffffffffffff1660448b0152602081015160648b01526040015160848a0152565b517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc8883030160a4890152612ae5565b95511660c4850152515f60e485015b600882106130c95750505092613077839261306b5f96516101e4860190612e42565b51610224840190612e42565b038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115610736575f916120c15750611f7581613a6e565b825181528795506020928301926001929092019101613049565b8135815260209182019101612f40565b67ffffffffffffffff165f52600560205260405f205480156131125790565b7f5b48b457000000000000000000000000000000000000000000000000000000005f5260045ffd5b805182101561314e5760209160051b010190565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b335f9081527f4c48d1d2b0f4f7485fc28e3db22341d96a20aa29e6efa8149da9751603abd4e0602052604090205460ff16156131b357565b7fe2517d3f000000000000000000000000000000000000000000000000000000005f52336004527fbd893629a699470e4ec82a5715bb4981fdaacc5d0a728bf5f55b801d8f4ef10b60245260445ffd5b805f525f60205260405f2073ffffffffffffffffffffffffffffffffffffffff33165f5260205260ff60405f2054161561323a5750565b7fe2517d3f000000000000000000000000000000000000000000000000000000005f523360045260245260445ffd5b959693919092936132798161140c565b806139f65750801515806139ea575b156139b457939290846040519586957f13542b7a000000000000000000000000000000000000000000000000000000008752606487019060048801526060602488015252608485019060848160051b8701019480925f915b8383106138d3575050505050600319848403016044850152808352602083019260208260051b82010193835f927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe182360301905b85851061374e575050505050505090805f9203818373ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165af1908115610736575f916135a3575b506020815192019160206133a384613eee565b6134036133b0368861233f565b91610612604051858101906133eb8287604080916fffffffffffffffffffffffffffffffff8151168452602081015160208501520151910152565b606081526133fa6080826113e9565b519020916130f3565b015180820361357557505060200151906001825111613446575b50506fffffffffffffffffffffffffffffffff613440633b9aca0092369061233f565b51160490565b61345290939193613eee565b918035916fffffffffffffffffffffffffffffffff83168093036101e957909267ffffffffffffffff16905f5b855181101561355757613492818761313a565b519084604051602081019086825260408082015260a081019480519560406060840152865180915260c08301602060c08360051b8601019801915f5b81811061350257505050508160019660206134f8930151608083015203601f1981018352826113e9565b5190205d0161347f565b919394965091949697602080613542837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff408b600196030188528951611246565b970194019101918b96949392989795986134ce565b50935050506fffffffffffffffffffffffffffffffff61344061341d565b7f4b2dfe98000000000000000000000000000000000000000000000000000000005f5260045260245260445ffd5b90503d805f833e6135b481836113e9565b8101906020818303126101e95780519067ffffffffffffffff82116101e95701906040828203126101e957604051916135ec83611379565b8051835260208101519067ffffffffffffffff82116101e9570181601f820112156101e95780519061361d826123ca565b9261362b60405194856113e9565b82845260208085019360051b830101918183116101e95760208101935b83851061365f57505050505060208201525f613390565b845167ffffffffffffffff81116101e95782016040601f1982860301126101e9576040519061368d82611379565b602081015167ffffffffffffffff81116101e95760209082010185601f820112156101e95780516136bd816123ca565b916136cb60405193846113e9565b81835260208084019260051b820101908882116101e95760208101925b82841061370f57505050509160406020949285948352015183820152815201940193613648565b835167ffffffffffffffff81116101e95782018a603f820112156101e9576020916137438c83604086809601519101611486565b8152019301926136e8565b9193959750919395601f198282030185528735838112156101e95784019061377a602082019280611885565b8291936020829452526040810160408360051b8301019380935f915b8183106137bd57505050505050602080600192990195019501929091889796949592613334565b9091929394957fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08682030185526137f48783613d49565b9081359160028310156101e9576138c560209283928561381560019761140c565b8252613844613839613829868401846118d8565b6060878601526060850190613d7b565b9160408101906117fb565b9160408183039101526138b761389961386e6138608580611719565b60a0865260a0860191611769565b6138798786016117d1565b15158785015261388c60408601866118d8565b8482036040860152613d7b565b926138a6606082016117d1565b1515606084015260808101906118d8565b906080818403910152613d7b565b980196950193019190613796565b91939690929495977fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7c90820301835261390c878361182d565b90604081019161391c8180611885565b809460408552526060830160608560051b85010194825f5b82811061396257505050505060019260209283808094013591015298019301930190928897959492936132e0565b90919293966020806139a7837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa08b600196030189526139a18c88611719565b90611769565b9901950193929101613934565b7fb0369c31000000000000000000000000000000000000000000000000000000005f52600452600160245261ffff60445260645ffd5b5061ffff811115613288565b60ff90613a028161140c565b613a0b8161140c565b7f112d89cc000000000000000000000000000000000000000000000000000000005f521660045260245ffd5b15613a40575050565b7f6d23e8a3000000000000000000000000000000000000000000000000000000005f5260045260245260445ffd5b613aec90613a9481516fffffffffffffffffffffffffffffffff60608401511690613f03565b61061267ffffffffffffffff6020608081850151604051613ade8482018093604080916fffffffffffffffffffffffffffffffff8151168452602081015160208501520151910152565b6060815261065f83826113e9565b565b67ffffffffffffffff602060a08301510151165f52600560205260405f205480155f14613b1b5750505f90565b60408201908151604051613b59602082018093604080916fffffffffffffffffffffffffffffffff8151168452602081015160208501520151910152565b60608152613b686080826113e9565b5190201491821592613b86575b505015613b8157600190565b600290565b6fffffffffffffffffffffffffffffffff91925060208291015151169151511611155f80613b75565b805f525f60205260405f2073ffffffffffffffffffffffffffffffffffffffff83165f5260205260ff60405f205416155f14613c7b57805f525f60205260405f2073ffffffffffffffffffffffffffffffffffffffff83165f5260205260405f2060017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0082541617905573ffffffffffffffffffffffffffffffffffffffff339216907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d5f80a4600190565b50505f90565b805f525f60205260405f2073ffffffffffffffffffffffffffffffffffffffff83165f5260205260ff60405f2054165f14613c7b57805f525f60205260405f2073ffffffffffffffffffffffffffffffffffffffff83165f5260205260405f207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00815416905573ffffffffffffffffffffffffffffffffffffffff339216907ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b5f80a4600190565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa1823603018112156101e9570190565b613d96613d888280611719565b608085526080850191611769565b60208201356020840152613dad60408301836118d8565b908381036040850152813560038110156101e957613dca81611176565b8152602082013560038110156101e957613de381611176565b602082015260408201359060038210156101e9576080613e1f613e3a9484613e10613e2f96999899611176565b60408501526060810190611719565b9190928160608201520191611769565b926060810190611885565b90916060818503910152808352602083019060208160051b85010193835f915b838310613e6a5750505050505090565b909192939495601f19828203018652613e838784613d49565b80359160038310156101e957613ee0602092839285613ea3600197611176565b8152613ed2613ec7613eb786850185611719565b6060888601526060850191611769565b926040810190611719565b916040818503910152611769565b980196019493019190613e5a565b3567ffffffffffffffff811681036101e95790565b906fffffffffffffffffffffffffffffffff633b9aca009116044281116141565780420342811161412957610708106140fa5750805151613f4560015461126b565b14806140d3575b8151901561407d5750602081015160ff815116906002549160ff8316928382149283614064575b6020015160ff16921561401c575050505063ffffffff606082015116906004549163ffffffff8316808203613fee57505063ffffffff608081920151169160201c16808203613fc0575050565b7f1eb5c1eb000000000000000000000000000000000000000000000000000000005f5260045260245260445ffd5b7f73b1bde3000000000000000000000000000000000000000000000000000000005f5260045260245260445ffd5b6084945060ff90604051947fd382033a000000000000000000000000000000000000000000000000000000008652600486015260081c16602484015260448301526064820152fd5b6020810151600883901c60ff9081169116149350613f73565b6140cf906040519182917ff6b6676b000000000000000000000000000000000000000000000000000000008352604060048401526140bd604484016112bc565b90600319848303016024850152611246565b0390fd5b508051602081519101206040516140ed816107fc816112bc565b6020815191012014613f4c565b7f12e74add000000000000000000000000000000000000000000000000000000005f524260045260245260445ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b7f20daedb6000000000000000000000000000000000000000000000000000000005f524260045260245260445ffd5b906141c2575080511561419a57602081519101fd5b7fd6bda275000000000000000000000000000000000000000000000000000000005f5260045ffd5b81511580614215575b6141d3575090565b73ffffffffffffffffffffffffffffffffffffffff907f9996b315000000000000000000000000000000000000000000000000000000005f521660045260245ffd5b50803b156141cb56fea164736f6c634300081c000a2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d4c48d1d2b0f4f7485fc28e3db22341d96a20aa29e6efa8149da9751603abd4e06e0c24a6e293ff9b755263dbaa15ba3796b0b8d3fe17cfb4ddf8143b268eac47bd893629a699470e4ec82a5715bb4981fdaacc5d0a728bf5f55b801d8f4ef10bad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5",
}

// ContractSP1ICS07TendermintABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractSP1ICS07TendermintMetaData.ABI instead.
var ContractSP1ICS07TendermintABI = ContractSP1ICS07TendermintMetaData.ABI

// ContractSP1ICS07TendermintBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractSP1ICS07TendermintMetaData.Bin instead.
var ContractSP1ICS07TendermintBin = ContractSP1ICS07TendermintMetaData.Bin

// DeployContractSP1ICS07Tendermint deploys a new Ethereum contract, binding an instance of ContractSP1ICS07Tendermint to it.
func DeployContractSP1ICS07Tendermint(auth *bind.TransactOpts, backend bind.ContractBackend, verifier common.Address, membership_ common.Address, misbehaviour_ common.Address, updateClient_ common.Address, _clientState []byte, _consensusState [32]byte, roleManager common.Address) (common.Address, *types.Transaction, *ContractSP1ICS07Tendermint, error) {
	parsed, err := ContractSP1ICS07TendermintMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractSP1ICS07TendermintBin), backend, verifier, membership_, misbehaviour_, updateClient_, _clientState, _consensusState, roleManager)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractSP1ICS07Tendermint{ContractSP1ICS07TendermintCaller: ContractSP1ICS07TendermintCaller{contract: contract}, ContractSP1ICS07TendermintTransactor: ContractSP1ICS07TendermintTransactor{contract: contract}, ContractSP1ICS07TendermintFilterer: ContractSP1ICS07TendermintFilterer{contract: contract}}, nil
}

// ContractSP1ICS07Tendermint is an auto generated Go binding around an Ethereum contract.
type ContractSP1ICS07Tendermint struct {
	ContractSP1ICS07TendermintCaller     // Read-only binding to the contract
	ContractSP1ICS07TendermintTransactor // Write-only binding to the contract
	ContractSP1ICS07TendermintFilterer   // Log filterer for contract events
}

// ContractSP1ICS07TendermintCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractSP1ICS07TendermintCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSP1ICS07TendermintTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractSP1ICS07TendermintTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSP1ICS07TendermintFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractSP1ICS07TendermintFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSP1ICS07TendermintSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSP1ICS07TendermintSession struct {
	Contract     *ContractSP1ICS07Tendermint // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ContractSP1ICS07TendermintCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractSP1ICS07TendermintCallerSession struct {
	Contract *ContractSP1ICS07TendermintCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// ContractSP1ICS07TendermintTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractSP1ICS07TendermintTransactorSession struct {
	Contract     *ContractSP1ICS07TendermintTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// ContractSP1ICS07TendermintRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractSP1ICS07TendermintRaw struct {
	Contract *ContractSP1ICS07Tendermint // Generic contract binding to access the raw methods on
}

// ContractSP1ICS07TendermintCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractSP1ICS07TendermintCallerRaw struct {
	Contract *ContractSP1ICS07TendermintCaller // Generic read-only contract binding to access the raw methods on
}

// ContractSP1ICS07TendermintTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractSP1ICS07TendermintTransactorRaw struct {
	Contract *ContractSP1ICS07TendermintTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractSP1ICS07Tendermint creates a new instance of ContractSP1ICS07Tendermint, bound to a specific deployed contract.
func NewContractSP1ICS07Tendermint(address common.Address, backend bind.ContractBackend) (*ContractSP1ICS07Tendermint, error) {
	contract, err := bindContractSP1ICS07Tendermint(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractSP1ICS07Tendermint{ContractSP1ICS07TendermintCaller: ContractSP1ICS07TendermintCaller{contract: contract}, ContractSP1ICS07TendermintTransactor: ContractSP1ICS07TendermintTransactor{contract: contract}, ContractSP1ICS07TendermintFilterer: ContractSP1ICS07TendermintFilterer{contract: contract}}, nil
}

// NewContractSP1ICS07TendermintCaller creates a new read-only instance of ContractSP1ICS07Tendermint, bound to a specific deployed contract.
func NewContractSP1ICS07TendermintCaller(address common.Address, caller bind.ContractCaller) (*ContractSP1ICS07TendermintCaller, error) {
	contract, err := bindContractSP1ICS07Tendermint(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractSP1ICS07TendermintCaller{contract: contract}, nil
}

// NewContractSP1ICS07TendermintTransactor creates a new write-only instance of ContractSP1ICS07Tendermint, bound to a specific deployed contract.
func NewContractSP1ICS07TendermintTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractSP1ICS07TendermintTransactor, error) {
	contract, err := bindContractSP1ICS07Tendermint(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractSP1ICS07TendermintTransactor{contract: contract}, nil
}

// NewContractSP1ICS07TendermintFilterer creates a new log filterer instance of ContractSP1ICS07Tendermint, bound to a specific deployed contract.
func NewContractSP1ICS07TendermintFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractSP1ICS07TendermintFilterer, error) {
	contract, err := bindContractSP1ICS07Tendermint(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractSP1ICS07TendermintFilterer{contract: contract}, nil
}

// bindContractSP1ICS07Tendermint binds a generic wrapper to an already deployed contract.
func bindContractSP1ICS07Tendermint(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractSP1ICS07TendermintMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractSP1ICS07Tendermint *ContractSP1ICS07TendermintRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractSP1ICS07Tendermint.Contract.ContractSP1ICS07TendermintCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractSP1ICS07Tendermint *ContractSP1ICS07TendermintRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractSP1ICS07Tendermint.Contract.ContractSP1ICS07TendermintTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractSP1ICS07Tendermint *ContractSP1ICS07TendermintRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractSP1ICS07Tendermint.Contract.ContractSP1ICS07TendermintTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractSP1ICS07Tendermint *ContractSP1ICS07TendermintCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractSP1ICS07Tendermint.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractSP1ICS07Tendermint *ContractSP1ICS07TendermintTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractSP1ICS07Tendermint.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractSP1ICS07Tendermint *ContractSP1ICS07TendermintTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractSP1ICS07Tendermint.Contract.contract.Transact(opts, method, params...)
}

// ALLOWEDSP1CLOCKDRIFT is a free data retrieval call binding the contract method 0x2c3ee474.
//
// Solidity: function ALLOWED_SP1_CLOCK_DRIFT() view returns(uint16)
func (_ContractSP1ICS07Tendermint *ContractSP1ICS07TendermintCaller) ALLOWEDSP1CLOCKDRIFT(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _ContractSP1ICS07Tendermint.contract.Call(opts, &out, "ALLOWED_SP1_CLOCK_DRIFT")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// ALLOWEDSP1CLOCKDRIFT is a free data retrieval call binding the contract method 0x2c3ee474.
//
// Solidity: function ALLOWED_SP1_CLOCK_DRIFT() view returns(uint16)
func (_ContractSP1ICS07Tendermint *ContractSP1ICS07TendermintSession) ALLOWEDSP1CLOCKDRIFT() (uint16, error) {
	return _ContractSP1ICS07Tendermint.Contract.ALLOWEDSP1CLOCKDRIFT(&_ContractSP1ICS07Tendermint.CallOpts)
}

// ALLOWEDSP1CLOCKDRIFT is a free data retrieval call binding the contract method 0x2c3ee474.
//
// Solidity: function ALLOWED_SP1_CLOCK_DRIFT() view returns(uint16)
func (_ContractSP1ICS07Tendermint *ContractSP1ICS07TendermintCallerSession) ALLOWEDSP1CLOCKDRIFT() (uint16, error) {
	return _ContractSP1ICS07Tendermint.Contract.ALLOWEDSP1CLOCKDRIFT(&_ContractSP1ICS07Tendermint.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ContractSP1ICS07Tendermint *ContractSP1ICS07TendermintCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ContractSP1ICS07Tendermint.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ContractSP1ICS07Tendermint *ContractSP1ICS07TendermintSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ContractSP1ICS07Tendermint.Contract.DEFAULTADMINROLE(&_ContractSP1ICS07Tendermint.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ContractSP1ICS07Tendermint *ContractSP1ICS07TendermintCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ContractSP1ICS07Tendermint.Contract.DEFAULTADMINROLE(&_ContractSP1ICS07Tendermint.CallOpts)
}

// MEMBERSHIP is a free data retrieval call binding the contract method 0x89df51f1.
//
// Solidity: function MEMBERSHIP() view returns(address)
func (_ContractSP1ICS07Tendermint *ContractSP1ICS07TendermintCaller) MEMBERSHIP(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractSP1ICS07Tendermint.contract.Call(opts, &out, "MEMBERSHIP")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MEMBERSHIP is a free data retrieval call binding the contract method 0x89df51f1.
//
// Solidity: function MEMBERSHIP() view returns(address)
func (_ContractSP1ICS07Tendermint *ContractSP1ICS07TendermintSession) MEMBERSHIP() (common.Address, error) {
	return _ContractSP1ICS07Tendermint.Contract.MEMBERSHIP(&_ContractSP1ICS07Tendermint.CallOpts)
}

// MEMBERSHIP is a free data retrieval call binding the contract method 0x89df51f1.
//
// Solidity: function MEMBERSHIP() view returns(address)
func (_ContractSP1ICS07Tendermint *ContractSP1ICS07TendermintCallerSession) MEMBERSHIP() (common.Address, error) {
	return _ContractSP1ICS07Tendermint.Contract.MEMBERSHIP(&_ContractSP1ICS07Tendermint.CallOpts)
}

// MISBEHAVIOUR is a free data retrieval call binding the contract method 0x02cf2952.
//
// Solidity: function MISBEHAVIOUR() view returns(address)
func (_ContractSP1ICS07Tendermint *ContractSP1ICS07TendermintCaller) MISBEHAVIOUR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractSP1ICS07Tendermint.contract.Call(opts, &out, "MISBEHAVIOUR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MISBEHAVIOUR is a free data retrieval call binding the contract method 0x02cf2952.
//
// Solidity: function MISBEHAVIOUR() view returns(address)
func (_ContractSP1ICS07Tendermint *ContractSP1ICS07TendermintSession) MISBEHAVIOUR() (common.Address, error) {
	return _ContractSP1ICS07Tendermint.Contract.MISBEHAVIOUR(&_ContractSP1ICS07Tendermint.CallOpts)
}

// MISBEHAVIOUR is a free data retrieval call binding the contract method 0x02cf2952.
//
// Solidity: function MISBEHAVIOUR() view returns(address)
func (_ContractSP1ICS07Tendermint *ContractSP1ICS07TendermintCallerSession) MISBEHAVIOUR() (common.Address, error) {
	return _ContractSP1ICS07Tendermint.Contract.MISBEHAVIOUR(&_ContractSP1ICS07Tendermint.CallOpts)
}

// PROOFSUBMITTERROLE is a free data retrieval call binding the contract method 0x5972185a.
//
// Solidity: function PROOF_SUBMITTER_ROLE() view returns(bytes32)
func (_ContractSP1ICS07Tendermint *ContractSP1ICS07TendermintCaller) PROOFSUBMITTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ContractSP1ICS07Tendermint.contract.Call(opts, &out, "PROOF_SUBMITTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PROOFSUBMITTERROLE is a free data retrieval call binding the contract method 0x5972185a.
//
// Solidity: function PROOF_SUBMITTER_ROLE() view returns(bytes32)
func (_ContractSP1ICS07Tendermint *ContractSP1ICS07TendermintSession) PROOFSUBMITTERROLE() ([32]byte, error) {
	return _ContractSP1ICS07Tendermint.Contract.PROOFSUBMITTERROLE(&_ContractSP1ICS07Tendermint.CallOpts)
}

// PROOFSUBMITTERROLE is a free data retrieval call binding the contract method 0x5972185a.
//
// Solidity: function PROOF_SUBMITTER_ROLE() view returns(bytes32)
func (_ContractSP1ICS07Tendermint *ContractSP1ICS07TendermintCallerSession) PROOFSUBMITTERROLE() ([32]byte, error) {
	return _ContractSP1ICS07Tendermint.Contract.PROOFSUBMITTERROLE(&_ContractSP1ICS07Tendermint.CallOpts)
}

// UPDATECLIENT is a free data retrieval call binding the contract method 0x87d4332f.
//
// Solidity: function UPDATE_CLIENT() view returns(address)
func (_ContractSP1ICS07Tendermint *ContractSP1ICS07TendermintCaller) UPDATECLIENT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractSP1ICS07Tendermint.contract.Call(opts, &out, "UPDATE_CLIENT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UPDATECLIENT is a free data retrieval call binding the contract method 0x87d4332f.
//
// Solidity: function UPDATE_CLIENT() view returns(address)
func (_ContractSP1ICS07Tendermint *ContractSP1ICS07TendermintSession) UPDATECLIENT() (common.Address, error) {
	return _ContractSP1ICS07Tendermint.Contract.UPDATECLIENT(&_ContractSP1ICS07Tendermint.CallOpts)
}

// UPDATECLIENT is a free data retrieval call binding the contract method 0x87d4332f.
//
// Solidity: function UPDATE_CLIENT() view returns(address)
func (_ContractSP1ICS07Tendermint *ContractSP1ICS07TendermintCallerSession) UPDATECLIENT() (common.Address, error) {
	return _ContractSP1ICS07Tendermint.Contract.UPDATECLIENT(&_ContractSP1ICS07Tendermint.CallOpts)
}

// VERIFIER is a free data retrieval call binding the contract method 0x08c84e70.
//
// Solidity: function VERIFIER() view returns(address)
func (_ContractSP1ICS07Tendermint *ContractSP1ICS07TendermintCaller) VERIFIER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractSP1ICS07Tendermint.contract.Call(opts, &out, "VERIFIER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VERIFIER is a free data retrieval call binding the contract method 0x08c84e70.
//
// Solidity: function VERIFIER() view returns(address)
func (_ContractSP1ICS07Tendermint *ContractSP1ICS07TendermintSession) VERIFIER() (common.Address, error) {
	return _ContractSP1ICS07Tendermint.Contract.VERIFIER(&_ContractSP1ICS07Tendermint.CallOpts)
}

// VERIFIER is a free data retrieval call binding the contract method 0x08c84e70.
//
// Solidity: function VERIFIER() view returns(address)
func (_ContractSP1ICS07Tendermint *ContractSP1ICS07TendermintCallerSession) VERIFIER() (common.Address, error) {
	return _ContractSP1ICS07Tendermint.Contract.VERIFIER(&_ContractSP1ICS07Tendermint.CallOpts)
}

// ClientState is a free data retrieval call binding the contract method 0xbd3ce6b0.
//
// Solidity: function clientState() view returns(string chainId, (uint8,uint8) trustLevel, (uint64,uint64) latestHeight, uint32 trustingPeriod, uint32 unbondingPeriod, bool isFrozen, uint8 zkAlgorithm)
func (_ContractSP1ICS07Tendermint *ContractSP1ICS07TendermintCaller) ClientState(opts *bind.CallOpts) (struct {
	ChainId         string
	TrustLevel      IICS07TendermintMsgsTrustThreshold
	LatestHeight    IICS02ClientMsgsHeight
	TrustingPeriod  uint32
	UnbondingPeriod uint32
	IsFrozen        bool
	ZkAlgorithm     uint8
}, error) {
	var out []interface{}
	err := _ContractSP1ICS07Tendermint.contract.Call(opts, &out, "clientState")

	outstruct := new(struct {
		ChainId         string
		TrustLevel      IICS07TendermintMsgsTrustThreshold
		LatestHeight    IICS02ClientMsgsHeight
		TrustingPeriod  uint32
		UnbondingPeriod uint32
		IsFrozen        bool
		ZkAlgorithm     uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ChainId = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.TrustLevel = *abi.ConvertType(out[1], new(IICS07TendermintMsgsTrustThreshold)).(*IICS07TendermintMsgsTrustThreshold)
	outstruct.LatestHeight = *abi.ConvertType(out[2], new(IICS02ClientMsgsHeight)).(*IICS02ClientMsgsHeight)
	outstruct.TrustingPeriod = *abi.ConvertType(out[3], new(uint32)).(*uint32)
	outstruct.UnbondingPeriod = *abi.ConvertType(out[4], new(uint32)).(*uint32)
	outstruct.IsFrozen = *abi.ConvertType(out[5], new(bool)).(*bool)
	outstruct.ZkAlgorithm = *abi.ConvertType(out[6], new(uint8)).(*uint8)

	return *outstruct, err

}

// ClientState is a free data retrieval call binding the contract method 0xbd3ce6b0.
//
// Solidity: function clientState() view returns(string chainId, (uint8,uint8) trustLevel, (uint64,uint64) latestHeight, uint32 trustingPeriod, uint32 unbondingPeriod, bool isFrozen, uint8 zkAlgorithm)
func (_ContractSP1ICS07Tendermint *ContractSP1ICS07TendermintSession) ClientState() (struct {
	ChainId         string
	TrustLevel      IICS07TendermintMsgsTrustThreshold
	LatestHeight    IICS02ClientMsgsHeight
	TrustingPeriod  uint32
	UnbondingPeriod uint32
	IsFrozen        bool
	ZkAlgorithm     uint8
}, error) {
	return _ContractSP1ICS07Tendermint.Contract.ClientState(&_ContractSP1ICS07Tendermint.CallOpts)
}

// ClientState is a free data retrieval call binding the contract method 0xbd3ce6b0.
//
// Solidity: function clientState() view returns(string chainId, (uint8,uint8) trustLevel, (uint64,uint64) latestHeight, uint32 trustingPeriod, uint32 unbondingPeriod, bool isFrozen, uint8 zkAlgorithm)
func (_ContractSP1ICS07Tendermint *ContractSP1ICS07TendermintCallerSession) ClientState() (struct {
	ChainId         string
	TrustLevel      IICS07TendermintMsgsTrustThreshold
	LatestHeight    IICS02ClientMsgsHeight
	TrustingPeriod  uint32
	UnbondingPeriod uint32
	IsFrozen        bool
	ZkAlgorithm     uint8
}, error) {
	return _ContractSP1ICS07Tendermint.Contract.ClientState(&_ContractSP1ICS07Tendermint.CallOpts)
}

// GetClientState is a free data retrieval call binding the contract method 0xef913a4b.
//
// Solidity: function getClientState() view returns(bytes)
func (_ContractSP1ICS07Tendermint *ContractSP1ICS07TendermintCaller) GetClientState(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _ContractSP1ICS07Tendermint.contract.Call(opts, &out, "getClientState")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetClientState is a free data retrieval call binding the contract method 0xef913a4b.
//
// Solidity: function getClientState() view returns(bytes)
func (_ContractSP1ICS07Tendermint *ContractSP1ICS07TendermintSession) GetClientState() ([]byte, error) {
	return _ContractSP1ICS07Tendermint.Contract.GetClientState(&_ContractSP1ICS07Tendermint.CallOpts)
}

// GetClientState is a free data retrieval call binding the contract method 0xef913a4b.
//
// Solidity: function getClientState() view returns(bytes)
func (_ContractSP1ICS07Tendermint *ContractSP1ICS07TendermintCallerSession) GetClientState() ([]byte, error) {
	return _ContractSP1ICS07Tendermint.Contract.GetClientState(&_ContractSP1ICS07Tendermint.CallOpts)
}

// GetConsensusStateHash is a free data retrieval call binding the contract method 0x23842fb8.
//
// Solidity: function getConsensusStateHash(uint64 revisionHeight) view returns(bytes32)
func (_ContractSP1ICS07Tendermint *ContractSP1ICS07TendermintCaller) GetConsensusStateHash(opts *bind.CallOpts, revisionHeight uint64) ([32]byte, error) {
	var out []interface{}
	err := _ContractSP1ICS07Tendermint.contract.Call(opts, &out, "getConsensusStateHash", revisionHeight)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetConsensusStateHash is a free data retrieval call binding the contract method 0x23842fb8.
//
// Solidity: function getConsensusStateHash(uint64 revisionHeight) view returns(bytes32)
func (_ContractSP1ICS07Tendermint *ContractSP1ICS07TendermintSession) GetConsensusStateHash(revisionHeight uint64) ([32]byte, error) {
	return _ContractSP1ICS07Tendermint.Contract.GetConsensusStateHash(&_ContractSP1ICS07Tendermint.CallOpts, revisionHeight)
}

// GetConsensusStateHash is a free data retrieval call binding the contract method 0x23842fb8.
//
// Solidity: function getConsensusStateHash(uint64 revisionHeight) view returns(bytes32)
func (_ContractSP1ICS07Tendermint *ContractSP1ICS07TendermintCallerSession) GetConsensusStateHash(revisionHeight uint64) ([32]byte, error) {
	return _ContractSP1ICS07Tendermint.Contract.GetConsensusStateHash(&_ContractSP1ICS07Tendermint.CallOpts, revisionHeight)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ContractSP1ICS07Tendermint *ContractSP1ICS07TendermintCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ContractSP1ICS07Tendermint.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ContractSP1ICS07Tendermint *ContractSP1ICS07TendermintSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ContractSP1ICS07Tendermint.Contract.GetRoleAdmin(&_ContractSP1ICS07Tendermint.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ContractSP1ICS07Tendermint *ContractSP1ICS07TendermintCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ContractSP1ICS07Tendermint.Contract.GetRoleAdmin(&_ContractSP1ICS07Tendermint.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ContractSP1ICS07Tendermint *ContractSP1ICS07TendermintCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _ContractSP1ICS07Tendermint.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ContractSP1ICS07Tendermint *ContractSP1ICS07TendermintSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ContractSP1ICS07Tendermint.Contract.HasRole(&_ContractSP1ICS07Tendermint.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ContractSP1ICS07Tendermint *ContractSP1ICS07TendermintCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ContractSP1ICS07Tendermint.Contract.HasRole(&_ContractSP1ICS07Tendermint.CallOpts, role, account)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ContractSP1ICS07Tendermint *ContractSP1ICS07TendermintCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ContractSP1ICS07Tendermint.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ContractSP1ICS07Tendermint *ContractSP1ICS07TendermintSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ContractSP1ICS07Tendermint.Contract.SupportsInterface(&_ContractSP1ICS07Tendermint.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ContractSP1ICS07Tendermint *ContractSP1ICS07TendermintCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ContractSP1ICS07Tendermint.Contract.SupportsInterface(&_ContractSP1ICS07Tendermint.CallOpts, interfaceId)
}

// UpgradeClient is a free data retrieval call binding the contract method 0x8a8e4c5d.
//
// Solidity: function upgradeClient(bytes ) view returns()
func (_ContractSP1ICS07Tendermint *ContractSP1ICS07TendermintCaller) UpgradeClient(opts *bind.CallOpts, arg0 []byte) error {
	var out []interface{}
	err := _ContractSP1ICS07Tendermint.contract.Call(opts, &out, "upgradeClient", arg0)

	if err != nil {
		return err
	}

	return err

}

// UpgradeClient is a free data retrieval call binding the contract method 0x8a8e4c5d.
//
// Solidity: function upgradeClient(bytes ) view returns()
func (_ContractSP1ICS07Tendermint *ContractSP1ICS07TendermintSession) UpgradeClient(arg0 []byte) error {
	return _ContractSP1ICS07Tendermint.Contract.UpgradeClient(&_ContractSP1ICS07Tendermint.CallOpts, arg0)
}

// UpgradeClient is a free data retrieval call binding the contract method 0x8a8e4c5d.
//
// Solidity: function upgradeClient(bytes ) view returns()
func (_ContractSP1ICS07Tendermint *ContractSP1ICS07TendermintCallerSession) UpgradeClient(arg0 []byte) error {
	return _ContractSP1ICS07Tendermint.Contract.UpgradeClient(&_ContractSP1ICS07Tendermint.CallOpts, arg0)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ContractSP1ICS07Tendermint *ContractSP1ICS07TendermintTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ContractSP1ICS07Tendermint.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ContractSP1ICS07Tendermint *ContractSP1ICS07TendermintSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ContractSP1ICS07Tendermint.Contract.GrantRole(&_ContractSP1ICS07Tendermint.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ContractSP1ICS07Tendermint *ContractSP1ICS07TendermintTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ContractSP1ICS07Tendermint.Contract.GrantRole(&_ContractSP1ICS07Tendermint.TransactOpts, role, account)
}

// Misbehaviour is a paid mutator transaction binding the contract method 0xddba6537.
//
// Solidity: function misbehaviour(bytes misbehaviourMsg) returns()
func (_ContractSP1ICS07Tendermint *ContractSP1ICS07TendermintTransactor) Misbehaviour(opts *bind.TransactOpts, misbehaviourMsg []byte) (*types.Transaction, error) {
	return _ContractSP1ICS07Tendermint.contract.Transact(opts, "misbehaviour", misbehaviourMsg)
}

// Misbehaviour is a paid mutator transaction binding the contract method 0xddba6537.
//
// Solidity: function misbehaviour(bytes misbehaviourMsg) returns()
func (_ContractSP1ICS07Tendermint *ContractSP1ICS07TendermintSession) Misbehaviour(misbehaviourMsg []byte) (*types.Transaction, error) {
	return _ContractSP1ICS07Tendermint.Contract.Misbehaviour(&_ContractSP1ICS07Tendermint.TransactOpts, misbehaviourMsg)
}

// Misbehaviour is a paid mutator transaction binding the contract method 0xddba6537.
//
// Solidity: function misbehaviour(bytes misbehaviourMsg) returns()
func (_ContractSP1ICS07Tendermint *ContractSP1ICS07TendermintTransactorSession) Misbehaviour(misbehaviourMsg []byte) (*types.Transaction, error) {
	return _ContractSP1ICS07Tendermint.Contract.Misbehaviour(&_ContractSP1ICS07Tendermint.TransactOpts, misbehaviourMsg)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_ContractSP1ICS07Tendermint *ContractSP1ICS07TendermintTransactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _ContractSP1ICS07Tendermint.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_ContractSP1ICS07Tendermint *ContractSP1ICS07TendermintSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _ContractSP1ICS07Tendermint.Contract.Multicall(&_ContractSP1ICS07Tendermint.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_ContractSP1ICS07Tendermint *ContractSP1ICS07TendermintTransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _ContractSP1ICS07Tendermint.Contract.Multicall(&_ContractSP1ICS07Tendermint.TransactOpts, data)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ContractSP1ICS07Tendermint *ContractSP1ICS07TendermintTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ContractSP1ICS07Tendermint.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ContractSP1ICS07Tendermint *ContractSP1ICS07TendermintSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ContractSP1ICS07Tendermint.Contract.RenounceRole(&_ContractSP1ICS07Tendermint.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ContractSP1ICS07Tendermint *ContractSP1ICS07TendermintTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ContractSP1ICS07Tendermint.Contract.RenounceRole(&_ContractSP1ICS07Tendermint.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ContractSP1ICS07Tendermint *ContractSP1ICS07TendermintTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ContractSP1ICS07Tendermint.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ContractSP1ICS07Tendermint *ContractSP1ICS07TendermintSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ContractSP1ICS07Tendermint.Contract.RevokeRole(&_ContractSP1ICS07Tendermint.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ContractSP1ICS07Tendermint *ContractSP1ICS07TendermintTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ContractSP1ICS07Tendermint.Contract.RevokeRole(&_ContractSP1ICS07Tendermint.TransactOpts, role, account)
}

// UpdateClient is a paid mutator transaction binding the contract method 0x0bece356.
//
// Solidity: function updateClient(bytes updateClientMsg) returns(uint8)
func (_ContractSP1ICS07Tendermint *ContractSP1ICS07TendermintTransactor) UpdateClient(opts *bind.TransactOpts, updateClientMsg []byte) (*types.Transaction, error) {
	return _ContractSP1ICS07Tendermint.contract.Transact(opts, "updateClient", updateClientMsg)
}

// UpdateClient is a paid mutator transaction binding the contract method 0x0bece356.
//
// Solidity: function updateClient(bytes updateClientMsg) returns(uint8)
func (_ContractSP1ICS07Tendermint *ContractSP1ICS07TendermintSession) UpdateClient(updateClientMsg []byte) (*types.Transaction, error) {
	return _ContractSP1ICS07Tendermint.Contract.UpdateClient(&_ContractSP1ICS07Tendermint.TransactOpts, updateClientMsg)
}

// UpdateClient is a paid mutator transaction binding the contract method 0x0bece356.
//
// Solidity: function updateClient(bytes updateClientMsg) returns(uint8)
func (_ContractSP1ICS07Tendermint *ContractSP1ICS07TendermintTransactorSession) UpdateClient(updateClientMsg []byte) (*types.Transaction, error) {
	return _ContractSP1ICS07Tendermint.Contract.UpdateClient(&_ContractSP1ICS07Tendermint.TransactOpts, updateClientMsg)
}

// UpdateClientMsg is a paid mutator transaction binding the contract method 0x04354dac.
//
// Solidity: function updateClientMsg(((string,(uint8,uint8),(uint64,uint64),uint32,uint32,bool,uint8),(uint128,bytes32,bytes32),((((uint64,uint64),string,uint64,uint128,bool,(bytes32,(uint32,bytes32)),bool,bytes32,bool,bytes32,bytes32,bytes32,bytes32,bytes32,bool,bytes32,bool,bytes32,bytes),(uint64,uint32,(bytes32,(uint32,bytes32)),(uint8,(bytes,uint128,bool,bytes))[])),((bytes,bytes32,uint64,int64)[],bool,(bytes,bytes32,uint64,int64),uint64),(uint64,uint64),((bytes,bytes32,uint64,int64)[],bool,(bytes,bytes32,uint64,int64),uint64)),uint128,uint256[8],uint256[2],uint256[2]) msg_) returns(uint8)
func (_ContractSP1ICS07Tendermint *ContractSP1ICS07TendermintTransactor) UpdateClientMsg(opts *bind.TransactOpts, msg_ IUpdateClientMsgsMsgUpdateClient) (*types.Transaction, error) {
	return _ContractSP1ICS07Tendermint.contract.Transact(opts, "updateClientMsg", msg_)
}

// UpdateClientMsg is a paid mutator transaction binding the contract method 0x04354dac.
//
// Solidity: function updateClientMsg(((string,(uint8,uint8),(uint64,uint64),uint32,uint32,bool,uint8),(uint128,bytes32,bytes32),((((uint64,uint64),string,uint64,uint128,bool,(bytes32,(uint32,bytes32)),bool,bytes32,bool,bytes32,bytes32,bytes32,bytes32,bytes32,bool,bytes32,bool,bytes32,bytes),(uint64,uint32,(bytes32,(uint32,bytes32)),(uint8,(bytes,uint128,bool,bytes))[])),((bytes,bytes32,uint64,int64)[],bool,(bytes,bytes32,uint64,int64),uint64),(uint64,uint64),((bytes,bytes32,uint64,int64)[],bool,(bytes,bytes32,uint64,int64),uint64)),uint128,uint256[8],uint256[2],uint256[2]) msg_) returns(uint8)
func (_ContractSP1ICS07Tendermint *ContractSP1ICS07TendermintSession) UpdateClientMsg(msg_ IUpdateClientMsgsMsgUpdateClient) (*types.Transaction, error) {
	return _ContractSP1ICS07Tendermint.Contract.UpdateClientMsg(&_ContractSP1ICS07Tendermint.TransactOpts, msg_)
}

// UpdateClientMsg is a paid mutator transaction binding the contract method 0x04354dac.
//
// Solidity: function updateClientMsg(((string,(uint8,uint8),(uint64,uint64),uint32,uint32,bool,uint8),(uint128,bytes32,bytes32),((((uint64,uint64),string,uint64,uint128,bool,(bytes32,(uint32,bytes32)),bool,bytes32,bool,bytes32,bytes32,bytes32,bytes32,bytes32,bool,bytes32,bool,bytes32,bytes),(uint64,uint32,(bytes32,(uint32,bytes32)),(uint8,(bytes,uint128,bool,bytes))[])),((bytes,bytes32,uint64,int64)[],bool,(bytes,bytes32,uint64,int64),uint64),(uint64,uint64),((bytes,bytes32,uint64,int64)[],bool,(bytes,bytes32,uint64,int64),uint64)),uint128,uint256[8],uint256[2],uint256[2]) msg_) returns(uint8)
func (_ContractSP1ICS07Tendermint *ContractSP1ICS07TendermintTransactorSession) UpdateClientMsg(msg_ IUpdateClientMsgsMsgUpdateClient) (*types.Transaction, error) {
	return _ContractSP1ICS07Tendermint.Contract.UpdateClientMsg(&_ContractSP1ICS07Tendermint.TransactOpts, msg_)
}

// VerifyMembership is a paid mutator transaction binding the contract method 0x0c6faf59.
//
// Solidity: function verifyMembership(((uint64,uint64),(bytes[],bytes32)[],((uint8,(bytes,bytes32,(uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),(bytes,bool,(bytes,bytes32,(uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),bool,(bytes,bytes32,(uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[])))[])[],bytes32,(uint128,bytes32,bytes32),uint8) msg_) returns(uint256)
func (_ContractSP1ICS07Tendermint *ContractSP1ICS07TendermintTransactor) VerifyMembership(opts *bind.TransactOpts, msg_ ILightClientMsgsMsgVerifyMembership) (*types.Transaction, error) {
	return _ContractSP1ICS07Tendermint.contract.Transact(opts, "verifyMembership", msg_)
}

// VerifyMembership is a paid mutator transaction binding the contract method 0x0c6faf59.
//
// Solidity: function verifyMembership(((uint64,uint64),(bytes[],bytes32)[],((uint8,(bytes,bytes32,(uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),(bytes,bool,(bytes,bytes32,(uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),bool,(bytes,bytes32,(uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[])))[])[],bytes32,(uint128,bytes32,bytes32),uint8) msg_) returns(uint256)
func (_ContractSP1ICS07Tendermint *ContractSP1ICS07TendermintSession) VerifyMembership(msg_ ILightClientMsgsMsgVerifyMembership) (*types.Transaction, error) {
	return _ContractSP1ICS07Tendermint.Contract.VerifyMembership(&_ContractSP1ICS07Tendermint.TransactOpts, msg_)
}

// VerifyMembership is a paid mutator transaction binding the contract method 0x0c6faf59.
//
// Solidity: function verifyMembership(((uint64,uint64),(bytes[],bytes32)[],((uint8,(bytes,bytes32,(uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),(bytes,bool,(bytes,bytes32,(uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),bool,(bytes,bytes32,(uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[])))[])[],bytes32,(uint128,bytes32,bytes32),uint8) msg_) returns(uint256)
func (_ContractSP1ICS07Tendermint *ContractSP1ICS07TendermintTransactorSession) VerifyMembership(msg_ ILightClientMsgsMsgVerifyMembership) (*types.Transaction, error) {
	return _ContractSP1ICS07Tendermint.Contract.VerifyMembership(&_ContractSP1ICS07Tendermint.TransactOpts, msg_)
}

// VerifyNonMembership is a paid mutator transaction binding the contract method 0x038d5cb3.
//
// Solidity: function verifyNonMembership(((uint64,uint64),(bytes[],bytes32)[],((uint8,(bytes,bytes32,(uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),(bytes,bool,(bytes,bytes32,(uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),bool,(bytes,bytes32,(uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[])))[])[],bytes32,(uint128,bytes32,bytes32),uint8) msg_) returns(uint256)
func (_ContractSP1ICS07Tendermint *ContractSP1ICS07TendermintTransactor) VerifyNonMembership(opts *bind.TransactOpts, msg_ ILightClientMsgsMsgVerifyNonMembership) (*types.Transaction, error) {
	return _ContractSP1ICS07Tendermint.contract.Transact(opts, "verifyNonMembership", msg_)
}

// VerifyNonMembership is a paid mutator transaction binding the contract method 0x038d5cb3.
//
// Solidity: function verifyNonMembership(((uint64,uint64),(bytes[],bytes32)[],((uint8,(bytes,bytes32,(uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),(bytes,bool,(bytes,bytes32,(uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),bool,(bytes,bytes32,(uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[])))[])[],bytes32,(uint128,bytes32,bytes32),uint8) msg_) returns(uint256)
func (_ContractSP1ICS07Tendermint *ContractSP1ICS07TendermintSession) VerifyNonMembership(msg_ ILightClientMsgsMsgVerifyNonMembership) (*types.Transaction, error) {
	return _ContractSP1ICS07Tendermint.Contract.VerifyNonMembership(&_ContractSP1ICS07Tendermint.TransactOpts, msg_)
}

// VerifyNonMembership is a paid mutator transaction binding the contract method 0x038d5cb3.
//
// Solidity: function verifyNonMembership(((uint64,uint64),(bytes[],bytes32)[],((uint8,(bytes,bytes32,(uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),(bytes,bool,(bytes,bytes32,(uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),bool,(bytes,bytes32,(uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[])))[])[],bytes32,(uint128,bytes32,bytes32),uint8) msg_) returns(uint256)
func (_ContractSP1ICS07Tendermint *ContractSP1ICS07TendermintTransactorSession) VerifyNonMembership(msg_ ILightClientMsgsMsgVerifyNonMembership) (*types.Transaction, error) {
	return _ContractSP1ICS07Tendermint.Contract.VerifyNonMembership(&_ContractSP1ICS07Tendermint.TransactOpts, msg_)
}

// ContractSP1ICS07TendermintRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the ContractSP1ICS07Tendermint contract.
type ContractSP1ICS07TendermintRoleAdminChangedIterator struct {
	Event *ContractSP1ICS07TendermintRoleAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractSP1ICS07TendermintRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSP1ICS07TendermintRoleAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractSP1ICS07TendermintRoleAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractSP1ICS07TendermintRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSP1ICS07TendermintRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSP1ICS07TendermintRoleAdminChanged represents a RoleAdminChanged event raised by the ContractSP1ICS07Tendermint contract.
type ContractSP1ICS07TendermintRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ContractSP1ICS07Tendermint *ContractSP1ICS07TendermintFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*ContractSP1ICS07TendermintRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _ContractSP1ICS07Tendermint.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &ContractSP1ICS07TendermintRoleAdminChangedIterator{contract: _ContractSP1ICS07Tendermint.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ContractSP1ICS07Tendermint *ContractSP1ICS07TendermintFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *ContractSP1ICS07TendermintRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _ContractSP1ICS07Tendermint.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSP1ICS07TendermintRoleAdminChanged)
				if err := _ContractSP1ICS07Tendermint.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ContractSP1ICS07Tendermint *ContractSP1ICS07TendermintFilterer) ParseRoleAdminChanged(log types.Log) (*ContractSP1ICS07TendermintRoleAdminChanged, error) {
	event := new(ContractSP1ICS07TendermintRoleAdminChanged)
	if err := _ContractSP1ICS07Tendermint.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSP1ICS07TendermintRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the ContractSP1ICS07Tendermint contract.
type ContractSP1ICS07TendermintRoleGrantedIterator struct {
	Event *ContractSP1ICS07TendermintRoleGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractSP1ICS07TendermintRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSP1ICS07TendermintRoleGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractSP1ICS07TendermintRoleGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractSP1ICS07TendermintRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSP1ICS07TendermintRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSP1ICS07TendermintRoleGranted represents a RoleGranted event raised by the ContractSP1ICS07Tendermint contract.
type ContractSP1ICS07TendermintRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ContractSP1ICS07Tendermint *ContractSP1ICS07TendermintFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ContractSP1ICS07TendermintRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ContractSP1ICS07Tendermint.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ContractSP1ICS07TendermintRoleGrantedIterator{contract: _ContractSP1ICS07Tendermint.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ContractSP1ICS07Tendermint *ContractSP1ICS07TendermintFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *ContractSP1ICS07TendermintRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ContractSP1ICS07Tendermint.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSP1ICS07TendermintRoleGranted)
				if err := _ContractSP1ICS07Tendermint.contract.UnpackLog(event, "RoleGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ContractSP1ICS07Tendermint *ContractSP1ICS07TendermintFilterer) ParseRoleGranted(log types.Log) (*ContractSP1ICS07TendermintRoleGranted, error) {
	event := new(ContractSP1ICS07TendermintRoleGranted)
	if err := _ContractSP1ICS07Tendermint.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSP1ICS07TendermintRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the ContractSP1ICS07Tendermint contract.
type ContractSP1ICS07TendermintRoleRevokedIterator struct {
	Event *ContractSP1ICS07TendermintRoleRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractSP1ICS07TendermintRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSP1ICS07TendermintRoleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractSP1ICS07TendermintRoleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractSP1ICS07TendermintRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSP1ICS07TendermintRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSP1ICS07TendermintRoleRevoked represents a RoleRevoked event raised by the ContractSP1ICS07Tendermint contract.
type ContractSP1ICS07TendermintRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ContractSP1ICS07Tendermint *ContractSP1ICS07TendermintFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ContractSP1ICS07TendermintRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ContractSP1ICS07Tendermint.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ContractSP1ICS07TendermintRoleRevokedIterator{contract: _ContractSP1ICS07Tendermint.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ContractSP1ICS07Tendermint *ContractSP1ICS07TendermintFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *ContractSP1ICS07TendermintRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ContractSP1ICS07Tendermint.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSP1ICS07TendermintRoleRevoked)
				if err := _ContractSP1ICS07Tendermint.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ContractSP1ICS07Tendermint *ContractSP1ICS07TendermintFilterer) ParseRoleRevoked(log types.Log) (*ContractSP1ICS07TendermintRoleRevoked, error) {
	event := new(ContractSP1ICS07TendermintRoleRevoked)
	if err := _ContractSP1ICS07Tendermint.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
