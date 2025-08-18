// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractICS20Transfer

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

// IIBCAppCallbacksOnAcknowledgementPacketCallback is an auto generated low-level Go binding around an user-defined struct.
type IIBCAppCallbacksOnAcknowledgementPacketCallback struct {
	SourceClient      string
	DestinationClient string
	Sequence          uint64
	Payload           IICS26RouterMsgsPayload
	Acknowledgement   []byte
	Relayer           common.Address
}

// IIBCAppCallbacksOnRecvPacketCallback is an auto generated low-level Go binding around an user-defined struct.
type IIBCAppCallbacksOnRecvPacketCallback struct {
	SourceClient      string
	DestinationClient string
	Sequence          uint64
	Payload           IICS26RouterMsgsPayload
	Relayer           common.Address
}

// IIBCAppCallbacksOnTimeoutPacketCallback is an auto generated low-level Go binding around an user-defined struct.
type IIBCAppCallbacksOnTimeoutPacketCallback struct {
	SourceClient      string
	DestinationClient string
	Sequence          uint64
	Payload           IICS26RouterMsgsPayload
	Relayer           common.Address
}

// IICS20TransferMsgsSendTransferMsg is an auto generated low-level Go binding around an user-defined struct.
type IICS20TransferMsgsSendTransferMsg struct {
	Denom            common.Address
	Amount           *big.Int
	Receiver         string
	SourceClient     string
	DestPort         string
	TimeoutTimestamp uint64
	Memo             string
}

// IICS26RouterMsgsPayload is an auto generated low-level Go binding around an user-defined struct.
type IICS26RouterMsgsPayload struct {
	SourcePort string
	DestPort   string
	Version    string
	Encoding   string
	Value      []byte
}

// ISignatureTransferPermitTransferFrom is an auto generated low-level Go binding around an user-defined struct.
type ISignatureTransferPermitTransferFrom struct {
	Permitted ISignatureTransferTokenPermissions
	Nonce     *big.Int
	Deadline  *big.Int
}

// ISignatureTransferTokenPermissions is an auto generated low-level Go binding around an user-defined struct.
type ISignatureTransferTokenPermissions struct {
	Token  common.Address
	Amount *big.Int
}

