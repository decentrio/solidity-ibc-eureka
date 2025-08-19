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

// ContractUpdateClientMetaData contains all meta data concerning the ContractUpdateClient contract.
var ContractUpdateClientMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"updateClient\",\"inputs\":[{\"name\":\"clientState\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.ClientState\",\"components\":[{\"name\":\"chainId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"trustLevel\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.TrustThreshold\",\"components\":[{\"name\":\"numerator\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"denominator\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"name\":\"latestHeight\",\"type\":\"tuple\",\"internalType\":\"structIICS02ClientMsgs.Height\",\"components\":[{\"name\":\"revisionNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revisionHeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"trustingPeriod\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"unbondingPeriod\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isFrozen\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"zkAlgorithm\",\"type\":\"uint8\",\"internalType\":\"enumIICS07TendermintMsgs.SupportedZkAlgorithm\"}]},{\"name\":\"trustedConsensusState\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.ConsensusState\",\"components\":[{\"name\":\"timestamp\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"nextValidatorsHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"proposedHeader\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.Header\",\"components\":[{\"name\":\"signedHeader\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.SignedHeader\",\"components\":[{\"name\":\"header\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.BlockHeader\",\"components\":[{\"name\":\"version\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.Version\",\"components\":[{\"name\":\"blockVersion\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"appVersion\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"chainId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"height\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"time\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"hasLastBlockId\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"lastBlockId\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.BlockId\",\"components\":[{\"name\":\"hashData\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"partSetHeader\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.PartSetHeader\",\"components\":[{\"name\":\"total\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashData\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]},{\"name\":\"hasLastCommitHash\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"lastCommitHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"hasDataHash\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"dataHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"validatorsHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"nextValidatorsHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"consensusHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"appHash\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"hasLastResultsHash\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"lastResultsHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"hasEvidenceHash\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"evidenceHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"proposerAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"commit\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.BlockCommit\",\"components\":[{\"name\":\"height\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"round\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"blockId\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.BlockId\",\"components\":[{\"name\":\"hashData\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"partSetHeader\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.PartSetHeader\",\"components\":[{\"name\":\"total\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashData\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]},{\"name\":\"commitSigs\",\"type\":\"tuple[]\",\"internalType\":\"structIICS07TendermintMsgs.CommitSig[]\",\"components\":[{\"name\":\"flag\",\"type\":\"uint8\",\"internalType\":\"enumIICS07TendermintMsgs.CommitSigFlag\"},{\"name\":\"data\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.CommitSigData\",\"components\":[{\"name\":\"validatorAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"timestamp\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"hasSignature\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]}]}]},{\"name\":\"validatorSet\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.ValidatorSet\",\"components\":[{\"name\":\"validators\",\"type\":\"tuple[]\",\"internalType\":\"structIICS07TendermintMsgs.ValidatorInfo[]\",\"components\":[{\"name\":\"valAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"pubKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"votingPower\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"proposerPriority\",\"type\":\"int64\",\"internalType\":\"int64\"}]},{\"name\":\"hasProposer\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"proposer\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.ValidatorInfo\",\"components\":[{\"name\":\"valAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"pubKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"votingPower\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"proposerPriority\",\"type\":\"int64\",\"internalType\":\"int64\"}]},{\"name\":\"totalVotingPower\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"trustedHeight\",\"type\":\"tuple\",\"internalType\":\"structIICS02ClientMsgs.Height\",\"components\":[{\"name\":\"revisionNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revisionHeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"trustedNextValidatorSet\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.ValidatorSet\",\"components\":[{\"name\":\"validators\",\"type\":\"tuple[]\",\"internalType\":\"structIICS07TendermintMsgs.ValidatorInfo[]\",\"components\":[{\"name\":\"valAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"pubKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"votingPower\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"proposerPriority\",\"type\":\"int64\",\"internalType\":\"int64\"}]},{\"name\":\"hasProposer\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"proposer\",\"type\":\"tuple\",\"internalType\":\"structIICS07TendermintMsgs.ValidatorInfo\",\"components\":[{\"name\":\"valAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"pubKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"votingPower\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"proposerPriority\",\"type\":\"int64\",\"internalType\":\"int64\"}]},{\"name\":\"totalVotingPower\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}]},{\"name\":\"time\",\"type\":\"uint128\",\"internalType\":\"uint128\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"error\",\"name\":\"FailedToVerifyHeader\",\"inputs\":[{\"name\":\"reason\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"HeaderChainIdMismatch\",\"inputs\":[{\"name\":\"expected\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"actual\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"InvalidHeaderHeight\",\"inputs\":[{\"name\":\"height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"MismatchedRevisionHeight\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"actual\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ValSetHashMismatch\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"actual\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]",
	Bin: "0x6080806040523460b9575f54600181811c9116801560b0575b6020821014609c57601f81116058575b507f30372d74656e6465726d696e742d30000000000000000000000000000000001e5f5561210290816100be8239f35b5f8052601f0160051c7f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e563908101905b818110609257506028565b5f81556001016087565b634e487b7160e01b5f52602260045260245ffd5b90607f16906018565b5f80fdfe6080806040526004361015610012575f80fd5b5f3560e01c635b9d601314610025575f80fd5b346111b85760c07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126111b8576004359067ffffffffffffffff82116111b8578136036101207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8201126111b85760e0820182811067ffffffffffffffff82111761126b57604052826004013567ffffffffffffffff81116111b8576040916100f77fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc926004369188010161164c565b845201126111b8576040519061010c826115b9565b61011860248401611690565b825261012660448401611690565b60208301526020810191825261013f36606485016116b3565b604082015261010461015360a485016116e9565b936060830194855261016760c482016116e9565b608084015261017860e482016116fa565b60a0840152013560028110156111b85760c082015260607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc3601126111b8576040516101c3816115d5565b602435906fffffffffffffffffffffffffffffffff821682036111b857604091815260443560208201520160643581526084359167ffffffffffffffff83116111b85760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc84360301126111b8576040519361023f856115f1565b836004013567ffffffffffffffff81116111b857840160407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82360301126111b8576040519061028e826115b9565b600481013567ffffffffffffffff81116111b8576004908201016102c0813603126111b85760405190610260820182811067ffffffffffffffff82111761126b576040526102dc36826116b3565b8252604081013567ffffffffffffffff81116111b8576102ff903690830161164c565b60208301526103106060820161169e565b604083015261032160808201611707565b606083015261033260a082016116fa565b60808301526103443660c08301611724565b60a083015261035661012082016116fa565b60c083015261014081013560e083015261037361016082016116fa565b6101008301526101808101356101208301526101a08101356101408301526101c08101356101608301526101e081013561018083015261020081013567ffffffffffffffff81116111b8576103cb903690830161164c565b6101a08301526103de61022082016116fa565b6101c08301526102408101356101e08301526103fd61026082016116fa565b6102008301526102808101356102208301526102a08101359067ffffffffffffffff82116111b8576104319136910161164c565b6102408201528252602481013567ffffffffffffffff81116111b857600491010160c0813603126111b85760405190610469826115f1565b6104728161169e565b8252610480602082016116e9565b60208301526104923660408301611724565b604083015260a08101359067ffffffffffffffff82116111b8570136601f820112156111b85780356104c381611776565b916104d1604051938461160d565b81835260208084019260051b820101903682116111b85760208101925b8284106114cd5750505050606082015260208201528552602484013567ffffffffffffffff81116111b85761052990600436918701016117f8565b916020860192835261053e36604487016116b3565b946040870195865260848101359067ffffffffffffffff82116111b857600461056a92369201016117f8565b916060870192835260a435916fffffffffffffffffffffffffffffffff83168093036111b85763ffffffff6105a160409251611a7a565b9251995116988151906105b3826115d5565b815260208101998a520190600f82528651604067ffffffffffffffff60208184511693015116918151906105e6826115d5565b6105ee6118e8565b8252602082015201526105ff6118e8565b506020610610818a51510151611a7a565b0167ffffffffffffffff81511667ffffffffffffffff895151169081810361149f57505067ffffffffffffffff90511661067267ffffffffffffffff60408b51510151169160405190610662826115b9565b8152602081019283528951612075565b61067b81611f52565b6002811490811561148b575b50611454575073__$7440a880b7578767f72184d998805816e4$__956106dd60208751604051809381927fc87f1f6900000000000000000000000000000000000000000000000000000000835260048301611e0b565b03818b5af49081156111c4575f91611422575b506101408a515101518082036113f457505061071160208a51510151611a7a565b67ffffffffffffffff6020818186015116920151160361137657906107679160208651604051809581927fc87f1f6900000000000000000000000000000000000000000000000000000000835260048301611e0b565b03818b5af49283156111c4575f93611342575b5051820361129857602067ffffffffffffffff915198510151169351976040519760a0890189811067ffffffffffffffff82111761126b5760405288526020880198848a526040890195865260608901526080880191825251945195604051956107e3876115b9565b8652610824602080880198808a52604051809381927fc87f1f6900000000000000000000000000000000000000000000000000000000835260048301611e0b565b0381855af49081156111c4575f91611239575b506101408751510151036111cf576020610a12918751519060405180809581947f95b25f7900000000000000000000000000000000000000000000000000000000835286600484015267ffffffffffffffff87825182815116602487015201511660448401526102406109ab6108bd898401516102c060648801526102e4870190611da5565b67ffffffffffffffff60408501511660848701526fffffffffffffffffffffffffffffffff60608501511660a48701526080840151151560c4870152898060a0860151805160e48a0152015163ffffffff815116610104890152015161012487015260c0840151151561014487015260e084015161016487015261010084015115156101848701526101208401516101a48701526101408401516101c48701526101608401516101e48701526101808401516102048701526101a08401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc87830301610224880152611da5565b916101c081015115156102448601526101e081015161026486015261020081015115156102848601526102208101516102a486015201517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc848303016102c4850152611da5565b03915af49081156111c4575f9161118e575b506040602087510151015151036111245760606020865197519701510151965f985f5b8951811015610a8457610a5a818b6120e1565b5151610a6581611f52565b610a6e81611f52565b610a7b575b600101610a47565b60019a50610a73565b508a99989798156110ba5787518951510361102a575f5b8851811015610bde57610aae818a6120e1565b518051610aba81611f52565b610ac381611f52565b610ad257506001905b01610a9b565b6020909b9692979398949995919b015151985f985f5b8c518051821015610bcc5781610afd916120e1565b5151604051610b2b6020828180820195805191829101875e81015f838201520301601f19810183528261160d565b5190208c604051610b5b6020828180820195805191829101875e81015f838201520301601f19810183528261160d565b51902014610b6b57600101610ae8565b50939792969b9195995093975060015b15610b8857600190610acc565b606460405162461bcd60e51b815260206004820152601d60248201527f696e76616c696420636f6d6d69743a206661756c7479207369676e65720000006044820152fd5b5050939792969b919599509397610b7b565b5063ffffffff8a51166fffffffffffffffffffffffffffffffff82511690818710918215610ff4575b5050610f8a576fffffffffffffffffffffffffffffffff8060608951510151169151161015610f205760208651510151604051610c636020828180820195805191829101875e81015f838201520301601f19810183528261160d565b5190209051604051610c946020828180820195805191829101875e81015f838201520301601f19810183528261160d565b51902003610edc57610cb067ffffffffffffffff855116611f89565b8551516040015167ffffffffffffffff9182169116818103610e605750506101408551510151905103610df65763ffffffff905b5116016fffffffffffffffffffffffffffffffff8111610dc9576fffffffffffffffffffffffffffffffff80606085515101511691161115610d5f5767ffffffffffffffff6040610d388280945116611f89565b935151015116911614610d4757005b60036020604051610d57816115b9565b600281520152005b608460405162461bcd60e51b815260206004820152602860248201527f696e76616c696420626c6f636b3a206865616465722069732066726f6d20746860448201527f65206675747572650000000000000000000000000000000000000000000000006064820152fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b608460405162461bcd60e51b815260206004820152602f60248201527f696e76616c696420626c6f636b3a206e6578742076616c696461746f7220736560448201527f742068617368206d69736d6174636800000000000000000000000000000000006064820152fd5b11159050610e735763ffffffff90610ce4565b608460405162461bcd60e51b8152602060048201526024808201527f696e76616c696420626c6f636b3a206e6f6e20696e6372656173696e6720686560448201527f69676874000000000000000000000000000000000000000000000000000000006064820152fd5b606460405162461bcd60e51b815260206004820152602060248201527f696e76616c696420626c6f636b3a20636861696e2d6964206d69736d617463686044820152fd5b608460405162461bcd60e51b815260206004820152602560248201527f696e76616c696420626c6f636b3a206e6f6e206d6f6e6f746f6e69632062667460448201527f2074696d650000000000000000000000000000000000000000000000000000006064820152fd5b608460405162461bcd60e51b815260206004820152603c60248201527f696e76616c696420626c6f636b3a20756e74727573746564207374617465206960448201527f73206f757473696465206f66207472757374696e6720706572696f64000000006064820152fd5b90915086036fffffffffffffffffffffffffffffffff8111610dc9576fffffffffffffffffffffffffffffffff16118880610c07565b60a460405162461bcd60e51b815260206004820152604860248201527f696e76616c696420636f6d6d69743a206e756d626572206f66207369676e617460448201527f7572657320646f6573206e6f74206d61746368206e756d626572206f6620766160648201527f6c696461746f72730000000000000000000000000000000000000000000000006084820152fd5b608460405162461bcd60e51b815260206004820152602560248201527f696e76616c696420636f6d6d69743a206e6f2070726573656e74207369676e6160448201527f74757265730000000000000000000000000000000000000000000000000000006064820152fd5b608460405162461bcd60e51b815260206004820152602360248201527f696e76616c696420626c6f636b3a206865616465722068617368206d69736d6160448201527f74636800000000000000000000000000000000000000000000000000000000006064820152fd5b90506020813d6020116111bc575b816111a96020938361160d565b810103126111b857515f610a24565b5f80fd5b3d915061119c565b6040513d5f823e3d90fd5b608460405162461bcd60e51b815260206004820152602a60248201527f696e76616c696420626c6f636b3a2076616c696461746f72207365742068617360448201527f68206d69736d61746368000000000000000000000000000000000000000000006064820152fd5b90506020813d602011611263575b816112546020938361160d565b810103126111b857515f610837565b3d9150611247565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b60a46040517ff492ef2b00000000000000000000000000000000000000000000000000000000815260206004820152604360248201527f74727573746564206e6578742076616c696461746f722073657420686173682060448201527f646f6573206e6f74206d6174636820686173682073746f726564206f6e20636860648201527f61696e00000000000000000000000000000000000000000000000000000000006084820152fd5b9092506020813d60201161136e575b8161135e6020938361160d565b810103126111b85751915f61077a565b3d9150611351565b6113c0826113f060208c5151015191516040519384937fc6913e9f000000000000000000000000000000000000000000000000000000008552604060048601526044850190611da5565b907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc848303016024850152611da5565b0390fd5b7f4d6d8014000000000000000000000000000000000000000000000000000000005f5260045260245260445ffd5b90506020813d60201161144c575b8161143d6020938361160d565b810103126111b857515f6106f0565b3d9150611430565b67ffffffffffffffff9051167f90f4dbed000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b6001915061149881611f52565b145f610687565b7f05b4c7a3000000000000000000000000000000000000000000000000000000005f5260045260245260445ffd5b833567ffffffffffffffff81116111b8578201906040601f1983360301126111b857604051916114fc836115b9565b602081013560038110156111b8578352604081013567ffffffffffffffff81116111b8576020910101906080823603126111b8576040519261153d846115f1565b823567ffffffffffffffff81116111b85761155b903690850161164c565b845261156960208401611707565b602085015261157a604084016116fa565b604085015260608301359367ffffffffffffffff85116111b8576115a560209594869536910161164c565b6060820152838201528152019301926104ee565b6040810190811067ffffffffffffffff82111761126b57604052565b6060810190811067ffffffffffffffff82111761126b57604052565b6080810190811067ffffffffffffffff82111761126b57604052565b90601f601f19910116810190811067ffffffffffffffff82111761126b57604052565b67ffffffffffffffff811161126b57601f01601f191660200190565b81601f820112156111b85760208135910161166682611630565b92611674604051948561160d565b828452828201116111b857815f92602092838601378301015290565b359060ff821682036111b857565b359067ffffffffffffffff821682036111b857565b91908260409103126111b8576040516116cb816115b9565b60206116e48183956116dc8161169e565b85520161169e565b910152565b359063ffffffff821682036111b857565b359081151582036111b857565b35906fffffffffffffffffffffffffffffffff821682036111b857565b8092910391606083126111b85760405161173d816115b9565b6040601f1982958435845201126111b857602090604080519361175f856115b9565b61176a8482016116e9565b85520135828401520152565b67ffffffffffffffff811161126b5760051b60200190565b91906080838203126111b857604051906117a7826115f1565b8193803567ffffffffffffffff81116111b8576060926117c891830161164c565b8352602081013560208401526117e06040820161169e565b60408401520135908160070b82036111b85760600152565b9190916080818403126111b85760405190611812826115f1565b8193813567ffffffffffffffff81116111b857820181601f820112156111b857803561183d81611776565b9161184b604051938461160d565b81835260208084019260051b820101918483116111b85760208201905b8382106118ba57505050508352611881602083016116fa565b602084015260408201359067ffffffffffffffff82116111b857826118af606094926116e49486940161178e565b60408601520161169e565b813567ffffffffffffffff81116111b8576020916118dd8884809488010161178e565b815201910190611868565b604051905f5f548060011c91600182169182156119f4575b6020841083146119c757838652859290811561198a575060011461192d575b61192b9250038361160d565b565b505f80805290917f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e5635b81831061196e57505090602061192b9282010161191f565b6020919350806001915483858901015201910190918492611956565b6020925061192b9491507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001682840152151560051b82010161191f565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b92607f1692611900565b908151811015611a0f570160200190565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b90611a4682611630565b611a53604051918261160d565b828152601f19611a638294611630565b0190602036910137565b91908201809211610dc957565b604051611a86816115b9565b606081525f60209091015280515f198101908111610dc9575b7f2d000000000000000000000000000000000000000000000000000000000000007fff00000000000000000000000000000000000000000000000000000000000000611aeb83856119fe565b511614611b00578015610dc9575f1901611a9f565b91905f198314611d82578051838103908111610dc9575f198101908111610dc957611b2a90611a3c565b905f5b8251811015611b8d576001850190818611610dc9577fff00000000000000000000000000000000000000000000000000000000000000611b78611b7283600195611a6d565b856119fe565b51165f1a611b8682866119fe565b5301611b2d565b509192908051158015611d24575b611c95578051611baa91611ec8565b9190158015611d13575b611c9557611bc181611a3c565b905f5b818110611c52575050516001811190811591611c46575b50611c025767ffffffffffffffff9060405192611bf7846115b9565b835216602082015290565b606460405162461bcd60e51b815260206004820152601b60248201527f496e76616c696420636861696e20707265666978206c656e67746800000000006044820152fd5b602b915010155f611bdb565b807fff00000000000000000000000000000000000000000000000000000000000000611c80600193886119fe565b51165f1a611c8e82866119fe565b5301611bc4565b5050805180600111159081611d08575b5015611cc35760405190611cb8826115b9565b81525f602082015290565b60405162461bcd60e51b815260206004820152601760248201527f496e76616c696420636861696e204944206c656e6774680000000000000000006044820152606490fd5b60409150105f611ca5565b5067ffffffffffffffff8211611bb4565b611a0f577f30000000000000000000000000000000000000000000000000000000000000007fff00000000000000000000000000000000000000000000000000000000000000602083015116148015611b9b57506001815111611b9b565b809192505180600111159081611d08575015611cc35760405190611cb8826115b9565b90601f19601f602080948051918291828752018686015e5f8582860101520116010190565b90606080611de18451608085526080850190611da5565b936020810151602085015267ffffffffffffffff6040820151166040850152015160070b91015290565b6020815260a0810182519060806020840152815180915260c0830190602060c08260051b8601019301915f905b828210611e7f575050505067ffffffffffffffff6060611e756080936020870151151560408701526040870151601f198783030184880152611dca565b9401511691015290565b90919293602080611eba837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff408a600196030186528851611dca565b960192019201909291611e38565b5f929183915b818310611ede5750505060019190565b90919360ff611f147fff000000000000000000000000000000000000000000000000000000000000006020888601015116611fa9565b169060098211611f4657600a810290808204600a1490151715610dc957600191611f3d91611a6d565b94019190611ece565b5050505090505f905f90565b60031115611f5c57565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b67ffffffffffffffff60019116019067ffffffffffffffff8211610dc957565b60f81c602f81118061206b575b15611fe3577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffd00160ff1690565b6060811180612061575b1561201a577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa90160ff1690565b6040811180612057575b15612051577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc90160ff1690565b5060ff90565b5060478110612024565b5060678110611fed565b50603a8110611fb6565b67ffffffffffffffff81511667ffffffffffffffff835116908181105f1461209f57505050505f90565b11156120ac575050600290565b602067ffffffffffffffff81819301511692015116908181105f146120d15750505f90565b11156120dc57600290565b600190565b8051821015611a0f5760209160051b01019056fea164736f6c634300081c000a",
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

