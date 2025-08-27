// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractUpdateClient

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

// IUpdateClientMsgsUpdateClientOutput is an auto generated low-level Go binding around an user-defined struct.
type IUpdateClientMsgsUpdateClientOutput struct {
	ClientState           IICS07TendermintMsgsClientState
	TrustedConsensusState IICS07TendermintMsgsConsensusState
	NewConsensusState     IICS07TendermintMsgsConsensusState
	Time                  *big.Int
	TrustedHeight         IICS02ClientMsgsHeight
	NewHeight             IICS02ClientMsgsHeight
}

// ContractUpdateClientMetaData contains all meta data concerning the ContractUpdateClient contract.
var ContractUpdateClientMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"updateClient\",\"inputs\":[{\"name\":\"clientState\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.ClientState\",\"components\":[{\"name\":\"chainId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"trustLevel\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.TrustThreshold\",\"components\":[{\"name\":\"numerator\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"denominator\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"name\":\"latestHeight\",\"type\":\"tuple\",\"internalType\":\"structIICS02ClientMsgs.Height\",\"components\":[{\"name\":\"revisionNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revisionHeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"trustingPeriod\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"unbondingPeriod\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isFrozen\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"zkAlgorithm\",\"type\":\"uint8\",\"internalType\":\"enumIICS07TendermintMsgs.SupportedZkAlgorithm\"}]},{\"name\":\"trustedConsensusState\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.ConsensusState\",\"components\":[{\"name\":\"timestamp\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"nextValidatorsHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"proposedHeader\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.Header\",\"components\":[{\"name\":\"signedHeader\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.SignedHeader\",\"components\":[{\"name\":\"header\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.BlockHeader\",\"components\":[{\"name\":\"version\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.Version\",\"components\":[{\"name\":\"blockVersion\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"appVersion\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"chainId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"height\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"time\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"hasLastBlockId\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"lastBlockId\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.BlockId\",\"components\":[{\"name\":\"hashData\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"partSetHeader\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.PartSetHeader\",\"components\":[{\"name\":\"total\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashData\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]},{\"name\":\"hasLastCommitHash\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"lastCommitHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"hasDataHash\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"dataHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"validatorsHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"nextValidatorsHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"consensusHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"appHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"hasLastResultsHash\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"lastResultsHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"hasEvidenceHash\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"evidenceHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"proposerAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"commit\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.BlockCommit\",\"components\":[{\"name\":\"height\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"round\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"blockId\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.BlockId\",\"components\":[{\"name\":\"hashData\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"partSetHeader\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.PartSetHeader\",\"components\":[{\"name\":\"total\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashData\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]},{\"name\":\"commitSigs\",\"type\":\"tuple[]\",\"internalType\":\"structIICS07TendermintMsgs.CommitSig[]\",\"components\":[{\"name\":\"flag\",\"type\":\"uint8\",\"internalType\":\"enumIICS07TendermintMsgs.CommitSigFlag\"},{\"name\":\"data\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.CommitSigData\",\"components\":[{\"name\":\"validatorAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"timestamp\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"hasSignature\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]}]}]},{\"name\":\"validatorSet\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.ValidatorSet\",\"components\":[{\"name\":\"validators\",\"type\":\"tuple[]\",\"internalType\":\"structIICS07TendermintMsgs.ValidatorInfo[]\",\"components\":[{\"name\":\"valAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"pubKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"votingPower\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"proposerPriority\",\"type\":\"int64\",\"internalType\":\"int64\"}]},{\"name\":\"hasProposer\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"proposer\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.ValidatorInfo\",\"components\":[{\"name\":\"valAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"pubKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"votingPower\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"proposerPriority\",\"type\":\"int64\",\"internalType\":\"int64\"}]},{\"name\":\"totalVotingPower\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"trustedHeight\",\"type\":\"tuple\",\"internalType\":\"structIICS02ClientMsgs.Height\",\"components\":[{\"name\":\"revisionNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revisionHeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"trustedNextValidatorSet\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.ValidatorSet\",\"components\":[{\"name\":\"validators\",\"type\":\"tuple[]\",\"internalType\":\"structIICS07TendermintMsgs.ValidatorInfo[]\",\"components\":[{\"name\":\"valAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"pubKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"votingPower\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"proposerPriority\",\"type\":\"int64\",\"internalType\":\"int64\"}]},{\"name\":\"hasProposer\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"proposer\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.ValidatorInfo\",\"components\":[{\"name\":\"valAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"pubKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"votingPower\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"proposerPriority\",\"type\":\"int64\",\"internalType\":\"int64\"}]},{\"name\":\"totalVotingPower\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}]},{\"name\":\"time\",\"type\":\"uint128\",\"internalType\":\"uint128\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIUpdateClientMsgs.UpdateClientOutput\",\"components\":[{\"name\":\"clientState\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.ClientState\",\"components\":[{\"name\":\"chainId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"trustLevel\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.TrustThreshold\",\"components\":[{\"name\":\"numerator\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"denominator\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"name\":\"latestHeight\",\"type\":\"tuple\",\"internalType\":\"structIICS02ClientMsgs.Height\",\"components\":[{\"name\":\"revisionNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revisionHeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"trustingPeriod\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"unbondingPeriod\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isFrozen\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"zkAlgorithm\",\"type\":\"uint8\",\"internalType\":\"enumIICS07TendermintMsgs.SupportedZkAlgorithm\"}]},{\"name\":\"trustedConsensusState\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.ConsensusState\",\"components\":[{\"name\":\"timestamp\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"nextValidatorsHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"newConsensusState\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.ConsensusState\",\"components\":[{\"name\":\"timestamp\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"nextValidatorsHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"time\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"trustedHeight\",\"type\":\"tuple\",\"internalType\":\"structIICS02ClientMsgs.Height\",\"components\":[{\"name\":\"revisionNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revisionHeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"newHeight\",\"type\":\"tuple\",\"internalType\":\"structIICS02ClientMsgs.Height\",\"components\":[{\"name\":\"revisionNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revisionHeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}]}],\"stateMutability\":\"view\"},{\"type\":\"error\",\"name\":\"FailedToVerifyHeader\",\"inputs\":[{\"name\":\"reason\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"HeaderChainIdMismatch\",\"inputs\":[{\"name\":\"expected\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"actual\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"InvalidHeaderHeight\",\"inputs\":[{\"name\":\"height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"MismatchedRevisionHeight\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"actual\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ValSetHashMismatch\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"actual\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]",
	Bin: "0x6080806040523460b9575f54600181811c9116801560b0575b6020821014609c57601f81116058575b507f30372d74656e6465726d696e742d30000000000000000000000000000000001e5f5561265f90816100be8239f35b5f8052601f0160051c7f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e563908101905b818110609257506028565b5f81556001016087565b634e487b7160e01b5f52602260045260245ffd5b90607f16906018565b5f80fdfe6080806040526004361015610012575f80fd5b5f3560e01c63cb1e26cf14610025575f80fd5b346111415760c07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126111415767ffffffffffffffff600435116111415760043536036101207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8201126111415760607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc360112611141576084359167ffffffffffffffff83116111415760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc84360301126111415760a435916fffffffffffffffffffffffffffffffff8316830361114157816101467fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdd93611a0b565b60405161015281611a27565b6060815260405161016281611a7b565b5f81525f6020820152602082015260405161017c81611a7b565b5f81525f602082015260408201525f60608201525f60808201525f60a08201525f60c082015281526101ac611aba565b60208201526101b9611aba565b60408201525f60608201526040516101d081611a7b565b5f81525f6020820152608082015260a0604051916101ed83611a7b565b5f83525f602084015201520160043560040135121561114157600435600401356004350190600482013567ffffffffffffffff811161114157602483019080360382136111415761024891610243913691611af4565b612029565b9260a460043501359263ffffffff8416809403611141576040519361026c85611a43565b61027b36602460043501611b38565b85526020850152600f604085015261029560448301611b6e565b604067ffffffffffffffff6102ac60648601611b6e565b818351946102b986611a43565b6102c1611b83565b865216602085015216910152604051926102da84611a5f565b826004013567ffffffffffffffff811161114157830160407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8236030112611141576040519061032982611a7b565b600481013567ffffffffffffffff8111611141576004908201016102c0813603126111415760405190610260820182811067ffffffffffffffff82111761167a576040526103773682611cae565b8252604081013567ffffffffffffffff81116111415761039a9036908301611cdf565b60208301526103ab60608201611c99565b60408301526103bc608082016119c9565b60608301526103cd60a08201611cfd565b60808301526103df3660c08301611d1b565b60a08301526103f16101208201611cfd565b60c083015261014081013560e083015261040e6101608201611cfd565b6101008301526101808101356101208301526101a08101356101408301526101c08101356101608301526101e08101356101808301526102008101356101a083015261045d6102208201611cfd565b6101c08301526102408101356101e083015261047c6102608201611cfd565b6102008301526102808101356102208301526102a08101359067ffffffffffffffff8211611141576104b091369101611cdf565b6102408201528252602481013567ffffffffffffffff811161114157600491010160c08136031261114157604051906104e882611a5f565b6104f181611c99565b82526104ff60208201611d0a565b60208301526105113660408301611d1b565b604083015260a08101359067ffffffffffffffff8211611141570136601f8201121561114157803561054281611d6d565b916105506040519384611a97565b81835260208084019260051b820101903682116111415760208101925b8284106118dd5750505050606082015260208201528452602483013567ffffffffffffffff8111611141576105a89060043691860101611def565b60208501526105ba3660448501611cae565b6040850152608483013567ffffffffffffffff8111611141576105e39060043691860101611def565b60608501526105f0611b83565b506105fa36611edf565b95602061060b818751510151612029565b0167ffffffffffffffff81511667ffffffffffffffff60408801515116908181036118af57505067ffffffffffffffff90511661067367ffffffffffffffff6040885151015116916040519061066082611a7b565b81526020810192835260408801516125d2565b61067c816124dc565b6002811490811561189b575b5061186457506106ca602080870151604051809381927fc87f1f6900000000000000000000000000000000000000000000000000000000835260048301612395565b038173__$7440a880b7578767f72184d998805816e4$__5af49081156115d3575f91611832575b50610140865151015180820361180457505061071260208651510151612029565b67ffffffffffffffff602081818501511692015116036117865761076b959660206060870151604051809981927fc87f1f6900000000000000000000000000000000000000000000000000000000835260048301612395565b038173__$7440a880b7578767f72184d998805816e4$__5af49687156115d3575f97611751575b506040015186036116a75780519567ffffffffffffffff60206040880151015116606087015190604051988960a081011067ffffffffffffffff60a08c01111761167a5760a08a0160405289526fffffffffffffffffffffffffffffffff851660208a015260408901526060880152608087015261085860208087519701516040519761081e89611a7b565b88528082890152604051809381927fc87f1f6900000000000000000000000000000000000000000000000000000000835260048301612395565b038173__$7440a880b7578767f72184d998805816e4$__5af49081156115d3575f91611648575b506101408651510151036115de57610a316020865151604051809381927f95b25f790000000000000000000000000000000000000000000000000000000083528460048401526108e960248401825167ffffffffffffffff60208092828151168552015116910152565b610240610906868301516102c060648701526102e48601906119e6565b9167ffffffffffffffff60408201511660848601526fffffffffffffffffffffffffffffffff60608201511660a48601526080810151151560c4860152868060a0830151805160e4890152015163ffffffff815116610104880152015161012486015260c0810151151561014486015260e081015161016486015261010081015115156101848601526101208101516101a48601526101408101516101c48601526101608101516101e48601526101808101516102048601526101a08101516102248601526101c081015115156102448601526101e081015161026486015261020081015115156102848601526102208101516102a486015201517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc848303016102c48501526119e6565b038173__$7440a880b7578767f72184d998805816e4$__5af49081156115d3575f916115a1575b506040602087510151015151036115375784519360606020808801519601510151955f975f5b8851811015610abb57610a91818a61263e565b5151610a9c816124dc565b610aa5816124dc565b610ab2575b600101610a7e565b60019950610aaa565b508998979697156114cd5786518851510361143d575f5b8751811015610c1857610ae5818961263e565b518051610af1816124dc565b610afa816124dc565b610b0957506001905b01610ad2565b6020909a95919a989498979397969296015151975f975f5b8b518051821015610c045781610b369161263e565b5151604051610b646020828180820195805191829101875e81015f838201520301601f198101835282611a97565b5190208b604051610b946020828180820195805191829101875e81015f838201520301601f198101835282611a97565b51902014610ba457600101610b21565b50939750939750939860019591955b15610bc057600190610b03565b606460405162461bcd60e51b815260206004820152601d60248201527f696e76616c696420636f6d6d69743a206661756c7479207369676e65720000006044820152fd5b505093979094985099909499959195610bb3565b508863ffffffff6020820151166fffffffffffffffffffffffffffffffff60208401511690816fffffffffffffffffffffffffffffffff8916109182156113f5575b505061138b576fffffffffffffffffffffffffffffffff60608451510151166fffffffffffffffffffffffffffffffff60208401511610156113215760208351510151604051610cc96020828180820195805191829101875e81015f838201520301601f198101835282611a97565b5190208251604051610cfa6020828180820195805191829101875e81015f838201520301601f198101835282611a97565b519020036112dd57610d1967ffffffffffffffff6040840151166124e6565b8351516040015167ffffffffffffffff918216911681810361126157505061014083515101516080830151036111f757604063ffffffff915b0151166fffffffffffffffffffffffffffffffff8616016fffffffffffffffffffffffffffffffff81116111ca576fffffffffffffffffffffffffffffffff806060855151015116911611156111605767ffffffffffffffff6040610dbc828281950151166124e6565b935151015116911614611145575b6080610de2610ddc6004840180611f47565b80611f7a565b0135916fffffffffffffffffffffffffffffffff83168093036111415767ffffffffffffffff90602090610200610e1f610ddc6004870180611f47565b01356101c0610e34610ddc6004880180611f47565b01359060405196610e4488611a43565b875283870152604086015201511667ffffffffffffffff610e766060610e70610ddc6004870180611f47565b01611b6e565b60405192610e8384611a7b565b835216602082015260405194610e9886611a0b565b60405190610ea582611a27565b67ffffffffffffffff600435600401351161114157610ec8906004369101611cdf565b8152610ed936602460043501611b38565b6020820152610eed36606460043501611cae565b6040820152610f0060a460043501611d0a565b6060820152610f1360c460043501611d0a565b6080820152610f2660e460043501611cfd565b60a0820152610104600435013560028110156111415760c08201528552610f81610f4f36611edf565b9260208701938452604087019485526fffffffffffffffffffffffffffffffff60608801961686526044369101611cae565b926080860193845260a0860191825260405195602087525193610180602088015260c0610fbd86516101206101a08b01526102c08a01906119e6565b9560ff602080830151828151166101c08d01520151166101e08a015261100260408201516102008b019067ffffffffffffffff60208092828151168552015116910152565b63ffffffff6060820151166102408a015263ffffffff6080820151166102608a015260a081015115156102808a0152015160028110156111145787966110c46110ee94611095611110986fffffffffffffffffffffffffffffffff956102a08d01525160408c0190604080916fffffffffffffffffffffffffffffffff8151168452602081015160208501520151910152565b5180516fffffffffffffffffffffffffffffffff1660a08b0152602081015160c08b01526040015160e08a0152565b51166101008701525161012086019067ffffffffffffffff60208092828151168552015116910152565b5161016084019067ffffffffffffffff60208092828151168552015116910152565b0390f35b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b5f80fd5b6003602060405161115581611a7b565b600281520152610dca565b608460405162461bcd60e51b815260206004820152602860248201527f696e76616c696420626c6f636b3a206865616465722069732066726f6d20746860448201527f65206675747572650000000000000000000000000000000000000000000000006064820152fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b608460405162461bcd60e51b815260206004820152602f60248201527f696e76616c696420626c6f636b3a206e6578742076616c696461746f7220736560448201527f742068617368206d69736d6174636800000000000000000000000000000000006064820152fd5b111561127457604063ffffffff91610d52565b608460405162461bcd60e51b8152602060048201526024808201527f696e76616c696420626c6f636b3a206e6f6e20696e6372656173696e6720686560448201527f69676874000000000000000000000000000000000000000000000000000000006064820152fd5b606460405162461bcd60e51b815260206004820152602060248201527f696e76616c696420626c6f636b3a20636861696e2d6964206d69736d617463686044820152fd5b608460405162461bcd60e51b815260206004820152602560248201527f696e76616c696420626c6f636b3a206e6f6e206d6f6e6f746f6e69632062667460448201527f2074696d650000000000000000000000000000000000000000000000000000006064820152fd5b608460405162461bcd60e51b815260206004820152603c60248201527f696e76616c696420626c6f636b3a20756e74727573746564207374617465206960448201527f73206f757473696465206f66207472757374696e6720706572696f64000000006064820152fd5b9091506fffffffffffffffffffffffffffffffff8816036fffffffffffffffffffffffffffffffff81116111ca576fffffffffffffffffffffffffffffffff16118880610c5a565b60a460405162461bcd60e51b815260206004820152604860248201527f696e76616c696420636f6d6d69743a206e756d626572206f66207369676e617460448201527f7572657320646f6573206e6f74206d61746368206e756d626572206f6620766160648201527f6c696461746f72730000000000000000000000000000000000000000000000006084820152fd5b608460405162461bcd60e51b815260206004820152602560248201527f696e76616c696420636f6d6d69743a206e6f2070726573656e74207369676e6160448201527f74757265730000000000000000000000000000000000000000000000000000006064820152fd5b608460405162461bcd60e51b815260206004820152602360248201527f696e76616c696420626c6f636b3a206865616465722068617368206d69736d6160448201527f74636800000000000000000000000000000000000000000000000000000000006064820152fd5b90506020813d6020116115cb575b816115bc60209383611a97565b8101031261114157515f610a58565b3d91506115af565b6040513d5f823e3d90fd5b608460405162461bcd60e51b815260206004820152602a60248201527f696e76616c696420626c6f636b3a2076616c696461746f72207365742068617360448201527f68206d69736d61746368000000000000000000000000000000000000000000006064820152fd5b90506020813d602011611672575b8161166360209383611a97565b8101031261114157515f61087f565b3d9150611656565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b60a46040517ff492ef2b00000000000000000000000000000000000000000000000000000000815260206004820152604360248201527f74727573746564206e6578742076616c696461746f722073657420686173682060448201527f646f6573206e6f74206d6174636820686173682073746f726564206f6e20636860648201527f61696e00000000000000000000000000000000000000000000000000000000006084820152fd5b9096506020813d60201161177e575b8161176d60209383611a97565b810103126111415751956040610792565b3d9150611760565b6117d0906118006020875151015191516040519384937fc6913e9f0000000000000000000000000000000000000000000000000000000085526040600486015260448501906119e6565b907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8483030160248501526119e6565b0390fd5b7f4d6d8014000000000000000000000000000000000000000000000000000000005f5260045260245260445ffd5b90506020813d60201161185c575b8161184d60209383611a97565b8101031261114157515f6106f1565b3d9150611840565b67ffffffffffffffff9051167f90f4dbed000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b600191506118a8816124dc565b145f610688565b7f05b4c7a3000000000000000000000000000000000000000000000000000000005f5260045260245260445ffd5b833567ffffffffffffffff8111611141578201906040601f198336030112611141576040519161190c83611a7b565b60208101356003811015611141578352604081013567ffffffffffffffff811161114157602091010190608082360312611141576040519261194d84611a5f565b823567ffffffffffffffff81116111415761196b9036908501611cdf565b8452611979602084016119c9565b602085015261198a60408401611cfd565b604085015260608301359367ffffffffffffffff8511611141576119b5602095948695369101611cdf565b60608201528382015281520193019261056d565b35906fffffffffffffffffffffffffffffffff8216820361114157565b90601f19601f602080948051918291828752018686015e5f8582860101520116010190565b60c0810190811067ffffffffffffffff82111761167a57604052565b60e0810190811067ffffffffffffffff82111761167a57604052565b6060810190811067ffffffffffffffff82111761167a57604052565b6080810190811067ffffffffffffffff82111761167a57604052565b6040810190811067ffffffffffffffff82111761167a57604052565b90601f601f19910116810190811067ffffffffffffffff82111761167a57604052565b60405190611ac782611a43565b5f6040838281528260208201520152565b67ffffffffffffffff811161167a57601f01601f191660200190565b929192611b0082611ad8565b91611b0e6040519384611a97565b829481845281830111611141578281602093845f960137010152565b359060ff8216820361114157565b919082604091031261114157604051611b5081611a7b565b6020611b69818395611b6181611b2a565b855201611b2a565b910152565b3567ffffffffffffffff811681036111415790565b604051905f5f548060011c9160018216918215611c8f575b602084108314611c62578386528592908115611c255750600114611bc8575b611bc692500383611a97565b565b505f80805290917f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e5635b818310611c09575050906020611bc692820101611bba565b6020919350806001915483858901015201910190918492611bf1565b60209250611bc69491507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001682840152151560051b820101611bba565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b92607f1692611b9b565b359067ffffffffffffffff8216820361114157565b919082604091031261114157604051611cc681611a7b565b6020611b69818395611cd781611c99565b855201611c99565b9080601f8301121561114157816020611cfa93359101611af4565b90565b3590811515820361114157565b359063ffffffff8216820361114157565b80929103916060831261114157604051611d3481611a7b565b6040601f198295843584520112611141576020906040805193611d5685611a7b565b611d61848201611d0a565b85520135828401520152565b67ffffffffffffffff811161167a5760051b60200190565b91906080838203126111415760405190611d9e82611a5f565b8193803567ffffffffffffffff811161114157606092611dbf918301611cdf565b835260208101356020840152611dd760408201611c99565b60408401520135908160070b82036111415760600152565b9190916080818403126111415760405190611e0982611a5f565b8193813567ffffffffffffffff811161114157820181601f82011215611141578035611e3481611d6d565b91611e426040519384611a97565b81835260208084019260051b820101918483116111415760208201905b838210611eb157505050508352611e7860208301611cfd565b602084015260408201359067ffffffffffffffff82116111415782611ea660609492611b6994869401611d85565b604086015201611c99565b813567ffffffffffffffff811161114157602091611ed488848094880101611d85565b815201910190611e5f565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc60609101126111415760405190611f1682611a43565b816024356fffffffffffffffffffffffffffffffff8116810361114157815260443560208201526040606435910152565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc181360301821215611141570190565b9035907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffd4181360301821215611141570190565b908151811015611fbe570160200190565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b90611ff582611ad8565b6120026040519182611a97565b828152601f196120128294611ad8565b0190602036910137565b919082018092116111ca57565b60405161203581611a7b565b606081525f60209091015280515f1981019081116111ca575b7f2d000000000000000000000000000000000000000000000000000000000000007fff0000000000000000000000000000000000000000000000000000000000000061209a8385611fad565b5116146120af5780156111ca575f190161204e565b91905f1983146123315780518381039081116111ca575f1981019081116111ca576120d990611feb565b905f5b825181101561213c5760018501908186116111ca577fff000000000000000000000000000000000000000000000000000000000000006121276121218360019561201c565b85611fad565b51165f1a6121358286611fad565b53016120dc565b5091929080511580156122d3575b61224457805161215991612452565b91901580156122c2575b6122445761217081611feb565b905f5b8181106122015750505160018111908115916121f5575b506121b15767ffffffffffffffff90604051926121a684611a7b565b835216602082015290565b606460405162461bcd60e51b815260206004820152601b60248201527f496e76616c696420636861696e20707265666978206c656e67746800000000006044820152fd5b602b915010155f61218a565b807fff0000000000000000000000000000000000000000000000000000000000000061222f60019388611fad565b51165f1a61223d8286611fad565b5301612173565b50508051806001111590816122b7575b5015612272576040519061226782611a7b565b81525f602082015290565b60405162461bcd60e51b815260206004820152601760248201527f496e76616c696420636861696e204944206c656e6774680000000000000000006044820152606490fd5b60409150105f612254565b5067ffffffffffffffff8211612163565b611fbe577f30000000000000000000000000000000000000000000000000000000000000007fff0000000000000000000000000000000000000000000000000000000000000060208301511614801561214a5750600181511161214a565b8091925051806001111590816122b7575015612272576040519061226782611a7b565b9060608061236b84516080855260808501906119e6565b936020810151602085015267ffffffffffffffff6040820151166040850152015160070b91015290565b6020815260a0810182519060806020840152815180915260c0830190602060c08260051b8601019301915f905b828210612409575050505067ffffffffffffffff60606123ff6080936020870151151560408701526040870151601f198783030184880152612354565b9401511691015290565b90919293602080612444837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff408a600196030186528851612354565b9601920192019092916123c2565b5f929183915b8183106124685750505060019190565b90919360ff61249e7fff000000000000000000000000000000000000000000000000000000000000006020888601015116612506565b1690600982116124d057600a810290808204600a14901517156111ca576001916124c79161201c565b94019190612458565b5050505090505f905f90565b6003111561111457565b67ffffffffffffffff60019116019067ffffffffffffffff82116111ca57565b60f81c602f8111806125c8575b15612540577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffd00160ff1690565b60608111806125be575b15612577577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa90160ff1690565b60408111806125b4575b156125ae577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc90160ff1690565b5060ff90565b5060478110612581565b506067811061254a565b50603a8110612513565b67ffffffffffffffff81511667ffffffffffffffff835116908181105f146125fc57505050505f90565b1115612609575050600290565b602067ffffffffffffffff81819301511692015116908181105f1461262e5750505f90565b111561263957600290565b600190565b8051821015611fbe5760209160051b01019056fea164736f6c634300081c000a",
}

// ContractUpdateClientABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractUpdateClientMetaData.ABI instead.
var ContractUpdateClientABI = ContractUpdateClientMetaData.ABI

// ContractUpdateClientBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractUpdateClientMetaData.Bin instead.
var ContractUpdateClientBin = ContractUpdateClientMetaData.Bin

// DeployContractUpdateClient deploys a new Ethereum contract, binding an instance of ContractUpdateClient to it.
func DeployContractUpdateClient(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ContractUpdateClient, error) {
	parsed, err := ContractUpdateClientMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractUpdateClientBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractUpdateClient{ContractUpdateClientCaller: ContractUpdateClientCaller{contract: contract}, ContractUpdateClientTransactor: ContractUpdateClientTransactor{contract: contract}, ContractUpdateClientFilterer: ContractUpdateClientFilterer{contract: contract}}, nil
}

// ContractUpdateClient is an auto generated Go binding around an Ethereum contract.
type ContractUpdateClient struct {
	ContractUpdateClientCaller     // Read-only binding to the contract
	ContractUpdateClientTransactor // Write-only binding to the contract
	ContractUpdateClientFilterer   // Log filterer for contract events
}

// ContractUpdateClientCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractUpdateClientCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractUpdateClientTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractUpdateClientTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractUpdateClientFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractUpdateClientFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractUpdateClientSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractUpdateClientSession struct {
	Contract     *ContractUpdateClient // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ContractUpdateClientCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractUpdateClientCallerSession struct {
	Contract *ContractUpdateClientCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// ContractUpdateClientTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractUpdateClientTransactorSession struct {
	Contract     *ContractUpdateClientTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// ContractUpdateClientRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractUpdateClientRaw struct {
	Contract *ContractUpdateClient // Generic contract binding to access the raw methods on
}

// ContractUpdateClientCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractUpdateClientCallerRaw struct {
	Contract *ContractUpdateClientCaller // Generic read-only contract binding to access the raw methods on
}

// ContractUpdateClientTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractUpdateClientTransactorRaw struct {
	Contract *ContractUpdateClientTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractUpdateClient creates a new instance of ContractUpdateClient, bound to a specific deployed contract.
func NewContractUpdateClient(address common.Address, backend bind.ContractBackend) (*ContractUpdateClient, error) {
	contract, err := bindContractUpdateClient(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractUpdateClient{ContractUpdateClientCaller: ContractUpdateClientCaller{contract: contract}, ContractUpdateClientTransactor: ContractUpdateClientTransactor{contract: contract}, ContractUpdateClientFilterer: ContractUpdateClientFilterer{contract: contract}}, nil
}

// NewContractUpdateClientCaller creates a new read-only instance of ContractUpdateClient, bound to a specific deployed contract.
func NewContractUpdateClientCaller(address common.Address, caller bind.ContractCaller) (*ContractUpdateClientCaller, error) {
	contract, err := bindContractUpdateClient(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractUpdateClientCaller{contract: contract}, nil
}

// NewContractUpdateClientTransactor creates a new write-only instance of ContractUpdateClient, bound to a specific deployed contract.
func NewContractUpdateClientTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractUpdateClientTransactor, error) {
	contract, err := bindContractUpdateClient(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractUpdateClientTransactor{contract: contract}, nil
}

// NewContractUpdateClientFilterer creates a new log filterer instance of ContractUpdateClient, bound to a specific deployed contract.
func NewContractUpdateClientFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractUpdateClientFilterer, error) {
	contract, err := bindContractUpdateClient(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractUpdateClientFilterer{contract: contract}, nil
}

// bindContractUpdateClient binds a generic wrapper to an already deployed contract.
func bindContractUpdateClient(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractUpdateClientMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractUpdateClient *ContractUpdateClientRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractUpdateClient.Contract.ContractUpdateClientCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractUpdateClient *ContractUpdateClientRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractUpdateClient.Contract.ContractUpdateClientTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractUpdateClient *ContractUpdateClientRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractUpdateClient.Contract.ContractUpdateClientTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractUpdateClient *ContractUpdateClientCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractUpdateClient.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractUpdateClient *ContractUpdateClientTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractUpdateClient.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractUpdateClient *ContractUpdateClientTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractUpdateClient.Contract.contract.Transact(opts, method, params...)
}

// UpdateClient is a free data retrieval call binding the contract method 0xcb1e26cf.
//
// Solidity: function updateClient((string,(uint8,uint8),(uint64,uint64),uint32,uint32,bool,uint8) clientState, (uint128,bytes32,bytes32) trustedConsensusState, ((((uint64,uint64),string,uint64,uint128,bool,(bytes32,(uint32,bytes32)),bool,bytes32,bool,bytes32,bytes32,bytes32,bytes32,bytes32,bool,bytes32,bool,bytes32,bytes),(uint64,uint32,(bytes32,(uint32,bytes32)),(uint8,(bytes,uint128,bool,bytes))[])),((bytes,bytes32,uint64,int64)[],bool,(bytes,bytes32,uint64,int64),uint64),(uint64,uint64),((bytes,bytes32,uint64,int64)[],bool,(bytes,bytes32,uint64,int64),uint64)) proposedHeader, uint128 time) view returns(((string,(uint8,uint8),(uint64,uint64),uint32,uint32,bool,uint8),(uint128,bytes32,bytes32),(uint128,bytes32,bytes32),uint128,(uint64,uint64),(uint64,uint64)))
func (_ContractUpdateClient *ContractUpdateClientCaller) UpdateClient(opts *bind.CallOpts, clientState IICS07TendermintMsgsClientState, trustedConsensusState IICS07TendermintMsgsConsensusState, proposedHeader IICS07TendermintMsgsHeader, time *big.Int) (IUpdateClientMsgsUpdateClientOutput, error) {
	var out []interface{}
	err := _ContractUpdateClient.contract.Call(opts, &out, "updateClient", clientState, trustedConsensusState, proposedHeader, time)

	if err != nil {
		return *new(IUpdateClientMsgsUpdateClientOutput), err
	}

	out0 := *abi.ConvertType(out[0], new(IUpdateClientMsgsUpdateClientOutput)).(*IUpdateClientMsgsUpdateClientOutput)

	return out0, err

}

// UpdateClient is a free data retrieval call binding the contract method 0xcb1e26cf.
//
// Solidity: function updateClient((string,(uint8,uint8),(uint64,uint64),uint32,uint32,bool,uint8) clientState, (uint128,bytes32,bytes32) trustedConsensusState, ((((uint64,uint64),string,uint64,uint128,bool,(bytes32,(uint32,bytes32)),bool,bytes32,bool,bytes32,bytes32,bytes32,bytes32,bytes32,bool,bytes32,bool,bytes32,bytes),(uint64,uint32,(bytes32,(uint32,bytes32)),(uint8,(bytes,uint128,bool,bytes))[])),((bytes,bytes32,uint64,int64)[],bool,(bytes,bytes32,uint64,int64),uint64),(uint64,uint64),((bytes,bytes32,uint64,int64)[],bool,(bytes,bytes32,uint64,int64),uint64)) proposedHeader, uint128 time) view returns(((string,(uint8,uint8),(uint64,uint64),uint32,uint32,bool,uint8),(uint128,bytes32,bytes32),(uint128,bytes32,bytes32),uint128,(uint64,uint64),(uint64,uint64)))
func (_ContractUpdateClient *ContractUpdateClientSession) UpdateClient(clientState IICS07TendermintMsgsClientState, trustedConsensusState IICS07TendermintMsgsConsensusState, proposedHeader IICS07TendermintMsgsHeader, time *big.Int) (IUpdateClientMsgsUpdateClientOutput, error) {
	return _ContractUpdateClient.Contract.UpdateClient(&_ContractUpdateClient.CallOpts, clientState, trustedConsensusState, proposedHeader, time)
}

// UpdateClient is a free data retrieval call binding the contract method 0xcb1e26cf.
//
// Solidity: function updateClient((string,(uint8,uint8),(uint64,uint64),uint32,uint32,bool,uint8) clientState, (uint128,bytes32,bytes32) trustedConsensusState, ((((uint64,uint64),string,uint64,uint128,bool,(bytes32,(uint32,bytes32)),bool,bytes32,bool,bytes32,bytes32,bytes32,bytes32,bytes32,bool,bytes32,bool,bytes32,bytes),(uint64,uint32,(bytes32,(uint32,bytes32)),(uint8,(bytes,uint128,bool,bytes))[])),((bytes,bytes32,uint64,int64)[],bool,(bytes,bytes32,uint64,int64),uint64),(uint64,uint64),((bytes,bytes32,uint64,int64)[],bool,(bytes,bytes32,uint64,int64),uint64)) proposedHeader, uint128 time) view returns(((string,(uint8,uint8),(uint64,uint64),uint32,uint32,bool,uint8),(uint128,bytes32,bytes32),(uint128,bytes32,bytes32),uint128,(uint64,uint64),(uint64,uint64)))
func (_ContractUpdateClient *ContractUpdateClientCallerSession) UpdateClient(clientState IICS07TendermintMsgsClientState, trustedConsensusState IICS07TendermintMsgsConsensusState, proposedHeader IICS07TendermintMsgsHeader, time *big.Int) (IUpdateClientMsgsUpdateClientOutput, error) {
	return _ContractUpdateClient.Contract.UpdateClient(&_ContractUpdateClient.CallOpts, clientState, trustedConsensusState, proposedHeader, time)
}