// ContractICS20TransferMetaData contains all meta data concerning the ContractICS20Transfer contract.
var ContractICS20TransferMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DELEGATE_SENDER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ERC20_CUSTOMIZER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PAUSER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TOKEN_OPERATOR_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UNPAUSER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getEscrow\",\"inputs\":[{\"name\":\"clientId\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getEscrowBeacon\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getIBCERC20Beacon\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPermit2\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantDelegateSenderRole\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"grantERC20CustomizerRole\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"grantPauserRole\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"grantTokenOperatorRole\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"grantUnpauserRole\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ibcERC20Contract\",\"inputs\":[{\"name\":\"denom\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ibcERC20Denom\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ics26\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"ics26Router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"escrowLogic\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"ibcERC20Logic\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"permit2\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isTokenOperator\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"multicall\",\"inputs\":[{\"name\":\"data\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"outputs\":[{\"name\":\"results\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onAcknowledgementPacket\",\"inputs\":[{\"name\":\"msg_\",\"type\":\"tuple\",\"internalType\":\"structIIBCAppCallbacks.OnAcknowledgementPacketCallback\",\"components\":[{\"name\":\"sourceClient\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"destinationClient\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"payload\",\"type\":\"tuple\",\"internalType\":\"structIICS26RouterMsgs.Payload\",\"components\":[{\"name\":\"sourcePort\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"destPort\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"encoding\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"acknowledgement\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"relayer\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onRecvPacket\",\"inputs\":[{\"name\":\"msg_\",\"type\":\"tuple\",\"internalType\":\"structIIBCAppCallbacks.OnRecvPacketCallback\",\"components\":[{\"name\":\"sourceClient\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"destinationClient\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"payload\",\"type\":\"tuple\",\"internalType\":\"structIICS26RouterMsgs.Payload\",\"components\":[{\"name\":\"sourcePort\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"destPort\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"encoding\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"relayer\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onTimeoutPacket\",\"inputs\":[{\"name\":\"msg_\",\"type\":\"tuple\",\"internalType\":\"structIIBCAppCallbacks.OnTimeoutPacketCallback\",\"components\":[{\"name\":\"sourceClient\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"destinationClient\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"payload\",\"type\":\"tuple\",\"internalType\":\"structIICS26RouterMsgs.Payload\",\"components\":[{\"name\":\"sourcePort\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"destPort\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"encoding\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"relayer\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeDelegateSenderRole\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeERC20CustomizerRole\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokePauserRole\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeTokenOperatorRole\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeUnpauserRole\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"sendTransfer\",\"inputs\":[{\"name\":\"msg_\",\"type\":\"tuple\",\"internalType\":\"structIICS20TransferMsgs.SendTransferMsg\",\"components\":[{\"name\":\"denom\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"sourceClient\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"destPort\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"timeoutTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"memo\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"sendTransferWithPermit2\",\"inputs\":[{\"name\":\"msg_\",\"type\":\"tuple\",\"internalType\":\"structIICS20TransferMsgs.SendTransferMsg\",\"components\":[{\"name\":\"denom\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"sourceClient\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"destPort\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"timeoutTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"memo\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"permit\",\"type\":\"tuple\",\"internalType\":\"structISignatureTransfer.PermitTransferFrom\",\"components\":[{\"name\":\"permitted\",\"type\":\"tuple\",\"internalType\":\"structISignatureTransfer.TokenPermissions\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"sendTransferWithSender\",\"inputs\":[{\"name\":\"msg_\",\"type\":\"tuple\",\"internalType\":\"structIICS20TransferMsgs.SendTransferMsg\",\"components\":[{\"name\":\"denom\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"sourceClient\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"destPort\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"timeoutTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"memo\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCustomERC20\",\"inputs\":[{\"name\":\"denom\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeEscrowTo\",\"inputs\":[{\"name\":\"newEscrowLogic\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeIBCERC20To\",\"inputs\":[{\"name\":\"newIBCERC20Logic\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"IBCERC20ContractCreated\",\"inputs\":[{\"name\":\"contractAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"fullDenomPath\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpectedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ICS20DenomAlreadyExists\",\"inputs\":[{\"name\":\"denom\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"ICS20DenomNotFound\",\"inputs\":[{\"name\":\"denom\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"ICS20EscrowNotFound\",\"inputs\":[{\"name\":\"clientID\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"ICS20InvalidAddress\",\"inputs\":[{\"name\":\"addr\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"ICS20InvalidAmount\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ICS20InvalidPort\",\"inputs\":[{\"name\":\"expected\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"actual\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"ICS20Permit2TokenMismatch\",\"inputs\":[{\"name\":\"permitToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sentToken\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ICS20TokenAlreadyExists\",\"inputs\":[{\"name\":\"denom\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"ICS20Unauthorized\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ICS20UnauthorizedPacketSender\",\"inputs\":[{\"name\":\"packetSender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ICS20UnexpectedERC20Balance\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"actual\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ICS20UnexpectedEncoding\",\"inputs\":[{\"name\":\"expected\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"actual\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"ICS20UnexpectedVersion\",\"inputs\":[{\"name\":\"expected\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"StringsInsufficientHexLength\",\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]",
	Bin: "0x60a080604052346100c257306080525f516020615e3d5f395f51905f525460ff8160401c166100b3576002600160401b03196001600160401b03821601610060575b604051615d7690816100c78239608051818181611ceb0152611db50152f35b6001600160401b0319166001600160401b039081175f516020615e3d5f395f51905f525581527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a15f80610041565b63f92ee8a960e01b5f5260045ffd5b5f80fdfe6080806040526004361015610012575f80fd5b5f905f3560e01c90816301ffc9a7146130045750806302769d7314612fca57806306ab20bc14612e99578063078c4a79146126f65780631bbf2e23146126b05780631e5150e41461266a57806323993870146125cd578063248a9ca31461257a5780632540e2da146125505780632ac3dc38146124965780632f2ff15d14612438578063329687821461240e57806333daf980146123d357806336568abe146123755780633f4ba83a14612207578063428e4e17146120e65780634551d2b0146120495780634f1ef28614611d6357806352d1902d14611cd057806353816a7c146119a75780635c975abb146119655780635e32b6b61461189957806366780af9146117fc5780636ad7ba491461175f5780636c0e87e1146116c25780636c11c21c14611698578063826cae7a146116525780638456cb591461150a57806391d14854146114a0578063969631d514611424578063a1d28f57146110b3578063a217fddf14611097578063a50ee2b414610ffc578063aaa2c34314610ec8578063ac9650d814610d1a578063ad3cb1cc14610cb9578063b29c715d14610a6b578063c46023a31461099e578063d413227d14610958578063d547741f146108f5578063da0f21b5146108ba578063e163b1af14610833578063e63ab1e9146107f8578063f17a7e6614610770578063f865af0814610742578063f8c8765e146102605763fb1bb9de14610223575f80fd5b3461025d578060031936011261025d5760206040517f427da25fe773164f88948d3e215c94b6554e2ed5e5f203a821c9f2f6131cf75a8152f35b80fd5b503461025d57608060031936011261025d5761027a6130a2565b6102826130b8565b604435906001600160a01b038216820361073e57606435916001600160a01b03831680930361073a577ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00549360ff8560401c16159467ffffffffffffffff811680159081610732575b6001149081610728575b15908161071f575b506106f75790818660017fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000006001600160a01b039516177ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00556106a2575b50610363615148565b61036b615148565b610373615148565b61037b615148565b610383615148565b167fffffffffffffffffffffffff00000000000000000000000000000000000000007f823f7a8ea9ae6df0eb03ec5e1682d7a2839417ad8a91774118e6acf2e8d2f8035416177f823f7a8ea9ae6df0eb03ec5e1682d7a2839417ad8a91774118e6acf2e8d2f80355604051916105129283810181811067ffffffffffffffff82111761066a5761042d829161537495878785396001600160a01b0316815230602082015260400190565b039087f08015610697576001600160a01b03167fffffffffffffffffffffffff00000000000000000000000000000000000000007f823f7a8ea9ae6df0eb03ec5e1682d7a2839417ad8a91774118e6acf2e8d2f8045416177f823f7a8ea9ae6df0eb03ec5e1682d7a2839417ad8a91774118e6acf2e8d2f80455604051928084019284841067ffffffffffffffff85111761066a57918493916104e39385396001600160a01b0316815230602082015260400190565b039084f0801561065f576001600160a01b03167fffffffffffffffffffffffff00000000000000000000000000000000000000007f823f7a8ea9ae6df0eb03ec5e1682d7a2839417ad8a91774118e6acf2e8d2f8055416177f823f7a8ea9ae6df0eb03ec5e1682d7a2839417ad8a91774118e6acf2e8d2f805557fffffffffffffffffffffffff00000000000000000000000000000000000000007f823f7a8ea9ae6df0eb03ec5e1682d7a2839417ad8a91774118e6acf2e8d2f8065416177f823f7a8ea9ae6df0eb03ec5e1682d7a2839417ad8a91774118e6acf2e8d2f806556105cb5780f35b7fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054167ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00557fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2602060405160018152a180f35b6040513d85823e3d90fd5b6024887f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b6040513d88823e3d90fd5b7fffffffffffffffffffffffffffffffffffffffffffffff0000000000000000001668010000000000000001177ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00555f61035a565b6004877ff92ee8a9000000000000000000000000000000000000000000000000000000008152fd5b9050155f6102fd565b303b1591506102f5565b8791506102eb565b8480fd5b8380fd5b503461025d57602060031936011261025d5761076c61075f6130a2565b610767613fe2565b613e4f565b5080f35b503461025d57602060031936011261025d576001600160a01b0360406107946130a2565b927f62150a51582c26f4255242a3c4ca35fb04250e7315069523d650676aed01a56a81527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020522091165f52602052602060ff60405f2054166040519015158152f35b503461025d578060031936011261025d5760206040517f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a8152f35b503461025d57602060031936011261025d576108b661089b6108a261088f6108596130a2565b6001600160a01b03165f527f823f7a8ea9ae6df0eb03ec5e1682d7a2839417ad8a91774118e6acf2e8d2f80260205260405f2090565b60405192838092613647565b038261316f565b604051918291602083526020830190613101565b0390f35b503461025d578060031936011261025d5760206040517f04c6d412b3efc8acc11f9a5a5dc9f2a5d9ae1f6d869984b47386662ab77133cd8152f35b503461025d57604060031936011261025d5761076c6004356109156130b8565b9061095361094e825f527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052600160405f20015490565b614727565b613f19565b503461025d578060031936011261025d5760206001600160a01b037f823f7a8ea9ae6df0eb03ec5e1682d7a2839417ad8a91774118e6acf2e8d2f8035416604051908152f35b503461025d57602060031936011261025d5760246109ba6130a2565b60206001600160a01b037f823f7a8ea9ae6df0eb03ec5e1682d7a2839417ad8a91774118e6acf2e8d2f8035416604051938480927f24d7806c0000000000000000000000000000000000000000000000000000000082523360048301525afa91821561065f5761076c92610a37918591610a3c575b50339061322a565b613d85565b610a5e915060203d602011610a64575b610a56818361316f565b810190613212565b5f610a2f565b503d610a4c565b503461025d57604060031936011261025d5760043567ffffffffffffffff8111610c3557806004019060e06003198236030112610cb557610aaa6130b8565b91610ab3613850565b610abb6137dc565b7f04c6d412b3efc8acc11f9a5a5dc9f2a5d9ae1f6d869984b47386662ab77133cd84527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604084206001600160a01b0333165f5260205260ff60405f20541615610c65576024820135918215610c3957610b55610b50610b4960646001600160a01b0394018561329a565b36916131ae565b613917565b1691610b6b81610b64846135aa565b853361406a565b84610b75836135aa565b91843b15610c35576040517fb4f22eb70000000000000000000000000000000000000000000000000000000081526001600160a01b0393909316600484015233602484015260448301528160648183875af18015610c2a57610c15575b602085610be0868686614278565b907f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f005d67ffffffffffffffff60405191168152f35b610c2085809261316f565b61073e575f610bd2565b6040513d87823e3d90fd5b5080fd5b6024857f4f6df8d000000000000000000000000000000000000000000000000000000000815280600452fd5b6044847fe2517d3f000000000000000000000000000000000000000000000000000000008152336004527f04c6d412b3efc8acc11f9a5a5dc9f2a5d9ae1f6d869984b47386662ab77133cd602452fd5b8280fd5b503461025d578060031936011261025d57506108b6604051610cdc60408261316f565b600581527f352e302e300000000000000000000000000000000000000000000000000000006020820152604051918291602083526020830190613101565b503461025d57602060031936011261025d5760043567ffffffffffffffff8111610c355736602382011215610c355780600401359067ffffffffffffffff8211610cb557602481013660248460051b8401011161073e576040516020610d80818361316f565b85825280820192601f198201368537610d9886613783565b94610da6604051968761316f565b868652601f19610db588613783565b0183895b828110610eb857505050875b87811015610e3b57600190610e1f610e19610de860248460051b8701018761329a565b91908d898c856040519687958487013784018281018481528e519283915e010190815203601f19810183528261316f565b30614e81565b610e29828a61379b565b52610e34818961379b565b5001610dc5565b83898860405191838301848452825180915260408401948060408360051b870101940192955b828710610e6e5785850386f35b909192938280610ea8837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08a600196030186528851613101565b9601920196019592919092610e61565b606082828b010152018490610db9565b503461025d57602060031936011261025d5780610ee36130a2565b602460206001600160a01b037f823f7a8ea9ae6df0eb03ec5e1682d7a2839417ad8a91774118e6acf2e8d2f8035416604051928380927f24d7806c0000000000000000000000000000000000000000000000000000000082523360048301525afa801561065f57610f5c918491610a3c5750339061322a565b6001600160a01b037f823f7a8ea9ae6df0eb03ec5e1682d7a2839417ad8a91774118e6acf2e8d2f805541690813b15610ff8576001600160a01b03602484928360405195869485937f3659cfe60000000000000000000000000000000000000000000000000000000085521660048401525af18015610fed57610fdc5750f35b81610fe69161316f565b61025d5780f35b6040513d84823e3d90fd5b5050fd5b503461025d57602060031936011261025d5760043567ffffffffffffffff8111610c355761102e9036906004016131e4565b6001600160a01b0361104082846135be565b541690811561105457602082604051908152f35b90506110936040519283927fe1275e2f000000000000000000000000000000000000000000000000000000008452602060048501526024840191613326565b0390fd5b503461025d578060031936011261025d57602090604051908152f35b503461025d57604060031936011261025d5760043567ffffffffffffffff8111610c35576110e59036906004016131e4565b6110f09291926130b8565b7f70e300420270129b335fa3e16cefa8088dcfc9433472e4725448f7fc7a82856c83527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604083206001600160a01b0333165f5260205260ff60405f205416156113d4576001600160a01b0361116983866135be565b5416611398576111b26111ac826001600160a01b03165f527f823f7a8ea9ae6df0eb03ec5e1682d7a2839417ad8a91774118e6acf2e8d2f80260205260405f2090565b546135f6565b156111ed826001600160a01b03165f527f823f7a8ea9ae6df0eb03ec5e1682d7a2839417ad8a91774118e6acf2e8d2f80260205260405f2090565b901561135a575061126b9061120283866135be565b6001600160a01b03808316167fffffffffffffffffffffffff00000000000000000000000000000000000000008254161790556001600160a01b03165f527f823f7a8ea9ae6df0eb03ec5e1682d7a2839417ad8a91774118e6acf2e8d2f80260205260405f2090565b9067ffffffffffffffff811161132d5761128f8161128984546135f6565b846136c9565b82601f82116001146112cd57819084956112bd9495926112c2575b50505f198260011b9260031b1c19161790565b905580f35b013590505f806112aa565b601f198216948385526020852091855b8781106113155750836001959697106112fc575b505050811b01905580f35b5f1960f88560031b161c199101351690555f80806112f1565b909260206001819286860135815501940191016112dd565b6024837f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b611093906040519182917f778769c4000000000000000000000000000000000000000000000000000000008352602060048401526024830190613647565b6040517f0c0ef5340000000000000000000000000000000000000000000000000000000081526020600482015280611093602482018588613326565b6044837fe2517d3f000000000000000000000000000000000000000000000000000000008152336004527f70e300420270129b335fa3e16cefa8088dcfc9433472e4725448f7fc7a82856c602452fd5b503461025d57602060031936011261025d5760043567ffffffffffffffff8111610c3557602091826114626001600160a01b039336906004016131e4565b925082604051938492833781017f823f7a8ea9ae6df0eb03ec5e1682d7a2839417ad8a91774118e6acf2e8d2f8008152030190205416604051908152f35b503461025d57604060031936011261025d576001600160a01b0360406114c46130b8565b9260043581527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020522091165f52602052602060ff60405f2054166040519015158152f35b503461025d578060031936011261025d577f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604081206001600160a01b0333165f5260205260ff60405f205416156116025761158a613850565b600160ff197fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005416177fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586020604051338152a180f35b807fe2517d3f0000000000000000000000000000000000000000000000000000000060449252336004527f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a602452fd5b503461025d578060031936011261025d5760206001600160a01b037f823f7a8ea9ae6df0eb03ec5e1682d7a2839417ad8a91774118e6acf2e8d2f8045416604051908152f35b503461025d57602060031936011261025d5761076c6116b56130a2565b6116bd613fe2565b614ad0565b503461025d57602060031936011261025d5760246116de6130a2565b60206001600160a01b037f823f7a8ea9ae6df0eb03ec5e1682d7a2839417ad8a91774118e6acf2e8d2f8035416604051938480927f24d7806c0000000000000000000000000000000000000000000000000000000082523360048301525afa91821561065f5761076c9261175a918591610a3c5750339061322a565b614a04565b503461025d57602060031936011261025d57602461177b6130a2565b60206001600160a01b037f823f7a8ea9ae6df0eb03ec5e1682d7a2839417ad8a91774118e6acf2e8d2f8035416604051938480927f24d7806c0000000000000000000000000000000000000000000000000000000082523360048301525afa91821561065f5761076c926117f7918591610a3c5750339061322a565b613cbb565b503461025d57602060031936011261025d5760246118186130a2565b60206001600160a01b037f823f7a8ea9ae6df0eb03ec5e1682d7a2839417ad8a91774118e6acf2e8d2f8035416604051938480927f24d7806c0000000000000000000000000000000000000000000000000000000082523360048301525afa91821561065f5761076c92611894918591610a3c5750339061322a565b614938565b503461025d5761193f6118ab366130ce565b6118e1336001600160a01b037f823f7a8ea9ae6df0eb03ec5e1682d7a2839417ad8a91774118e6acf2e8d2f8035416331461322a565b6118e96137dc565b6118f1613850565b611937606082019161192e61192861192161191961190f8786613267565b608081019061329a565b810190613443565b9483613267565b8061329a565b9290918061329a565b929091614c63565b807f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f005d80f35b503461025d578060031936011261025d57602060ff7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330054166040519015158152f35b503461025d5760c060031936011261025d5760043567ffffffffffffffff8111610c3557806004019060e06003198236030112610cb55760807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc360112610cb55760a43567ffffffffffffffff811161073e57611a289036906004016131e4565b611a33929192613850565b611a3b6137dc565b6024820135918215611ca457611a4f613594565b6001600160a01b0380611a61886135aa565b16911614611a6d613594565b611a76876135aa565b9115611c6a575050610b50610b496064611a9193018761329a565b926001600160a01b03807f823f7a8ea9ae6df0eb03ec5e1682d7a2839417ad8a91774118e6acf2e8d2f8065416941693604051906040820182811067ffffffffffffffff821117611c3d579088949392916040528682526020820191868352813b15611c3957856001600160a01b0391611b808296604051988997889687957f30f28b7a00000000000000000000000000000000000000000000000000000000875281611b3c6130b8565b166004880152604435602488015260643560448801526084356064880152511660848601525160a48501523360c485015261010060e4850152610104840191613326565b03925af18015610fed57611c24575b50611b99846135aa565b91833b15610c35576040517fb4f22eb70000000000000000000000000000000000000000000000000000000081526001600160a01b0393909316600484015233602484015260448301528160648183865af18015611c1957611c04575b602084610be0338587614278565b611c0f84809261316f565b610cb5575f611bf6565b6040513d86823e3d90fd5b81611c2e9161316f565b61073e57835f611b8f565b8580fd5b6024897f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b7fe36164780000000000000000000000000000000000000000000000000000000088526001600160a01b0390811660045216602452604486fd5b6024867f4f6df8d000000000000000000000000000000000000000000000000000000000815280600452fd5b503461025d578060031936011261025d576001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000163003611d3b5760206040517f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc8152f35b807fe07c8dba0000000000000000000000000000000000000000000000000000000060049252fd5b50604060031936011261025d57611d786130a2565b9060243567ffffffffffffffff8111610c355736602382011215610c3557611daa9036906024816004013591016131ae565b916001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016803014908115612014575b50611fec57611ded613fe2565b6001600160a01b03811690604051937f52d1902d000000000000000000000000000000000000000000000000000000008552602085600481865afa80958596611fb8575b50611e6257602484847f4c9c8ce3000000000000000000000000000000000000000000000000000000008252600452fd5b9091847f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc8103611f8d5750823b15611f6257807fffffffffffffffffffffffff00000000000000000000000000000000000000007f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5416177f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc557fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b8480a2805115611f305761076c91614e81565b505034611f3a5780f35b807fb398979f0000000000000000000000000000000000000000000000000000000060049252fd5b7f4c9c8ce3000000000000000000000000000000000000000000000000000000008452600452602483fd5b7faa1d49a4000000000000000000000000000000000000000000000000000000008552600452602484fd5b9095506020813d602011611fe4575b81611fd46020938361316f565b8101031261073a5751945f611e31565b3d9150611fc7565b6004827fe07c8dba000000000000000000000000000000000000000000000000000000008152fd5b90506001600160a01b037f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc541614155f611de0565b503461025d57602060031936011261025d5760246120656130a2565b60206001600160a01b037f823f7a8ea9ae6df0eb03ec5e1682d7a2839417ad8a91774118e6acf2e8d2f8035416604051938480927f24d7806c0000000000000000000000000000000000000000000000000000000082523360048301525afa91821561065f5761076c926120e1918591610a3c5750339061322a565b61486c565b503461025d57602060031936011261025d5760043567ffffffffffffffff8111610c35578060040160c06003198336030112610cb557612152336001600160a01b037f823f7a8ea9ae6df0eb03ec5e1682d7a2839417ad8a91774118e6acf2e8d2f8035416331461322a565b61215a6137dc565b612162613850565b612172610b49608484018361329a565b6020815191012060208060405161218a60408261316f565b818152017f4774d4a575993f963b1c06573736617a457abef8589178db8d10c94b4ab511ab815220146121df575b82807f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f005d80f35b611937606461220093019161192e61192861192161191961190f8786613267565b5f806121b8565b503461025d578060031936011261025d577f427da25fe773164f88948d3e215c94b6554e2ed5e5f203a821c9f2f6131cf75a81527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604081206001600160a01b0333165f5260205260ff60405f20541615612325577fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff8116156122fd5760ff19167fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa6020604051338152a180f35b6004827f8dfc202b000000000000000000000000000000000000000000000000000000008152fd5b807fe2517d3f0000000000000000000000000000000000000000000000000000000060449252336004527f427da25fe773164f88948d3e215c94b6554e2ed5e5f203a821c9f2f6131cf75a602452fd5b503461025d57604060031936011261025d5761238f6130b8565b336001600160a01b038216036123ab5761076c90600435613f19565b6004827f6697b232000000000000000000000000000000000000000000000000000000008152fd5b503461025d578060031936011261025d5760206040517f62150a51582c26f4255242a3c4ca35fb04250e7315069523d650676aed01a56a8152f35b503461025d57602060031936011261025d5761076c61242b6130a2565b612433613fe2565b6147a0565b503461025d57604060031936011261025d5761076c6004356124586130b8565b9061249161094e825f527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052600160405f20015490565b614b9c565b503461025d57602060031936011261025d5760043567ffffffffffffffff8111610c3557806004019060e06003198236030112610cb5576124d5613850565b6124dd6137dc565b602481013590811561252457612504610b50610b4960646001600160a01b0394018661329a565b169061251a81612513856135aa565b843361406a565b83611b99846135aa565b6024847f4f6df8d000000000000000000000000000000000000000000000000000000000815280600452fd5b503461025d57602060031936011261025d5761076c61256d6130a2565b612575613fe2565b613bf1565b503461025d57602060031936011261025d5760206125c56004355f527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052600160405f20015490565b604051908152f35b503461025d57602060031936011261025d5760246125e96130a2565b60206001600160a01b037f823f7a8ea9ae6df0eb03ec5e1682d7a2839417ad8a91774118e6acf2e8d2f8035416604051938480927f24d7806c0000000000000000000000000000000000000000000000000000000082523360048301525afa91821561065f5761076c92612665918591610a3c5750339061322a565b613b22565b503461025d578060031936011261025d5760206001600160a01b037f823f7a8ea9ae6df0eb03ec5e1682d7a2839417ad8a91774118e6acf2e8d2f8055416604051908152f35b503461025d578060031936011261025d5760206001600160a01b037f823f7a8ea9ae6df0eb03ec5e1682d7a2839417ad8a91774118e6acf2e8d2f8065416604051908152f35b503461025d57612705366130ce565b9061273c336001600160a01b037f823f7a8ea9ae6df0eb03ec5e1682d7a2839417ad8a91774118e6acf2e8d2f8035416331461322a565b6127446137dc565b61274c613850565b606082019161276b610b496127618584613267565b604081019061329a565b6020815191012061277a6132eb565b602081519101201461278a6132eb565b6127976127618685613267565b909215612e63575050506127e96127b4610b496119288685613267565b602081519101206127c361336e565b60208151910120146127d361336e565b906127e16119288786613267565b9290916133a9565b612803610b496127f98584613267565b606081019061329a565b602081519101206128126133ed565b60208151910120146128226133ed565b61282f6127f98685613267565b909215612e2d57505050612883612856610b4961284c8685613267565b602081019061329a565b6020815191012061286561336e565b602081519101201461287561336e565b906127e161284c8786613267565b61289361191961190f8584613267565b9260608401805115612524576128ac60408601516138a3565b9160208401936128c2610b50610b49878461329a565b9651946128f16128d56119288585613267565b6128ec6128e2868061329a565b93909236916131ae565b613a83565b926128fc8488614ec5565b15612a895750509051845195968795909250906020908781189088110287188061292e612929828b613718565b613752565b9803920101602087015e612943855186614f17565b959015868115612a77575b50612a4e575b506001600160a01b03905b16905193813b1561073e576040517f0779afe60000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152921660248301526044820193909352918290606490829084905af18015610fed57612a39575b6108b682604051906129d960408361316f565b601182527f7b22726573756c74223a2241513d3d227d00000000000000000000000000000060208301527f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f005d604051918291602083526020830190613101565b612a4482809261316f565b61025d575f6129c6565b94506001600160a01b0390612a7182612a6688613510565b54169687151561354e565b90612954565b6001600160a01b03915016155f61294e565b612af39350612abe8360209894936128ec6128e2612ab6612aad8d98978998613267565b8781019061329a565b93909461329a565b926040519784899551918291018487015e8401908282018a8152815193849201905e010186815203601f19810185528461316f565b6001600160a01b038516946001600160a01b03612b0f85613510565b5416938415612bb4575b5084956001600160a01b038596959616908351823b15612bb0576040517f40c10f190000000000000000000000000000000000000000000000000000000081526001600160a01b0392909216600483015260248201529085908290604490829084905af1908115610c2a578591612b9b575b50506001600160a01b039061295f565b81612ba59161316f565b61073e57835f612b8b565b8680fd5b93506001600160a01b037f823f7a8ea9ae6df0eb03ec5e1682d7a2839417ad8a91774118e6acf2e8d2f80454166040517f4571e3a600000000000000000000000000000000000000000000000000000000602082015230602482015287604482015260606064820152612c3c81612c2e6084820189613101565b03601f19810183528261316f565b604051916104e48084019084821067ffffffffffffffff831117612e005791849391612c6c9361588686396138f7565b039086f08015610c2a576001600160a01b031695612c8985613510565b6001600160a01b0388167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055612cf4876001600160a01b03165f527f823f7a8ea9ae6df0eb03ec5e1682d7a2839417ad8a91774118e6acf2e8d2f80260205260405f2090565b9685519767ffffffffffffffff891161066a57612d1b89612d1583546135f6565b836136c9565b602098601f8111600114612d9e5780612d4b918a9b8b9a9b91612d93575b505f198260011b9260031b1c19161790565b90555b7f6031fab685dd6d86e4dbac9a69eae347145f332c95b3a0d728d3730fc5233d62612d888298604051918291602083526020830190613101565b0390a2959493612b19565b90508a01515f612d39565b818952898920601f1982168a5b818110612de85750908a9b83600194939c9b9c10612dd0575b5050811b019055612d4e565b8b01515f1960f88460031b161c191690555f80612dc4565b8a8d0151835560209c8d019c60019093019201612dab565b60248a7f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b611093906040519384937fd1ca953a00000000000000000000000000000000000000000000000000000000855260048501613346565b611093906040519384937f094af3b800000000000000000000000000000000000000000000000000000000855260048501613346565b5034612fc6576020600319360112612fc657612eb36130a2565b602460206001600160a01b037f823f7a8ea9ae6df0eb03ec5e1682d7a2839417ad8a91774118e6acf2e8d2f8035416604051928380927f24d7806c0000000000000000000000000000000000000000000000000000000082523360048301525afa8015612fbb57612f2c915f91610a3c5750339061322a565b6001600160a01b037f823f7a8ea9ae6df0eb03ec5e1682d7a2839417ad8a91774118e6acf2e8d2f804541690813b15612fc6576001600160a01b0360245f928360405195869485937f3659cfe60000000000000000000000000000000000000000000000000000000085521660048401525af18015612fbb57612fad575080f35b612fb991505f9061316f565b005b6040513d5f823e3d90fd5b5f80fd5b34612fc6575f600319360112612fc65760206040517f70e300420270129b335fa3e16cefa8088dcfc9433472e4725448f7fc7a82856c8152f35b34612fc6576020600319360112612fc657600435907fffffffff000000000000000000000000000000000000000000000000000000008216809203612fc657817f7965db0b0000000000000000000000000000000000000000000000000000000060209314908115613078575b5015158152f35b7f01ffc9a70000000000000000000000000000000000000000000000000000000091501483613071565b600435906001600160a01b0382168203612fc657565b602435906001600160a01b0382168203612fc657565b6020600319820112612fc6576004359067ffffffffffffffff8211612fc6576003198260a092030112612fc65760040190565b90601f19601f602080948051918291828752018686015e5f8582860101520116010190565b60a0810190811067ffffffffffffffff82111761314257604052565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b90601f601f19910116810190811067ffffffffffffffff82111761314257604052565b67ffffffffffffffff811161314257601f01601f191660200190565b9291926131ba82613192565b916131c8604051938461316f565b829481845281830111612fc6578281602093845f960137010152565b9181601f84011215612fc65782359167ffffffffffffffff8311612fc65760208381860195010111612fc657565b90816020910312612fc657518015158103612fc65790565b156132325750565b6001600160a01b03907f2ecb3242000000000000000000000000000000000000000000000000000000005f521660045260245ffd5b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6181360301821215612fc6570190565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215612fc6570180359067ffffffffffffffff8211612fc657602001918136038313612fc657565b604051906132fa60408361316f565b600782527f69637332302d31000000000000000000000000000000000000000000000000006020830152565b601f8260209493601f1993818652868601375f8582860101520116010190565b9161335d61336b9492604085526040850190613101565b926020818503910152613326565b90565b6040519061337d60408361316f565b600882527f7472616e736665720000000000000000000000000000000000000000000000006020830152565b92909192156133b757505050565b611093906040519384937f5d3a3cdd00000000000000000000000000000000000000000000000000000000855260048501613346565b604051906133fc60408361316f565b601a82527f6170706c69636174696f6e2f782d736f6c69646974792d6162690000000000006020830152565b9080601f83011215612fc65781602061336b933591016131ae565b602081830312612fc65780359067ffffffffffffffff8211612fc6570160a081830312612fc6576040519161347783613126565b813567ffffffffffffffff8111612fc65781613494918401613428565b8352602082013567ffffffffffffffff8111612fc657816134b6918401613428565b6020840152604082013567ffffffffffffffff8111612fc657816134db918401613428565b604084015260608201356060840152608082013567ffffffffffffffff8111612fc6576135089201613428565b608082015290565b60208091604051928184925191829101835e81017f823f7a8ea9ae6df0eb03ec5e1682d7a2839417ad8a91774118e6acf2e8d2f80181520301902090565b156135565750565b611093906040519182917fe1275e2f000000000000000000000000000000000000000000000000000000008352602060048401526024830190613101565b6024356001600160a01b0381168103612fc65790565b356001600160a01b0381168103612fc65790565b60209082604051938492833781017f823f7a8ea9ae6df0eb03ec5e1682d7a2839417ad8a91774118e6acf2e8d2f80181520301902090565b90600182811c9216801561363d575b602083101461361057565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b91607f1691613605565b5f9291815491613656836135f6565b80835292600181169081156136ab575060011461367257505050565b5f9081526020812093945091925b838310613691575060209250010190565b600181602092949394548385870101520191019190613680565b9050602094955060ff1991509291921683830152151560051b010190565b601f82116136d657505050565b5f5260205f20906020601f840160051c8301931061370e575b601f0160051c01905b818110613703575050565b5f81556001016136f8565b90915081906136ef565b9190820391821161372557565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b9061375c82613192565b613769604051918261316f565b828152601f196137798294613192565b0190602036910137565b67ffffffffffffffff81116131425760051b60200190565b80518210156137af5760209160051b010190565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f005c6138285760017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f005d565b7f3ee5aeb5000000000000000000000000000000000000000000000000000000005f5260045ffd5b60ff7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300541661387b57565b7fd93c0665000000000000000000000000000000000000000000000000000000005f5260045ffd5b6138ae815182614f17565b9190156138b9575090565b611093906040519182917f3fed5d87000000000000000000000000000000000000000000000000000000008352602060048401526024830190613101565b6040906001600160a01b0361336b94931681528160208201520190613101565b604051906001600160a01b03815192602081818501958087835e81017f823f7a8ea9ae6df0eb03ec5e1682d7a2839417ad8a91774118e6acf2e8d2f800815203019020541691821561396857505090565b9091506001600160a01b037f823f7a8ea9ae6df0eb03ec5e1682d7a2839417ad8a91774118e6acf2e8d2f8055416604051907fc4d66de8000000000000000000000000000000000000000000000000000000006020830152306024830152602482526139d560448361316f565b604051916104e4908184019284841067ffffffffffffffff851117613142578493613a049361588686396138f7565b03905ff08015612fbb576001600160a01b036020911692604051928391518091835e81017f823f7a8ea9ae6df0eb03ec5e1682d7a2839417ad8a91774118e6acf2e8d2f8008152030190206001600160a01b0382167fffffffffffffffffffffffff000000000000000000000000000000000000000082541617905590565b600180916020809561336b958160405198858a9651918291018688015e8501917f2f0000000000000000000000000000000000000000000000000000000000000085840152602183013701017f2f000000000000000000000000000000000000000000000000000000000000008382015203017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181018452018261316f565b6001600160a01b0381165f9081527f1ba63a4f1661e263bd1a14541416660b96cbd9dd24c2baecb03cee56c5403a14602052604090205460ff1615613bec576001600160a01b03165f8181527f1ba63a4f1661e263bd1a14541416660b96cbd9dd24c2baecb03cee56c5403a1460205260408120805460ff191690553391907f70e300420270129b335fa3e16cefa8088dcfc9433472e4725448f7fc7a82856c907ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9080a4600190565b505f90565b6001600160a01b0381165f9081527f475b312747b0505dfd322c59063ca43b615bb2e6d10fdf52fb58877bbcde4011602052604090205460ff1615613bec576001600160a01b03165f8181527f475b312747b0505dfd322c59063ca43b615bb2e6d10fdf52fb58877bbcde401160205260408120805460ff191690553391907f427da25fe773164f88948d3e215c94b6554e2ed5e5f203a821c9f2f6131cf75a907ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9080a4600190565b6001600160a01b0381165f9081527fd62723b73fed5fb0cae58785027a02356c6bdb68a107a140d6093b6871b6d9f4602052604090205460ff1615613bec576001600160a01b03165f8181527fd62723b73fed5fb0cae58785027a02356c6bdb68a107a140d6093b6871b6d9f460205260408120805460ff191690553391907f62150a51582c26f4255242a3c4ca35fb04250e7315069523d650676aed01a56a907ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9080a4600190565b6001600160a01b0381165f9081527f6af8e566f3cdd77e5fe3b944a8a260d555073fdf834bbffeb5e923cd960fa9b7602052604090205460ff1615613bec576001600160a01b03165f8181527f6af8e566f3cdd77e5fe3b944a8a260d555073fdf834bbffeb5e923cd960fa9b760205260408120805460ff191690553391907f04c6d412b3efc8acc11f9a5a5dc9f2a5d9ae1f6d869984b47386662ab77133cd907ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9080a4600190565b6001600160a01b0381165f9081527f75442b0a96088b5456bc4ed01394c96a4feec0f883c9494257d76b96ab1c9b6b602052604090205460ff1615613bec576001600160a01b03165f8181527f75442b0a96088b5456bc4ed01394c96a4feec0f883c9494257d76b96ab1c9b6b60205260408120805460ff191690553391907f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a907ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9080a4600190565b805f527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b62680060205260405f206001600160a01b0383165f5260205260ff60405f2054165f14613fdc57805f527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b62680060205260405f206001600160a01b0383165f5260205260405f2060ff1981541690556001600160a01b03339216907ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b5f80a4600190565b50505f90565b602460206001600160a01b037f823f7a8ea9ae6df0eb03ec5e1682d7a2839417ad8a91774118e6acf2e8d2f8035416604051928380927f24d7806c0000000000000000000000000000000000000000000000000000000082523360048301525afa8015612fbb5761405b915f91610a3c5750339061322a565b565b9190820180921161372557565b916001600160a01b03909391931692604051927f70a082310000000000000000000000000000000000000000000000000000000084526001600160a01b03821691826004860152602085602481895afa948515612fbb575f95614243575b506040517f23b872dd0000000000000000000000000000000000000000000000000000000060208281019182526001600160a01b039485166024840152939092166044820152606481018590525f91906141258160848101612c2e565b519082885af115612fbb575f513d61423a5750833b155b61420e576020906024604051809681937f70a0823100000000000000000000000000000000000000000000000000000000835260048301525afa928315612fbb575f936141d8575b5061418f908261405d565b908211806141cf575b156141a1575050565b7f2fb30cfc000000000000000000000000000000000000000000000000000000005f5260045260245260445ffd5b50808214614198565b9092506020813d602011614206575b816141f46020938361316f565b81010312612fc657519161418f614184565b3d91506141e7565b837f5274afe7000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b6001141561413c565b9094506020813d602011614270575b8161425f6020938361316f565b81010312612fc657519360206140c8565b3d9150614252565b614299614287610859836135aa565b926142a05f9460405193848092613647565b038361316f565b818051155f146146625750505060a06142c96142c36142be846135aa565b614fcd565b94614fcd565b936142d7604084018461329a565b93909561432861430d6142ed60c085018561329a565b999097604051966142fd88613126565b87526020870194855236916131ae565b9560408501968752606085019860208501358a5236916131ae565b94608084019586526001600160a01b037f823f7a8ea9ae6df0eb03ec5e1682d7a2839417ad8a91774118e6acf2e8d2f803541695614369606085018561329a565b96909401359467ffffffffffffffff861680960361465e579061447d6143d694939261446f61439661336e565b9c61439f61336e565b946144386143ab6132eb565b976144076143b76133ed565b9a6040519c8d986020808b01525160a060408b015260e08a0190613101565b90517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08983030160608a0152613101565b90517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0878303016080880152613101565b915160a0850152517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08483030160c0850152613101565b03601f19810186528561316f565b6040519961448a8b613126565b8a5260208a0152604089015260608801526080870152604051926060840184811067ffffffffffffffff82111761463157936145cb602096946144e0614534958a9567ffffffffffffffff9960405236916131ae565b835287830190815260408301998a52604051998a97889687957f4d6e7ce30000000000000000000000000000000000000000000000000000000087528b600488015251606060248801526084870190613101565b925116604485015251907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc84820301606485015260806145ba6145a8614596614586865160a0875260a0870190613101565b8d8701518682038f880152613101565b60408601518582036040870152613101565b60608501518482036060860152613101565b920151906080818403910152613101565b03925af19182156146245781926145e157505090565b9091506020813d60201161461c575b816145fd6020938361316f565b81010312610c3557519067ffffffffffffffff8216820361025d575090565b3d91506145f0565b50604051903d90823e3d90fd5b6024877f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b8880fd5b61468d9061468761467497949761336e565b614681606088018861329a565b91613a83565b90614ec5565b61469e575b506142c960a091614fcd565b6001600160a01b036146af846135aa565b16803b15612fc6576040517f9dc29fac0000000000000000000000000000000000000000000000000000000081526001600160a01b03929092166004830152602084013560248301525f908290604490829084905af18015612fbb57156146925761471d9193505f9061316f565b5f916142c9614692565b805f527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b62680060205260405f206001600160a01b0333165f5260205260ff60405f205416156147715750565b7fe2517d3f000000000000000000000000000000000000000000000000000000005f523360045260245260445ffd5b6001600160a01b0381165f9081527f475b312747b0505dfd322c59063ca43b615bb2e6d10fdf52fb58877bbcde4011602052604090205460ff16613bec576001600160a01b03165f8181527f475b312747b0505dfd322c59063ca43b615bb2e6d10fdf52fb58877bbcde401160205260408120805460ff191660011790553391907f427da25fe773164f88948d3e215c94b6554e2ed5e5f203a821c9f2f6131cf75a907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d9080a4600190565b6001600160a01b0381165f9081527f6af8e566f3cdd77e5fe3b944a8a260d555073fdf834bbffeb5e923cd960fa9b7602052604090205460ff16613bec576001600160a01b03165f8181527f6af8e566f3cdd77e5fe3b944a8a260d555073fdf834bbffeb5e923cd960fa9b760205260408120805460ff191660011790553391907f04c6d412b3efc8acc11f9a5a5dc9f2a5d9ae1f6d869984b47386662ab77133cd907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d9080a4600190565b6001600160a01b0381165f9081527f1ba63a4f1661e263bd1a14541416660b96cbd9dd24c2baecb03cee56c5403a14602052604090205460ff16613bec576001600160a01b03165f8181527f1ba63a4f1661e263bd1a14541416660b96cbd9dd24c2baecb03cee56c5403a1460205260408120805460ff191660011790553391907f70e300420270129b335fa3e16cefa8088dcfc9433472e4725448f7fc7a82856c907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d9080a4600190565b6001600160a01b0381165f9081527fd62723b73fed5fb0cae58785027a02356c6bdb68a107a140d6093b6871b6d9f4602052604090205460ff16613bec576001600160a01b03165f8181527fd62723b73fed5fb0cae58785027a02356c6bdb68a107a140d6093b6871b6d9f460205260408120805460ff191660011790553391907f62150a51582c26f4255242a3c4ca35fb04250e7315069523d650676aed01a56a907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d9080a4600190565b6001600160a01b0381165f9081527f75442b0a96088b5456bc4ed01394c96a4feec0f883c9494257d76b96ab1c9b6b602052604090205460ff16613bec576001600160a01b03165f8181527f75442b0a96088b5456bc4ed01394c96a4feec0f883c9494257d76b96ab1c9b6b60205260408120805460ff191660011790553391907f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d9080a4600190565b805f527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b62680060205260405f206001600160a01b0383165f5260205260ff60405f205416155f14613fdc57805f527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b62680060205260405f206001600160a01b0383165f5260205260405f20600160ff198254161790556001600160a01b03339216907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d5f80a4600190565b909193925f946001600160a01b03604051838382376020818581017f823f7a8ea9ae6df0eb03ec5e1682d7a2839417ad8a91774118e6acf2e8d2f8008152030190205416928315614e425791614ccf916128ec614cd694614cc760208a01516138a3565b9736916131ae565b8451614ec5565b15614df2576001600160a01b03614ced8451613510565b541692614cfd815185151561354e565b6060810151843b15612fc6576040517f40c10f190000000000000000000000000000000000000000000000000000000081526001600160a01b038416600482015260248101919091525f8160448183895af18015612fbb57614ddc575b506060905b015192813b1561073a576040517f0779afe60000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201529216602483015260448201929092529082908290606490829084905af18015610fed57614dca575050565b614dd582809261316f565b61025d5750565b614de99195505f9061316f565b5f936060614d5a565b6001600160a01b03614e048451613510565b5416928315614e16575b606090614d5f565b92506060614e2484516138a3565b93614e3b81516001600160a01b038716151561354e565b9050614e0e565b50906110936040519283927f5778f378000000000000000000000000000000000000000000000000000000008452602060048501526024840191613326565b5f8061336b93602081519101845af43d15614ebd573d91614ea183613192565b92614eaf604051948561316f565b83523d5f602085013e6150bc565b6060916150bc565b8051908251809210614f10578051918280821091180280831892141582028218906020614ef56129298486613718565b92808285019503920101835e51902090602081519101201490565b5050505f90565b805182118015614fc6575b614f6f576001821180614f77575b158015908160011b91820460021417156137255760280180602811613725578203614f6f576001600160a01b0392915f614f699261519f565b90921690565b50505f905f90565b507f30780000000000000000000000000000000000000000000000000000000000007fffff00000000000000000000000000000000000000000000000000000000000060208301511614614f30565b505f614f22565b6001600160a01b031680614fe1602a613192565b91614fef604051938461316f565b602a8352614ffd602a613192565b601f196020850191013682378351156137af57603090538251600110156137af576078602184015360295b600181116150695750615039575090565b7fe22e27eb000000000000000000000000000000000000000000000000000000005f52600452601460245260445ffd5b90600f811660108110156137af5784518310156137af577f3031323334353637383961626364656600000000000000000000000000000000901a8483016020015360041c908015613725575f1901615028565b906150f957508051156150d157805190602001fd5b7fd6bda275000000000000000000000000000000000000000000000000000000005f5260045ffd5b8151158061513f575b61510a575090565b6001600160a01b03907f9996b315000000000000000000000000000000000000000000000000000000005f521660045260245ffd5b50803b15615102565b60ff7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005460401c161561517757565b7fd7e6bcf8000000000000000000000000000000000000000000000000000000005f5260045ffd5b9290926001840180851161372557831180615256575b158015908160011b9182046002141715613725576151d8905f949293949561405d565b915b8183106151ea5750505060019190565b9092919360ff6152217fff0000000000000000000000000000000000000000000000000000000000000060208886010151166152a7565b16600f811161524b578160041b9180830460101490151715613725576001910194019192906151da565b505f94508493505050565b507f30780000000000000000000000000000000000000000000000000000000000007fffff0000000000000000000000000000000000000000000000000000000000006020868401015116146151b5565b60f81c602f811180615369575b156152e1577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffd00160ff1690565b606081118061535f575b15615318577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa90160ff1690565b6040811180615355575b1561534f577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc90160ff1690565b5060ff90565b5060478110615322565b50606781106152eb565b50603a81106152b456fe60803461013457601f61051238819003918201601f19168301916001600160401b03831184841017610138578084926040948552833981010312610134576100468161014c565b906001600160a01b039061005c9060200161014c565b16908115610121575f80546001600160a01b031981168417825560405193916001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09080a3803b1561010157600180546001600160a01b0319166001600160a01b039290921691821790557fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b5f80a26103b190816101618239f35b63211eb15960e21b5f9081526001600160a01b0391909116600452602490fd5b631e4fbdf760e01b5f525f60045260245ffd5b5f80fd5b634e487b7160e01b5f52604160045260245ffd5b51906001600160a01b03821682036101345756fe60806040526004361015610011575f80fd5b5f3560e01c80633659cfe61461027e5780635c60da1b1461022d578063715018a6146101935780638da5cb5b146101435763f2fde38b14610050575f80fd5b3461013f5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261013f5760043573ffffffffffffffffffffffffffffffffffffffff811680910361013f576100a8610358565b80156101135773ffffffffffffffffffffffffffffffffffffffff5f54827fffffffffffffffffffffffff00000000000000000000000000000000000000008216175f55167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e05f80a3005b7f1e4fbdf7000000000000000000000000000000000000000000000000000000005f525f60045260245ffd5b5f80fd5b3461013f575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261013f57602073ffffffffffffffffffffffffffffffffffffffff5f5416604051908152f35b3461013f575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261013f576101c9610358565b5f73ffffffffffffffffffffffffffffffffffffffff81547fffffffffffffffffffffffff000000000000000000000000000000000000000081168355167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a3005b3461013f575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261013f57602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b3461013f5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261013f5760043573ffffffffffffffffffffffffffffffffffffffff81169081810361013f576102d7610358565b3b1561032d57807fffffffffffffffffffffffff000000000000000000000000000000000000000060015416176001557fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b5f80a2005b7f847ac564000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b73ffffffffffffffffffffffffffffffffffffffff5f5416330361037857565b7f118cdaa7000000000000000000000000000000000000000000000000000000005f523360045260245ffdfea164736f6c634300081c000a60a0806040526104e480380380916100178285610292565b833981016040828203126101eb5761002e826102c9565b602083015190926001600160401b0382116101eb57019080601f830112156101eb57815161005b816102dd565b926100696040519485610292565b8184526020840192602083830101116101eb57815f926020809301855e84010152823b15610274577fa3f0ad74e5423aebfd80d3ef4346578335a9a72aeaee59ff6cb3582b35133d5080546001600160a01b0319166001600160a01b038516908117909155604051635c60da1b60e01b8152909190602081600481865afa9081156101f7575f9161023a575b50803b1561021a5750817f1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e5f80a282511561020257602060049260405193848092635c60da1b60e01b82525afa9182156101f7575f926101ae575b505f809161018a945190845af43d156101a6573d9161016e836102dd565b9261017c6040519485610292565b83523d5f602085013e6102f8565b505b60805260405161018d908161035782396080518160460152f35b6060916102f8565b9291506020833d6020116101ef575b816101ca60209383610292565b810103126101eb575f80916101e161018a956102c9565b9394509150610150565b5f80fd5b3d91506101bd565b6040513d5f823e3d90fd5b505050341561018c5763b398979f60e01b5f5260045ffd5b634c9c8ce360e01b5f9081526001600160a01b0391909116600452602490fd5b90506020813d60201161026c575b8161025560209383610292565b810103126101eb57610266906102c9565b5f6100f5565b3d9150610248565b631933b43b60e21b5f9081526001600160a01b038416600452602490fd5b601f909101601f19168101906001600160401b038211908210176102b557604052565b634e487b7160e01b5f52604160045260245ffd5b51906001600160a01b03821682036101eb57565b6001600160401b0381116102b557601f01601f191660200190565b9061031c575080511561030d57805190602001fd5b63d6bda27560e01b5f5260045ffd5b8151158061034d575b61032d575090565b639996b31560e01b5f9081526001600160a01b0391909116600452602490fd5b50803b1561032556fe60806040527f5c60da1b000000000000000000000000000000000000000000000000000000006080526020608060048173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa8015610107575f9015610163575060203d602011610100575b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f820116608001906080821067ffffffffffffffff8311176100d3576100ce91604052608001610112565b610163565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b503d610081565b6040513d5f823e3d90fd5b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80602091011261015f5760805173ffffffffffffffffffffffffffffffffffffffff8116810361015f5790565b5f80fd5b5f8091368280378136915af43d5f803e1561017c573d5ff35b3d5ffdfea164736f6c634300081c000aa164736f6c634300081c000af0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00",
}