// UpdateClient is a free data retrieval call binding the contract method 0x5b9d6013.
//
// Solidity: function updateClient((string,(uint8,uint8),(uint64,uint64),uint32,uint32,bool,uint8) clientState, (uint128,bytes32,bytes32) trustedConsensusState, ((((uint64,uint64),string,uint64,uint128,bool,(bytes32,(uint32,bytes32)),bool,bytes32,bool,bytes32,bytes32,bytes32,bytes32,bytes,bool,bytes32,bool,bytes32,bytes),(uint64,uint32,(bytes32,(uint32,bytes32)),(uint8,(bytes,uint128,bool,bytes))[])),((bytes,bytes32,uint64,int64)[],bool,(bytes,bytes32,uint64,int64),uint64),(uint64,uint64),((bytes,bytes32,uint64,int64)[],bool,(bytes,bytes32,uint64,int64),uint64)) proposedHeader, uint128 time) view returns()
func (_ContractUpdateClient *ContractUpdateClientCaller) UpdateClient(opts *bind.CallOpts, clientState IICS07TendermintMsgsClientState, trustedConsensusState IICS07TendermintMsgsConsensusState, proposedHeader IICS07TendermintMsgsHeader, time *big.Int) error {
	var out []interface{}
	err := _ContractUpdateClient.contract.Call(opts, &out, "updateClient", clientState, trustedConsensusState, proposedHeader, time)

	if err != nil {
		return err
	}

	return err

}

