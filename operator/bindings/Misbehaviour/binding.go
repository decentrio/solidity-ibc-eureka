// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractMisbehaviour

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
	AppHash            []byte
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

// IICS07TendermintMsgsChainId is an auto generated low-level Go binding around an user-defined struct.
type IICS07TendermintMsgsChainId struct {
	Id             string
	RevisionNumber uint64
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

// IMisbehaviourMsgsMisbehaviour is an auto generated low-level Go binding around an user-defined struct.
type IMisbehaviourMsgsMisbehaviour struct {
	ClientId IICS07TendermintMsgsChainId
	Header1  IICS07TendermintMsgsHeader
	Header2  IICS07TendermintMsgsHeader
}

// IMisbehaviourMsgsMisbehaviourOutput is an auto generated low-level Go binding around an user-defined struct.
type IMisbehaviourMsgsMisbehaviourOutput struct {
	ClientState            IICS07TendermintMsgsClientState
	Time                   *big.Int
	TrustedHeight1         IICS02ClientMsgsHeight
	TrustedHeight2         IICS02ClientMsgsHeight
	TrustedConsensusState1 IICS07TendermintMsgsConsensusState
	TrustedConsensusState2 IICS07TendermintMsgsConsensusState
}

// ContractMisbehaviourMetaData contains all meta data concerning the ContractMisbehaviour contract.
var ContractMisbehaviourMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"checkForMisbehaviour\",\"inputs\":[{\"name\":\"clientState\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.ClientState\",\"components\":[{\"name\":\"chainId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"trustLevel\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.TrustThreshold\",\"components\":[{\"name\":\"numerator\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"denominator\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"name\":\"latestHeight\",\"type\":\"tuple\",\"internalType\":\"structIICS02ClientMsgs.Height\",\"components\":[{\"name\":\"revisionNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revisionHeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"trustingPeriod\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"unbondingPeriod\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isFrozen\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"zkAlgorithm\",\"type\":\"uint8\",\"internalType\":\"enumIICS07TendermintMsgs.SupportedZkAlgorithm\"}]},{\"name\":\"misbehaviour\",\"type\":\"tuple\",\"internalType\":\"structIMisbehaviourMsgs.Misbehaviour\",\"components\":[{\"name\":\"client_id\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.ChainId\",\"components\":[{\"name\":\"id\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"revisionNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"header1\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.Header\",\"components\":[{\"name\":\"signedHeader\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.SignedHeader\",\"components\":[{\"name\":\"header\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.BlockHeader\",\"components\":[{\"name\":\"version\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.Version\",\"components\":[{\"name\":\"blockVersion\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"appVersion\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"chainId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"height\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"time\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"hasLastBlockId\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"lastBlockId\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.BlockId\",\"components\":[{\"name\":\"hashData\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"partSetHeader\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.PartSetHeader\",\"components\":[{\"name\":\"total\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashData\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]},{\"name\":\"hasLastCommitHash\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"lastCommitHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"hasDataHash\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"dataHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"validatorsHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"nextValidatorsHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"consensusHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"appHash\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"hasLastResultsHash\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"lastResultsHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"hasEvidenceHash\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"evidenceHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"proposerAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"commit\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.BlockCommit\",\"components\":[{\"name\":\"height\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"round\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"blockId\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.BlockId\",\"components\":[{\"name\":\"hashData\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"partSetHeader\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.PartSetHeader\",\"components\":[{\"name\":\"total\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashData\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]},{\"name\":\"commitSigs\",\"type\":\"tuple[]\",\"internalType\":\"structIICS07TendermintMsgs.CommitSig[]\",\"components\":[{\"name\":\"flag\",\"type\":\"uint8\",\"internalType\":\"enumIICS07TendermintMsgs.CommitSigFlag\"},{\"name\":\"data\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.CommitSigData\",\"components\":[{\"name\":\"validatorAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"timestamp\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"hasSignature\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]}]}]},{\"name\":\"validatorSet\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.ValidatorSet\",\"components\":[{\"name\":\"validators\",\"type\":\"tuple[]\",\"internalType\":\"structIICS07TendermintMsgs.ValidatorInfo[]\",\"components\":[{\"name\":\"valAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"pubKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"votingPower\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"proposerPriority\",\"type\":\"int64\",\"internalType\":\"int64\"}]},{\"name\":\"hasProposer\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"proposer\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.ValidatorInfo\",\"components\":[{\"name\":\"valAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"pubKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"votingPower\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"proposerPriority\",\"type\":\"int64\",\"internalType\":\"int64\"}]},{\"name\":\"totalVotingPower\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"trustedHeight\",\"type\":\"tuple\",\"internalType\":\"structIICS02ClientMsgs.Height\",\"components\":[{\"name\":\"revisionNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revisionHeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"trustedNextValidatorSet\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.ValidatorSet\",\"components\":[{\"name\":\"validators\",\"type\":\"tuple[]\",\"internalType\":\"structIICS07TendermintMsgs.ValidatorInfo[]\",\"components\":[{\"name\":\"valAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"pubKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"votingPower\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"proposerPriority\",\"type\":\"int64\",\"internalType\":\"int64\"}]},{\"name\":\"hasProposer\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"proposer\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.ValidatorInfo\",\"components\":[{\"name\":\"valAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"pubKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"votingPower\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"proposerPriority\",\"type\":\"int64\",\"internalType\":\"int64\"}]},{\"name\":\"totalVotingPower\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}]},{\"name\":\"header2\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.Header\",\"components\":[{\"name\":\"signedHeader\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.SignedHeader\",\"components\":[{\"name\":\"header\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.BlockHeader\",\"components\":[{\"name\":\"version\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.Version\",\"components\":[{\"name\":\"blockVersion\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"appVersion\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"chainId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"height\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"time\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"hasLastBlockId\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"lastBlockId\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.BlockId\",\"components\":[{\"name\":\"hashData\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"partSetHeader\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.PartSetHeader\",\"components\":[{\"name\":\"total\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashData\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]},{\"name\":\"hasLastCommitHash\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"lastCommitHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"hasDataHash\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"dataHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"validatorsHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"nextValidatorsHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"consensusHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"appHash\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"hasLastResultsHash\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"lastResultsHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"hasEvidenceHash\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"evidenceHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"proposerAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"commit\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.BlockCommit\",\"components\":[{\"name\":\"height\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"round\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"blockId\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.BlockId\",\"components\":[{\"name\":\"hashData\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"partSetHeader\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.PartSetHeader\",\"components\":[{\"name\":\"total\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashData\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]},{\"name\":\"commitSigs\",\"type\":\"tuple[]\",\"internalType\":\"structIICS07TendermintMsgs.CommitSig[]\",\"components\":[{\"name\":\"flag\",\"type\":\"uint8\",\"internalType\":\"enumIICS07TendermintMsgs.CommitSigFlag\"},{\"name\":\"data\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.CommitSigData\",\"components\":[{\"name\":\"validatorAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"timestamp\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"hasSignature\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]}]}]},{\"name\":\"validatorSet\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.ValidatorSet\",\"components\":[{\"name\":\"validators\",\"type\":\"tuple[]\",\"internalType\":\"structIICS07TendermintMsgs.ValidatorInfo[]\",\"components\":[{\"name\":\"valAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"pubKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"votingPower\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"proposerPriority\",\"type\":\"int64\",\"internalType\":\"int64\"}]},{\"name\":\"hasProposer\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"proposer\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.ValidatorInfo\",\"components\":[{\"name\":\"valAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"pubKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"votingPower\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"proposerPriority\",\"type\":\"int64\",\"internalType\":\"int64\"}]},{\"name\":\"totalVotingPower\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"trustedHeight\",\"type\":\"tuple\",\"internalType\":\"structIICS02ClientMsgs.Height\",\"components\":[{\"name\":\"revisionNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revisionHeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"trustedNextValidatorSet\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.ValidatorSet\",\"components\":[{\"name\":\"validators\",\"type\":\"tuple[]\",\"internalType\":\"structIICS07TendermintMsgs.ValidatorInfo[]\",\"components\":[{\"name\":\"valAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"pubKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"votingPower\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"proposerPriority\",\"type\":\"int64\",\"internalType\":\"int64\"}]},{\"name\":\"hasProposer\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"proposer\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.ValidatorInfo\",\"components\":[{\"name\":\"valAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"pubKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"votingPower\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"proposerPriority\",\"type\":\"int64\",\"internalType\":\"int64\"}]},{\"name\":\"totalVotingPower\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}]}]},{\"name\":\"trustedConsensusState1\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.ConsensusState\",\"components\":[{\"name\":\"timestamp\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"nextValidatorsHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"trustedConsensusState2\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.ConsensusState\",\"components\":[{\"name\":\"timestamp\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"nextValidatorsHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"time\",\"type\":\"uint128\",\"internalType\":\"uint128\"}],\"outputs\":[{\"name\":\"output\",\"type\":\"tuple\",\"internalType\":\"structIMisbehaviourMsgs.MisbehaviourOutput\",\"components\":[{\"name\":\"clientState\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.ClientState\",\"components\":[{\"name\":\"chainId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"trustLevel\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.TrustThreshold\",\"components\":[{\"name\":\"numerator\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"denominator\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"name\":\"latestHeight\",\"type\":\"tuple\",\"internalType\":\"structIICS02ClientMsgs.Height\",\"components\":[{\"name\":\"revisionNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revisionHeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"trustingPeriod\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"unbondingPeriod\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isFrozen\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"zkAlgorithm\",\"type\":\"uint8\",\"internalType\":\"enumIICS07TendermintMsgs.SupportedZkAlgorithm\"}]},{\"name\":\"time\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"trustedHeight1\",\"type\":\"tuple\",\"internalType\":\"structIICS02ClientMsgs.Height\",\"components\":[{\"name\":\"revisionNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revisionHeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"trustedHeight2\",\"type\":\"tuple\",\"internalType\":\"structIICS02ClientMsgs.Height\",\"components\":[{\"name\":\"revisionNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revisionHeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"trustedConsensusState1\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.ConsensusState\",\"components\":[{\"name\":\"timestamp\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"nextValidatorsHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"trustedConsensusState2\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.ConsensusState\",\"components\":[{\"name\":\"timestamp\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"nextValidatorsHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"error\",\"name\":\"ChainIdMismatch\",\"inputs\":[{\"name\":\"expected\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"actual\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"ChainIdMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CheckForMisbehaviourFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedToVerifyHeader\",\"inputs\":[{\"name\":\"description\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"InsufficientMisbehaviourHeaderHeight\",\"inputs\":[{\"name\":\"height1\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"height2\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InsufficientTrustingPeriod\",\"inputs\":[{\"name\":\"durationSinceConsensusState\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"trustingPeriod\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"type\":\"error\",\"name\":\"InvalidChainId\",\"inputs\":[{\"name\":\"id\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"InvalidClientId\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidConsensusStateTimestamp\",\"inputs\":[{\"name\":\"timestamp\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"type\":\"error\",\"name\":\"InvalidHeaderHeight\",\"inputs\":[{\"name\":\"height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"MisbehaviourNotDetected\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MisbehaviourVerificationFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MismatchedRevisionHeight\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"actual\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ValSetHashMismatch\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"actual\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]",
	Bin: "0x60808060405234601557612a17908161001a8239f35b5f80fdfe6080806040526004361015610012575f80fd5b5f3560e01c63abb6074214610025575f80fd5b346108f5576101207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126108f55760043567ffffffffffffffff81116108f5578036036101207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8201126108f55761009e836108f9565b816004013567ffffffffffffffff81116108f5576040916100e67fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc92600436918701016109e1565b855201126108f557604051916100fb83610915565b61010760248301610a25565b835261011560448301610a25565b60208401526020810192835261012e3660648401610a48565b604082015261010461014260a48401610a7e565b926060830193845261015660c48201610a7e565b608084015261016760e48201610a8f565b60a0840152013560028110156108f55760c082015260243567ffffffffffffffff81116108f55760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82360301126108f557604051936101c785610931565b816004013567ffffffffffffffff81116108f557820160407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82360301126108f5576040519061021682610915565b600481013567ffffffffffffffff81116108f55761024891610240602492600436918401016109e1565b845201610a33565b60208201528552602482013567ffffffffffffffff81116108f5576102739060043691850101610c9d565b9160208601928352604481013567ffffffffffffffff81116108f55760409160046102a19236920101610c9d565b950194855260607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffbc3601126108f557604051916102dd83610931565b6044356fffffffffffffffffffffffffffffffff811681036108f557835260643560208401526040830190608435825260607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff5c3601126108f5576040519061034482610931565b60a4356fffffffffffffffffffffffffffffffff811681036108f557825260c4356020830152604082019060e435825261037c610a9c565b506040519860c08a01978a891067ffffffffffffffff8a11176108c8578a986040526040516103aa816108f9565b606081526040516103ba81610915565b5f81525f602082015260208201526040516103d481610915565b5f81525f602082015260408201525f60608201525f60808201525f60a08201525f60c0820152895260208901975f895260408a019660405161041581610915565b5f81525f6020820152885260608b019b60405161043181610915565b5f81525f60208201528d5260a060808d019c8d61044c6110a3565b905201996104586110a3565b8b52845160405161048781610479602082019460208652604083019061107e565b03601f1981018352826109a2565b51902060208851515101516040516104af81610479602082019460208652604083019061107e565b519020036108a0576104c187516125b6565b6104cb86516125b6565b60208751515101516040516104f081610479602082019460208652604083019061107e565b519020602087515151015160405161051881610479602082019460208652604083019061107e565b5190200361081c576105a0610533602089515151015161116a565b67ffffffffffffffff6020818161054f818d515151015161116a565b94015116928260408d5151510151166040519461056b86610915565b8552828501520151169067ffffffffffffffff60408a5151510151166040519261059484610915565b83526020830152612834565b6105a98161279c565b156107d1576106209896946fffffffffffffffffffffffffffffffff9485946106059a989463ffffffff610614955191511698604051998a926105eb84610931565b83526020830152600f60408301528742169c8d955161116a565b988991519451169251936114c7565b519451169251936114c7565b60405195602087525193610180602088015260c061064d86516101206101a08b01526102c08a019061107e565b9560ff602080830151828151166101c08d01520151166101e08a015261069260408201516102008b019067ffffffffffffffff60208092828151168552015116910152565b63ffffffff6060820151166102408a015263ffffffff6080820151166102608a015260a081015115156102808a0152015160028110156107a457879661071c61076e946fffffffffffffffffffffffffffffffff6107a09861073d956102a08d0152511660408b01525160608a019067ffffffffffffffff60208092828151168552015116910152565b5160a088019067ffffffffffffffff60208092828151168552015116910152565b5180516fffffffffffffffffffffffffffffffff1660e0870152602081015161010087015260400151610120860152565b5180516fffffffffffffffffffffffffffffffff16610140850152602081015161016085015260400151610180840152565b0390f35b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b8567ffffffffffffffff604081818b51515101511692515151015116907f4447469a000000000000000000000000000000000000000000000000000000005f5260045260245260445ffd5b61086c8661089c6020808b51515101519251515101516040519384937ff6b6676b00000000000000000000000000000000000000000000000000000000855260406004860152604485019061107e565b907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc84830301602485015261107e565b0390fd5b7fa179f8c9000000000000000000000000000000000000000000000000000000005f5260045ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f80fd5b60e0810190811067ffffffffffffffff8211176108c857604052565b6040810190811067ffffffffffffffff8211176108c857604052565b6060810190811067ffffffffffffffff8211176108c857604052565b6080810190811067ffffffffffffffff8211176108c857604052565b610260810190811067ffffffffffffffff8211176108c857604052565b60a0810190811067ffffffffffffffff8211176108c857604052565b90601f601f19910116810190811067ffffffffffffffff8211176108c857604052565b67ffffffffffffffff81116108c857601f01601f191660200190565b81601f820112156108f5576020813591016109fb826109c5565b92610a0960405194856109a2565b828452828201116108f557815f92602092838601378301015290565b359060ff821682036108f557565b359067ffffffffffffffff821682036108f557565b91908260409103126108f557604051610a6081610915565b6020610a79818395610a7181610a33565b855201610a33565b910152565b359063ffffffff821682036108f557565b359081151582036108f557565b61010435906fffffffffffffffffffffffffffffffff821682036108f557565b35906fffffffffffffffffffffffffffffffff821682036108f557565b8092910391606083126108f557604051610af281610915565b6040601f1982958435845201126108f5576020906040805193610b1485610915565b610b1f848201610a7e565b85520135828401520152565b67ffffffffffffffff81116108c85760051b60200190565b91906080838203126108f55760405190610b5c8261094d565b8193803567ffffffffffffffff81116108f557606092610b7d9183016109e1565b835260208101356020840152610b9560408201610a33565b60408401520135908160070b82036108f55760600152565b9190916080818403126108f55760405190610bc78261094d565b8193813567ffffffffffffffff81116108f557820181601f820112156108f5578035610bf281610b2b565b91610c0060405193846109a2565b81835260208084019260051b820101918483116108f55760208201905b838210610c6f57505050508352610c3660208301610a8f565b602084015260408201359067ffffffffffffffff82116108f55782610c6460609492610a7994869401610b43565b604086015201610a33565b813567ffffffffffffffff81116108f557602091610c9288848094880101610b43565b815201910190610c1d565b919060a0838203126108f55760405190610cb68261094d565b8193803567ffffffffffffffff81116108f55781016040818403126108f55760405190610ce282610915565b803567ffffffffffffffff81116108f55781016102c0818603126108f55760405190610d0d82610969565b610d178682610a48565b8252604081013567ffffffffffffffff81116108f55786610d399183016109e1565b6020830152610d4a60608201610a33565b6040830152610d5b60808201610abc565b6060830152610d6c60a08201610a8f565b6080830152610d7e8660c08301610ad9565b60a0830152610d906101208201610a8f565b60c083015261014081013560e0830152610dad6101608201610a8f565b6101008301526101808101356101208301526101a08101356101408301526101c08101356101608301526101e081013561018083015261020081013567ffffffffffffffff81116108f55786610e049183016109e1565b6101a0830152610e176102208201610a8f565b6101c08301526102408101356101e0830152610e366102608201610a8f565b6102008301526102808101356102208301526102a08101359067ffffffffffffffff82116108f557610e6a918791016109e1565b610240820152825260208101359067ffffffffffffffff82116108f5570160c0818503126108f55760405190610e9f8261094d565b610ea881610a33565b8252610eb660208201610a7e565b6020830152610ec88560408301610ad9565b604083015260a08101359067ffffffffffffffff82116108f5570184601f820112156108f557803590610efa82610b2b565b91610f0860405193846109a2565b80835260208084019160051b830101918783116108f55760208101915b838310610f955750505050606082015260208201528352602081013567ffffffffffffffff81116108f55782610f5c918301610bad565b6020840152610f6e8260408301610a48565b604084015260808101359167ffffffffffffffff83116108f557606092610a799201610bad565b823567ffffffffffffffff81116108f5578201906040601f19838c0301126108f55760405191610fc483610915565b602081013560038110156108f5578352604081013567ffffffffffffffff81116108f5576020910101906080828c03126108f557604051926110058461094d565b823567ffffffffffffffff81116108f5578c6110229185016109e1565b845261103060208401610abc565b602085015261104160408401610a8f565b604085015260608301359367ffffffffffffffff85116108f55761106a8d6020968796016109e1565b606082015283820152815201920191610f25565b90601f19601f602080948051918291828752018686015e5f8582860101520116010190565b604051906110b082610931565b5f6040838281528260208201520152565b9081518110156110d2570160200190565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b90611109826109c5565b61111660405191826109a2565b828152601f1961112682946109c5565b0190602036910137565b9190820180921161113d57565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b60405161117681610915565b606081525f60209091015280515f19810190811161113d575b7f2d000000000000000000000000000000000000000000000000000000000000007fff000000000000000000000000000000000000000000000000000000000000006111db83856110c1565b5116146111f057801561113d575f190161118f565b91905f19831461147257805183810390811161113d575f19810190811161113d5761121a906110ff565b905f5b825181101561127d57600185019081861161113d577fff0000000000000000000000000000000000000000000000000000000000000061126861126283600195611130565b856110c1565b51165f1a61127682866110c1565b530161121d565b509192908051158015611414575b61138557805161129a916128a0565b9190158015611403575b611385576112b1816110ff565b905f5b818110611342575050516001811190811591611336575b506112f25767ffffffffffffffff90604051926112e784610915565b835216602082015290565b606460405162461bcd60e51b815260206004820152601b60248201527f496e76616c696420636861696e20707265666978206c656e67746800000000006044820152fd5b602b915010155f6112cb565b807fff00000000000000000000000000000000000000000000000000000000000000611370600193886110c1565b51165f1a61137e82866110c1565b53016112b4565b50508051806001111590816113f8575b50156113b357604051906113a882610915565b81525f602082015290565b60405162461bcd60e51b815260206004820152601760248201527f496e76616c696420636861696e204944206c656e6774680000000000000000006044820152606490fd5b60409150105f611395565b5067ffffffffffffffff82116112a4565b6110d2577f30000000000000000000000000000000000000000000000000000000000000007fff0000000000000000000000000000000000000000000000000000000000000060208301511614801561128b5750600181511161128b565b8091925051806001111590816113f85750156113b357604051906113a882610915565b906fffffffffffffffffffffffffffffffff809116911603906fffffffffffffffffffffffffffffffff821161113d57565b9491939573__$7440a880b7578767f72184d998805816e4$__6060870161151e60208251604051809381927fc87f1f69000000000000000000000000000000000000000000000000000000008352600483016124f9565b0381865af48015612289578a915f91612483575b50036123d9576fffffffffffffffffffffffffffffffff8616986fffffffffffffffffffffffffffffffff8616968a88106123ad5761157360209188611495565b9801976fffffffffffffffffffffffffffffffff63ffffffff8a511691168181101561237f57505084519889518015908115612374575b50612330575f5b8a518110156117e8577fff000000000000000000000000000000000000000000000000000000000000006115e5828d6110c1565b51167f610000000000000000000000000000000000000000000000000000000000000081101590816117bd575b8115611761575b8115611705575b81156116db575b81156116b1575b8115611687575b5015611643576001016115b1565b606460405162461bcd60e51b815260206004820152601860248201527f696e76616c696420636861696e206964206368617273657400000000000000006044820152fd5b7f2e000000000000000000000000000000000000000000000000000000000000009150145f611635565b7f5f000000000000000000000000000000000000000000000000000000000000008114915061162e565b7f2d0000000000000000000000000000000000000000000000000000000000000081149150611627565b90507f300000000000000000000000000000000000000000000000000000000000000081101580611737575b90611620565b507f3900000000000000000000000000000000000000000000000000000000000000811115611731565b90507f410000000000000000000000000000000000000000000000000000000000000081101580611793575b90611619565b507f5a0000000000000000000000000000000000000000000000000000000000000081111561178d565b7f7a000000000000000000000000000000000000000000000000000000000000008111159150611612565b50919395979990929496985060405161180081610915565b60405161180c81610915565b60405161181881610969565b60405161182481610915565b5f81525f60208201528152606060208201525f60408201525f60608201525f60808201526118506127a6565b60a08201525f60c08201525f60e08201525f6101008201525f6101208201525f6101408201525f6101608201525f61018082015260606101a08201525f6101c08201525f6101e08201525f6102008201525f610220820152606061024082015281526040516118be8161094d565b5f81525f60208201526118cf6127a6565b60408201526060808201526020820152815260206118eb6127d1565b91015281519467ffffffffffffffff60206040818601519a81519961190f8b610915565b8a52828a019b8c5251955f6080835161192781610986565b606081528286820152828582015261193d6127d1565b60608201520152015101511694516040519361195885610986565b845260208401928352604084019586526060840152608083019384526119ae60208951604051809381927fc87f1f69000000000000000000000000000000000000000000000000000000008352600483016124f9565b0381855af4908115612289575f916122fe575b50610140875151015103612294576020611ba3918751519060405180809581947f95b25f79000000000000000000000000000000000000000000000000000000008352866004840152611a2e60248401825167ffffffffffffffff60208092828151168552015116910152565b610240611b3c611a4e898401516102c060648801526102e487019061107e565b67ffffffffffffffff60408501511660848701526fffffffffffffffffffffffffffffffff60608501511660a48701526080840151151560c4870152898060a0860151805160e48a0152015163ffffffff815116610104890152015161012487015260c0840151151561014487015260e084015161016487015261010084015115156101848701526101208401516101a48701526101408401516101c48701526101608401516101e48701526101808401516102048701526101a08401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc8783030161022488015261107e565b916101c081015115156102448601526101e081015161026486015261020081015115156102848601526102208101516102a486015201517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc848303016102c485015261107e565b03915af4908115612289575f91612257575b506040602087510151015151036121ed5760606020865198519801510151975f995f5b8a51811015611c1557611beb818c61292a565b5151611bf68161279c565b611bff8161279c565b611c0c575b600101611bd8565b60019b50611c04565b50909192939495969799156121835788518a5151036120f3575f5b8951811015611d7857611c43818b61292a565b518051611c4f8161279c565b611c588161279c565b611c6757506001905b01611c30565b6020909c97939995919c9a969298949a015151995f995f5b8d518051821015611d645781611c949161292a565b5151604051611cc26020828180820195805191829101875e81015f838201520301601f1981018352826109a2565b5190208d604051611cf26020828180820195805191829101875e81015f838201520301601f1981018352826109a2565b51902014611d0257600101611c7f565b5093979c9195995093979195995060015b15611d2057600190611c61565b606460405162461bcd60e51b815260206004820152601d60248201527f696e76616c696420636f6d6d69743a206661756c7479207369676e65720000006044820152fd5b505093979c91959990949892969a50611d13565b509295985092959863ffffffff919497505116906fffffffffffffffffffffffffffffffff8451168091109283156120cb575b505050612061576fffffffffffffffffffffffffffffffff8060608751510151169151161015611ff75760208451510151604051611e086020828180820195805191829101875e81015f838201520301601f1981018352826109a2565b5190209051604051611e396020828180820195805191829101875e81015f838201520301601f1981018352826109a2565b51902003611fb357611e5567ffffffffffffffff835116612814565b8351516040015167ffffffffffffffff9182169116818103611f2c5750506101408351510151905103611ec25767ffffffffffffffff6040611e9b8280945b5116612814565b935151015116911614611eaa57565b60036020604051611eba81610915565b600281520152565b608460405162461bcd60e51b815260206004820152602f60248201527f696e76616c696420626c6f636b3a206e6578742076616c696461746f7220736560448201527f742068617368206d69736d6174636800000000000000000000000000000000006064820152fd5b11159050611f4a5767ffffffffffffffff6040611e9b828094611e94565b608460405162461bcd60e51b8152602060048201526024808201527f696e76616c696420626c6f636b3a206e6f6e20696e6372656173696e6720686560448201527f69676874000000000000000000000000000000000000000000000000000000006064820152fd5b606460405162461bcd60e51b815260206004820152602060248201527f696e76616c696420626c6f636b3a20636861696e2d6964206d69736d617463686044820152fd5b608460405162461bcd60e51b815260206004820152602560248201527f696e76616c696420626c6f636b3a206e6f6e206d6f6e6f746f6e69632062667460448201527f2074696d650000000000000000000000000000000000000000000000000000006064820152fd5b608460405162461bcd60e51b815260206004820152603c60248201527f696e76616c696420626c6f636b3a20756e74727573746564207374617465206960448201527f73206f757473696465206f66207472757374696e6720706572696f64000000006064820152fd5b6fffffffffffffffffffffffffffffffff929350906120e991611495565b16115f8080611dab565b60a460405162461bcd60e51b815260206004820152604860248201527f696e76616c696420636f6d6d69743a206e756d626572206f66207369676e617460448201527f7572657320646f6573206e6f74206d61746368206e756d626572206f6620766160648201527f6c696461746f72730000000000000000000000000000000000000000000000006084820152fd5b608460405162461bcd60e51b815260206004820152602560248201527f696e76616c696420636f6d6d69743a206e6f2070726573656e74207369676e6160448201527f74757265730000000000000000000000000000000000000000000000000000006064820152fd5b608460405162461bcd60e51b815260206004820152602360248201527f696e76616c696420626c6f636b3a206865616465722068617368206d69736d6160448201527f74636800000000000000000000000000000000000000000000000000000000006064820152fd5b90506020813d602011612281575b81612272602093836109a2565b810103126108f557515f611bb5565b3d9150612265565b6040513d5f823e3d90fd5b608460405162461bcd60e51b815260206004820152602a60248201527f696e76616c696420626c6f636b3a2076616c696461746f72207365742068617360448201527f68206d69736d61746368000000000000000000000000000000000000000000006064820152fd5b90506020813d602011612328575b81612319602093836109a2565b810103126108f557515f6119c1565b3d915061230c565b606460405162461bcd60e51b815260206004820152601760248201527f496e76616c696420636861696e206964206c656e6774680000000000000000006044820152fd5b60329150115f6115aa565b7fdb020436000000000000000000000000000000000000000000000000000000005f5260045260245260445ffd5b8a7fe42a1980000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b60a46040517ff492ef2b00000000000000000000000000000000000000000000000000000000815260206004820152604360248201527f74727573746564206e6578742076616c696461746f722073657420686173682060448201527f646f6573206e6f74206d6174636820686173682073746f726564206f6e20636860648201527f61696e00000000000000000000000000000000000000000000000000000000006084820152fd5b9150506020813d6020116124b0575b8161249f602093836109a2565b810103126108f5578990515f611532565b3d9150612492565b906060806124cf845160808552608085019061107e565b936020810151602085015267ffffffffffffffff6040820151166040850152015160070b91015290565b6020815260a0810182519060806020840152815180915260c0830190602060c08260051b8601019301915f905b82821061256d575050505067ffffffffffffffff60606125636080936020870151151560408701526040870151601f1987830301848801526124b8565b9401511691015290565b909192936020806125a8837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff408a6001960301865288516124b8565b960192019201909291612526565b60206125c681835151015161116a565b0167ffffffffffffffff81511690604083019167ffffffffffffffff835151169081810361276e57505067ffffffffffffffff61262d91511667ffffffffffffffff6040855151015116926040519161261e83610915565b82526020820193845251612834565b6126368161279c565b6002811490811561275a575b50612723575080602080612685930151604051809481927fc87f1f69000000000000000000000000000000000000000000000000000000008352600483016124f9565b038173__$7440a880b7578767f72184d998805816e4$__5af4918215612289575f926126ed575b50515161014001518082036126bf575050565b7f4d6d8014000000000000000000000000000000000000000000000000000000005f5260045260245260445ffd5b9091506020813d60201161271b575b81612709602093836109a2565b810103126108f55751906101406126ac565b3d91506126fc565b67ffffffffffffffff9051167f90f4dbed000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b600191506127678161279c565b145f612642565b7f05b4c7a3000000000000000000000000000000000000000000000000000000005f5260045260245260445ffd5b600311156107a457565b604051906127b382610915565b5f82526040516020836127c583610915565b5f83525f828401520152565b604051906127de8261094d565b5f6060838181528260208201526040516127f78161094d565b828152836020820152836040820152838382015260408201520152565b67ffffffffffffffff60019116019067ffffffffffffffff821161113d57565b67ffffffffffffffff81511667ffffffffffffffff835116908181105f1461285e57505050505f90565b111561286b575050600290565b602067ffffffffffffffff81819301511692015116908181105f146128905750505f90565b111561289b57600290565b600190565b5f929183915b8183106128b65750505060019190565b90919360ff6128ec7fff00000000000000000000000000000000000000000000000000000000000000602088860101511661293e565b16906009821161291e57600a810290808204600a149015171561113d5760019161291591611130565b940191906128a6565b5050505090505f905f90565b80518210156110d25760209160051b010190565b60f81c602f811180612a00575b15612978577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffd00160ff1690565b60608111806129f6575b156129af577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa90160ff1690565b60408111806129ec575b156129e6577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc90160ff1690565b5060ff90565b50604781106129b9565b5060678110612982565b50603a811061294b56fea164736f6c634300081c000a",
}

// ContractMisbehaviourABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMisbehaviourMetaData.ABI instead.
var ContractMisbehaviourABI = ContractMisbehaviourMetaData.ABI

// ContractMisbehaviourBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractMisbehaviourMetaData.Bin instead.
var ContractMisbehaviourBin = ContractMisbehaviourMetaData.Bin

// DeployContractMisbehaviour deploys a new Ethereum contract, binding an instance of ContractMisbehaviour to it.
func DeployContractMisbehaviour(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ContractMisbehaviour, error) {
	parsed, err := ContractMisbehaviourMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractMisbehaviourBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractMisbehaviour{ContractMisbehaviourCaller: ContractMisbehaviourCaller{contract: contract}, ContractMisbehaviourTransactor: ContractMisbehaviourTransactor{contract: contract}, ContractMisbehaviourFilterer: ContractMisbehaviourFilterer{contract: contract}}, nil
}

// ContractMisbehaviour is an auto generated Go binding around an Ethereum contract.
type ContractMisbehaviour struct {
	ContractMisbehaviourCaller     // Read-only binding to the contract
	ContractMisbehaviourTransactor // Write-only binding to the contract
	ContractMisbehaviourFilterer   // Log filterer for contract events
}

// ContractMisbehaviourCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractMisbehaviourCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractMisbehaviourTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractMisbehaviourTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractMisbehaviourFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractMisbehaviourFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractMisbehaviourSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractMisbehaviourSession struct {
	Contract     *ContractMisbehaviour // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ContractMisbehaviourCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractMisbehaviourCallerSession struct {
	Contract *ContractMisbehaviourCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// ContractMisbehaviourTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractMisbehaviourTransactorSession struct {
	Contract     *ContractMisbehaviourTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// ContractMisbehaviourRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractMisbehaviourRaw struct {
	Contract *ContractMisbehaviour // Generic contract binding to access the raw methods on
}

// ContractMisbehaviourCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractMisbehaviourCallerRaw struct {
	Contract *ContractMisbehaviourCaller // Generic read-only contract binding to access the raw methods on
}

// ContractMisbehaviourTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractMisbehaviourTransactorRaw struct {
	Contract *ContractMisbehaviourTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractMisbehaviour creates a new instance of ContractMisbehaviour, bound to a specific deployed contract.
func NewContractMisbehaviour(address common.Address, backend bind.ContractBackend) (*ContractMisbehaviour, error) {
	contract, err := bindContractMisbehaviour(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractMisbehaviour{ContractMisbehaviourCaller: ContractMisbehaviourCaller{contract: contract}, ContractMisbehaviourTransactor: ContractMisbehaviourTransactor{contract: contract}, ContractMisbehaviourFilterer: ContractMisbehaviourFilterer{contract: contract}}, nil
}

// NewContractMisbehaviourCaller creates a new read-only instance of ContractMisbehaviour, bound to a specific deployed contract.
func NewContractMisbehaviourCaller(address common.Address, caller bind.ContractCaller) (*ContractMisbehaviourCaller, error) {
	contract, err := bindContractMisbehaviour(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractMisbehaviourCaller{contract: contract}, nil
}

// NewContractMisbehaviourTransactor creates a new write-only instance of ContractMisbehaviour, bound to a specific deployed contract.
func NewContractMisbehaviourTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractMisbehaviourTransactor, error) {
	contract, err := bindContractMisbehaviour(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractMisbehaviourTransactor{contract: contract}, nil
}

// NewContractMisbehaviourFilterer creates a new log filterer instance of ContractMisbehaviour, bound to a specific deployed contract.
func NewContractMisbehaviourFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractMisbehaviourFilterer, error) {
	contract, err := bindContractMisbehaviour(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractMisbehaviourFilterer{contract: contract}, nil
}

// bindContractMisbehaviour binds a generic wrapper to an already deployed contract.
func bindContractMisbehaviour(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractMisbehaviourMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractMisbehaviour *ContractMisbehaviourRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractMisbehaviour.Contract.ContractMisbehaviourCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractMisbehaviour *ContractMisbehaviourRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractMisbehaviour.Contract.ContractMisbehaviourTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractMisbehaviour *ContractMisbehaviourRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractMisbehaviour.Contract.ContractMisbehaviourTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractMisbehaviour *ContractMisbehaviourCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractMisbehaviour.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractMisbehaviour *ContractMisbehaviourTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractMisbehaviour.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractMisbehaviour *ContractMisbehaviourTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractMisbehaviour.Contract.contract.Transact(opts, method, params...)
}

// CheckForMisbehaviour is a paid mutator transaction binding the contract method 0xabb60742.
//
// Solidity: function checkForMisbehaviour((string,(uint8,uint8),(uint64,uint64),uint32,uint32,bool,uint8) clientState, ((string,uint64),((((uint64,uint64),string,uint64,uint128,bool,(bytes32,(uint32,bytes32)),bool,bytes32,bool,bytes32,bytes32,bytes32,bytes32,bytes,bool,bytes32,bool,bytes32,bytes),(uint64,uint32,(bytes32,(uint32,bytes32)),(uint8,(bytes,uint128,bool,bytes))[])),((bytes,bytes32,uint64,int64)[],bool,(bytes,bytes32,uint64,int64),uint64),(uint64,uint64),((bytes,bytes32,uint64,int64)[],bool,(bytes,bytes32,uint64,int64),uint64)),((((uint64,uint64),string,uint64,uint128,bool,(bytes32,(uint32,bytes32)),bool,bytes32,bool,bytes32,bytes32,bytes32,bytes32,bytes,bool,bytes32,bool,bytes32,bytes),(uint64,uint32,(bytes32,(uint32,bytes32)),(uint8,(bytes,uint128,bool,bytes))[])),((bytes,bytes32,uint64,int64)[],bool,(bytes,bytes32,uint64,int64),uint64),(uint64,uint64),((bytes,bytes32,uint64,int64)[],bool,(bytes,bytes32,uint64,int64),uint64))) misbehaviour, (uint128,bytes32,bytes32) trustedConsensusState1, (uint128,bytes32,bytes32) trustedConsensusState2, uint128 time) returns(((string,(uint8,uint8),(uint64,uint64),uint32,uint32,bool,uint8),uint128,(uint64,uint64),(uint64,uint64),(uint128,bytes32,bytes32),(uint128,bytes32,bytes32)) output)
func (_ContractMisbehaviour *ContractMisbehaviourTransactor) CheckForMisbehaviour(opts *bind.TransactOpts, clientState IICS07TendermintMsgsClientState, misbehaviour IMisbehaviourMsgsMisbehaviour, trustedConsensusState1 IICS07TendermintMsgsConsensusState, trustedConsensusState2 IICS07TendermintMsgsConsensusState, time *big.Int) (*types.Transaction, error) {
	return _ContractMisbehaviour.contract.Transact(opts, "checkForMisbehaviour", clientState, misbehaviour, trustedConsensusState1, trustedConsensusState2, time)
}

// CheckForMisbehaviour is a paid mutator transaction binding the contract method 0xabb60742.
//
// Solidity: function checkForMisbehaviour((string,(uint8,uint8),(uint64,uint64),uint32,uint32,bool,uint8) clientState, ((string,uint64),((((uint64,uint64),string,uint64,uint128,bool,(bytes32,(uint32,bytes32)),bool,bytes32,bool,bytes32,bytes32,bytes32,bytes32,bytes,bool,bytes32,bool,bytes32,bytes),(uint64,uint32,(bytes32,(uint32,bytes32)),(uint8,(bytes,uint128,bool,bytes))[])),((bytes,bytes32,uint64,int64)[],bool,(bytes,bytes32,uint64,int64),uint64),(uint64,uint64),((bytes,bytes32,uint64,int64)[],bool,(bytes,bytes32,uint64,int64),uint64)),((((uint64,uint64),string,uint64,uint128,bool,(bytes32,(uint32,bytes32)),bool,bytes32,bool,bytes32,bytes32,bytes32,bytes32,bytes,bool,bytes32,bool,bytes32,bytes),(uint64,uint32,(bytes32,(uint32,bytes32)),(uint8,(bytes,uint128,bool,bytes))[])),((bytes,bytes32,uint64,int64)[],bool,(bytes,bytes32,uint64,int64),uint64),(uint64,uint64),((bytes,bytes32,uint64,int64)[],bool,(bytes,bytes32,uint64,int64),uint64))) misbehaviour, (uint128,bytes32,bytes32) trustedConsensusState1, (uint128,bytes32,bytes32) trustedConsensusState2, uint128 time) returns(((string,(uint8,uint8),(uint64,uint64),uint32,uint32,bool,uint8),uint128,(uint64,uint64),(uint64,uint64),(uint128,bytes32,bytes32),(uint128,bytes32,bytes32)) output)
func (_ContractMisbehaviour *ContractMisbehaviourSession) CheckForMisbehaviour(clientState IICS07TendermintMsgsClientState, misbehaviour IMisbehaviourMsgsMisbehaviour, trustedConsensusState1 IICS07TendermintMsgsConsensusState, trustedConsensusState2 IICS07TendermintMsgsConsensusState, time *big.Int) (*types.Transaction, error) {
	return _ContractMisbehaviour.Contract.CheckForMisbehaviour(&_ContractMisbehaviour.TransactOpts, clientState, misbehaviour, trustedConsensusState1, trustedConsensusState2, time)
}

// CheckForMisbehaviour is a paid mutator transaction binding the contract method 0xabb60742.
//
// Solidity: function checkForMisbehaviour((string,(uint8,uint8),(uint64,uint64),uint32,uint32,bool,uint8) clientState, ((string,uint64),((((uint64,uint64),string,uint64,uint128,bool,(bytes32,(uint32,bytes32)),bool,bytes32,bool,bytes32,bytes32,bytes32,bytes32,bytes,bool,bytes32,bool,bytes32,bytes),(uint64,uint32,(bytes32,(uint32,bytes32)),(uint8,(bytes,uint128,bool,bytes))[])),((bytes,bytes32,uint64,int64)[],bool,(bytes,bytes32,uint64,int64),uint64),(uint64,uint64),((bytes,bytes32,uint64,int64)[],bool,(bytes,bytes32,uint64,int64),uint64)),((((uint64,uint64),string,uint64,uint128,bool,(bytes32,(uint32,bytes32)),bool,bytes32,bool,bytes32,bytes32,bytes32,bytes32,bytes,bool,bytes32,bool,bytes32,bytes),(uint64,uint32,(bytes32,(uint32,bytes32)),(uint8,(bytes,uint128,bool,bytes))[])),((bytes,bytes32,uint64,int64)[],bool,(bytes,bytes32,uint64,int64),uint64),(uint64,uint64),((bytes,bytes32,uint64,int64)[],bool,(bytes,bytes32,uint64,int64),uint64))) misbehaviour, (uint128,bytes32,bytes32) trustedConsensusState1, (uint128,bytes32,bytes32) trustedConsensusState2, uint128 time) returns(((string,(uint8,uint8),(uint64,uint64),uint32,uint32,bool,uint8),uint128,(uint64,uint64),(uint64,uint64),(uint128,bytes32,bytes32),(uint128,bytes32,bytes32)) output)
func (_ContractMisbehaviour *ContractMisbehaviourTransactorSession) CheckForMisbehaviour(clientState IICS07TendermintMsgsClientState, misbehaviour IMisbehaviourMsgsMisbehaviour, trustedConsensusState1 IICS07TendermintMsgsConsensusState, trustedConsensusState2 IICS07TendermintMsgsConsensusState, time *big.Int) (*types.Transaction, error) {
	return _ContractMisbehaviour.Contract.CheckForMisbehaviour(&_ContractMisbehaviour.TransactOpts, clientState, misbehaviour, trustedConsensusState1, trustedConsensusState2, time)
}