// ContractICS20TransferABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractICS20TransferMetaData.ABI instead.
var ContractICS20TransferABI = ContractICS20TransferMetaData.ABI

// ContractICS20TransferBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractICS20TransferMetaData.Bin instead.
var ContractICS20TransferBin = ContractICS20TransferMetaData.Bin

// DeployContractICS20Transfer deploys a new Ethereum contract, binding an instance of ContractICS20Transfer to it.
func DeployContractICS20Transfer(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ContractICS20Transfer, error) {
	parsed, err := ContractICS20TransferMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractICS20TransferBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractICS20Transfer{ContractICS20TransferCaller: ContractICS20TransferCaller{contract: contract}, ContractICS20TransferTransactor: ContractICS20TransferTransactor{contract: contract}, ContractICS20TransferFilterer: ContractICS20TransferFilterer{contract: contract}}, nil
}

// ContractICS20Transfer is an auto generated Go binding around an Ethereum contract.
type ContractICS20Transfer struct {
	ContractICS20TransferCaller     // Read-only binding to the contract
	ContractICS20TransferTransactor // Write-only binding to the contract
	ContractICS20TransferFilterer   // Log filterer for contract events
}

// ContractICS20TransferCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractICS20TransferCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractICS20TransferTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractICS20TransferTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractICS20TransferFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractICS20TransferFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractICS20TransferSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractICS20TransferSession struct {
	Contract     *ContractICS20Transfer // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ContractICS20TransferCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractICS20TransferCallerSession struct {
	Contract *ContractICS20TransferCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// ContractICS20TransferTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractICS20TransferTransactorSession struct {
	Contract     *ContractICS20TransferTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// ContractICS20TransferRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractICS20TransferRaw struct {
	Contract *ContractICS20Transfer // Generic contract binding to access the raw methods on
}