// UpdateClient is a free data retrieval call binding the contract method 0x5b9d6013.
//
// Solidity: function updateClient((string,(uint8,uint8),(uint64,uint64),uint32,uint32,bool,uint8) clientState, (uint128,bytes32,bytes32) trustedConsensusState, ((((uint64,uint64),string,uint64,uint128,bool,(bytes32,(uint32,bytes32)),bool,bytes32,bool,bytes32,bytes32,bytes32,bytes32,bytes,bool,bytes32,bool,bytes32,bytes),(uint64,uint32,(bytes32,(uint32,bytes32)),(uint8,(bytes,uint128,bool,bytes))[])),((bytes,bytes32,uint64,int64)[],bool,(bytes,bytes32,uint64,int64),uint64),(uint64,uint64),((bytes,bytes32,uint64,int64)[],bool,(bytes,bytes32,uint64,int64),uint64)) proposedHeader, uint128 time) view returns()
func (_ContractUpdateClient *ContractUpdateClientSession) UpdateClient(clientState IICS07TendermintMsgsClientState, trustedConsensusState IICS07TendermintMsgsConsensusState, proposedHeader IICS07TendermintMsgsHeader, time *big.Int) error {
	return _ContractUpdateClient.Contract.UpdateClient(&_ContractUpdateClient.CallOpts, clientState, trustedConsensusState, proposedHeader, time)
}

