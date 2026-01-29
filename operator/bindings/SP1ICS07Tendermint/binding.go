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

// IICS07TendermintMsgsConsensusState is an auto generated low-level Go binding around an user-defined struct.
type IICS07TendermintMsgsConsensusState struct {
	Timestamp          *big.Int
	Root               [32]byte
	NextValidatorsHash [32]byte
}

// IICS07TendermintMsgsTrustThreshold is an auto generated low-level Go binding around an user-defined struct.
type IICS07TendermintMsgsTrustThreshold struct {
	Numerator   uint8
	Denominator uint8
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

// ContractSP1ICS07TendermintMetaData contains all meta data concerning the ContractSP1ICS07Tendermint contract.
var ContractSP1ICS07TendermintMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"verifier\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"membership_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"misbehaviour_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"updateClient_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_clientState\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_consensusState\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"roleManager\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ALLOWED_SP1_CLOCK_DRIFT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MEMBERSHIP\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIMembership\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MISBEHAVIOUR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIMisbehaviour\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PROOF_SUBMITTER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UPDATE_CLIENT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIUpdateClient\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"VERIFIER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIVerifier\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"clientState\",\"inputs\":[],\"outputs\":[{\"name\":\"chainId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"trustLevel\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.TrustThreshold\",\"components\":[{\"name\":\"numerator\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"denominator\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"name\":\"latestHeight\",\"type\":\"tuple\",\"internalType\":\"structIICS02ClientMsgs.Height\",\"components\":[{\"name\":\"revisionNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revisionHeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"trustingPeriod\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"unbondingPeriod\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isFrozen\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"zkAlgorithm\",\"type\":\"uint8\",\"internalType\":\"enumIICS07TendermintMsgs.SupportedZkAlgorithm\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getClientState\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getConsensusStateHash\",\"inputs\":[{\"name\":\"revisionHeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"misbehaviour\",\"inputs\":[{\"name\":\"misbehaviourMsg\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"multicall\",\"inputs\":[{\"name\":\"data\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"outputs\":[{\"name\":\"results\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateClient\",\"inputs\":[{\"name\":\"updateClientMsg\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumILightClientMsgs.UpdateResult\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeClient\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyMembership\",\"inputs\":[{\"name\":\"msg_\",\"type\":\"tuple\",\"internalType\":\"structILightClientMsgs.MsgVerifyMembership\",\"components\":[{\"name\":\"height\",\"type\":\"tuple\",\"internalType\":\"structIICS02ClientMsgs.Height\",\"components\":[{\"name\":\"revisionNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revisionHeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"kvPairs\",\"type\":\"tuple[]\",\"internalType\":\"structIMembershipMsgs.KVPair[]\",\"components\":[{\"name\":\"path\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"value\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"merkleProofs\",\"type\":\"tuple[]\",\"internalType\":\"structIMembershipMsgs.MerkleProof[]\",\"components\":[{\"name\":\"proofs\",\"type\":\"tuple[]\",\"internalType\":\"structIMembershipMsgs.CommitmentProof[]\",\"components\":[{\"name\":\"proofType\",\"type\":\"uint8\",\"internalType\":\"enumIMembershipMsgs.ProofType\"},{\"name\":\"existenceProof\",\"type\":\"tuple\",\"internalType\":\"structIMembershipMsgs.ExistenceProof\",\"components\":[{\"name\":\"key\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"leaf\",\"type\":\"tuple\",\"internalType\":\"structIMembershipMsgs.LeafOp\",\"components\":[{\"name\":\"hashOp\",\"type\":\"uint8\",\"internalType\":\"enumIMembershipMsgs.HashOp\"},{\"name\":\"prehashKey\",\"type\":\"uint8\",\"internalType\":\"enumIMembershipMsgs.HashOp\"},{\"name\":\"prehashValue\",\"type\":\"uint8\",\"internalType\":\"enumIMembershipMsgs.HashOp\"},{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"path\",\"type\":\"tuple[]\",\"internalType\":\"structIMembershipMsgs.InnerOp[]\",\"components\":[{\"name\":\"hashOp\",\"type\":\"uint8\",\"internalType\":\"enumIMembershipMsgs.HashOp\"},{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"suffix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]},{\"name\":\"nonExistenceProof\",\"type\":\"tuple\",\"internalType\":\"structIMembershipMsgs.NonExistenceProof\",\"components\":[{\"name\":\"key\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"hasLeft\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"left\",\"type\":\"tuple\",\"internalType\":\"structIMembershipMsgs.ExistenceProof\",\"components\":[{\"name\":\"key\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"leaf\",\"type\":\"tuple\",\"internalType\":\"structIMembershipMsgs.LeafOp\",\"components\":[{\"name\":\"hashOp\",\"type\":\"uint8\",\"internalType\":\"enumIMembershipMsgs.HashOp\"},{\"name\":\"prehashKey\",\"type\":\"uint8\",\"internalType\":\"enumIMembershipMsgs.HashOp\"},{\"name\":\"prehashValue\",\"type\":\"uint8\",\"internalType\":\"enumIMembershipMsgs.HashOp\"},{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"path\",\"type\":\"tuple[]\",\"internalType\":\"structIMembershipMsgs.InnerOp[]\",\"components\":[{\"name\":\"hashOp\",\"type\":\"uint8\",\"internalType\":\"enumIMembershipMsgs.HashOp\"},{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"suffix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]},{\"name\":\"hasRight\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"right\",\"type\":\"tuple\",\"internalType\":\"structIMembershipMsgs.ExistenceProof\",\"components\":[{\"name\":\"key\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"leaf\",\"type\":\"tuple\",\"internalType\":\"structIMembershipMsgs.LeafOp\",\"components\":[{\"name\":\"hashOp\",\"type\":\"uint8\",\"internalType\":\"enumIMembershipMsgs.HashOp\"},{\"name\":\"prehashKey\",\"type\":\"uint8\",\"internalType\":\"enumIMembershipMsgs.HashOp\"},{\"name\":\"prehashValue\",\"type\":\"uint8\",\"internalType\":\"enumIMembershipMsgs.HashOp\"},{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"path\",\"type\":\"tuple[]\",\"internalType\":\"structIMembershipMsgs.InnerOp[]\",\"components\":[{\"name\":\"hashOp\",\"type\":\"uint8\",\"internalType\":\"enumIMembershipMsgs.HashOp\"},{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"suffix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]}]}]}]},{\"name\":\"appHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"trustedConsensusState\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.ConsensusState\",\"components\":[{\"name\":\"timestamp\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"nextValidatorsHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"membershipType\",\"type\":\"uint8\",\"internalType\":\"enumIMembershipMsgs.MembershipType\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifyNonMembership\",\"inputs\":[{\"name\":\"msg_\",\"type\":\"tuple\",\"internalType\":\"structILightClientMsgs.MsgVerifyNonMembership\",\"components\":[{\"name\":\"height\",\"type\":\"tuple\",\"internalType\":\"structIICS02ClientMsgs.Height\",\"components\":[{\"name\":\"revisionNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revisionHeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"kvPairs\",\"type\":\"tuple[]\",\"internalType\":\"structIMembershipMsgs.KVPair[]\",\"components\":[{\"name\":\"path\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"value\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"merkleProofs\",\"type\":\"tuple[]\",\"internalType\":\"structIMembershipMsgs.MerkleProof[]\",\"components\":[{\"name\":\"proofs\",\"type\":\"tuple[]\",\"internalType\":\"structIMembershipMsgs.CommitmentProof[]\",\"components\":[{\"name\":\"proofType\",\"type\":\"uint8\",\"internalType\":\"enumIMembershipMsgs.ProofType\"},{\"name\":\"existenceProof\",\"type\":\"tuple\",\"internalType\":\"structIMembershipMsgs.ExistenceProof\",\"components\":[{\"name\":\"key\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"leaf\",\"type\":\"tuple\",\"internalType\":\"structIMembershipMsgs.LeafOp\",\"components\":[{\"name\":\"hashOp\",\"type\":\"uint8\",\"internalType\":\"enumIMembershipMsgs.HashOp\"},{\"name\":\"prehashKey\",\"type\":\"uint8\",\"internalType\":\"enumIMembershipMsgs.HashOp\"},{\"name\":\"prehashValue\",\"type\":\"uint8\",\"internalType\":\"enumIMembershipMsgs.HashOp\"},{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"path\",\"type\":\"tuple[]\",\"internalType\":\"structIMembershipMsgs.InnerOp[]\",\"components\":[{\"name\":\"hashOp\",\"type\":\"uint8\",\"internalType\":\"enumIMembershipMsgs.HashOp\"},{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"suffix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]},{\"name\":\"nonExistenceProof\",\"type\":\"tuple\",\"internalType\":\"structIMembershipMsgs.NonExistenceProof\",\"components\":[{\"name\":\"key\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"hasLeft\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"left\",\"type\":\"tuple\",\"internalType\":\"structIMembershipMsgs.ExistenceProof\",\"components\":[{\"name\":\"key\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"leaf\",\"type\":\"tuple\",\"internalType\":\"structIMembershipMsgs.LeafOp\",\"components\":[{\"name\":\"hashOp\",\"type\":\"uint8\",\"internalType\":\"enumIMembershipMsgs.HashOp\"},{\"name\":\"prehashKey\",\"type\":\"uint8\",\"internalType\":\"enumIMembershipMsgs.HashOp\"},{\"name\":\"prehashValue\",\"type\":\"uint8\",\"internalType\":\"enumIMembershipMsgs.HashOp\"},{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"path\",\"type\":\"tuple[]\",\"internalType\":\"structIMembershipMsgs.InnerOp[]\",\"components\":[{\"name\":\"hashOp\",\"type\":\"uint8\",\"internalType\":\"enumIMembershipMsgs.HashOp\"},{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"suffix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]},{\"name\":\"hasRight\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"right\",\"type\":\"tuple\",\"internalType\":\"structIMembershipMsgs.ExistenceProof\",\"components\":[{\"name\":\"key\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"leaf\",\"type\":\"tuple\",\"internalType\":\"structIMembershipMsgs.LeafOp\",\"components\":[{\"name\":\"hashOp\",\"type\":\"uint8\",\"internalType\":\"enumIMembershipMsgs.HashOp\"},{\"name\":\"prehashKey\",\"type\":\"uint8\",\"internalType\":\"enumIMembershipMsgs.HashOp\"},{\"name\":\"prehashValue\",\"type\":\"uint8\",\"internalType\":\"enumIMembershipMsgs.HashOp\"},{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"path\",\"type\":\"tuple[]\",\"internalType\":\"structIMembershipMsgs.InnerOp[]\",\"components\":[{\"name\":\"hashOp\",\"type\":\"uint8\",\"internalType\":\"enumIMembershipMsgs.HashOp\"},{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"suffix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]}]}]}]},{\"name\":\"appHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"trustedConsensusState\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.ConsensusState\",\"components\":[{\"name\":\"timestamp\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"nextValidatorsHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"membershipType\",\"type\":\"uint8\",\"internalType\":\"enumIMembershipMsgs.MembershipType\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotHandleMisbehavior\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChainIdMismatch\",\"inputs\":[{\"name\":\"expected\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"actual\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"ClientStateMismatch\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"actual\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"ConsensusStateHashMismatch\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"actual\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ConsensusStateNotFound\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ConsensusStateRootMismatch\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"actual\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"EmptyValue\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedToVerifyHeader\",\"inputs\":[{\"name\":\"description\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"FeatureNotSupported\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FrozenClientState\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientMisbehaviourHeaderHeight\",\"inputs\":[{\"name\":\"height1\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"height2\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InsufficientTrustingPeriod\",\"inputs\":[{\"name\":\"durationSinceConsensusState\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"trustingPeriod\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"type\":\"error\",\"name\":\"InvalidConsensusStateTimestamp\",\"inputs\":[{\"name\":\"timestamp\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"type\":\"error\",\"name\":\"InvalidHeaderHeight\",\"inputs\":[{\"name\":\"height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidMembershipProof\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"KeyValuePairNotInCache\",\"inputs\":[{\"name\":\"path\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"value\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"LengthIsOutOfRange\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"min\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"max\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MembershipProofKeyNotFound\",\"inputs\":[{\"name\":\"path\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}]},{\"type\":\"error\",\"name\":\"MembershipProofValueMismatch\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"actual\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"MismatchedRevisionHeights\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"actual\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"MismatchedValidatorHashes\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"actual\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ProofHeightMismatch\",\"inputs\":[{\"name\":\"expectedRevisionNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"expectedRevisionHeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"actualRevisionNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"actualRevisionHeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ProofIsInTheFuture\",\"inputs\":[{\"name\":\"now\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proofTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ProofIsTooOld\",\"inputs\":[{\"name\":\"now\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proofTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"TrustThresholdMismatch\",\"inputs\":[{\"name\":\"expectedNumerator\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expectedDenominator\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"actualNumerator\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"actualDenominator\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"TrustingPeriodMismatch\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"actual\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"TrustingPeriodTooLong\",\"inputs\":[{\"name\":\"trustingPeriod\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"unbondingPeriod\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"UnbondingPeriodMismatch\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"actual\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"UnknownMembershipType\",\"inputs\":[{\"name\":\"membershipType\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"UnknownZkAlgorithm\",\"inputs\":[{\"name\":\"algorithm\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"VerificationKeyMismatch\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"actual\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]",
	Bin: "0x6101006040523461054b5761412c8038038061001a8161056e565b928339810160e08282031261054b5761003282610593565b61003e60208401610593565b61004a60408501610593565b9161005760608601610593565b60808601519094906001600160401b03811161054b5786019080601f8301121561054b578151610089926020016105a7565b9461009b60c060a08301519201610593565b958051810190602082019060208184031261054b576020810151906001600160401b03821161054b570191829003601f198101906101201361054b576040519160e083016001600160401b038111848210176105375760405260208401516001600160401b03811161054b576020908501019080601f8301121561054b578151610127926020016105a7565b82526040811261054b57604061013b61054f565b916101478286016105e7565b8352610155606086016105e7565b602084015260208401928352603f19011261054b5761017261054f565b61017e608085016105f5565b815261018c60a085016105f5565b6020820152604083019081526101a460c08501610609565b90606084019182526101b860e08601610609565b926080850193845261010086015195861515870361054b5760a08601968752610120015194600286101561054b5760c08101958652518051906001600160401b03821161053757600154600181811c9116801561052d575b602082101461051957601f81116104b6575b50602090601f83116001146104495763ffffffff95949392915f918361043e575b50508160011b915f199060031b1c1916176001555b5160ff81511661ff00602060025493015160081b169161ffff191617176002555160018060401b0381511660035491602068010000000000000000600160801b0391015160401b169160018060801b031916171760035551169267ffffffff00000000600454925160201b169051151560401b925191600283101561042a5769ff00000000000000000068ff00000000000000009360481b169469ff000000000000000000199160018060481b0319161716179116171760045560018060401b0360035460401c165f52600560205260405f205560018060a01b031660805260018060a01b031660a05260018060a01b031660c05260018060a01b031660e05260045463ffffffff8116610708810163ffffffff81116104165763ffffffff809360201c16928391161161040157826001600160a01b0381166103e8575061039e610710565b505b6040516138f99081610793823960805181610f77015260a051818181610c3b015261298e015260c0518181816105170152610fc7015260e051818181610c8b015261232b0152f35b806103f56103fb9261061a565b50610690565b506103a0565b6333fae18560e21b5f5260045260245260445ffd5b634e487b7160e01b5f52601160045260245ffd5b634e487b7160e01b5f52602160045260245ffd5b015190505f80610243565b90601f1983169160015f52815f20925f5b81811061049e575091600193918563ffffffff999897969410610486575b505050811b01600155610258565b01515f1960f88460031b161c191690555f8080610478565b9293602060018192878601518155019501930161045a565b60015f527fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf6601f840160051c8101916020851061050f575b601f0160051c01905b8181106105045750610222565b5f81556001016104f7565b90915081906104ee565b634e487b7160e01b5f52602260045260245ffd5b90607f1690610210565b634e487b7160e01b5f52604160045260245ffd5b5f80fd5b60408051919082016001600160401b0381118382101761053757604052565b6040519190601f01601f191682016001600160401b0381118382101761053757604052565b51906001600160a01b038216820361054b57565b9192916001600160401b038211610537576105cb601f8301601f191660200161056e565b938285528282011161054b57815f926020928387015e84010152565b519060ff8216820361054b57565b51906001600160401b038216820361054b57565b519063ffffffff8216820361054b57565b6001600160a01b0381165f9081525f51602061410c5f395f51905f52602052604090205460ff1661068b576001600160a01b03165f8181525f51602061410c5f395f51905f5260205260408120805460ff191660011790553391905f51602061408c5f395f51905f528180a4600190565b505f90565b6001600160a01b0381165f9081525f5160206140ac5f395f51905f52602052604090205460ff1661068b576001600160a01b03165f8181525f5160206140ac5f395f51905f5260205260408120805460ff191660011790553391905f5160206140ec5f395f51905f52905f51602061408c5f395f51905f529080a4600190565b5f80525f5160206140ac5f395f51905f526020525f5160206140cc5f395f51905f525460ff1661078e575f8080525f5160206140ac5f395f51905f526020525f5160206140cc5f395f51905f52805460ff1916600117905533905f5160206140ec5f395f51905f525f51602061408c5f395f51905f528280a4600190565b5f9056fe6080806040526004361015610012575f80fd5b5f3560e01c90816301ffc9a714610feb5750806302cf295214610f9b578063038d5cb314610df257806308c84e7014610f4b5780630bece35614610eb15780630c6faf5914610df257806323842fb814610dc2578063248a9ca314610d905780632c3ee47414610d745780632f2ff15d14610d4557806336568abe14610ce95780635972185a14610caf57806387d4332f14610c5f57806389df51f114610c0f5780638a8e4c5d14610b6b57806391d1485414610b22578063a217fddf14610b08578063ac9650d8146108d8578063bd3ce6b0146107ce578063d547741f14610798578063ddba6537146101e25763ef913a4b1461010e575f80fd5b346101de575f6003193601126101de576101da60405160208082015261012060408201526101c6816101436101608201611203565b60ff600254818116606085015260081c16608083015267ffffffffffffffff60035481811660a085015260401c1660c083015260ff60045463ffffffff811660e085015263ffffffff8160201c16610100850152818160401c16151561012085015260481c166101b281611353565b61014083015203601f198101835282611330565b60405191829160208352602083019061118d565b0390f35b5f80fd5b346101de576101f0366110bd565b6004549060ff8260401c16610770575f80527f4c48d1d2b0f4f7485fc28e3db22341d96a20aa29e6efa8149da9751603abd4e06020527f6e0c24a6e293ff9b755263dbaa15ba3796b0b8d3fe17cfb4ddf8143b268eac475460ff1615610763575b8201916020818403126101de5780359067ffffffffffffffff82116101de5701610120818403126101de5760405160a0810181811067ffffffffffffffff82111761073657604052813567ffffffffffffffff81116101de57846102b6918401611473565b8152602082013567ffffffffffffffff81116101de578201916060838603126101de57604051926102e6846112dc565b803567ffffffffffffffff81116101de5781016040818803126101de5760405190610310826112c0565b803567ffffffffffffffff81116101de57816103338a60209361033b95016113cd565b845201611145565b60208201528452602081013567ffffffffffffffff81116101de578661036291830161175f565b602085015260408101359067ffffffffffffffff82116101de576103889187910161175f565b6040840152602082019283526103b56103a48660408401611562565b956040840196875260a08301611562565b9160806104376103cf610100606085019587875201611545565b95828401958787526fffffffffffffffffffffffffffffffff85519251986104f78c51936104c96104996040519d8e998a997fa6fe8f56000000000000000000000000000000000000000000000000000000008b5261012060048c01526101248b0190611c70565b6003198a82030160248b0152604061048883516060845267ffffffffffffffff602061046e835186606089015260a088019061118d565b920151168f85015260208501518482036020860152611dfe565b920151906040818403910152611dfe565b86516fffffffffffffffffffffffffffffffff166044890152602087015160648901526040909601516084880152565b80516fffffffffffffffffffffffffffffffff1660a4870152602081015160c48701526040015160e4860152565b16610104830152038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa93841561072b575f94610669575b5067ffffffffffffffff6020806106629561060f7fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff9998966105b0680100000000000000009c6fffffffffffffffffffffffffffffffff610607995191519351955116906135d1565b6040516105e68582018093604080916fffffffffffffffffffffffffffffffff8151168452602081015160208501520151910152565b606081526105f5608082611330565b51902061060786858a510151166126f9565b808214613090565b6040516106458382018093604080916fffffffffffffffffffffffffffffffff8151168452602081015160208501520151910152565b60608152610654608082611330565b5190209401510151166126f9565b1617600455005b959493509060803d608011610724575b6106838188611330565b8601906080878303126101de576020806106629561060f67ffffffffffffffff946105b07fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff9b6fffffffffffffffffffffffffffffffff680100000000000000009e61070c6106079b60408051936106fa856112c0565b6107048382611bd8565b855201611bd8565b888201529c9d50509c50509695505095505050610547565b503d610679565b6040513d5f823e3d90fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b61076b612781565b610251565b7f928b1233000000000000000000000000000000000000000000000000000000005f5260045ffd5b346101de576107cc6107a93661115a565b906107c76107c2825f525f602052600160405f20015490565b612809565b61325a565b005b346101de575f6003193601126101de576108656040516107f8816107f181611203565b0382611330565b604051610804816112c0565b60ff600254818116835260081c166020820152604051610823816112c0565b67ffffffffffffffff600354818116835260401c16602082015260ff60045461089f828260481c169361087f6040519889986101208a526101208a019061118d565b96602089019060ff60208092828151168552015116910152565b606087019067ffffffffffffffff60208092828151168552015116910152565b63ffffffff811660a086015263ffffffff8160201c1660c086015260401c16151560e08401526108ce81611353565b6101008301520390f35b346101de5760206003193601126101de5760043567ffffffffffffffff81116101de57366023820112156101de5780600401359067ffffffffffffffff82116101de573660248360051b830101116101de579060206040519061093b8183611330565b5f825280820193601f198201368637610953846115ed565b906109616040519283611330565b848252601f19610970866115ed565b015f5b818110610af95750505f907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffbd81360301915b86811015610a7a5760248160051b83010135838112156101de5782019060248201359167ffffffffffffffff83116101de576044019180360383136101de575f806001948a610a23610a56958f8d906040519483869484860198893784019083820190898252519283915e010185815203601f198101835282611330565b5190305af43d15610a72573d90610a39826113b1565b91610a476040519384611330565b82523d5f8a84013e5b30613853565b610a608287612740565b52610a6b8186612740565b50016109a5565b606090610a50565b5050506040519082820192808352815180945260408301938160408260051b8601019301915f955b828710610aaf5785850386f35b909192938280610ae9837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08a60019603018652885161118d565b9601920196019592919092610aa2565b60608482018601528401610973565b346101de575f6003193601126101de5760206040515f8152f35b346101de57610b303661115a565b905f525f60205273ffffffffffffffffffffffffffffffffffffffff60405f2091165f52602052602060ff60405f2054166040519015158152f35b346101de57610b79366110bd565b505060ff60045460401c16610770575f80527f4c48d1d2b0f4f7485fc28e3db22341d96a20aa29e6efa8149da9751603abd4e06020527f6e0c24a6e293ff9b755263dbaa15ba3796b0b8d3fe17cfb4ddf8143b268eac475460ff1615610c02575b7fda81d7c2000000000000000000000000000000000000000000000000000000005f5260045ffd5b610c0a612781565b610bda565b346101de575f6003193601126101de57602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346101de575f6003193601126101de57602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346101de575f6003193601126101de5760206040517fbd893629a699470e4ec82a5715bb4981fdaacc5d0a728bf5f55b801d8f4ef10b8152f35b346101de57610cf73661115a565b3373ffffffffffffffffffffffffffffffffffffffff821603610d1d576107cc9161325a565b7f6697b232000000000000000000000000000000000000000000000000000000005f5260045ffd5b346101de576107cc610d563661115a565b90610d6f6107c2825f525f602052600160405f20015490565b613188565b346101de575f6003193601126101de5760206040516107088152f35b346101de5760206003193601126101de576020610dba6004355f525f602052600160405f20015490565b604051908152f35b346101de5760206003193601126101de5760043567ffffffffffffffff811681036101de57610dba6020916126f9565b346101de57610e0036611089565b60ff60045460401c16610770575f80527f4c48d1d2b0f4f7485fc28e3db22341d96a20aa29e6efa8149da9751603abd4e06020527f6e0c24a6e293ff9b755263dbaa15ba3796b0b8d3fe17cfb4ddf8143b268eac475460ff1615610ea4575b610e6c604082018261135d565b90610e7a606084018461135d565b9190936101008101359260028410156101de57602095610dba9560a084019460808501359461286f565b610eac612781565b610e5f565b346101de57610ebf366110bd565b9060ff60045460401c16610770575f80527f4c48d1d2b0f4f7485fc28e3db22341d96a20aa29e6efa8149da9751603abd4e060209081527f6e0c24a6e293ff9b755263dbaa15ba3796b0b8d3fe17cfb4ddf8143b268eac47549092610f2d92909160ff1615610f3e5761215e565b60405190610f3a8161110e565b8152f35b610f46612781565b61215e565b346101de575f6003193601126101de57602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346101de575f6003193601126101de57602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346101de5760206003193601126101de57600435907fffffffff0000000000000000000000000000000000000000000000000000000082168092036101de57817f7965db0b000000000000000000000000000000000000000000000000000000006020931490811561105f575b5015158152f35b7f01ffc9a70000000000000000000000000000000000000000000000000000000091501483611058565b60206003198201126101de576004359067ffffffffffffffff82116101de5760031982610120920301126101de5760040190565b9060206003198301126101de5760043567ffffffffffffffff81116101de57826023820112156101de5780600401359267ffffffffffffffff84116101de57602484830101116101de576024019190565b6003111561111857565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b359067ffffffffffffffff821682036101de57565b60031960409101126101de576004359060243573ffffffffffffffffffffffffffffffffffffffff811681036101de5790565b90601f19601f602080948051918291828752018686015e5f8582860101520116010190565b90600182811c921680156111f9575b60208310146111cc57565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b91607f16916111c1565b6001545f9291611212826111b2565b8082529160018116908115611286575060011461122d575050565b60015f9081529293509091907fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf65b83831061126c575060209250010190565b60018160209294939454838587010152019101919061125b565b60209495507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091509291921683830152151560051b010190565b6040810190811067ffffffffffffffff82111761073657604052565b6060810190811067ffffffffffffffff82111761073657604052565b60e0810190811067ffffffffffffffff82111761073657604052565b6080810190811067ffffffffffffffff82111761073657604052565b90601f601f19910116810190811067ffffffffffffffff82111761073657604052565b6002111561111857565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156101de570180359067ffffffffffffffff82116101de57602001918160051b360383136101de57565b67ffffffffffffffff811161073657601f01601f191660200190565b81601f820112156101de576020813591016113e7826113b1565b926113f56040519485611330565b828452828201116101de57815f92602092838601378301015290565b359060ff821682036101de57565b91908260409103126101de57604051611437816112c0565b602061145081839561144881611145565b855201611145565b910152565b359063ffffffff821682036101de57565b359081151582036101de57565b808203929161012084126101de576040519161148e836112f8565b8294813567ffffffffffffffff81116101de576040916114b385601f199386016113cd565b865201126101de576114fb610100926040516114ce816112c0565b6114da60208501611411565b81526114e860408501611411565b602082015260208601526060830161141f565b604084015261150c60a08201611455565b606084015261151d60c08201611455565b608084015261152e60e08201611466565b60a084015201359060028210156101de5760c00152565b35906fffffffffffffffffffffffffffffffff821682036101de57565b91908260609103126101de5760405161157a816112dc565b604080829461158881611545565b8452602081013560208501520135910152565b8092910391606083126101de576040516115b4816112c0565b6040601f1982958435845201126101de5760209060408051936115d6856112c0565b6115e1848201611455565b85520135828401520152565b67ffffffffffffffff81116107365760051b60200190565b91906080838203126101de576040519061161e82611314565b8193803567ffffffffffffffff81116101de5760609261163f9183016113cd565b83526020810135602084015261165760408201611145565b60408401520135908160070b82036101de5760600152565b9190916080818403126101de576040519061168982611314565b8193813567ffffffffffffffff81116101de57820181601f820112156101de5780356116b4816115ed565b916116c26040519384611330565b81835260208084019260051b820101918483116101de5760208201905b838210611731575050505083526116f860208301611466565b602084015260408201359067ffffffffffffffff82116101de57826117266060949261145094869401611605565b604086015201611145565b813567ffffffffffffffff81116101de5760209161175488848094880101611605565b8152019101906116df565b919060a0838203126101de576040519061177882611314565b8193803567ffffffffffffffff81116101de5781016040818403126101de57604051906117a4826112c0565b803567ffffffffffffffff81116101de5781016102c0818603126101de5760405190610260820182811067ffffffffffffffff821117610736576040526117eb868261141f565b8252604081013567ffffffffffffffff81116101de578661180d9183016113cd565b602083015261181e60608201611145565b604083015261182f60808201611545565b606083015261184060a08201611466565b60808301526118528660c0830161159b565b60a08301526118646101208201611466565b60c083015261014081013560e08301526118816101608201611466565b6101008301526101808101356101208301526101a08101356101408301526101c08101356101608301526101e08101356101808301526102008101356101a08301526118d06102208201611466565b6101c08301526102408101356101e08301526118ef6102608201611466565b6102008301526102808101356102208301526102a08101359067ffffffffffffffff82116101de57611923918791016113cd565b610240820152825260208101359067ffffffffffffffff82116101de570160c0818503126101de576040519061195882611314565b61196181611145565b825261196f60208201611455565b6020830152611981856040830161159b565b604083015260a08101359067ffffffffffffffff82116101de570184601f820112156101de578035906119b3826115ed565b916119c16040519384611330565b80835260208084019160051b830101918783116101de5760208101915b838310611a4e5750505050606082015260208201528352602081013567ffffffffffffffff81116101de5782611a1591830161166f565b6020840152611a27826040830161141f565b604084015260808101359167ffffffffffffffff83116101de57606092611450920161166f565b823567ffffffffffffffff81116101de578201906040601f19838c0301126101de5760405191611a7d836112c0565b602081013560038110156101de578352604081013567ffffffffffffffff81116101de576020910101906080828c03126101de5760405192611abe84611314565b823567ffffffffffffffff81116101de578c611adb9185016113cd565b8452611ae960208401611545565b6020850152611afa60408401611466565b604085015260608301359367ffffffffffffffff85116101de57611b238d6020968796016113cd565b6060820152838201528152019201916119de565b9080601f830112156101de57604080519290611b539084611330565b8290604081019283116101de57905b828210611b6f5750505090565b8135815260209182019101611b62565b929192611b8b826113b1565b91611b996040519384611330565b8294818452818301116101de578281602093845f96015e010152565b519060ff821682036101de57565b519067ffffffffffffffff821682036101de57565b91908260409103126101de57604051611bf0816112c0565b6020611450818395611c0181611bc3565b855201611bc3565b519063ffffffff821682036101de57565b51906fffffffffffffffffffffffffffffffff821682036101de57565b91908260609103126101de57604051611c4f816112dc565b6040808294611c5d81611c1a565b8452602081015160208501520151910152565b9061010060c0611c8b8451610120855261012085019061118d565b602080860151805160ff9081168784015291015116604085015293611cce6040820151606086019067ffffffffffffffff60208092828151168552015116910152565b63ffffffff60608201511660a085015263ffffffff6080820151168285015260a0810151151560e0850152015191611d0583611353565b015290565b90606080611d21845160808552608085019061118d565b936020810151602085015267ffffffffffffffff6040820151166040850152015160070b91015290565b906080810182519060808352815180915260a0830190602060a08260051b8601019301915f905b828210611db5575050505067ffffffffffffffff6060611dab819360208701511515602087015260408701518682036040880152611d0a565b9401511691015290565b90919293602080611df0837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff608a600196030186528851611d0a565b960192019201909291611d72565b91909180519260a081526020611f878551604060a0850152611e3a60e08501825167ffffffffffffffff60208092828151168552015116910152565b610240611e58848301516102c06101208801526103a087019061118d565b604083015167ffffffffffffffff1661014087015260608301516fffffffffffffffffffffffffffffffff166101608701526080830151151561018087015260a083015180516101a0880152602090810151805163ffffffff166101c089015201516101e08701529160c0810151151561020087015260e08101516102208701526101008101511515828701526101208101516102608701526101408101516102808701526101608101516102a08701526101808101516102c08701526101a08101516102e08701526101c081015115156103008701526101e0810151610320870152610200810151151561034087015261022081015161036087015201517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff208583030161038086015261118d565b940151937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff608282030160c0830152606060c082019567ffffffffffffffff815116835263ffffffff60208201511660208401526120066040820151604085019060208060409280518552015163ffffffff815116828501520151910152565b01519460c060a0830152855180915260e0820190602060e08260051b8501019701915f905b82821061208f575050505050606061205361208c949560208501518482036020860152611d4b565b9261207c6040820151604085019067ffffffffffffffff60208092828151168552015116910152565b0151906080818403910152611d4b565b90565b909192939760208061214e837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff208d6001960301865282895180516120d28161110e565b835201519060408482015260606120f583516080604085015260c084019061118d565b926fffffffffffffffffffffffffffffffff86820151168284015260408101511515608084015201519060a07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08285030191015261118d565b9a9601949391909101910161202b565b908101906020818303126101de5780359067ffffffffffffffff82116101de5701610240818303126101de5760405191612197836112f8565b813567ffffffffffffffff81116101de57816121b4918401611473565b83526121c38160208401611562565b9260208101938452608083013567ffffffffffffffff81116101de57826121eb91850161175f565b91604082019283526121ff60a08501611545565b93606083019485528160df820112156101de5760405161222161010082611330565b806101c08301918483116101de5760c08401905b8382106126e9575050846122818561020061230c966122765f9c996122cb9c986122fb9f9c9860806fffffffffffffffffffffffffffffffff9a0152611b37565b60a086015201611b37565b60c082015251945191519351169260405197889687967fcb1e26cf00000000000000000000000000000000000000000000000000000000885260c0600489015260c4880190611c70565b83516fffffffffffffffffffffffffffffffff166024880152602084015160448801526040909301516064870152565b600319858303016084860152611dfe565b9060a4830152038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa90811561072b575f91612535575b5061237e81516fffffffffffffffffffffffffffffffff606084015116906135d1565b6123ea60208201516040516123bd602082018093604080916fffffffffffffffffffffffffffffffff8151168452602081015160208501520151910152565b606081526123cc608082611330565b51902061060767ffffffffffffffff602060808601510151166126f9565b6123f3816130c7565b906123fd8261110e565b816124d65767ffffffffffffffff6020604060a0840193845183810151600354918683861c168783161161248d575b50505001516040516124678382018093604080916fffffffffffffffffffffffffffffffff8151168452602081015160208501520151910152565b60608152612476608082611330565b51902092510151165f52600560205260405f205590565b6fffffffffffffffff0000000000000000877fffffffffffffffffffffffffffffffff0000000000000000000000000000000092511692861b16921617176003555f808061242c565b506124e08161110e565b6001810361251e57680100000000000000007fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff600454161760045590565b6125278161110e565b6002810361208c5750600290565b90503d805f833e6125468183611330565b8101906020818303126101de5780519067ffffffffffffffff82116101de5701610180818303126101de576040519160c0830183811067ffffffffffffffff82111761073657604052815167ffffffffffffffff81116101de578201918282039261012084126101de57604051936125bd856112f8565b815167ffffffffffffffff81116101de57820184601f820112156101de576040916125f186836020601f1995519101611b7f565b875201126101de57604051612605816112c0565b61261160208301611bb5565b815261261f60408301611bb5565b602082015260208501526126368360608301611bd8565b604085015261264760a08201611c09565b606085015261265860c08201611c09565b608085015260e08101519081151582036101de576101009160a086015201519060028210156101de57836101409260c06126de960152855261269d8360208301611c37565b60208601526126af8360808301611c37565b60408601526126c060e08201611c1a565b60608601526126d3836101008301611bd8565b608086015201611bd8565b60a08201525f61235b565b8135815260209182019101612235565b67ffffffffffffffff165f52600560205260405f205480156127185790565b7f5b48b457000000000000000000000000000000000000000000000000000000005f5260045ffd5b80518210156127545760209160051b010190565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b335f9081527f4c48d1d2b0f4f7485fc28e3db22341d96a20aa29e6efa8149da9751603abd4e0602052604090205460ff16156127b957565b7fe2517d3f000000000000000000000000000000000000000000000000000000005f52336004527fbd893629a699470e4ec82a5715bb4981fdaacc5d0a728bf5f55b801d8f4ef10b60245260445ffd5b805f525f60205260405f2073ffffffffffffffffffffffffffffffffffffffff33165f5260205260ff60405f205416156128405750565b7fe2517d3f000000000000000000000000000000000000000000000000000000005f523360045260245260445ffd5b9596939190929361287f81611353565b8061304f575080151580613043575b1561300d5793929084926040519586957f13542b7a000000000000000000000000000000000000000000000000000000008752606487019060048801526060602488015252608485019060848560051b8701019481925f907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc184360301935b838310612f2757505050505050600319848403016044850152808352602083019260208260051b82010193835f927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe182360301905b858510612d7c575050505050505090805f9203818373ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165af190811561072b575f91612bd1575b506020815192019160206129d1846135bc565b612a316129de3688611562565b9161060760405185810190612a198287604080916fffffffffffffffffffffffffffffffff8151168452602081015160208501520151910152565b60608152612a28608082611330565b519020916126f9565b0151808203612ba357505060200151906001825111612a74575b50506fffffffffffffffffffffffffffffffff612a6e633b9aca00923690611562565b51160490565b612a80909391936135bc565b918035916fffffffffffffffffffffffffffffffff83168093036101de57909267ffffffffffffffff16905f5b8551811015612b8557612ac08187612740565b519084604051602081019086825260408082015260a081019480519560406060840152865180915260c08301602060c08360051b8601019801915f5b818110612b305750505050816001966020612b26930151608083015203601f198101835282611330565b5190205d01612aad565b919394965091949697602080612b70837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff408b60019603018852895161118d565b970194019101918b9694939298979598612afc565b50935050506fffffffffffffffffffffffffffffffff612a6e612a4b565b7f4b2dfe98000000000000000000000000000000000000000000000000000000005f5260045260245260445ffd5b90503d805f833e612be28183611330565b8101906020818303126101de5780519067ffffffffffffffff82116101de5701906040828203126101de5760405191612c1a836112c0565b8051835260208101519067ffffffffffffffff82116101de570181601f820112156101de57805190612c4b826115ed565b92612c596040519485611330565b82845260208085019360051b830101918183116101de5760208101935b838510612c8d57505050505060208201525f6129be565b845167ffffffffffffffff81116101de5782016040601f1982860301126101de5760405190612cbb826112c0565b602081015167ffffffffffffffff81116101de5760209082010185601f820112156101de578051612ceb816115ed565b91612cf96040519384611330565b81835260208084019260051b820101908882116101de5760208101925b828410612d3d57505050509160406020949285948352015183820152815201940193612c76565b835167ffffffffffffffff81116101de5782018a603f820112156101de57602091612d718c83604086809601519101611b7f565b815201930192612d16565b9193959750919395601f198282030185528735838112156101de57840190612da8602082019280613322565b8091936020845252604082019060408160051b8401019380935f915b838310612deb57505050505050602080600192990195019501929091889796949592612962565b9091929394957fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0838203018652612e228783613417565b803560028110156101de57612e3681611353565b8252612e59612e4860208301836133e5565b606060208501526060840190613449565b9060408101357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff61823603018112156101de576001936020938493612f199301916040818303910152612f0b612eed612ec2612eb48580613395565b60a0865260a0860191613375565b612ecd878601611466565b151587850152612ee060408601866133e5565b8482036040860152613449565b92612efa60608201611466565b1515606084015260808101906133e5565b906080818403910152613449565b980196019493019190612dc4565b91939596987fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7c908992949699030183528735868112156101de578201906040810191612f738180613322565b809460408552526060830160608560051b85010194825f5b828110612fbb5750505050506001926020928380809401359101529901930193019092899896959394929461290d565b9091929396602080613000837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa08b60019603018952612ffa8c88613395565b90613375565b9901950193929101612f8b565b7fb0369c31000000000000000000000000000000000000000000000000000000005f52600452600160245261ffff60445260645ffd5b5061ffff81111561288e565b60ff9061305b81611353565b61306481611353565b7f112d89cc000000000000000000000000000000000000000000000000000000005f521660045260245ffd5b15613099575050565b7f6d23e8a3000000000000000000000000000000000000000000000000000000005f5260045260245260445ffd5b67ffffffffffffffff602060a08301510151165f52600560205260405f205480155f146130f45750505f90565b60408201908151604051613132602082018093604080916fffffffffffffffffffffffffffffffff8151168452602081015160208501520151910152565b60608152613141608082611330565b519020149182159261315f575b50501561315a57600190565b600290565b6fffffffffffffffffffffffffffffffff91925060208291015151169151511611155f8061314e565b805f525f60205260405f2073ffffffffffffffffffffffffffffffffffffffff83165f5260205260ff60405f205416155f1461325457805f525f60205260405f2073ffffffffffffffffffffffffffffffffffffffff83165f5260205260405f2060017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0082541617905573ffffffffffffffffffffffffffffffffffffffff339216907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d5f80a4600190565b50505f90565b805f525f60205260405f2073ffffffffffffffffffffffffffffffffffffffff83165f5260205260ff60405f2054165f1461325457805f525f60205260405f2073ffffffffffffffffffffffffffffffffffffffff83165f5260205260405f207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00815416905573ffffffffffffffffffffffffffffffffffffffff339216907ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b5f80a4600190565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1823603018112156101de57016020813591019167ffffffffffffffff82116101de578160051b360383136101de57565b601f8260209493601f1993818652868601375f8582860101520116010190565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1823603018112156101de57016020813591019167ffffffffffffffff82116101de5781360383136101de57565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81823603018112156101de570190565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa1823603018112156101de570190565b6134646134568280613395565b608085526080850191613375565b6020820135602084015261347b60408301836133e5565b908381036040850152813560038110156101de576134988161110e565b8152602082013560038110156101de576134b18161110e565b602082015260408201359060038210156101de5760806134ed61350894846134de6134fd9699989961110e565b60408501526060810190613395565b9190928160608201520191613375565b926060810190613322565b90916060818503910152808352602083019060208160051b85010193835f915b8383106135385750505050505090565b909192939495601f198282030186526135518784613417565b80359160038310156101de576135ae60209283928561357160019761110e565b81526135a061359561358586850185613395565b6060888601526060850191613375565b926040810190613395565b916040818503910152613375565b980196019493019190613528565b3567ffffffffffffffff811681036101de5790565b906fffffffffffffffffffffffffffffffff633b9aca00911604428111613824578042034281116137f757610708106137c857508051516136136001546111b2565b14806137a1575b8151901561374b5750602081015160ff815116906002549160ff8316928382149283613732575b6020015160ff1692156136ea575050505063ffffffff606082015116906004549163ffffffff83168082036136bc57505063ffffffff608081920151169160201c1680820361368e575050565b7f1eb5c1eb000000000000000000000000000000000000000000000000000000005f5260045260245260445ffd5b7f73b1bde3000000000000000000000000000000000000000000000000000000005f5260045260245260445ffd5b6084945060ff90604051947fd382033a000000000000000000000000000000000000000000000000000000008652600486015260081c16602484015260448301526064820152fd5b6020810151600883901c60ff9081169116149350613641565b61379d906040519182917ff6b6676b0000000000000000000000000000000000000000000000000000000083526040600484015261378b60448401611203565b9060031984830301602485015261118d565b0390fd5b508051602081519101206040516137bb816107f181611203565b602081519101201461361a565b7f12e74add000000000000000000000000000000000000000000000000000000005f524260045260245260445ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b7f20daedb6000000000000000000000000000000000000000000000000000000005f524260045260245260445ffd5b90613890575080511561386857805190602001fd5b7fd6bda275000000000000000000000000000000000000000000000000000000005f5260045ffd5b815115806138e3575b6138a1575090565b73ffffffffffffffffffffffffffffffffffffffff907f9996b315000000000000000000000000000000000000000000000000000000005f521660045260245ffd5b50803b1561389956fea164736f6c634300081c000a2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d4c48d1d2b0f4f7485fc28e3db22341d96a20aa29e6efa8149da9751603abd4e06e0c24a6e293ff9b755263dbaa15ba3796b0b8d3fe17cfb4ddf8143b268eac47bd893629a699470e4ec82a5715bb4981fdaacc5d0a728bf5f55b801d8f4ef10bad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5",
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