// ContractICS20TransferCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractICS20TransferCallerRaw struct {
	Contract *ContractICS20TransferCaller // Generic read-only contract binding to access the raw methods on
}

// ContractICS20TransferTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractICS20TransferTransactorRaw struct {
	Contract *ContractICS20TransferTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractICS20Transfer creates a new instance of ContractICS20Transfer, bound to a specific deployed contract.
func NewContractICS20Transfer(address common.Address, backend bind.ContractBackend) (*ContractICS20Transfer, error) {
	contract, err := bindContractICS20Transfer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractICS20Transfer{ContractICS20TransferCaller: ContractICS20TransferCaller{contract: contract}, ContractICS20TransferTransactor: ContractICS20TransferTransactor{contract: contract}, ContractICS20TransferFilterer: ContractICS20TransferFilterer{contract: contract}}, nil
}

// NewContractICS20TransferCaller creates a new read-only instance of ContractICS20Transfer, bound to a specific deployed contract.
func NewContractICS20TransferCaller(address common.Address, caller bind.ContractCaller) (*ContractICS20TransferCaller, error) {
	contract, err := bindContractICS20Transfer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractICS20TransferCaller{contract: contract}, nil
}

// NewContractICS20TransferTransactor creates a new write-only instance of ContractICS20Transfer, bound to a specific deployed contract.
func NewContractICS20TransferTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractICS20TransferTransactor, error) {
	contract, err := bindContractICS20Transfer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractICS20TransferTransactor{contract: contract}, nil
}

// NewContractICS20TransferFilterer creates a new log filterer instance of ContractICS20Transfer, bound to a specific deployed contract.
func NewContractICS20TransferFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractICS20TransferFilterer, error) {
	contract, err := bindContractICS20Transfer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractICS20TransferFilterer{contract: contract}, nil
}

// bindContractICS20Transfer binds a generic wrapper to an already deployed contract.
func bindContractICS20Transfer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractICS20TransferMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractICS20Transfer *ContractICS20TransferRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractICS20Transfer.Contract.ContractICS20TransferCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractICS20Transfer *ContractICS20TransferRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractICS20Transfer.Contract.ContractICS20TransferTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractICS20Transfer *ContractICS20TransferRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractICS20Transfer.Contract.ContractICS20TransferTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractICS20Transfer *ContractICS20TransferCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractICS20Transfer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractICS20Transfer *ContractICS20TransferTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractICS20Transfer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractICS20Transfer *ContractICS20TransferTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractICS20Transfer.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ContractICS20Transfer *ContractICS20TransferCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ContractICS20Transfer.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ContractICS20Transfer *ContractICS20TransferSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ContractICS20Transfer.Contract.DEFAULTADMINROLE(&_ContractICS20Transfer.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ContractICS20Transfer *ContractICS20TransferCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ContractICS20Transfer.Contract.DEFAULTADMINROLE(&_ContractICS20Transfer.CallOpts)
}