// UpdateClient is a free data retrieval call binding the contract method 0x5b9d6013.
//
// Solidity: function updateClient((string,(uint8,uint8),(uint64,uint64),uint32,uint32,bool,uint8) clientState, (uint128,bytes32,bytes32) trustedConsensusState, ((((uint64,uint64),string,uint64,uint128,bool,(bytes32,(uint32,bytes32)),bool,bytes32,bool,bytes32,bytes32,bytes32,bytes32,bytes,bool,bytes32,bool,bytes32,bytes),(uint64,uint32,(bytes32,(uint32,bytes32)),(uint8,(bytes,uint128,bool,bytes))[])),((bytes,bytes32,uint64,int64)[],bool,(bytes,bytes32,uint64,int64),uint64),(uint64,uint64),((bytes,bytes32,uint64,int64)[],bool,(bytes,bytes32,uint64,int64),uint64)) proposedHeader, uint128 time) view returns()
func (_ContractUpdateClient *ContractUpdateClientCallerSession) UpdateClient(clientState IICS07TendermintMsgsClientState, trustedConsensusState IICS07TendermintMsgsConsensusState, proposedHeader IICS07TendermintMsgsHeader, time *big.Int) error {
	return _ContractUpdateClient.Contract.UpdateClient(&_ContractUpdateClient.CallOpts, clientState, trustedConsensusState, proposedHeader, time)
}