// DELEGATESENDERROLE is a free data retrieval call binding the contract method 0xda0f21b5.
//
// Solidity: function DELEGATE_SENDER_ROLE() view returns(bytes32)
func (_ContractICS20Transfer *ContractICS20TransferCaller) DELEGATESENDERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ContractICS20Transfer.contract.Call(opts, &out, "DELEGATE_SENDER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DELEGATESENDERROLE is a free data retrieval call binding the contract method 0xda0f21b5.
//
// Solidity: function DELEGATE_SENDER_ROLE() view returns(bytes32)
func (_ContractICS20Transfer *ContractICS20TransferSession) DELEGATESENDERROLE() ([32]byte, error) {
	return _ContractICS20Transfer.Contract.DELEGATESENDERROLE(&_ContractICS20Transfer.CallOpts)
}

// DELEGATESENDERROLE is a free data retrieval call binding the contract method 0xda0f21b5.
//
// Solidity: function DELEGATE_SENDER_ROLE() view returns(bytes32)
func (_ContractICS20Transfer *ContractICS20TransferCallerSession) DELEGATESENDERROLE() ([32]byte, error) {
	return _ContractICS20Transfer.Contract.DELEGATESENDERROLE(&_ContractICS20Transfer.CallOpts)
}

// ERC20CUSTOMIZERROLE is a free data retrieval call binding the contract method 0x02769d73.
//
// Solidity: function ERC20_CUSTOMIZER_ROLE() view returns(bytes32)
func (_ContractICS20Transfer *ContractICS20TransferCaller) ERC20CUSTOMIZERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ContractICS20Transfer.contract.Call(opts, &out, "ERC20_CUSTOMIZER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ERC20CUSTOMIZERROLE is a free data retrieval call binding the contract method 0x02769d73.
//
// Solidity: function ERC20_CUSTOMIZER_ROLE() view returns(bytes32)
func (_ContractICS20Transfer *ContractICS20TransferSession) ERC20CUSTOMIZERROLE() ([32]byte, error) {
	return _ContractICS20Transfer.Contract.ERC20CUSTOMIZERROLE(&_ContractICS20Transfer.CallOpts)
}

// ERC20CUSTOMIZERROLE is a free data retrieval call binding the contract method 0x02769d73.
//
// Solidity: function ERC20_CUSTOMIZER_ROLE() view returns(bytes32)
func (_ContractICS20Transfer *ContractICS20TransferCallerSession) ERC20CUSTOMIZERROLE() ([32]byte, error) {
	return _ContractICS20Transfer.Contract.ERC20CUSTOMIZERROLE(&_ContractICS20Transfer.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_ContractICS20Transfer *ContractICS20TransferCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ContractICS20Transfer.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_ContractICS20Transfer *ContractICS20TransferSession) PAUSERROLE() ([32]byte, error) {
	return _ContractICS20Transfer.Contract.PAUSERROLE(&_ContractICS20Transfer.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_ContractICS20Transfer *ContractICS20TransferCallerSession) PAUSERROLE() ([32]byte, error) {
	return _ContractICS20Transfer.Contract.PAUSERROLE(&_ContractICS20Transfer.CallOpts)
}

// TOKENOPERATORROLE is a free data retrieval call binding the contract method 0x33daf980.
//
// Solidity: function TOKEN_OPERATOR_ROLE() view returns(bytes32)
func (_ContractICS20Transfer *ContractICS20TransferCaller) TOKENOPERATORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ContractICS20Transfer.contract.Call(opts, &out, "TOKEN_OPERATOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TOKENOPERATORROLE is a free data retrieval call binding the contract method 0x33daf980.
//
// Solidity: function TOKEN_OPERATOR_ROLE() view returns(bytes32)
func (_ContractICS20Transfer *ContractICS20TransferSession) TOKENOPERATORROLE() ([32]byte, error) {
	return _ContractICS20Transfer.Contract.TOKENOPERATORROLE(&_ContractICS20Transfer.CallOpts)
}

// TOKENOPERATORROLE is a free data retrieval call binding the contract method 0x33daf980.
//
// Solidity: function TOKEN_OPERATOR_ROLE() view returns(bytes32)
func (_ContractICS20Transfer *ContractICS20TransferCallerSession) TOKENOPERATORROLE() ([32]byte, error) {
	return _ContractICS20Transfer.Contract.TOKENOPERATORROLE(&_ContractICS20Transfer.CallOpts)
}

// UNPAUSERROLE is a free data retrieval call binding the contract method 0xfb1bb9de.
//
// Solidity: function UNPAUSER_ROLE() view returns(bytes32)
func (_ContractICS20Transfer *ContractICS20TransferCaller) UNPAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ContractICS20Transfer.contract.Call(opts, &out, "UNPAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// UNPAUSERROLE is a free data retrieval call binding the contract method 0xfb1bb9de.
//
// Solidity: function UNPAUSER_ROLE() view returns(bytes32)
func (_ContractICS20Transfer *ContractICS20TransferSession) UNPAUSERROLE() ([32]byte, error) {
	return _ContractICS20Transfer.Contract.UNPAUSERROLE(&_ContractICS20Transfer.CallOpts)
}

// UNPAUSERROLE is a free data retrieval call binding the contract method 0xfb1bb9de.
//
// Solidity: function UNPAUSER_ROLE() view returns(bytes32)
func (_ContractICS20Transfer *ContractICS20TransferCallerSession) UNPAUSERROLE() ([32]byte, error) {
	return _ContractICS20Transfer.Contract.UNPAUSERROLE(&_ContractICS20Transfer.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_ContractICS20Transfer *ContractICS20TransferCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ContractICS20Transfer.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_ContractICS20Transfer *ContractICS20TransferSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _ContractICS20Transfer.Contract.UPGRADEINTERFACEVERSION(&_ContractICS20Transfer.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_ContractICS20Transfer *ContractICS20TransferCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _ContractICS20Transfer.Contract.UPGRADEINTERFACEVERSION(&_ContractICS20Transfer.CallOpts)
}

// GetEscrow is a free data retrieval call binding the contract method 0x969631d5.
//
// Solidity: function getEscrow(string clientId) view returns(address)
func (_ContractICS20Transfer *ContractICS20TransferCaller) GetEscrow(opts *bind.CallOpts, clientId string) (common.Address, error) {
	var out []interface{}
	err := _ContractICS20Transfer.contract.Call(opts, &out, "getEscrow", clientId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetEscrow is a free data retrieval call binding the contract method 0x969631d5.
//
// Solidity: function getEscrow(string clientId) view returns(address)
func (_ContractICS20Transfer *ContractICS20TransferSession) GetEscrow(clientId string) (common.Address, error) {
	return _ContractICS20Transfer.Contract.GetEscrow(&_ContractICS20Transfer.CallOpts, clientId)
}

// GetEscrow is a free data retrieval call binding the contract method 0x969631d5.
//
// Solidity: function getEscrow(string clientId) view returns(address)
func (_ContractICS20Transfer *ContractICS20TransferCallerSession) GetEscrow(clientId string) (common.Address, error) {
	return _ContractICS20Transfer.Contract.GetEscrow(&_ContractICS20Transfer.CallOpts, clientId)
}

// GetEscrowBeacon is a free data retrieval call binding the contract method 0x1e5150e4.
//
// Solidity: function getEscrowBeacon() view returns(address)
func (_ContractICS20Transfer *ContractICS20TransferCaller) GetEscrowBeacon(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractICS20Transfer.contract.Call(opts, &out, "getEscrowBeacon")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetEscrowBeacon is a free data retrieval call binding the contract method 0x1e5150e4.
//
// Solidity: function getEscrowBeacon() view returns(address)
func (_ContractICS20Transfer *ContractICS20TransferSession) GetEscrowBeacon() (common.Address, error) {
	return _ContractICS20Transfer.Contract.GetEscrowBeacon(&_ContractICS20Transfer.CallOpts)
}

// GetEscrowBeacon is a free data retrieval call binding the contract method 0x1e5150e4.
//
// Solidity: function getEscrowBeacon() view returns(address)
func (_ContractICS20Transfer *ContractICS20TransferCallerSession) GetEscrowBeacon() (common.Address, error) {
	return _ContractICS20Transfer.Contract.GetEscrowBeacon(&_ContractICS20Transfer.CallOpts)
}

// GetIBCERC20Beacon is a free data retrieval call binding the contract method 0x826cae7a.
//
// Solidity: function getIBCERC20Beacon() view returns(address)
func (_ContractICS20Transfer *ContractICS20TransferCaller) GetIBCERC20Beacon(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractICS20Transfer.contract.Call(opts, &out, "getIBCERC20Beacon")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetIBCERC20Beacon is a free data retrieval call binding the contract method 0x826cae7a.
//
// Solidity: function getIBCERC20Beacon() view returns(address)
func (_ContractICS20Transfer *ContractICS20TransferSession) GetIBCERC20Beacon() (common.Address, error) {
	return _ContractICS20Transfer.Contract.GetIBCERC20Beacon(&_ContractICS20Transfer.CallOpts)
}

// GetIBCERC20Beacon is a free data retrieval call binding the contract method 0x826cae7a.
//
// Solidity: function getIBCERC20Beacon() view returns(address)
func (_ContractICS20Transfer *ContractICS20TransferCallerSession) GetIBCERC20Beacon() (common.Address, error) {
	return _ContractICS20Transfer.Contract.GetIBCERC20Beacon(&_ContractICS20Transfer.CallOpts)
}

// GetPermit2 is a free data retrieval call binding the contract method 0x1bbf2e23.
//
// Solidity: function getPermit2() view returns(address)
func (_ContractICS20Transfer *ContractICS20TransferCaller) GetPermit2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractICS20Transfer.contract.Call(opts, &out, "getPermit2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPermit2 is a free data retrieval call binding the contract method 0x1bbf2e23.
//
// Solidity: function getPermit2() view returns(address)
func (_ContractICS20Transfer *ContractICS20TransferSession) GetPermit2() (common.Address, error) {
	return _ContractICS20Transfer.Contract.GetPermit2(&_ContractICS20Transfer.CallOpts)
}

// GetPermit2 is a free data retrieval call binding the contract method 0x1bbf2e23.
//
// Solidity: function getPermit2() view returns(address)
func (_ContractICS20Transfer *ContractICS20TransferCallerSession) GetPermit2() (common.Address, error) {
	return _ContractICS20Transfer.Contract.GetPermit2(&_ContractICS20Transfer.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ContractICS20Transfer *ContractICS20TransferCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ContractICS20Transfer.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ContractICS20Transfer *ContractICS20TransferSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ContractICS20Transfer.Contract.GetRoleAdmin(&_ContractICS20Transfer.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ContractICS20Transfer *ContractICS20TransferCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ContractICS20Transfer.Contract.GetRoleAdmin(&_ContractICS20Transfer.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ContractICS20Transfer *ContractICS20TransferCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _ContractICS20Transfer.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ContractICS20Transfer *ContractICS20TransferSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ContractICS20Transfer.Contract.HasRole(&_ContractICS20Transfer.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ContractICS20Transfer *ContractICS20TransferCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ContractICS20Transfer.Contract.HasRole(&_ContractICS20Transfer.CallOpts, role, account)
}

// IbcERC20Contract is a free data retrieval call binding the contract method 0xa50ee2b4.
//
// Solidity: function ibcERC20Contract(string denom) view returns(address)
func (_ContractICS20Transfer *ContractICS20TransferCaller) IbcERC20Contract(opts *bind.CallOpts, denom string) (common.Address, error) {
	var out []interface{}
	err := _ContractICS20Transfer.contract.Call(opts, &out, "ibcERC20Contract", denom)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IbcERC20Contract is a free data retrieval call binding the contract method 0xa50ee2b4.
//
// Solidity: function ibcERC20Contract(string denom) view returns(address)
func (_ContractICS20Transfer *ContractICS20TransferSession) IbcERC20Contract(denom string) (common.Address, error) {
	return _ContractICS20Transfer.Contract.IbcERC20Contract(&_ContractICS20Transfer.CallOpts, denom)
}

// IbcERC20Contract is a free data retrieval call binding the contract method 0xa50ee2b4.
//
// Solidity: function ibcERC20Contract(string denom) view returns(address)
func (_ContractICS20Transfer *ContractICS20TransferCallerSession) IbcERC20Contract(denom string) (common.Address, error) {
	return _ContractICS20Transfer.Contract.IbcERC20Contract(&_ContractICS20Transfer.CallOpts, denom)
}

// IbcERC20Denom is a free data retrieval call binding the contract method 0xe163b1af.
//
// Solidity: function ibcERC20Denom(address token) view returns(string)
func (_ContractICS20Transfer *ContractICS20TransferCaller) IbcERC20Denom(opts *bind.CallOpts, token common.Address) (string, error) {
	var out []interface{}
	err := _ContractICS20Transfer.contract.Call(opts, &out, "ibcERC20Denom", token)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// IbcERC20Denom is a free data retrieval call binding the contract method 0xe163b1af.
//
// Solidity: function ibcERC20Denom(address token) view returns(string)
func (_ContractICS20Transfer *ContractICS20TransferSession) IbcERC20Denom(token common.Address) (string, error) {
	return _ContractICS20Transfer.Contract.IbcERC20Denom(&_ContractICS20Transfer.CallOpts, token)
}

// IbcERC20Denom is a free data retrieval call binding the contract method 0xe163b1af.
//
// Solidity: function ibcERC20Denom(address token) view returns(string)
func (_ContractICS20Transfer *ContractICS20TransferCallerSession) IbcERC20Denom(token common.Address) (string, error) {
	return _ContractICS20Transfer.Contract.IbcERC20Denom(&_ContractICS20Transfer.CallOpts, token)
}

// Ics26 is a free data retrieval call binding the contract method 0xd413227d.
//
// Solidity: function ics26() view returns(address)
func (_ContractICS20Transfer *ContractICS20TransferCaller) Ics26(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractICS20Transfer.contract.Call(opts, &out, "ics26")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Ics26 is a free data retrieval call binding the contract method 0xd413227d.
//
// Solidity: function ics26() view returns(address)
func (_ContractICS20Transfer *ContractICS20TransferSession) Ics26() (common.Address, error) {
	return _ContractICS20Transfer.Contract.Ics26(&_ContractICS20Transfer.CallOpts)
}

// Ics26 is a free data retrieval call binding the contract method 0xd413227d.
//
// Solidity: function ics26() view returns(address)
func (_ContractICS20Transfer *ContractICS20TransferCallerSession) Ics26() (common.Address, error) {
	return _ContractICS20Transfer.Contract.Ics26(&_ContractICS20Transfer.CallOpts)
}

// IsTokenOperator is a free data retrieval call binding the contract method 0xf17a7e66.
//
// Solidity: function isTokenOperator(address account) view returns(bool)
func (_ContractICS20Transfer *ContractICS20TransferCaller) IsTokenOperator(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _ContractICS20Transfer.contract.Call(opts, &out, "isTokenOperator", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTokenOperator is a free data retrieval call binding the contract method 0xf17a7e66.
//
// Solidity: function isTokenOperator(address account) view returns(bool)
func (_ContractICS20Transfer *ContractICS20TransferSession) IsTokenOperator(account common.Address) (bool, error) {
	return _ContractICS20Transfer.Contract.IsTokenOperator(&_ContractICS20Transfer.CallOpts, account)
}

// IsTokenOperator is a free data retrieval call binding the contract method 0xf17a7e66.
//
// Solidity: function isTokenOperator(address account) view returns(bool)
func (_ContractICS20Transfer *ContractICS20TransferCallerSession) IsTokenOperator(account common.Address) (bool, error) {
	return _ContractICS20Transfer.Contract.IsTokenOperator(&_ContractICS20Transfer.CallOpts, account)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ContractICS20Transfer *ContractICS20TransferCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ContractICS20Transfer.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ContractICS20Transfer *ContractICS20TransferSession) Paused() (bool, error) {
	return _ContractICS20Transfer.Contract.Paused(&_ContractICS20Transfer.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ContractICS20Transfer *ContractICS20TransferCallerSession) Paused() (bool, error) {
	return _ContractICS20Transfer.Contract.Paused(&_ContractICS20Transfer.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ContractICS20Transfer *ContractICS20TransferCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ContractICS20Transfer.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ContractICS20Transfer *ContractICS20TransferSession) ProxiableUUID() ([32]byte, error) {
	return _ContractICS20Transfer.Contract.ProxiableUUID(&_ContractICS20Transfer.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ContractICS20Transfer *ContractICS20TransferCallerSession) ProxiableUUID() ([32]byte, error) {
	return _ContractICS20Transfer.Contract.ProxiableUUID(&_ContractICS20Transfer.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ContractICS20Transfer *ContractICS20TransferCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ContractICS20Transfer.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ContractICS20Transfer *ContractICS20TransferSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ContractICS20Transfer.Contract.SupportsInterface(&_ContractICS20Transfer.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ContractICS20Transfer *ContractICS20TransferCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ContractICS20Transfer.Contract.SupportsInterface(&_ContractICS20Transfer.CallOpts, interfaceId)
}

// GrantDelegateSenderRole is a paid mutator transaction binding the contract method 0x4551d2b0.
//
// Solidity: function grantDelegateSenderRole(address account) returns()
func (_ContractICS20Transfer *ContractICS20TransferTransactor) GrantDelegateSenderRole(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _ContractICS20Transfer.contract.Transact(opts, "grantDelegateSenderRole", account)
}

// GrantDelegateSenderRole is a paid mutator transaction binding the contract method 0x4551d2b0.
//
// Solidity: function grantDelegateSenderRole(address account) returns()
func (_ContractICS20Transfer *ContractICS20TransferSession) GrantDelegateSenderRole(account common.Address) (*types.Transaction, error) {
	return _ContractICS20Transfer.Contract.GrantDelegateSenderRole(&_ContractICS20Transfer.TransactOpts, account)
}

// GrantDelegateSenderRole is a paid mutator transaction binding the contract method 0x4551d2b0.
//
// Solidity: function grantDelegateSenderRole(address account) returns()
func (_ContractICS20Transfer *ContractICS20TransferTransactorSession) GrantDelegateSenderRole(account common.Address) (*types.Transaction, error) {
	return _ContractICS20Transfer.Contract.GrantDelegateSenderRole(&_ContractICS20Transfer.TransactOpts, account)
}

// GrantERC20CustomizerRole is a paid mutator transaction binding the contract method 0x66780af9.
//
// Solidity: function grantERC20CustomizerRole(address account) returns()
func (_ContractICS20Transfer *ContractICS20TransferTransactor) GrantERC20CustomizerRole(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _ContractICS20Transfer.contract.Transact(opts, "grantERC20CustomizerRole", account)
}

// GrantERC20CustomizerRole is a paid mutator transaction binding the contract method 0x66780af9.
//
// Solidity: function grantERC20CustomizerRole(address account) returns()
func (_ContractICS20Transfer *ContractICS20TransferSession) GrantERC20CustomizerRole(account common.Address) (*types.Transaction, error) {
	return _ContractICS20Transfer.Contract.GrantERC20CustomizerRole(&_ContractICS20Transfer.TransactOpts, account)
}

// GrantERC20CustomizerRole is a paid mutator transaction binding the contract method 0x66780af9.
//
// Solidity: function grantERC20CustomizerRole(address account) returns()
func (_ContractICS20Transfer *ContractICS20TransferTransactorSession) GrantERC20CustomizerRole(account common.Address) (*types.Transaction, error) {
	return _ContractICS20Transfer.Contract.GrantERC20CustomizerRole(&_ContractICS20Transfer.TransactOpts, account)
}

// GrantPauserRole is a paid mutator transaction binding the contract method 0x6c11c21c.
//
// Solidity: function grantPauserRole(address account) returns()
func (_ContractICS20Transfer *ContractICS20TransferTransactor) GrantPauserRole(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _ContractICS20Transfer.contract.Transact(opts, "grantPauserRole", account)
}

// GrantPauserRole is a paid mutator transaction binding the contract method 0x6c11c21c.
//
// Solidity: function grantPauserRole(address account) returns()
func (_ContractICS20Transfer *ContractICS20TransferSession) GrantPauserRole(account common.Address) (*types.Transaction, error) {
	return _ContractICS20Transfer.Contract.GrantPauserRole(&_ContractICS20Transfer.TransactOpts, account)
}

// GrantPauserRole is a paid mutator transaction binding the contract method 0x6c11c21c.
//
// Solidity: function grantPauserRole(address account) returns()
func (_ContractICS20Transfer *ContractICS20TransferTransactorSession) GrantPauserRole(account common.Address) (*types.Transaction, error) {
	return _ContractICS20Transfer.Contract.GrantPauserRole(&_ContractICS20Transfer.TransactOpts, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ContractICS20Transfer *ContractICS20TransferTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ContractICS20Transfer.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ContractICS20Transfer *ContractICS20TransferSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ContractICS20Transfer.Contract.GrantRole(&_ContractICS20Transfer.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ContractICS20Transfer *ContractICS20TransferTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ContractICS20Transfer.Contract.GrantRole(&_ContractICS20Transfer.TransactOpts, role, account)
}

// GrantTokenOperatorRole is a paid mutator transaction binding the contract method 0x6c0e87e1.
//
// Solidity: function grantTokenOperatorRole(address account) returns()
func (_ContractICS20Transfer *ContractICS20TransferTransactor) GrantTokenOperatorRole(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _ContractICS20Transfer.contract.Transact(opts, "grantTokenOperatorRole", account)
}

// GrantTokenOperatorRole is a paid mutator transaction binding the contract method 0x6c0e87e1.
//
// Solidity: function grantTokenOperatorRole(address account) returns()
func (_ContractICS20Transfer *ContractICS20TransferSession) GrantTokenOperatorRole(account common.Address) (*types.Transaction, error) {
	return _ContractICS20Transfer.Contract.GrantTokenOperatorRole(&_ContractICS20Transfer.TransactOpts, account)
}

// GrantTokenOperatorRole is a paid mutator transaction binding the contract method 0x6c0e87e1.
//
// Solidity: function grantTokenOperatorRole(address account) returns()
func (_ContractICS20Transfer *ContractICS20TransferTransactorSession) GrantTokenOperatorRole(account common.Address) (*types.Transaction, error) {
	return _ContractICS20Transfer.Contract.GrantTokenOperatorRole(&_ContractICS20Transfer.TransactOpts, account)
}

// GrantUnpauserRole is a paid mutator transaction binding the contract method 0x32968782.
//
// Solidity: function grantUnpauserRole(address account) returns()
func (_ContractICS20Transfer *ContractICS20TransferTransactor) GrantUnpauserRole(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _ContractICS20Transfer.contract.Transact(opts, "grantUnpauserRole", account)
}

// GrantUnpauserRole is a paid mutator transaction binding the contract method 0x32968782.
//
// Solidity: function grantUnpauserRole(address account) returns()
func (_ContractICS20Transfer *ContractICS20TransferSession) GrantUnpauserRole(account common.Address) (*types.Transaction, error) {
	return _ContractICS20Transfer.Contract.GrantUnpauserRole(&_ContractICS20Transfer.TransactOpts, account)
}

// GrantUnpauserRole is a paid mutator transaction binding the contract method 0x32968782.
//
// Solidity: function grantUnpauserRole(address account) returns()
func (_ContractICS20Transfer *ContractICS20TransferTransactorSession) GrantUnpauserRole(account common.Address) (*types.Transaction, error) {
	return _ContractICS20Transfer.Contract.GrantUnpauserRole(&_ContractICS20Transfer.TransactOpts, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address ics26Router, address escrowLogic, address ibcERC20Logic, address permit2) returns()
func (_ContractICS20Transfer *ContractICS20TransferTransactor) Initialize(opts *bind.TransactOpts, ics26Router common.Address, escrowLogic common.Address, ibcERC20Logic common.Address, permit2 common.Address) (*types.Transaction, error) {
	return _ContractICS20Transfer.contract.Transact(opts, "initialize", ics26Router, escrowLogic, ibcERC20Logic, permit2)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address ics26Router, address escrowLogic, address ibcERC20Logic, address permit2) returns()
func (_ContractICS20Transfer *ContractICS20TransferSession) Initialize(ics26Router common.Address, escrowLogic common.Address, ibcERC20Logic common.Address, permit2 common.Address) (*types.Transaction, error) {
	return _ContractICS20Transfer.Contract.Initialize(&_ContractICS20Transfer.TransactOpts, ics26Router, escrowLogic, ibcERC20Logic, permit2)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address ics26Router, address escrowLogic, address ibcERC20Logic, address permit2) returns()
func (_ContractICS20Transfer *ContractICS20TransferTransactorSession) Initialize(ics26Router common.Address, escrowLogic common.Address, ibcERC20Logic common.Address, permit2 common.Address) (*types.Transaction, error) {
	return _ContractICS20Transfer.Contract.Initialize(&_ContractICS20Transfer.TransactOpts, ics26Router, escrowLogic, ibcERC20Logic, permit2)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_ContractICS20Transfer *ContractICS20TransferTransactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _ContractICS20Transfer.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_ContractICS20Transfer *ContractICS20TransferSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _ContractICS20Transfer.Contract.Multicall(&_ContractICS20Transfer.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_ContractICS20Transfer *ContractICS20TransferTransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _ContractICS20Transfer.Contract.Multicall(&_ContractICS20Transfer.TransactOpts, data)
}

// OnAcknowledgementPacket is a paid mutator transaction binding the contract method 0x428e4e17.
//
// Solidity: function onAcknowledgementPacket((string,string,uint64,(string,string,string,string,bytes),bytes,address) msg_) returns()
func (_ContractICS20Transfer *ContractICS20TransferTransactor) OnAcknowledgementPacket(opts *bind.TransactOpts, msg_ IIBCAppCallbacksOnAcknowledgementPacketCallback) (*types.Transaction, error) {
	return _ContractICS20Transfer.contract.Transact(opts, "onAcknowledgementPacket", msg_)
}

// OnAcknowledgementPacket is a paid mutator transaction binding the contract method 0x428e4e17.
//
// Solidity: function onAcknowledgementPacket((string,string,uint64,(string,string,string,string,bytes),bytes,address) msg_) returns()
func (_ContractICS20Transfer *ContractICS20TransferSession) OnAcknowledgementPacket(msg_ IIBCAppCallbacksOnAcknowledgementPacketCallback) (*types.Transaction, error) {
	return _ContractICS20Transfer.Contract.OnAcknowledgementPacket(&_ContractICS20Transfer.TransactOpts, msg_)
}

// OnAcknowledgementPacket is a paid mutator transaction binding the contract method 0x428e4e17.
//
// Solidity: function onAcknowledgementPacket((string,string,uint64,(string,string,string,string,bytes),bytes,address) msg_) returns()
func (_ContractICS20Transfer *ContractICS20TransferTransactorSession) OnAcknowledgementPacket(msg_ IIBCAppCallbacksOnAcknowledgementPacketCallback) (*types.Transaction, error) {
	return _ContractICS20Transfer.Contract.OnAcknowledgementPacket(&_ContractICS20Transfer.TransactOpts, msg_)
}

// OnRecvPacket is a paid mutator transaction binding the contract method 0x078c4a79.
//
// Solidity: function onRecvPacket((string,string,uint64,(string,string,string,string,bytes),address) msg_) returns(bytes)
func (_ContractICS20Transfer *ContractICS20TransferTransactor) OnRecvPacket(opts *bind.TransactOpts, msg_ IIBCAppCallbacksOnRecvPacketCallback) (*types.Transaction, error) {
	return _ContractICS20Transfer.contract.Transact(opts, "onRecvPacket", msg_)
}

// OnRecvPacket is a paid mutator transaction binding the contract method 0x078c4a79.
//
// Solidity: function onRecvPacket((string,string,uint64,(string,string,string,string,bytes),address) msg_) returns(bytes)
func (_ContractICS20Transfer *ContractICS20TransferSession) OnRecvPacket(msg_ IIBCAppCallbacksOnRecvPacketCallback) (*types.Transaction, error) {
	return _ContractICS20Transfer.Contract.OnRecvPacket(&_ContractICS20Transfer.TransactOpts, msg_)
}

// OnRecvPacket is a paid mutator transaction binding the contract method 0x078c4a79.
//
// Solidity: function onRecvPacket((string,string,uint64,(string,string,string,string,bytes),address) msg_) returns(bytes)
func (_ContractICS20Transfer *ContractICS20TransferTransactorSession) OnRecvPacket(msg_ IIBCAppCallbacksOnRecvPacketCallback) (*types.Transaction, error) {
	return _ContractICS20Transfer.Contract.OnRecvPacket(&_ContractICS20Transfer.TransactOpts, msg_)
}

// OnTimeoutPacket is a paid mutator transaction binding the contract method 0x5e32b6b6.
//
// Solidity: function onTimeoutPacket((string,string,uint64,(string,string,string,string,bytes),address) msg_) returns()
func (_ContractICS20Transfer *ContractICS20TransferTransactor) OnTimeoutPacket(opts *bind.TransactOpts, msg_ IIBCAppCallbacksOnTimeoutPacketCallback) (*types.Transaction, error) {
	return _ContractICS20Transfer.contract.Transact(opts, "onTimeoutPacket", msg_)
}

// OnTimeoutPacket is a paid mutator transaction binding the contract method 0x5e32b6b6.
//
// Solidity: function onTimeoutPacket((string,string,uint64,(string,string,string,string,bytes),address) msg_) returns()
func (_ContractICS20Transfer *ContractICS20TransferSession) OnTimeoutPacket(msg_ IIBCAppCallbacksOnTimeoutPacketCallback) (*types.Transaction, error) {
	return _ContractICS20Transfer.Contract.OnTimeoutPacket(&_ContractICS20Transfer.TransactOpts, msg_)
}

// OnTimeoutPacket is a paid mutator transaction binding the contract method 0x5e32b6b6.
//
// Solidity: function onTimeoutPacket((string,string,uint64,(string,string,string,string,bytes),address) msg_) returns()
func (_ContractICS20Transfer *ContractICS20TransferTransactorSession) OnTimeoutPacket(msg_ IIBCAppCallbacksOnTimeoutPacketCallback) (*types.Transaction, error) {
	return _ContractICS20Transfer.Contract.OnTimeoutPacket(&_ContractICS20Transfer.TransactOpts, msg_)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ContractICS20Transfer *ContractICS20TransferTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractICS20Transfer.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ContractICS20Transfer *ContractICS20TransferSession) Pause() (*types.Transaction, error) {
	return _ContractICS20Transfer.Contract.Pause(&_ContractICS20Transfer.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ContractICS20Transfer *ContractICS20TransferTransactorSession) Pause() (*types.Transaction, error) {
	return _ContractICS20Transfer.Contract.Pause(&_ContractICS20Transfer.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ContractICS20Transfer *ContractICS20TransferTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ContractICS20Transfer.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ContractICS20Transfer *ContractICS20TransferSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ContractICS20Transfer.Contract.RenounceRole(&_ContractICS20Transfer.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ContractICS20Transfer *ContractICS20TransferTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ContractICS20Transfer.Contract.RenounceRole(&_ContractICS20Transfer.TransactOpts, role, callerConfirmation)
}

// RevokeDelegateSenderRole is a paid mutator transaction binding the contract method 0xc46023a3.
//
// Solidity: function revokeDelegateSenderRole(address account) returns()
func (_ContractICS20Transfer *ContractICS20TransferTransactor) RevokeDelegateSenderRole(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _ContractICS20Transfer.contract.Transact(opts, "revokeDelegateSenderRole", account)
}

// RevokeDelegateSenderRole is a paid mutator transaction binding the contract method 0xc46023a3.
//
// Solidity: function revokeDelegateSenderRole(address account) returns()
func (_ContractICS20Transfer *ContractICS20TransferSession) RevokeDelegateSenderRole(account common.Address) (*types.Transaction, error) {
	return _ContractICS20Transfer.Contract.RevokeDelegateSenderRole(&_ContractICS20Transfer.TransactOpts, account)
}

// RevokeDelegateSenderRole is a paid mutator transaction binding the contract method 0xc46023a3.
//
// Solidity: function revokeDelegateSenderRole(address account) returns()
func (_ContractICS20Transfer *ContractICS20TransferTransactorSession) RevokeDelegateSenderRole(account common.Address) (*types.Transaction, error) {
	return _ContractICS20Transfer.Contract.RevokeDelegateSenderRole(&_ContractICS20Transfer.TransactOpts, account)
}

// RevokeERC20CustomizerRole is a paid mutator transaction binding the contract method 0x23993870.
//
// Solidity: function revokeERC20CustomizerRole(address account) returns()
func (_ContractICS20Transfer *ContractICS20TransferTransactor) RevokeERC20CustomizerRole(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _ContractICS20Transfer.contract.Transact(opts, "revokeERC20CustomizerRole", account)
}

// RevokeERC20CustomizerRole is a paid mutator transaction binding the contract method 0x23993870.
//
// Solidity: function revokeERC20CustomizerRole(address account) returns()
func (_ContractICS20Transfer *ContractICS20TransferSession) RevokeERC20CustomizerRole(account common.Address) (*types.Transaction, error) {
	return _ContractICS20Transfer.Contract.RevokeERC20CustomizerRole(&_ContractICS20Transfer.TransactOpts, account)
}

// RevokeERC20CustomizerRole is a paid mutator transaction binding the contract method 0x23993870.
//
// Solidity: function revokeERC20CustomizerRole(address account) returns()
func (_ContractICS20Transfer *ContractICS20TransferTransactorSession) RevokeERC20CustomizerRole(account common.Address) (*types.Transaction, error) {
	return _ContractICS20Transfer.Contract.RevokeERC20CustomizerRole(&_ContractICS20Transfer.TransactOpts, account)
}

// RevokePauserRole is a paid mutator transaction binding the contract method 0xf865af08.
//
// Solidity: function revokePauserRole(address account) returns()
func (_ContractICS20Transfer *ContractICS20TransferTransactor) RevokePauserRole(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _ContractICS20Transfer.contract.Transact(opts, "revokePauserRole", account)
}

// RevokePauserRole is a paid mutator transaction binding the contract method 0xf865af08.
//
// Solidity: function revokePauserRole(address account) returns()
func (_ContractICS20Transfer *ContractICS20TransferSession) RevokePauserRole(account common.Address) (*types.Transaction, error) {
	return _ContractICS20Transfer.Contract.RevokePauserRole(&_ContractICS20Transfer.TransactOpts, account)
}

// RevokePauserRole is a paid mutator transaction binding the contract method 0xf865af08.
//
// Solidity: function revokePauserRole(address account) returns()
func (_ContractICS20Transfer *ContractICS20TransferTransactorSession) RevokePauserRole(account common.Address) (*types.Transaction, error) {
	return _ContractICS20Transfer.Contract.RevokePauserRole(&_ContractICS20Transfer.TransactOpts, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ContractICS20Transfer *ContractICS20TransferTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ContractICS20Transfer.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ContractICS20Transfer *ContractICS20TransferSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ContractICS20Transfer.Contract.RevokeRole(&_ContractICS20Transfer.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ContractICS20Transfer *ContractICS20TransferTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ContractICS20Transfer.Contract.RevokeRole(&_ContractICS20Transfer.TransactOpts, role, account)
}

// RevokeTokenOperatorRole is a paid mutator transaction binding the contract method 0x6ad7ba49.
//
// Solidity: function revokeTokenOperatorRole(address account) returns()
func (_ContractICS20Transfer *ContractICS20TransferTransactor) RevokeTokenOperatorRole(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _ContractICS20Transfer.contract.Transact(opts, "revokeTokenOperatorRole", account)
}

// RevokeTokenOperatorRole is a paid mutator transaction binding the contract method 0x6ad7ba49.
//
// Solidity: function revokeTokenOperatorRole(address account) returns()
func (_ContractICS20Transfer *ContractICS20TransferSession) RevokeTokenOperatorRole(account common.Address) (*types.Transaction, error) {
	return _ContractICS20Transfer.Contract.RevokeTokenOperatorRole(&_ContractICS20Transfer.TransactOpts, account)
}

// RevokeTokenOperatorRole is a paid mutator transaction binding the contract method 0x6ad7ba49.
//
// Solidity: function revokeTokenOperatorRole(address account) returns()
func (_ContractICS20Transfer *ContractICS20TransferTransactorSession) RevokeTokenOperatorRole(account common.Address) (*types.Transaction, error) {
	return _ContractICS20Transfer.Contract.RevokeTokenOperatorRole(&_ContractICS20Transfer.TransactOpts, account)
}

// RevokeUnpauserRole is a paid mutator transaction binding the contract method 0x2540e2da.
//
// Solidity: function revokeUnpauserRole(address account) returns()
func (_ContractICS20Transfer *ContractICS20TransferTransactor) RevokeUnpauserRole(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _ContractICS20Transfer.contract.Transact(opts, "revokeUnpauserRole", account)
}

// RevokeUnpauserRole is a paid mutator transaction binding the contract method 0x2540e2da.
//
// Solidity: function revokeUnpauserRole(address account) returns()
func (_ContractICS20Transfer *ContractICS20TransferSession) RevokeUnpauserRole(account common.Address) (*types.Transaction, error) {
	return _ContractICS20Transfer.Contract.RevokeUnpauserRole(&_ContractICS20Transfer.TransactOpts, account)
}

// RevokeUnpauserRole is a paid mutator transaction binding the contract method 0x2540e2da.
//
// Solidity: function revokeUnpauserRole(address account) returns()
func (_ContractICS20Transfer *ContractICS20TransferTransactorSession) RevokeUnpauserRole(account common.Address) (*types.Transaction, error) {
	return _ContractICS20Transfer.Contract.RevokeUnpauserRole(&_ContractICS20Transfer.TransactOpts, account)
}

// SendTransfer is a paid mutator transaction binding the contract method 0x2ac3dc38.
//
// Solidity: function sendTransfer((address,uint256,string,string,string,uint64,string) msg_) returns(uint64)
func (_ContractICS20Transfer *ContractICS20TransferTransactor) SendTransfer(opts *bind.TransactOpts, msg_ IICS20TransferMsgsSendTransferMsg) (*types.Transaction, error) {
	return _ContractICS20Transfer.contract.Transact(opts, "sendTransfer", msg_)
}

// SendTransfer is a paid mutator transaction binding the contract method 0x2ac3dc38.
//
// Solidity: function sendTransfer((address,uint256,string,string,string,uint64,string) msg_) returns(uint64)
func (_ContractICS20Transfer *ContractICS20TransferSession) SendTransfer(msg_ IICS20TransferMsgsSendTransferMsg) (*types.Transaction, error) {
	return _ContractICS20Transfer.Contract.SendTransfer(&_ContractICS20Transfer.TransactOpts, msg_)
}

// SendTransfer is a paid mutator transaction binding the contract method 0x2ac3dc38.
//
// Solidity: function sendTransfer((address,uint256,string,string,string,uint64,string) msg_) returns(uint64)
func (_ContractICS20Transfer *ContractICS20TransferTransactorSession) SendTransfer(msg_ IICS20TransferMsgsSendTransferMsg) (*types.Transaction, error) {
	return _ContractICS20Transfer.Contract.SendTransfer(&_ContractICS20Transfer.TransactOpts, msg_)
}

// SendTransferWithPermit2 is a paid mutator transaction binding the contract method 0x53816a7c.
//
// Solidity: function sendTransferWithPermit2((address,uint256,string,string,string,uint64,string) msg_, ((address,uint256),uint256,uint256) permit, bytes signature) returns(uint64)
func (_ContractICS20Transfer *ContractICS20TransferTransactor) SendTransferWithPermit2(opts *bind.TransactOpts, msg_ IICS20TransferMsgsSendTransferMsg, permit ISignatureTransferPermitTransferFrom, signature []byte) (*types.Transaction, error) {
	return _ContractICS20Transfer.contract.Transact(opts, "sendTransferWithPermit2", msg_, permit, signature)
}

// SendTransferWithPermit2 is a paid mutator transaction binding the contract method 0x53816a7c.
//
// Solidity: function sendTransferWithPermit2((address,uint256,string,string,string,uint64,string) msg_, ((address,uint256),uint256,uint256) permit, bytes signature) returns(uint64)
func (_ContractICS20Transfer *ContractICS20TransferSession) SendTransferWithPermit2(msg_ IICS20TransferMsgsSendTransferMsg, permit ISignatureTransferPermitTransferFrom, signature []byte) (*types.Transaction, error) {
	return _ContractICS20Transfer.Contract.SendTransferWithPermit2(&_ContractICS20Transfer.TransactOpts, msg_, permit, signature)
}

// SendTransferWithPermit2 is a paid mutator transaction binding the contract method 0x53816a7c.
//
// Solidity: function sendTransferWithPermit2((address,uint256,string,string,string,uint64,string) msg_, ((address,uint256),uint256,uint256) permit, bytes signature) returns(uint64)
func (_ContractICS20Transfer *ContractICS20TransferTransactorSession) SendTransferWithPermit2(msg_ IICS20TransferMsgsSendTransferMsg, permit ISignatureTransferPermitTransferFrom, signature []byte) (*types.Transaction, error) {
	return _ContractICS20Transfer.Contract.SendTransferWithPermit2(&_ContractICS20Transfer.TransactOpts, msg_, permit, signature)
}

// SendTransferWithSender is a paid mutator transaction binding the contract method 0xb29c715d.
//
// Solidity: function sendTransferWithSender((address,uint256,string,string,string,uint64,string) msg_, address sender) returns(uint64)
func (_ContractICS20Transfer *ContractICS20TransferTransactor) SendTransferWithSender(opts *bind.TransactOpts, msg_ IICS20TransferMsgsSendTransferMsg, sender common.Address) (*types.Transaction, error) {
	return _ContractICS20Transfer.contract.Transact(opts, "sendTransferWithSender", msg_, sender)
}

// SendTransferWithSender is a paid mutator transaction binding the contract method 0xb29c715d.
//
// Solidity: function sendTransferWithSender((address,uint256,string,string,string,uint64,string) msg_, address sender) returns(uint64)
func (_ContractICS20Transfer *ContractICS20TransferSession) SendTransferWithSender(msg_ IICS20TransferMsgsSendTransferMsg, sender common.Address) (*types.Transaction, error) {
	return _ContractICS20Transfer.Contract.SendTransferWithSender(&_ContractICS20Transfer.TransactOpts, msg_, sender)
}

// SendTransferWithSender is a paid mutator transaction binding the contract method 0xb29c715d.
//
// Solidity: function sendTransferWithSender((address,uint256,string,string,string,uint64,string) msg_, address sender) returns(uint64)
func (_ContractICS20Transfer *ContractICS20TransferTransactorSession) SendTransferWithSender(msg_ IICS20TransferMsgsSendTransferMsg, sender common.Address) (*types.Transaction, error) {
	return _ContractICS20Transfer.Contract.SendTransferWithSender(&_ContractICS20Transfer.TransactOpts, msg_, sender)
}

// SetCustomERC20 is a paid mutator transaction binding the contract method 0xa1d28f57.
//
// Solidity: function setCustomERC20(string denom, address token) returns()
func (_ContractICS20Transfer *ContractICS20TransferTransactor) SetCustomERC20(opts *bind.TransactOpts, denom string, token common.Address) (*types.Transaction, error) {
	return _ContractICS20Transfer.contract.Transact(opts, "setCustomERC20", denom, token)
}

// SetCustomERC20 is a paid mutator transaction binding the contract method 0xa1d28f57.
//
// Solidity: function setCustomERC20(string denom, address token) returns()
func (_ContractICS20Transfer *ContractICS20TransferSession) SetCustomERC20(denom string, token common.Address) (*types.Transaction, error) {
	return _ContractICS20Transfer.Contract.SetCustomERC20(&_ContractICS20Transfer.TransactOpts, denom, token)
}

// SetCustomERC20 is a paid mutator transaction binding the contract method 0xa1d28f57.
//
// Solidity: function setCustomERC20(string denom, address token) returns()
func (_ContractICS20Transfer *ContractICS20TransferTransactorSession) SetCustomERC20(denom string, token common.Address) (*types.Transaction, error) {
	return _ContractICS20Transfer.Contract.SetCustomERC20(&_ContractICS20Transfer.TransactOpts, denom, token)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ContractICS20Transfer *ContractICS20TransferTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractICS20Transfer.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ContractICS20Transfer *ContractICS20TransferSession) Unpause() (*types.Transaction, error) {
	return _ContractICS20Transfer.Contract.Unpause(&_ContractICS20Transfer.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ContractICS20Transfer *ContractICS20TransferTransactorSession) Unpause() (*types.Transaction, error) {
	return _ContractICS20Transfer.Contract.Unpause(&_ContractICS20Transfer.TransactOpts)
}

// UpgradeEscrowTo is a paid mutator transaction binding the contract method 0xaaa2c343.
//
// Solidity: function upgradeEscrowTo(address newEscrowLogic) returns()
func (_ContractICS20Transfer *ContractICS20TransferTransactor) UpgradeEscrowTo(opts *bind.TransactOpts, newEscrowLogic common.Address) (*types.Transaction, error) {
	return _ContractICS20Transfer.contract.Transact(opts, "upgradeEscrowTo", newEscrowLogic)
}

// UpgradeEscrowTo is a paid mutator transaction binding the contract method 0xaaa2c343.
//
// Solidity: function upgradeEscrowTo(address newEscrowLogic) returns()
func (_ContractICS20Transfer *ContractICS20TransferSession) UpgradeEscrowTo(newEscrowLogic common.Address) (*types.Transaction, error) {
	return _ContractICS20Transfer.Contract.UpgradeEscrowTo(&_ContractICS20Transfer.TransactOpts, newEscrowLogic)
}

// UpgradeEscrowTo is a paid mutator transaction binding the contract method 0xaaa2c343.
//
// Solidity: function upgradeEscrowTo(address newEscrowLogic) returns()
func (_ContractICS20Transfer *ContractICS20TransferTransactorSession) UpgradeEscrowTo(newEscrowLogic common.Address) (*types.Transaction, error) {
	return _ContractICS20Transfer.Contract.UpgradeEscrowTo(&_ContractICS20Transfer.TransactOpts, newEscrowLogic)
}

// UpgradeIBCERC20To is a paid mutator transaction binding the contract method 0x06ab20bc.
//
// Solidity: function upgradeIBCERC20To(address newIBCERC20Logic) returns()
func (_ContractICS20Transfer *ContractICS20TransferTransactor) UpgradeIBCERC20To(opts *bind.TransactOpts, newIBCERC20Logic common.Address) (*types.Transaction, error) {
	return _ContractICS20Transfer.contract.Transact(opts, "upgradeIBCERC20To", newIBCERC20Logic)
}

// UpgradeIBCERC20To is a paid mutator transaction binding the contract method 0x06ab20bc.
//
// Solidity: function upgradeIBCERC20To(address newIBCERC20Logic) returns()
func (_ContractICS20Transfer *ContractICS20TransferSession) UpgradeIBCERC20To(newIBCERC20Logic common.Address) (*types.Transaction, error) {
	return _ContractICS20Transfer.Contract.UpgradeIBCERC20To(&_ContractICS20Transfer.TransactOpts, newIBCERC20Logic)
}

// UpgradeIBCERC20To is a paid mutator transaction binding the contract method 0x06ab20bc.
//
// Solidity: function upgradeIBCERC20To(address newIBCERC20Logic) returns()
func (_ContractICS20Transfer *ContractICS20TransferTransactorSession) UpgradeIBCERC20To(newIBCERC20Logic common.Address) (*types.Transaction, error) {
	return _ContractICS20Transfer.Contract.UpgradeIBCERC20To(&_ContractICS20Transfer.TransactOpts, newIBCERC20Logic)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ContractICS20Transfer *ContractICS20TransferTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ContractICS20Transfer.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ContractICS20Transfer *ContractICS20TransferSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ContractICS20Transfer.Contract.UpgradeToAndCall(&_ContractICS20Transfer.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ContractICS20Transfer *ContractICS20TransferTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ContractICS20Transfer.Contract.UpgradeToAndCall(&_ContractICS20Transfer.TransactOpts, newImplementation, data)
}

// ContractICS20TransferIBCERC20ContractCreatedIterator is returned from FilterIBCERC20ContractCreated and is used to iterate over the raw logs and unpacked data for IBCERC20ContractCreated events raised by the ContractICS20Transfer contract.
type ContractICS20TransferIBCERC20ContractCreatedIterator struct {
	Event *ContractICS20TransferIBCERC20ContractCreated // Event containing the contract specifics and raw log

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
func (it *ContractICS20TransferIBCERC20ContractCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractICS20TransferIBCERC20ContractCreated)
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
		it.Event = new(ContractICS20TransferIBCERC20ContractCreated)
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
func (it *ContractICS20TransferIBCERC20ContractCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractICS20TransferIBCERC20ContractCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractICS20TransferIBCERC20ContractCreated represents a IBCERC20ContractCreated event raised by the ContractICS20Transfer contract.
type ContractICS20TransferIBCERC20ContractCreated struct {
	ContractAddress common.Address
	FullDenomPath   string
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterIBCERC20ContractCreated is a free log retrieval operation binding the contract event 0x6031fab685dd6d86e4dbac9a69eae347145f332c95b3a0d728d3730fc5233d62.
//
// Solidity: event IBCERC20ContractCreated(address indexed contractAddress, string fullDenomPath)
func (_ContractICS20Transfer *ContractICS20TransferFilterer) FilterIBCERC20ContractCreated(opts *bind.FilterOpts, contractAddress []common.Address) (*ContractICS20TransferIBCERC20ContractCreatedIterator, error) {

	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _ContractICS20Transfer.contract.FilterLogs(opts, "IBCERC20ContractCreated", contractAddressRule)
	if err != nil {
		return nil, err
	}
	return &ContractICS20TransferIBCERC20ContractCreatedIterator{contract: _ContractICS20Transfer.contract, event: "IBCERC20ContractCreated", logs: logs, sub: sub}, nil
}

// WatchIBCERC20ContractCreated is a free log subscription operation binding the contract event 0x6031fab685dd6d86e4dbac9a69eae347145f332c95b3a0d728d3730fc5233d62.
//
// Solidity: event IBCERC20ContractCreated(address indexed contractAddress, string fullDenomPath)
func (_ContractICS20Transfer *ContractICS20TransferFilterer) WatchIBCERC20ContractCreated(opts *bind.WatchOpts, sink chan<- *ContractICS20TransferIBCERC20ContractCreated, contractAddress []common.Address) (event.Subscription, error) {

	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _ContractICS20Transfer.contract.WatchLogs(opts, "IBCERC20ContractCreated", contractAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractICS20TransferIBCERC20ContractCreated)
				if err := _ContractICS20Transfer.contract.UnpackLog(event, "IBCERC20ContractCreated", log); err != nil {
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

// ParseIBCERC20ContractCreated is a log parse operation binding the contract event 0x6031fab685dd6d86e4dbac9a69eae347145f332c95b3a0d728d3730fc5233d62.
//
// Solidity: event IBCERC20ContractCreated(address indexed contractAddress, string fullDenomPath)
func (_ContractICS20Transfer *ContractICS20TransferFilterer) ParseIBCERC20ContractCreated(log types.Log) (*ContractICS20TransferIBCERC20ContractCreated, error) {
	event := new(ContractICS20TransferIBCERC20ContractCreated)
	if err := _ContractICS20Transfer.contract.UnpackLog(event, "IBCERC20ContractCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractICS20TransferInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContractICS20Transfer contract.
type ContractICS20TransferInitializedIterator struct {
	Event *ContractICS20TransferInitialized // Event containing the contract specifics and raw log

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
func (it *ContractICS20TransferInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractICS20TransferInitialized)
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
		it.Event = new(ContractICS20TransferInitialized)
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
func (it *ContractICS20TransferInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractICS20TransferInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractICS20TransferInitialized represents a Initialized event raised by the ContractICS20Transfer contract.
type ContractICS20TransferInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ContractICS20Transfer *ContractICS20TransferFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractICS20TransferInitializedIterator, error) {

	logs, sub, err := _ContractICS20Transfer.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractICS20TransferInitializedIterator{contract: _ContractICS20Transfer.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ContractICS20Transfer *ContractICS20TransferFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractICS20TransferInitialized) (event.Subscription, error) {

	logs, sub, err := _ContractICS20Transfer.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractICS20TransferInitialized)
				if err := _ContractICS20Transfer.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ContractICS20Transfer *ContractICS20TransferFilterer) ParseInitialized(log types.Log) (*ContractICS20TransferInitialized, error) {
	event := new(ContractICS20TransferInitialized)
	if err := _ContractICS20Transfer.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractICS20TransferPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ContractICS20Transfer contract.
type ContractICS20TransferPausedIterator struct {
	Event *ContractICS20TransferPaused // Event containing the contract specifics and raw log

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
func (it *ContractICS20TransferPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractICS20TransferPaused)
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
		it.Event = new(ContractICS20TransferPaused)
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
func (it *ContractICS20TransferPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractICS20TransferPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractICS20TransferPaused represents a Paused event raised by the ContractICS20Transfer contract.
type ContractICS20TransferPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ContractICS20Transfer *ContractICS20TransferFilterer) FilterPaused(opts *bind.FilterOpts) (*ContractICS20TransferPausedIterator, error) {

	logs, sub, err := _ContractICS20Transfer.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &ContractICS20TransferPausedIterator{contract: _ContractICS20Transfer.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ContractICS20Transfer *ContractICS20TransferFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ContractICS20TransferPaused) (event.Subscription, error) {

	logs, sub, err := _ContractICS20Transfer.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractICS20TransferPaused)
				if err := _ContractICS20Transfer.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ContractICS20Transfer *ContractICS20TransferFilterer) ParsePaused(log types.Log) (*ContractICS20TransferPaused, error) {
	event := new(ContractICS20TransferPaused)
	if err := _ContractICS20Transfer.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractICS20TransferRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the ContractICS20Transfer contract.
type ContractICS20TransferRoleAdminChangedIterator struct {
	Event *ContractICS20TransferRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *ContractICS20TransferRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractICS20TransferRoleAdminChanged)
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
		it.Event = new(ContractICS20TransferRoleAdminChanged)
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
func (it *ContractICS20TransferRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractICS20TransferRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractICS20TransferRoleAdminChanged represents a RoleAdminChanged event raised by the ContractICS20Transfer contract.
type ContractICS20TransferRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ContractICS20Transfer *ContractICS20TransferFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*ContractICS20TransferRoleAdminChangedIterator, error) {

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

	logs, sub, err := _ContractICS20Transfer.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &ContractICS20TransferRoleAdminChangedIterator{contract: _ContractICS20Transfer.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ContractICS20Transfer *ContractICS20TransferFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *ContractICS20TransferRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _ContractICS20Transfer.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractICS20TransferRoleAdminChanged)
				if err := _ContractICS20Transfer.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_ContractICS20Transfer *ContractICS20TransferFilterer) ParseRoleAdminChanged(log types.Log) (*ContractICS20TransferRoleAdminChanged, error) {
	event := new(ContractICS20TransferRoleAdminChanged)
	if err := _ContractICS20Transfer.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractICS20TransferRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the ContractICS20Transfer contract.
type ContractICS20TransferRoleGrantedIterator struct {
	Event *ContractICS20TransferRoleGranted // Event containing the contract specifics and raw log

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
func (it *ContractICS20TransferRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractICS20TransferRoleGranted)
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
		it.Event = new(ContractICS20TransferRoleGranted)
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
func (it *ContractICS20TransferRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractICS20TransferRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractICS20TransferRoleGranted represents a RoleGranted event raised by the ContractICS20Transfer contract.
type ContractICS20TransferRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ContractICS20Transfer *ContractICS20TransferFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ContractICS20TransferRoleGrantedIterator, error) {

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

	logs, sub, err := _ContractICS20Transfer.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ContractICS20TransferRoleGrantedIterator{contract: _ContractICS20Transfer.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ContractICS20Transfer *ContractICS20TransferFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *ContractICS20TransferRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ContractICS20Transfer.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractICS20TransferRoleGranted)
				if err := _ContractICS20Transfer.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_ContractICS20Transfer *ContractICS20TransferFilterer) ParseRoleGranted(log types.Log) (*ContractICS20TransferRoleGranted, error) {
	event := new(ContractICS20TransferRoleGranted)
	if err := _ContractICS20Transfer.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractICS20TransferRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the ContractICS20Transfer contract.
type ContractICS20TransferRoleRevokedIterator struct {
	Event *ContractICS20TransferRoleRevoked // Event containing the contract specifics and raw log

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
func (it *ContractICS20TransferRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractICS20TransferRoleRevoked)
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
		it.Event = new(ContractICS20TransferRoleRevoked)
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
func (it *ContractICS20TransferRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractICS20TransferRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractICS20TransferRoleRevoked represents a RoleRevoked event raised by the ContractICS20Transfer contract.
type ContractICS20TransferRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ContractICS20Transfer *ContractICS20TransferFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ContractICS20TransferRoleRevokedIterator, error) {

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

	logs, sub, err := _ContractICS20Transfer.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ContractICS20TransferRoleRevokedIterator{contract: _ContractICS20Transfer.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ContractICS20Transfer *ContractICS20TransferFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *ContractICS20TransferRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ContractICS20Transfer.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractICS20TransferRoleRevoked)
				if err := _ContractICS20Transfer.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_ContractICS20Transfer *ContractICS20TransferFilterer) ParseRoleRevoked(log types.Log) (*ContractICS20TransferRoleRevoked, error) {
	event := new(ContractICS20TransferRoleRevoked)
	if err := _ContractICS20Transfer.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractICS20TransferUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ContractICS20Transfer contract.
type ContractICS20TransferUnpausedIterator struct {
	Event *ContractICS20TransferUnpaused // Event containing the contract specifics and raw log

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
func (it *ContractICS20TransferUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractICS20TransferUnpaused)
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
		it.Event = new(ContractICS20TransferUnpaused)
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
func (it *ContractICS20TransferUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractICS20TransferUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractICS20TransferUnpaused represents a Unpaused event raised by the ContractICS20Transfer contract.
type ContractICS20TransferUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ContractICS20Transfer *ContractICS20TransferFilterer) FilterUnpaused(opts *bind.FilterOpts) (*ContractICS20TransferUnpausedIterator, error) {

	logs, sub, err := _ContractICS20Transfer.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &ContractICS20TransferUnpausedIterator{contract: _ContractICS20Transfer.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ContractICS20Transfer *ContractICS20TransferFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ContractICS20TransferUnpaused) (event.Subscription, error) {

	logs, sub, err := _ContractICS20Transfer.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractICS20TransferUnpaused)
				if err := _ContractICS20Transfer.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ContractICS20Transfer *ContractICS20TransferFilterer) ParseUnpaused(log types.Log) (*ContractICS20TransferUnpaused, error) {
	event := new(ContractICS20TransferUnpaused)
	if err := _ContractICS20Transfer.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractICS20TransferUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the ContractICS20Transfer contract.
type ContractICS20TransferUpgradedIterator struct {
	Event *ContractICS20TransferUpgraded // Event containing the contract specifics and raw log

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
func (it *ContractICS20TransferUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractICS20TransferUpgraded)
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
		it.Event = new(ContractICS20TransferUpgraded)
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
func (it *ContractICS20TransferUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractICS20TransferUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractICS20TransferUpgraded represents a Upgraded event raised by the ContractICS20Transfer contract.
type ContractICS20TransferUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ContractICS20Transfer *ContractICS20TransferFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*ContractICS20TransferUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ContractICS20Transfer.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &ContractICS20TransferUpgradedIterator{contract: _ContractICS20Transfer.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ContractICS20Transfer *ContractICS20TransferFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *ContractICS20TransferUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ContractICS20Transfer.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractICS20TransferUpgraded)
				if err := _ContractICS20Transfer.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ContractICS20Transfer *ContractICS20TransferFilterer) ParseUpgraded(log types.Log) (*ContractICS20TransferUpgraded, error) {
	event := new(ContractICS20TransferUpgraded)
	if err := _ContractICS20Transfer.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
