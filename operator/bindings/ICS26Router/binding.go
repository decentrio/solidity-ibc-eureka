// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractICS26Router

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

// IICS02ClientMsgsCounterpartyInfo is an auto generated low-level Go binding around an user-defined struct.
type IICS02ClientMsgsCounterpartyInfo struct {
	ClientId     string
	MerklePrefix [][]byte
}

// IICS02ClientMsgsHeight is an auto generated low-level Go binding around an user-defined struct.
type IICS02ClientMsgsHeight struct {
	RevisionNumber uint64
	RevisionHeight uint64
}

// IICS26RouterMsgsMsgAckPacket is an auto generated low-level Go binding around an user-defined struct.
type IICS26RouterMsgsMsgAckPacket struct {
	Packet          IICS26RouterMsgsPacket
	Acknowledgement []byte
	ProofAcked      []byte
	ProofHeight     IICS02ClientMsgsHeight
}

// IICS26RouterMsgsMsgRecvPacket is an auto generated low-level Go binding around an user-defined struct.
type IICS26RouterMsgsMsgRecvPacket struct {
	Packet          IICS26RouterMsgsPacket
	ProofCommitment []byte
	ProofHeight     IICS02ClientMsgsHeight
}

// IICS26RouterMsgsMsgSendPacket is an auto generated low-level Go binding around an user-defined struct.
type IICS26RouterMsgsMsgSendPacket struct {
	SourceClient     string
	TimeoutTimestamp uint64
	Payload          IICS26RouterMsgsPayload
}

// IICS26RouterMsgsMsgTimeoutPacket is an auto generated low-level Go binding around an user-defined struct.
type IICS26RouterMsgsMsgTimeoutPacket struct {
	Packet       IICS26RouterMsgsPacket
	ProofTimeout []byte
	ProofHeight  IICS02ClientMsgsHeight
}

// IICS26RouterMsgsPacket is an auto generated low-level Go binding around an user-defined struct.
type IICS26RouterMsgsPacket struct {
	Sequence         uint64
	SourceClient     string
	DestClient       string
	TimeoutTimestamp uint64
	Payloads         []IICS26RouterMsgsPayload
}

// IICS26RouterMsgsPayload is an auto generated low-level Go binding around an user-defined struct.
type IICS26RouterMsgsPayload struct {
	SourcePort string
	DestPort   string
	Version    string
	Encoding   string
	Value      []byte
}

// ContractICS26RouterMetaData contains all meta data concerning the ContractICS26Router contract.
var ContractICS26RouterMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ackPacket\",\"inputs\":[{\"name\":\"msg_\",\"type\":\"tuple\",\"internalType\":\"structIICS26RouterMsgs.MsgAckPacket\",\"components\":[{\"name\":\"packet\",\"type\":\"tuple\",\"internalType\":\"structIICS26RouterMsgs.Packet\",\"components\":[{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sourceClient\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"destClient\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"timeoutTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"payloads\",\"type\":\"tuple[]\",\"internalType\":\"structIICS26RouterMsgs.Payload[]\",\"components\":[{\"name\":\"sourcePort\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"destPort\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"encoding\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]},{\"name\":\"acknowledgement\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proofAcked\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proofHeight\",\"type\":\"tuple\",\"internalType\":\"structIICS02ClientMsgs.Height\",\"components\":[{\"name\":\"revisionNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revisionHeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addClient\",\"inputs\":[{\"name\":\"clientId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"counterpartyInfo\",\"type\":\"tuple\",\"internalType\":\"structIICS02ClientMsgs.CounterpartyInfo\",\"components\":[{\"name\":\"clientId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"merklePrefix\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}]},{\"name\":\"client\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addClient\",\"inputs\":[{\"name\":\"counterpartyInfo\",\"type\":\"tuple\",\"internalType\":\"structIICS02ClientMsgs.CounterpartyInfo\",\"components\":[{\"name\":\"clientId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"merklePrefix\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}]},{\"name\":\"client\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addIBCApp\",\"inputs\":[{\"name\":\"app\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addIBCApp\",\"inputs\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"app\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"authority\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getClient\",\"inputs\":[{\"name\":\"clientId\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractILightClient\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCommitment\",\"inputs\":[{\"name\":\"hashedPath\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCounterparty\",\"inputs\":[{\"name\":\"clientId\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIICS02ClientMsgs.CounterpartyInfo\",\"components\":[{\"name\":\"clientId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"merklePrefix\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getIBCApp\",\"inputs\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIIBCApp\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNextClientSeq\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"authority\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initializeV2\",\"inputs\":[{\"name\":\"authority\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isConsumingScheduledOp\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"migrateClient\",\"inputs\":[{\"name\":\"clientId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"counterpartyInfo\",\"type\":\"tuple\",\"internalType\":\"structIICS02ClientMsgs.CounterpartyInfo\",\"components\":[{\"name\":\"clientId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"merklePrefix\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}]},{\"name\":\"client\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"multicall\",\"inputs\":[{\"name\":\"data\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"outputs\":[{\"name\":\"results\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"recvPacket\",\"inputs\":[{\"name\":\"msg_\",\"type\":\"tuple\",\"internalType\":\"structIICS26RouterMsgs.MsgRecvPacket\",\"components\":[{\"name\":\"packet\",\"type\":\"tuple\",\"internalType\":\"structIICS26RouterMsgs.Packet\",\"components\":[{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sourceClient\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"destClient\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"timeoutTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"payloads\",\"type\":\"tuple[]\",\"internalType\":\"structIICS26RouterMsgs.Payload[]\",\"components\":[{\"name\":\"sourcePort\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"destPort\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"encoding\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]},{\"name\":\"proofCommitment\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proofHeight\",\"type\":\"tuple\",\"internalType\":\"structIICS02ClientMsgs.Height\",\"components\":[{\"name\":\"revisionNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revisionHeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"sendPacket\",\"inputs\":[{\"name\":\"msg_\",\"type\":\"tuple\",\"internalType\":\"structIICS26RouterMsgs.MsgSendPacket\",\"components\":[{\"name\":\"sourceClient\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"timeoutTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"payload\",\"type\":\"tuple\",\"internalType\":\"structIICS26RouterMsgs.Payload\",\"components\":[{\"name\":\"sourcePort\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"destPort\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"encoding\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setAuthority\",\"inputs\":[{\"name\":\"newAuthority\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitMisbehaviour\",\"inputs\":[{\"name\":\"clientId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"misbehaviourMsg\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"timeoutPacket\",\"inputs\":[{\"name\":\"msg_\",\"type\":\"tuple\",\"internalType\":\"structIICS26RouterMsgs.MsgTimeoutPacket\",\"components\":[{\"name\":\"packet\",\"type\":\"tuple\",\"internalType\":\"structIICS26RouterMsgs.Packet\",\"components\":[{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sourceClient\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"destClient\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"timeoutTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"payloads\",\"type\":\"tuple[]\",\"internalType\":\"structIICS26RouterMsgs.Payload[]\",\"components\":[{\"name\":\"sourcePort\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"destPort\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"encoding\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]},{\"name\":\"proofTimeout\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proofHeight\",\"type\":\"tuple\",\"internalType\":\"structIICS02ClientMsgs.Height\",\"components\":[{\"name\":\"revisionNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revisionHeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateClient\",\"inputs\":[{\"name\":\"clientId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"updateMsg\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumILightClientMsgs.UpdateResult\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"AckPacket\",\"inputs\":[{\"name\":\"clientId\",\"type\":\"string\",\"indexed\":true,\"internalType\":\"string\"},{\"name\":\"sequence\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"packet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIICS26RouterMsgs.Packet\",\"components\":[{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sourceClient\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"destClient\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"timeoutTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"payloads\",\"type\":\"tuple[]\",\"internalType\":\"structIICS26RouterMsgs.Payload[]\",\"components\":[{\"name\":\"sourcePort\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"destPort\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"encoding\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]},{\"name\":\"acknowledgement\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AuthorityUpdated\",\"inputs\":[{\"name\":\"authority\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"IBCAppAdded\",\"inputs\":[{\"name\":\"portId\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"app\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"IBCAppRecvPacketCallbackError\",\"inputs\":[{\"name\":\"reason\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ICS02ClientAdded\",\"inputs\":[{\"name\":\"clientId\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"counterpartyInfo\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIICS02ClientMsgs.CounterpartyInfo\",\"components\":[{\"name\":\"clientId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"merklePrefix\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}]},{\"name\":\"client\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ICS02ClientMigrated\",\"inputs\":[{\"name\":\"clientId\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"counterpartyInfo\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIICS02ClientMsgs.CounterpartyInfo\",\"components\":[{\"name\":\"clientId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"merklePrefix\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}]},{\"name\":\"client\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ICS02ClientUpdated\",\"inputs\":[{\"name\":\"clientId\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"result\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumILightClientMsgs.UpdateResult\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ICS02MisbehaviourSubmitted\",\"inputs\":[{\"name\":\"clientId\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Noop\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SendPacket\",\"inputs\":[{\"name\":\"clientId\",\"type\":\"string\",\"indexed\":true,\"internalType\":\"string\"},{\"name\":\"sequence\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"packet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIICS26RouterMsgs.Packet\",\"components\":[{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sourceClient\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"destClient\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"timeoutTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"payloads\",\"type\":\"tuple[]\",\"internalType\":\"structIICS26RouterMsgs.Payload[]\",\"components\":[{\"name\":\"sourcePort\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"destPort\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"encoding\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TimeoutPacket\",\"inputs\":[{\"name\":\"clientId\",\"type\":\"string\",\"indexed\":true,\"internalType\":\"string\"},{\"name\":\"sequence\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"packet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIICS26RouterMsgs.Packet\",\"components\":[{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sourceClient\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"destClient\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"timeoutTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"payloads\",\"type\":\"tuple[]\",\"internalType\":\"structIICS26RouterMsgs.Payload[]\",\"components\":[{\"name\":\"sourcePort\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"destPort\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"encoding\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WriteAcknowledgement\",\"inputs\":[{\"name\":\"clientId\",\"type\":\"string\",\"indexed\":true,\"internalType\":\"string\"},{\"name\":\"sequence\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"packet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIICS26RouterMsgs.Packet\",\"components\":[{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sourceClient\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"destClient\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"timeoutTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"payloads\",\"type\":\"tuple[]\",\"internalType\":\"structIICS26RouterMsgs.Payload[]\",\"components\":[{\"name\":\"sourcePort\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"destPort\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"encoding\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]},{\"name\":\"acknowledgements\",\"type\":\"bytes[]\",\"indexed\":false,\"internalType\":\"bytes[]\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessManagedInvalidAuthority\",\"inputs\":[{\"name\":\"authority\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AccessManagedRequiredDelay\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delay\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"AccessManagedUnauthorized\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"DefaultAdminRoleCannotBeGranted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"IBCAppNotFound\",\"inputs\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"IBCAsyncAcknowledgementNotSupported\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"IBCClientAlreadyExists\",\"inputs\":[{\"name\":\"clientId\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"IBCClientNotFound\",\"inputs\":[{\"name\":\"clientId\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"IBCCounterpartyClientNotFound\",\"inputs\":[{\"name\":\"counterpartyClientId\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"IBCErrorUniversalAcknowledgement\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"IBCFailedCallback\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"IBCInvalidClientId\",\"inputs\":[{\"name\":\"clientId\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"IBCInvalidCounterparty\",\"inputs\":[{\"name\":\"expected\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"actual\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"IBCInvalidPortIdentifier\",\"inputs\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"IBCInvalidTimeoutDuration\",\"inputs\":[{\"name\":\"maxTimeoutDuration\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"actualTimeoutDuration\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"IBCInvalidTimeoutTimestamp\",\"inputs\":[{\"name\":\"timeoutTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"comparedTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"IBCMultiPayloadPacketNotSupported\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"IBCPacketAcknowledgementAlreadyExists\",\"inputs\":[{\"name\":\"path\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"IBCPacketCommitmentAlreadyExists\",\"inputs\":[{\"name\":\"path\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"IBCPacketCommitmentMismatch\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"actual\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"IBCPacketReceiptMismatch\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"actual\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"IBCPortAlreadyExists\",\"inputs\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"IBCUnauthorizedSender\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidMerklePrefix\",\"inputs\":[{\"name\":\"prefix\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}]},{\"type\":\"error\",\"name\":\"NoAcknowledgements\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StringsInsufficientHexLength\",\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"Unreachable\",\"inputs\":[]}]",
	Bin: "0x60a080604052346100c257306080525f5160206157d65f395f51905f525460ff8160401c166100b3576002600160401b03196001600160401b03821601610060575b60405161570f90816100c7823960805181818161125401526113170152f35b6001600160401b0319166001600160401b039081175f5160206157d65f395f51905f525581527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a15f80610041565b63f92ee8a960e01b5f5260045ffd5b5f80fdfe60806040526004361015610011575f80fd5b5f5f3560e01c80631bca011a14611f355780631ec43e2314611eda5780632447af2914611ea457806327f146f314611e6857806329b6eca914611b185780634b720d5b146119ca5780634d6e7ce3146115af5780634f1ef286146112cc57806352d1902d1461122d5780635ebd10ca146112055780635f516889146111345780636fbf80791461100a5780637795820c14610fc15780637a9e5e4b14610eea5780637eb7893214610e905780638fb3603714610dfd5780639e2e5c8314610cec578063ac9650d814610b8f578063ad3cb1cc14610b2e578063b0777bfa14610abb578063b98c330a14610a6c578063bf7e214f14610a19578063c4d66de8146107b1578063cce0b265146103dc5763e3cb36a01461012d575f80fd5b346103d95760406003193601126103d9576004359067ffffffffffffffff82116103d957604060031983360301126103d957610167611fd5565b916101706146be565b927f515a8336edcaab4ae6524d41223c1782132890f89189ba6632107a7b5a44960254915f1983146103ac57600183017f515a8336edcaab4ae6524d41223c1782132890f89189ba6632107a7b5a449602558383807a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000811015610381575b50806d04ee2d6d415b85acef8100000000600a921015610366575b662386f26fc10000811015610352575b6305f5e100811015610341575b612710811015610332575b6064811015610324575b101561031c575b6001810193600a5f19602161026a61025489612138565b986102626040519a8b612115565b808a52612138565b94601f1960208a0196013687378801015b01917f30313233343536373839616263646566000000000000000000000000000000008282061a8353049081156102b757600a905f199061027b565b505061031895602095866102fb936103049760405199858b9651918291018588015e85019083820190858252519283915e010190815203601f198101865285612115565b600401836148c7565b60405191829160208352602083019061206f565b0390f35b60010161023d565b606460029104920191610236565b6127106004910492019161022c565b6305f5e10060089104920191610221565b662386f26fc1000060109104920191610214565b6d04ee2d6d415b85acef810000000060209104920191610204565b604092507a184f03e93ff9f4daa797ed6e38ed64bf6a1f01000000000000000090049050600a6101e9565b6024847f4e487b710000000000000000000000000000000000000000000000000000000081526011600452fd5b80fd5b50346103d9576103eb36611ff8565b9291906103f836336140e2565b610402828461388a565b5061040d8284612e5a565b6104178280612346565b9067ffffffffffffffff821161071e5761043b82610435855461391e565b85613f97565b8790601f831160011461074b57918061046c92600195948b92610603575b50505f198260011b9260031b1c19161790565b81555b0161047d60208301836122f2565b9168010000000000000000831161071e5780548382558084106106a4575b508752602087208791805b84841061059157505050505061058573ffffffffffffffffffffffffffffffffffffffff7f23c2e29d6ae84e79fa116b8afd6e28ddc1de7f473d3edb407fbd08093c3ed6bf951691604051848682376020818681017f515a8336edcaab4ae6524d41223c1782132890f89189ba6632107a7b5a44960081520301902073ffffffffffffffffffffffffffffffffffffffff84167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055610577604051958695606087526060870191612397565b908482036020860152613fdc565b9060408301520390a180f35b61059b8183612346565b9067ffffffffffffffff8211610677576105bf826105b9875461391e565b87613f97565b8b908c601f841160011461060e5783600195929460209487966105f494926106035750505f198260011b9260031b1c19161790565b86555b019301930192916104a6565b013590505f80610459565b91601f19841687845260208420935b81811061065f5750936020936001969387969383889510610646575b505050811b0186556105f7565b5f1960f88560031b161c199101351690555f8080610639565b9193602060018192878701358155019501920161061d565b60248c7f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b8189528360208a2091820191015b8181106106bf575061049b565b808a6106cd6001935461391e565b806106db575b5050016106b2565b601f811184146106f25750508a81555b8a5f6106d3565b83601f6020848661070d965220920160051c82019101613f81565b808b528a60208120818355556106eb565b6024887f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b8389526020892091601f1984168a5b8181106107995750916001959492918387959310610780575b505050811b01815561046f565b5f1960f88560031b161c199101351690555f8080610773565b9193602060018192878701358155019501920161075a565b50346103d95760206003193601126103d9576107cb611fb2565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005467ffffffffffffffff811690816109f15760401c60ff169081156109e5575b506109bd576109299061088360027fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000007ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005416177ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0055565b680100000000000000007fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005416177ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00556108fc61537c565b61090461537c565b61090c61537c565b61091461537c565b61091c61537c565b61092461537c565b6151e0565b7fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054167ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00557fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2602060405160028152a180f35b6004827ff92ee8a9000000000000000000000000000000000000000000000000000000008152fd5b6002915010155f61080c565b6004847ff92ee8a9000000000000000000000000000000000000000000000000000000008152fd5b50346103d957806003193601126103d957602073ffffffffffffffffffffffffffffffffffffffff7ff3177357ab46d8af007ab3fdb9af81da189e1068fefdc0073dca88a2cab40a005416604051908152f35b50346103d957610a95610a7e366121a8565b610a8661406e565b610a9036336140e2565b613ba3565b807f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f005d80f35b50346103d95760206003193601126103d9576004359067ffffffffffffffff82116103d957610318610af9610af33660048601611f84565b9061396f565b604051918291602083526020610b1a8251604083870152606086019061206f565b910151601f19848303016040850152612267565b50346103d957806003193601126103d95750610318604051610b51604082612115565b600581527f352e302e30000000000000000000000000000000000000000000000000000000602082015260405191829160208352602083019061206f565b50346103d95760206003193601126103d9576004359067ffffffffffffffff82116103d957366023830112156103d95781600401359067ffffffffffffffff82116103d957602483013660248460051b86010111610ce8576040516020610bf68183612115565b83825280820192601f198201368537610c0e8661242e565b96610c1c6040519889612115565b868852601f19610c2b8861242e565b0183875b828110610cd857505050855b87811015610cc557600190610ca988808989610c958a610c6360248960051b8c01018c612346565b9190946040519483869484860198893784019083820190898252519283915e010185815203601f198101835282612115565b5190305af4610ca2613249565b9030615496565b610cb3828c61248d565b52610cbe818b61248d565b5001610c3b565b604051848152806103188187018c612267565b606082828d010152018490610c2f565b5080fd5b5034610df957610cfb366121db565b73ffffffffffffffffffffffffffffffffffffffff610d1d848695979661388a565b1691823b15610df957610d6a925f92836040518096819582947fddba6537000000000000000000000000000000000000000000000000000000008452602060048501526024840191612397565b03925af18015610dee57610dba575b507fa263f0a976b2937a51fd2e416491cf0ca724d5499fa870715929dfde4ee4a4309192610db4604051928392602084526020840191612397565b0390a180f35b7fa263f0a976b2937a51fd2e416491cf0ca724d5499fa870715929dfde4ee4a43092505f610de791612115565b5f91610d79565b6040513d5f823e3d90fd5b5f80fd5b34610df9575f600319360112610df9577ff3177357ab46d8af007ab3fdb9af81da189e1068fefdc0073dca88a2cab40a005460a01c60ff1615610e885760207f8fb36037000000000000000000000000000000000000000000000000000000005b7fffffffff0000000000000000000000000000000000000000000000000000000060405191168152f35b60205f610e5e565b34610df9576020600319360112610df95760043567ffffffffffffffff8111610df957610ecc610ec66020923690600401611f84565b9061388a565b73ffffffffffffffffffffffffffffffffffffffff60405191168152f35b34610df9576020600319360112610df957610f03611fb2565b73ffffffffffffffffffffffffffffffffffffffff7ff3177357ab46d8af007ab3fdb9af81da189e1068fefdc0073dca88a2cab40a0054163303610f9557803b15610f5357610f51906151e0565b005b73ffffffffffffffffffffffffffffffffffffffff907fc2f31e5e000000000000000000000000000000000000000000000000000000005f521660045260245ffd5b7f068ca9d8000000000000000000000000000000000000000000000000000000005f523360045260245ffd5b34610df9576020600319360112610df9576004355f527f1260944489272988d9df285149b5aa1b0f48f2136d6f416159f840a3e0747600602052602060405f2054604051908152f35b34610df957602061108961101d366121db565b61102b9592939536336140e2565b73ffffffffffffffffffffffffffffffffffffffff61104a858861388a565b16905f6040518097819582947f0bece3560000000000000000000000000000000000000000000000000000000084528860048501526024840191612397565b03925af1918215610dee575f926110f6575b506020927f87bbef2779889a19f0435ddca81fda94132c06ffddb0ea73def256307a293aef916110d8604051928392604084526040840191612397565b6110e48683018661222d565b0390a16110f4604051809261222d565bf35b9091506020813d60201161112c575b8161111260209383612115565b81010312610df957516003811015610df95790602061109b565b3d9150611105565b34610df9576040600319360112610df95760043567ffffffffffffffff8111610df9576111db61116b6111e0923690600401611f84565b9190611175611fd5565b9261117e61406e565b61118836336140e2565b6111958183811515613841565b6111b781836111b06111a8368484612154565b8051906153d3565b5015613841565b6111d481836111cf6111ca368484612154565b61470a565b613841565b3691612154565b614cbc565b5f7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f005d005b34610df9576111e0611216366121a8565b61121e61406e565b61122836336140e2565b613278565b34610df9575f600319360112610df95773ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001630036112a45760206040517f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc8152f35b7fe07c8dba000000000000000000000000000000000000000000000000000000005f5260045ffd5b6040600319360112610df9576112e0611fb2565b60243567ffffffffffffffff8111610df95761130090369060040161218a565b73ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001680301490811561156d575b506112a45761135136336140e2565b73ffffffffffffffffffffffffffffffffffffffff8216916040517f52d1902d000000000000000000000000000000000000000000000000000000008152602081600481875afa5f9181611539575b506113d157837f4c9c8ce3000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b807f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc85920361150e5750813b156114e357807fffffffffffffffffffffffff00000000000000000000000000000000000000007f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5416177f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc557fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b5f80a28151156114b2575f80836020610f5195519101845af46114ac613249565b91615496565b5050346114bb57005b7fb398979f000000000000000000000000000000000000000000000000000000005f5260045ffd5b7f4c9c8ce3000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b7faa1d49a4000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b9091506020813d602011611565575b8161155560209383612115565b81010312610df9575190856113a0565b3d9150611548565b905073ffffffffffffffffffffffffffffffffffffffff7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5416141583611342565b34610df9576020600319360112610df95760043567ffffffffffffffff8111610df957806004019060606003198236030112610df9576115ed61406e565b604481019161163073ffffffffffffffffffffffffffffffffffffffff61162661162061161a87866122bf565b80612346565b90612e92565b1633903314612f26565b602461163f610af38380612346565b5192019061166d61164f83612419565b429061165a85612419565b9067ffffffffffffffff42911611612f70565b6201518061168d4267ffffffffffffffff61168786612419565b16612fb2565b116116a44267ffffffffffffffff61168786612419565b90611998575060206116b68280612346565b919082604051938492833781017f1260944489272988d9df285149b5aa1b0f48f2136d6f416159f840a3e07476018152030190209267ffffffffffffffff8454169467ffffffffffffffff861461196b5767ffffffffffffffff600161175397011694857fffffffffffffffffffffffffffffffffffffffffffffffff000000000000000082541617905561174b8380612346565b969094612419565b604094855197611763878a612115565b60018952601f1987015f5b8181106119355750509267ffffffffffffffff8899936117a46117cc946117db978b519d8e61179c81612094565b523691612154565b9660208c01978852898c01521660608a01528260808a01526117c73691876122bf565b61255d565b6117d582612480565b52612480565b506117f3815167ffffffffffffffff87511690614e1f565b6020815191012090815f527f1260944489272988d9df285149b5aa1b0f48f2136d6f416159f840a3e074760060205261183e845f205415915167ffffffffffffffff88511690614e1f565b90156118f45750937fab3a4458a269be61dfa43faa33aa7b1f5d570716f83ad078bc2ba5dab039abae6118c86118ab869460209861187b86614e99565b905f527f1260944489272988d9df285149b5aa1b0f48f2136d6f416159f840a3e07476008a52875f205580612346565b90818751928392833781015f815203902092855191829182612fbf565b0390a35f7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f005d51908152f35b6119319084519182917f91ffd92400000000000000000000000000000000000000000000000000000000835260206004840152602483019061206f565b0390fd5b8089602080938e606084519461194a86612094565b8186528185870152850152606080850152606060808501520101520161176e565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b7f715fed60000000000000000000000000000000000000000000000000000000005f526201518060045260245260445ffd5b34610df9576020600319360112610df9576119e3611fb2565b6119eb61406e565b73ffffffffffffffffffffffffffffffffffffffff81169081611a0e602a612138565b90611a1c6040519283612115565b602a8252611a2a602a612138565b601f19602084019101368237825115611aeb5760309053815160011015611aeb576078602183015360295b60018111611a9d5750611a6c576111e09250614cbc565b827fe22e27eb000000000000000000000000000000000000000000000000000000005f52600452601460245260445ffd5b90600f81166010811015611aeb577f3031323334353637383961626364656600000000000000000000000000000000901a611ad883856146f9565b5360041c90801561196b575f1901611a55565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b34610df9576020600319360112610df957611b31611fb2565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005467ffffffffffffffff81169060018203611e345760401c60ff16908115611e5c575b50611e3457611d5c90611bec60027fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000007ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005416177ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0055565b680100000000000000007fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005416177ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00557fffffffffffffffffffffffff00000000000000000000000000000000000000007fba83ed17c16070da0debaa680185af188d82c999a75962a12a40699ca48a2b005473ffffffffffffffffffffffffffffffffffffffff811633148015611def575b611ccb903390612f26565b167fba83ed17c16070da0debaa680185af188d82c999a75962a12a40699ca48a2b00557fffffffffffffffffffffffff00000000000000000000000000000000000000007fba83ed17c16070da0debaa680185af188d82c999a75962a12a40699ca48a2b0154167fba83ed17c16070da0debaa680185af188d82c999a75962a12a40699ca48a2b015561091461537c565b7fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054167ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00557fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2602060405160028152a1005b50611ccb73ffffffffffffffffffffffffffffffffffffffff7fba83ed17c16070da0debaa680185af188d82c999a75962a12a40699ca48a2b01541633149050611cc0565b7ff92ee8a9000000000000000000000000000000000000000000000000000000005f5260045ffd5b60029150101582611b75565b34610df9575f600319360112610df95760207f515a8336edcaab4ae6524d41223c1782132890f89189ba6632107a7b5a44960254604051908152f35b34610df9576020600319360112610df95760043567ffffffffffffffff8111610df957610ecc6116206020923690600401611f84565b34610df9576103186103046111d4611ef136611ff8565b90611f009492939436336140e2565b611f0d8486811515612e11565b611f258486611f206111ca368484612154565b612e11565b611f30368587612154565b6148c7565b34610df9576020600319360112610df95760043567ffffffffffffffff8111610df95760a06003198236030112610df9576111e090611f7261406e565b611f7c36336140e2565b6004016128d1565b9181601f84011215610df95782359167ffffffffffffffff8311610df95760208381860195010111610df957565b6004359073ffffffffffffffffffffffffffffffffffffffff82168203610df957565b6024359073ffffffffffffffffffffffffffffffffffffffff82168203610df957565b6060600319820112610df95760043567ffffffffffffffff8111610df9578161202391600401611f84565b929092916024359067ffffffffffffffff8211610df95760031982604092030112610df9576004019060443573ffffffffffffffffffffffffffffffffffffffff81168103610df95790565b90601f19601f602080948051918291828752018686015e5f8582860101520116010190565b60a0810190811067ffffffffffffffff8211176120b057604052565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6040810190811067ffffffffffffffff8211176120b057604052565b6080810190811067ffffffffffffffff8211176120b057604052565b90601f601f19910116810190811067ffffffffffffffff8211176120b057604052565b67ffffffffffffffff81116120b057601f01601f191660200190565b92919261216082612138565b9161216e6040519384612115565b829481845281830111610df9578281602093845f960137010152565b9080601f83011215610df9578160206121a593359101612154565b90565b6020600319820112610df9576004359067ffffffffffffffff8211610df95760031982608092030112610df95760040190565b6040600319820112610df95760043567ffffffffffffffff8111610df9578161220691600401611f84565b929092916024359067ffffffffffffffff8211610df95761222991600401611f84565b9091565b90600382101561223a5752565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b9080602083519182815201916020808360051b8301019401925f915b83831061229257505050505090565b90919293946020806122b083601f198660019603018752895161206f565b97019301930191939290612283565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6181360301821215610df9570190565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610df9570180359067ffffffffffffffff8211610df957602001918160051b36038313610df957565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610df9570180359067ffffffffffffffff8211610df957602001918136038313610df957565b601f8260209493601f1993818652868601375f8582860101520116010190565b929092156123c457505050565b6124079291611931916040519485947f9fff831f00000000000000000000000000000000000000000000000000000000865260406004870152604486019061206f565b91600319858403016024860152612397565b3567ffffffffffffffff81168103610df95790565b67ffffffffffffffff81116120b05760051b60200190565b604080519091906124578382612115565b6001815291601f1901825f5b82811061246f57505050565b806060602080938501015201612463565b805115611aeb5760200190565b8051821015611aeb5760209160051b010190565b359067ffffffffffffffff82168203610df957565b9190826040910312610df9576040516124ce816120dd565b60206124e78183956124df816124a1565b8552016124a1565b910152565b906121a591602081526060612548612510845160a0602086015260c085019061206f565b602085810151805167ffffffffffffffff90811660408801529101511660608501526040850151601f19858303016080860152612267565b9201519060a0601f198285030191015261206f565b919060a083820312610df9576040519061257682612094565b8193803567ffffffffffffffff8111610df9578261259591830161218a565b8352602081013567ffffffffffffffff8111610df957826125b791830161218a565b6020840152604081013567ffffffffffffffff8111610df957826125dc91830161218a565b6040840152606081013567ffffffffffffffff8111610df9578261260191830161218a565b606084015260808101359167ffffffffffffffff8311610df9576080926124e7920161218a565b6121a591608061268061266e61265c61264a865160a0875260a087019061206f565b6020870151868203602088015261206f565b6040860151858203604087015261206f565b6060850151848203606086015261206f565b92015190608081840391015261206f565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe182360301811215610df957016020813591019167ffffffffffffffff8211610df9578136038313610df957565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe182360301811215610df957016020813591019167ffffffffffffffff8211610df9578160051b36038313610df957565b9067ffffffffffffffff612747836124a1565b1681526127b261278c6127716127606020860186612691565b60a0602087015260a0860191612397565b61277e6040860186612691565b908583036040870152612397565b9267ffffffffffffffff6127a2606083016124a1565b16606084015260808101906126e1565b9290916080818303910152828152602081019260208160051b83010193835f917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6182360301945b84841061280a575050505050505090565b90919293949596601f19828203018352873587811215610df95760206128c060019387839401906128b26128a761288c61287161285861284a8780612691565b60a0885260a0880191612397565b61286489880188612691565b908783038b890152612397565b61287e6040870187612691565b908683036040880152612397565b6128996060860186612691565b908583036060870152612397565b926080810190612691565b916080818503910152612397565b9901930194019291959493906127f9565b60016128ea6128e083806122bf565b60808101906122f2565b905003612de9576128fe6128e082806122bf565b15611aeb5761290e8183926122bf565b5f6020612aa18161292e610af361292588806122bf565b83810190612346565b61297281518381519101206129536111d46129498b806122bf565b6040810190612346565b8481519101201482519061296a6129498b806122bf565b9290916123b7565b6129a66129a16129856129498a806122bf565b91906129996129948c806122bf565b612419565b923691612154565b614367565b90612a1f6129fb6129e38a6129ca6111d4886129c0612446565b93019e8f90612346565b6129d382612480565b526129dd81612480565b506143f8565b936129f160408c018c612346565b969094015161451d565b916040519387850152868452612a12604085612115565b604051946111d4866120f9565b8352612a2e3660608a016124b6565b858401526040830152606082015273ffffffffffffffffffffffffffffffffffffffff612a6a610ec6612a6189806122bf565b86810190612346565b16906040519485809481937f682ed5f0000000000000000000000000000000000000000000000000000000008352600483016124ec565b03925af18015610dee57612dba575b50612ac3612abe83806122bf565b6145db565b15612d915773ffffffffffffffffffffffffffffffffffffffff612aea6116208380612346565b16612b02612af884806122bf565b6020810190612346565b612b1261294986809694966122bf565b612b2261299488809594956122bf565b91612b2d8989612346565b9790916040519560c087019487861067ffffffffffffffff8711176120b057612b65612b8e94612b6e94612b9d986040523691612154565b88523691612154565b956020860196875267ffffffffffffffff6040870195168552369061255d565b96606085019788523691612154565b6080830190815260a0830191338352853b15610df95760405196879586957f428e4e170000000000000000000000000000000000000000000000000000000087526004870160209052516024870160c0905260e48701612bfc9161206f565b9051908681037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc016044880152612c329161206f565b915167ffffffffffffffff16606486015251908481037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc016084860152612c7891612628565b9051908381037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc0160a4850152612cae9161206f565b905173ffffffffffffffffffffffffffffffffffffffff1660c483015203815a5f948591f18015610dee57612d7c927ff9bab74bcdb634f4d3dd064cc42a13df056598e1c0336905d2f5750fbfb08b7b92612d6c92612d81575b50612d16612af882806122bf565b949095612d3a612d2961299485806122bf565b91612d3485806122bf565b94612346565b96909781604051928392833781015f81520390209567ffffffffffffffff604051958695604087526040870190612734565b9285840360208701521697612397565b0390a3565b5f612d8b91612115565b5f612d08565b5050507fd08bf58b0e4eec5bfc697a4fdbb6839057fbf4dd06f1b1ce07445c0e5a654caf5f80a1565b6020813d602011612de1575b81612dd360209383612115565b81010312610df95751612ab0565b3d9150612dc6565b7f356f4dbd000000000000000000000000000000000000000000000000000000005f5260045ffd5b91909115612e1d575050565b6119316040519283927f4870bd74000000000000000000000000000000000000000000000000000000008452602060048501526024840191612397565b60209082604051938492833781017f515a8336edcaab4ae6524d41223c1782132890f89189ba6632107a7b5a44960181520301902090565b73ffffffffffffffffffffffffffffffffffffffff604051838382376020818581017fc5779f3c2c21083eefa6d04f6a698bc0d8c10db124ad5e0df6ef394b6d7bf6008152030190205416918215612ee957505090565b6119316040519283927fa09dbf59000000000000000000000000000000000000000000000000000000008452602060048501526024840191612397565b15612f2e5750565b73ffffffffffffffffffffffffffffffffffffffff907fbe2f2b45000000000000000000000000000000000000000000000000000000005f521660045260245ffd5b15612f79575050565b67ffffffffffffffff907f65d30129000000000000000000000000000000000000000000000000000000005f521660045260245260445ffd5b9190820391821161196b57565b906020825267ffffffffffffffff81511660208301526080613009612ff3602084015160a0604087015260c086019061206f565b6040840151601f1986830301606087015261206f565b9167ffffffffffffffff6060820151168285015201519160a0601f1982840301910152815180825260208201916020808360051b8301019401925f915b83831061305557505050505090565b909192939460208061307383601f1986600196030187528951612628565b97019301930191939290613046565b919060a083820312610df95760405161309a81612094565b80936130a5816124a1565b8252602081013567ffffffffffffffff8111610df957836130c791830161218a565b6020830152604081013567ffffffffffffffff8111610df957836130ec91830161218a565b60408301526130fd606082016124a1565b606083015260808101359067ffffffffffffffff8211610df957019180601f84011215610df957823561312f8161242e565b9361313d6040519586612115565b81855260208086019260051b82010191838311610df95760208201905b83821061316c57505050505060800152565b813567ffffffffffffffff8111610df95760209161318f8784809488010161255d565b81520191019061315a565b90608073ffffffffffffffffffffffffffffffffffffffff816132046131de6131cc875160a0885260a088019061206f565b6020880151878203602089015261206f565b67ffffffffffffffff604088015116604087015260608701518682036060880152612628565b9401511691015290565b6040519061321d604083612115565b602082527f4774d4a575993f963b1c06573736617a457abef8589178db8d10c94b4ab511ab6020830152565b3d15613273573d9061325a82612138565b916132686040519384612115565b82523d5f602084013e565b606090565b60016132876128e083806122bf565b905003612de95761329b6128e082806122bf565b15611aeb57806132aa916122bf565b905f60206133be816132c2610af361294987806122bf565b6132f481518381519101206132dd6111d4612a618a806122bf565b8481519101201482519061296a612a618a806122bf565b61331d61330c606061330689806122bf565b01612419565b429061165a60606133068b806122bf565b61334d61334861333961333089806122bf565b85810190612346565b91906129996129948b806122bf565b614e1f565b9061337c6129fb61336f61336a366133658c806122bf565b613082565b614e99565b936129f1868b018b612346565b835261338b36604089016124b6565b858401526040830152606082015273ffffffffffffffffffffffffffffffffffffffff612a6a610ec661294988806122bf565b03925af18015610dee57613812575b506133e06133db82806122bf565b6150fe565b156137ea575f806134c06133f2612446565b9467ffffffffffffffff61347873ffffffffffffffffffffffffffffffffffffffff6134246116206020860186612346565b1692613433612af889806122bf565b93906134668a612994612b6561345861344f61294985806122bf565b939094806122bf565b94604051996111d48b612094565b6020860152166040840152369061255d565b60608201523360808201526040519485809481937f078c4a7900000000000000000000000000000000000000000000000000000000835260206004840152602483019061319a565b03925af15f918161376e575b506136e357506134da613249565b8051156136bb5761351a7fb9edb487876e8be10f54e377c1a815a54ad92a6db1c9561dfe8fad2f0d1da84f9160405191829160208352602083019061206f565b0390a161352561320e565b61352e83612480565b5261353882612480565b505b61354481806122bf565b604081016135b36129996129a161356c6129a16135618688612346565b919061299989612419565b6020815191012094855f527f1260944489272988d9df285149b5aa1b0f48f2136d6f416159f840a3e07476006020526135ab60405f2054159582612346565b939091612419565b901561367d57506135c3836143f8565b905f527f1260944489272988d9df285149b5aa1b0f48f2136d6f416159f840a3e074760060205260405f20557f76765590e2b799b0506100f8a6610cfecab2c71e8e1f8aa981b099aff0dfdb7461366d612d7c61362361294985806122bf565b61363d61363661299488809695966122bf565b96806122bf565b9281604051928392833781015f81520390209467ffffffffffffffff604051948594604086526040860190612734565b9184830360208601521696612267565b611931906040519182917f40470d7400000000000000000000000000000000000000000000000000000000835260206004840152602483019061206f565b7fadef7fb8000000000000000000000000000000000000000000000000000000005f5260045ffd5b80511561374657805160208201206136f961320e565b602081519101201461371e5761370e83612480565b5261371882612480565b5061353a565b7f6b2675e3000000000000000000000000000000000000000000000000000000005f5260045ffd5b7fecfef798000000000000000000000000000000000000000000000000000000005f5260045ffd5b9091503d805f833e6137808183612115565b810190602081830312610df95780519067ffffffffffffffff8211610df9570181601f82011215610df9578051906137b782612138565b926137c56040519485612115565b82845260208383010111610df957815f9260208093018386015e83010152905f6134cc565b50507fd08bf58b0e4eec5bfc697a4fdbb6839057fbf4dd06f1b1ce07445c0e5a654caf5f80a1565b6020813d602011613839575b8161382b60209383612115565b81010312610df957516133cd565b3d915061381e565b9190911561384d575050565b6119316040519283927f14d71247000000000000000000000000000000000000000000000000000000008452602060048501526024840191612397565b73ffffffffffffffffffffffffffffffffffffffff604051838382376020818581017f515a8336edcaab4ae6524d41223c1782132890f89189ba6632107a7b5a44960081520301902054169182156138e157505090565b6119316040519283927fa0db16fe000000000000000000000000000000000000000000000000000000008452602060048501526024840191612397565b90600182811c92168015613965575b602083101461393857565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b91607f169161392d565b6060602060405161397f816120dd565b828152015261398e8282612e5a565b916040519261399c846120dd565b6040515f82546139ab8161391e565b8084529060018116908115613b615750600114613b1e575b50906139d481600194930382612115565b8552018054906139e38261242e565b916139f16040519384612115565b80835260208301915f5260205f20915f905b828210613a5d5750505050602084015282515115613a2057505090565b6119316040519283927fdf95155a000000000000000000000000000000000000000000000000000000008452602060048501526024840191612397565b6040515f8554613a6c8161391e565b8084529060018116908115613add5750600114613aa6575b5060019282613a9885946020940382612115565b815201940191019092613a03565b5f878152602081209092505b818310613ac757505081016020016001613a84565b6001816020925483868801015201920191613ab2565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660208581019190915291151560051b8401909101915060019050613a84565b5f8481526020812094939250905b808210613b4557509192509081016020016139d46139c3565b9192936001816020925483858801015201910190939291613b2c565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660208086019190915291151560051b840190910191506139d490506139c3565b6001613bb26128e083806122bf565b905003612de957613bc66128e082806122bf565b15611aeb5780613bd5916122bf565b906020613be8610af361292584806122bf565b613c1a8151838151910120613c036111d461294987806122bf565b8481519101201482519061296a61294987806122bf565b613c5a613c44613c3f613c3061294987806122bf565b919061299961299489806122bf565b61528b565b613c5084860186612346565b949093015161451d565b916040516060810181811067ffffffffffffffff8211176120b057602093613d4593613c8a926040523691612154565b8152613d235f613c9d36604089016124b6565b9285810193845260408101968752613d7573ffffffffffffffffffffffffffffffffffffffff613cdc610ec6613cd38c806122bf565b8a810190612346565b1694604051988997889687957f4d6d9ffb0000000000000000000000000000000000000000000000000000000087528b6004880152516080602488015260a487019061206f565b9251604486019067ffffffffffffffff60208092828151168552015116910152565b517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc848303016084850152612267565b03925af18015610dee575f90613f4d575b613dbc915067ffffffffffffffff613da3606061330686806122bf565b16811015613db6606061330686806122bf565b90612f70565b613dc9612abe82806122bf565b156137ea5781613e5773ffffffffffffffffffffffffffffffffffffffff613dfd6116208467ffffffffffffffff97612346565b1691613e0c612af885806122bf565b9590613e45613e1e61294988806122bf565b613e3c613e2e6129948b806122bf565b946040519b6111d48d612094565b8a523691612154565b6020880152166040860152369061255d565b6060840152336080840152803b15610df957613eae5f939184926040519586809481937f5e32b6b600000000000000000000000000000000000000000000000000000000835260206004840152602483019061319a565b03925af1918215610dee5767ffffffffffffffff92613f3d575b507f01e5ed58494819ef3f6480dd08e433b7c08ed75c7abdf2c22c6f04b71340a168613ef7612af883806122bf565b613f11613f0a61299486809895986122bf565b94806122bf565b9481604051928392833781015f815203902092612d7c6040519283926020845216956020830190612734565b5f613f4791612115565b5f613ec8565b506020813d602011613f79575b81613f6760209383612115565b81010312610df957613dbc9051613d86565b3d9150613f5a565b818110613f8c575050565b5f8155600101613f81565b9190601f8111613fa657505050565b613fd0925f5260205f20906020601f840160051c83019310613fd2575b601f0160051c0190613f81565b565b9091508190613fc3565b90614006613ffb613fed8480612691565b604085526040850191612397565b9260208101906126e1565b90916020818503910152808352602083019260208260051b82010193835f925b8484106140365750505050505090565b90919293949560208061405e83601f1986600196030188526140588b88612691565b90612397565b9801940194019294939190614026565b7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f005c6140ba5760017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f005d565b7f3ee5aeb5000000000000000000000000000000000000000000000000000000005f5260045ffd5b7ff3177357ab46d8af007ab3fdb9af81da189e1068fefdc0073dca88a2cab40a00549173ffffffffffffffffffffffffffffffffffffffff83169281600411610df9575f5f9060405f81519673ffffffffffffffffffffffffffffffffffffffff60208901917fb700961300000000000000000000000000000000000000000000000000000000835216978860248201523060448201527fffffffff000000000000000000000000000000000000000000000000000000008335166064820152606481526141b1608482612115565b828052826020525190895afa614354575b156141cf575b5050505050565b63ffffffff1615614328577fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1674010000000000000000000000000000000000000000177ff3177357ab46d8af007ab3fdb9af81da189e1068fefdc0073dca88a2cab40a0055823b15610df9576020925f92836040518096819582947f94c7d7ee000000000000000000000000000000000000000000000000000000008452600484015260406024840152601f19601f6044850192808452808786860137868582860101520116010103925af18015610dee57614318575b507fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff7ff3177357ab46d8af007ab3fdb9af81da189e1068fefdc0073dca88a2cab40a0054167ff3177357ab46d8af007ab3fdb9af81da189e1068fefdc0073dca88a2cab40a00555f808080806141c8565b5f61432291612115565b5f6142a7565b827f068ca9d8000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b50505f516020518060201c1502906141c2565b60096121a5916020937fffffffffffffffff000000000000000000000000000000000000000000000000856040519687948051918291018387015e840101917f0300000000000000000000000000000000000000000000000000000000000000835260c01b1660018201520301601f198101835282612115565b60209291908391805192839101825e019081520190565b908151156144f5576020916040516144108482612115565b5f8152905f915b815183101561447c57845f8161442d868661248d565b51604051918183925191829101835e8101838152039060025afa15610dee576001906144745f51916144666040519384928a84016143e1565b03601f198101835282612115565b920191614417565b90505f91509291926040516144d560218286808201957f020000000000000000000000000000000000000000000000000000000000000087528051918291018484015e810186838201520301601f198101835282612115565b604051918291518091835e8101838152039060025afa15610dee575f5190565b7f760d6a9b000000000000000000000000000000000000000000000000000000005f5260045ffd5b908151156145a05781515f19810190811161196b57602091828061454461457a948761248d565b51926040519584879551918291018487015e8401908282015f8152815193849201905e01015f815203601f198101835282612115565b81515f19810190811161196b5761459c91614595828561248d565b528261248d565b5090565b6040517fa7c34e4f00000000000000000000000000000000000000000000000000000000815260206004820152806119316024820185612267565b6145f96133486145ee6020840184612346565b919061299985612419565b6020815191012090815f527f1260944489272988d9df285149b5aa1b0f48f2136d6f416159f840a3e074760060205260405f205480156146b75761465061336a61464661336a3686613082565b8314933690613082565b91156146895750505f527f1260944489272988d9df285149b5aa1b0f48f2136d6f416159f840a3e07476006020525f6040812055600190565b7f3f87a2ec000000000000000000000000000000000000000000000000000000005f5260045260245260445ffd5b5050505f90565b604051906146cd604083612115565b600782527f636c69656e742d000000000000000000000000000000000000000000000000006020830152565b908151811015611aeb570160200190565b8051600481109081156148bc575b506148a05761475e60405161472e604082612115565b600881527f6368616e6e656c2d000000000000000000000000000000000000000000000000602082015282615305565b80156148a5575b6148a0575f5b81518110156148995761477e81836146f9565b5160f81c6061811015908161488d575b811561486f575b8115614851575b8115614812575b81156147bd575b506147b55750505f90565b60010161476b565b6023811491508115614807575b81156147fc575b81156147f1575b81156147e6575b505f6147aa565b603e9150145f6147df565b603c811491506147d8565b605d811491506147d1565b605b811491506147ca565b9050602e81148015614847575b801561483d575b8015614833575b906147a3565b50602d811461482d565b50602b8114614826565b50605f811461481f565b9050604181101580614864575b9061479c565b50605a81111561485e565b9050603081101580614882575b90614795565b50603981111561487c565b607a811115915061478e565b5050600190565b505f90565b506148b76148b16146be565b82615305565b614765565b60809150115f614718565b91906040519173ffffffffffffffffffffffffffffffffffffffff845193602081818801968088835e81017f515a8336edcaab4ae6524d41223c1782132890f89189ba6632107a7b5a4496008152030190205416614c815773ffffffffffffffffffffffffffffffffffffffff169160405160208186518085835e81017f515a8336edcaab4ae6524d41223c1782132890f89189ba6632107a7b5a44960081520301902073ffffffffffffffffffffffffffffffffffffffff84167fffffffffffffffffffffffff00000000000000000000000000000000000000008254161790556020604051809286518091835e81017f515a8336edcaab4ae6524d41223c1782132890f89189ba6632107a7b5a4496018152030190206149e98280612346565b9067ffffffffffffffff82116120b057614a0782610435855461391e565b5f90601f8311600114614c1a579180614a3792600195945f926106035750505f198260011b9260031b1c19161790565b81555b01614a4860208301836122f2565b906801000000000000000082116120b0578254828455808310614ba3575b505f928352602083209290805b838310614ac85750505050509161057791614abd7f0ecded31ecd211a73abf0fb3bc09150bbe321a05550fbe29ea0f16b6e25fbfa89460405194859460608652606086019061206f565b9060408301520390a1565b614ad28183612346565b9067ffffffffffffffff82116120b057614af682614af0895461391e565b89613f97565b5f90601f8311600114614b395792614b2a836001959460209487965f926106035750505f198260011b9260031b1c19161790565b88555b01950192019193614a73565b601f19831691885f5260205f20925f5b818110614b8b5750936020936001969387969383889510614b72575b505050811b018855614b2d565b5f1960f88560031b161c199101351690555f8080614b65565b91936020600181928787013581550195019201614b49565b835f528260205f2091820191015b818110614bbe5750614a66565b80614bcb6001925461391e565b80614bd8575b5001614bb1565b601f81118314614bed57505f81555b5f614bd1565b614c0990825f5283601f60205f20920160051c82019101613f81565b805f525f6020812081835555614be7565b601f19831691845f5260205f20925f5b818110614c695750916001959492918387959310614c50575b505050811b018155614a3a565b5f1960f88560031b161c199101351690555f8080614c43565b91936020600181928787013581550195019201614c2a565b6040517f87dfb2670000000000000000000000000000000000000000000000000000000081526020600482015280611931602482018761206f565b906040519073ffffffffffffffffffffffffffffffffffffffff835192602081818701958087835e81017fc5779f3c2c21083eefa6d04f6a698bc0d8c10db124ad5e0df6ef394b6d7bf6008152030190205416614de45791614dd99173ffffffffffffffffffffffffffffffffffffffff7fa6ec8e860960e638347460dc632fbe0175c51a5ca130e336138bbe26ff3044999416906020604051809285518091835e81017fc5779f3c2c21083eefa6d04f6a698bc0d8c10db124ad5e0df6ef394b6d7bf60081520301902073ffffffffffffffffffffffffffffffffffffffff82167fffffffffffffffffffffffff000000000000000000000000000000000000000082541617905560405192839260408452604084019061206f565b9060208301520390a1565b6040517f837f46a60000000000000000000000000000000000000000000000000000000081526020600482015280611931602482018661206f565b60096121a5916020937fffffffffffffffff000000000000000000000000000000000000000000000000856040519687948051918291018387015e840101917f0100000000000000000000000000000000000000000000000000000000000000835260c01b1660018201520301601f198101835282612115565b90602091604051614eaa8482612115565b5f8152905f915b608082015180518410156150005783614ec99161248d565b51855f818351604051918183925191829101835e8101838152039060025afa15610dee575f5190865f8180840151604051918183925191829101835e8101838152039060025afa15610dee575f5191875f816040850151604051918183925191829101835e8101838152039060025afa15610dee575f5192885f816060860151604051918183925191829101835e8101838152039060025afa15610dee57885f8160808251960151604051918183925191829101835e8101838152039060025afa15610dee5788935f938451916040519387850195865260408501526060840152608083015260a082015260a08152614fc360c082612115565b604051918291518091835e8101838152039060025afa15610dee57600190614ff85f51916144666040519384928a84016143e1565b920191614eb1565b509150929192825f816040840151604051918183925191829101835e8101838152039060025afa15610dee57825f606081519301516040517fffffffffffffffff0000000000000000000000000000000000000000000000008482019260c01b16825260088152615072602882612115565b604051918291518091835e8101838152039060025afa15610dee57825f81815194604051918183925191829101835e8101838152039060025afa15610dee575f91825160405191858301937f02000000000000000000000000000000000000000000000000000000000000008552602184015260418301526061820152606181526144d5608182612115565b61512e61511f613c3f6151146040850185612346565b919061299986612419565b60208151910120913690613082565b60405161514381614466602082019485612fbf565b51902090805f527f1260944489272988d9df285149b5aa1b0f48f2136d6f416159f840a3e074760060205260405f20548281146146b757806151b057505f527f1260944489272988d9df285149b5aa1b0f48f2136d6f416159f840a3e074760060205260405f2055600190565b90507f657b94fe000000000000000000000000000000000000000000000000000000005f5260045260245260445ffd5b602073ffffffffffffffffffffffffffffffffffffffff7f2f658b440c35314f52658ea8a740e05b284cdc84dc9ae01e891f21b8933e7cad9216807fffffffffffffffffffffffff00000000000000000000000000000000000000007ff3177357ab46d8af007ab3fdb9af81da189e1068fefdc0073dca88a2cab40a005416177ff3177357ab46d8af007ab3fdb9af81da189e1068fefdc0073dca88a2cab40a0055604051908152a1565b60096121a5916020937fffffffffffffffff000000000000000000000000000000000000000000000000856040519687948051918291018387015e840101917f0200000000000000000000000000000000000000000000000000000000000000835260c01b1660018201520301601f198101835282612115565b805190825180831061537457828082109118028083189214158202821890602061532f8385612fb2565b928061535361533d86612138565b9561534b6040519788612115565b808752612138565b95601f19848701970136883703920101835e51902090602081519101201490565b505050505f90565b60ff7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005460401c16156153ab57565b7fd7e6bcf8000000000000000000000000000000000000000000000000000000005f5260045ffd5b80518211801561548f575b615438576001821180615440575b158015908160011b918204600214171561196b576028018060281161196b5782036154385773ffffffffffffffffffffffffffffffffffffffff92915f6154329261552f565b90921690565b50505f905f90565b507f30780000000000000000000000000000000000000000000000000000000000007fffff000000000000000000000000000000000000000000000000000000000000602083015116146153ec565b505f6153de565b906154d357508051156154ab57805190602001fd5b7fd6bda275000000000000000000000000000000000000000000000000000000005f5260045ffd5b81511580615526575b6154e4575090565b73ffffffffffffffffffffffffffffffffffffffff907f9996b315000000000000000000000000000000000000000000000000000000005f521660045260245ffd5b50803b156154dc565b9290926001840180851161196b578311806155e5575b15938415948560011b958604600214171561196b575f94810180911161196b579192905b8183106155795750505060019190565b9092919360ff6155b07fff000000000000000000000000000000000000000000000000000000000000006020888601015116615636565b16600f81116155da578160041b918083046010149015171561196b57600191019401919290615569565b505f94508493505050565b507f30780000000000000000000000000000000000000000000000000000000000007fffff000000000000000000000000000000000000000000000000000000000000602086840101511614615545565b60f81c602f8111806156f8575b15615670577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffd00160ff1690565b60608111806156ee575b156156a7577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa90160ff1690565b60408111806156e4575b156156de577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc90160ff1690565b5060ff90565b50604781106156b1565b506067811061567a565b50603a811061564356fea164736f6c634300081c000af0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00",
}

// ContractICS26RouterABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractICS26RouterMetaData.ABI instead.
var ContractICS26RouterABI = ContractICS26RouterMetaData.ABI

// ContractICS26RouterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractICS26RouterMetaData.Bin instead.
var ContractICS26RouterBin = ContractICS26RouterMetaData.Bin

// DeployContractICS26Router deploys a new Ethereum contract, binding an instance of ContractICS26Router to it.
func DeployContractICS26Router(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ContractICS26Router, error) {
	parsed, err := ContractICS26RouterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractICS26RouterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractICS26Router{ContractICS26RouterCaller: ContractICS26RouterCaller{contract: contract}, ContractICS26RouterTransactor: ContractICS26RouterTransactor{contract: contract}, ContractICS26RouterFilterer: ContractICS26RouterFilterer{contract: contract}}, nil
}

// ContractICS26Router is an auto generated Go binding around an Ethereum contract.
type ContractICS26Router struct {
	ContractICS26RouterCaller     // Read-only binding to the contract
	ContractICS26RouterTransactor // Write-only binding to the contract
	ContractICS26RouterFilterer   // Log filterer for contract events
}

// ContractICS26RouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractICS26RouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractICS26RouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractICS26RouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractICS26RouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractICS26RouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractICS26RouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractICS26RouterSession struct {
	Contract     *ContractICS26Router // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ContractICS26RouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractICS26RouterCallerSession struct {
	Contract *ContractICS26RouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// ContractICS26RouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractICS26RouterTransactorSession struct {
	Contract     *ContractICS26RouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// ContractICS26RouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractICS26RouterRaw struct {
	Contract *ContractICS26Router // Generic contract binding to access the raw methods on
}

// ContractICS26RouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractICS26RouterCallerRaw struct {
	Contract *ContractICS26RouterCaller // Generic read-only contract binding to access the raw methods on
}

// ContractICS26RouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractICS26RouterTransactorRaw struct {
	Contract *ContractICS26RouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractICS26Router creates a new instance of ContractICS26Router, bound to a specific deployed contract.
func NewContractICS26Router(address common.Address, backend bind.ContractBackend) (*ContractICS26Router, error) {
	contract, err := bindContractICS26Router(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractICS26Router{ContractICS26RouterCaller: ContractICS26RouterCaller{contract: contract}, ContractICS26RouterTransactor: ContractICS26RouterTransactor{contract: contract}, ContractICS26RouterFilterer: ContractICS26RouterFilterer{contract: contract}}, nil
}

// NewContractICS26RouterCaller creates a new read-only instance of ContractICS26Router, bound to a specific deployed contract.
func NewContractICS26RouterCaller(address common.Address, caller bind.ContractCaller) (*ContractICS26RouterCaller, error) {
	contract, err := bindContractICS26Router(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractICS26RouterCaller{contract: contract}, nil
}

// NewContractICS26RouterTransactor creates a new write-only instance of ContractICS26Router, bound to a specific deployed contract.
func NewContractICS26RouterTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractICS26RouterTransactor, error) {
	contract, err := bindContractICS26Router(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractICS26RouterTransactor{contract: contract}, nil
}

// NewContractICS26RouterFilterer creates a new log filterer instance of ContractICS26Router, bound to a specific deployed contract.
func NewContractICS26RouterFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractICS26RouterFilterer, error) {
	contract, err := bindContractICS26Router(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractICS26RouterFilterer{contract: contract}, nil
}

// bindContractICS26Router binds a generic wrapper to an already deployed contract.
func bindContractICS26Router(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractICS26RouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractICS26Router *ContractICS26RouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractICS26Router.Contract.ContractICS26RouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractICS26Router *ContractICS26RouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractICS26Router.Contract.ContractICS26RouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractICS26Router *ContractICS26RouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractICS26Router.Contract.ContractICS26RouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractICS26Router *ContractICS26RouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractICS26Router.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractICS26Router *ContractICS26RouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractICS26Router.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractICS26Router *ContractICS26RouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractICS26Router.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_ContractICS26Router *ContractICS26RouterCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ContractICS26Router.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_ContractICS26Router *ContractICS26RouterSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _ContractICS26Router.Contract.UPGRADEINTERFACEVERSION(&_ContractICS26Router.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_ContractICS26Router *ContractICS26RouterCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _ContractICS26Router.Contract.UPGRADEINTERFACEVERSION(&_ContractICS26Router.CallOpts)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() view returns(address)
func (_ContractICS26Router *ContractICS26RouterCaller) Authority(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractICS26Router.contract.Call(opts, &out, "authority")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() view returns(address)
func (_ContractICS26Router *ContractICS26RouterSession) Authority() (common.Address, error) {
	return _ContractICS26Router.Contract.Authority(&_ContractICS26Router.CallOpts)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() view returns(address)
func (_ContractICS26Router *ContractICS26RouterCallerSession) Authority() (common.Address, error) {
	return _ContractICS26Router.Contract.Authority(&_ContractICS26Router.CallOpts)
}

// GetClient is a free data retrieval call binding the contract method 0x7eb78932.
//
// Solidity: function getClient(string clientId) view returns(address)
func (_ContractICS26Router *ContractICS26RouterCaller) GetClient(opts *bind.CallOpts, clientId string) (common.Address, error) {
	var out []interface{}
	err := _ContractICS26Router.contract.Call(opts, &out, "getClient", clientId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetClient is a free data retrieval call binding the contract method 0x7eb78932.
//
// Solidity: function getClient(string clientId) view returns(address)
func (_ContractICS26Router *ContractICS26RouterSession) GetClient(clientId string) (common.Address, error) {
	return _ContractICS26Router.Contract.GetClient(&_ContractICS26Router.CallOpts, clientId)
}

// GetClient is a free data retrieval call binding the contract method 0x7eb78932.
//
// Solidity: function getClient(string clientId) view returns(address)
func (_ContractICS26Router *ContractICS26RouterCallerSession) GetClient(clientId string) (common.Address, error) {
	return _ContractICS26Router.Contract.GetClient(&_ContractICS26Router.CallOpts, clientId)
}

// GetCommitment is a free data retrieval call binding the contract method 0x7795820c.
//
// Solidity: function getCommitment(bytes32 hashedPath) view returns(bytes32)
func (_ContractICS26Router *ContractICS26RouterCaller) GetCommitment(opts *bind.CallOpts, hashedPath [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ContractICS26Router.contract.Call(opts, &out, "getCommitment", hashedPath)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetCommitment is a free data retrieval call binding the contract method 0x7795820c.
//
// Solidity: function getCommitment(bytes32 hashedPath) view returns(bytes32)
func (_ContractICS26Router *ContractICS26RouterSession) GetCommitment(hashedPath [32]byte) ([32]byte, error) {
	return _ContractICS26Router.Contract.GetCommitment(&_ContractICS26Router.CallOpts, hashedPath)
}

// GetCommitment is a free data retrieval call binding the contract method 0x7795820c.
//
// Solidity: function getCommitment(bytes32 hashedPath) view returns(bytes32)
func (_ContractICS26Router *ContractICS26RouterCallerSession) GetCommitment(hashedPath [32]byte) ([32]byte, error) {
	return _ContractICS26Router.Contract.GetCommitment(&_ContractICS26Router.CallOpts, hashedPath)
}

// GetCounterparty is a free data retrieval call binding the contract method 0xb0777bfa.
//
// Solidity: function getCounterparty(string clientId) view returns((string,bytes[]))
func (_ContractICS26Router *ContractICS26RouterCaller) GetCounterparty(opts *bind.CallOpts, clientId string) (IICS02ClientMsgsCounterpartyInfo, error) {
	var out []interface{}
	err := _ContractICS26Router.contract.Call(opts, &out, "getCounterparty", clientId)

	if err != nil {
		return *new(IICS02ClientMsgsCounterpartyInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IICS02ClientMsgsCounterpartyInfo)).(*IICS02ClientMsgsCounterpartyInfo)

	return out0, err

}

// GetCounterparty is a free data retrieval call binding the contract method 0xb0777bfa.
//
// Solidity: function getCounterparty(string clientId) view returns((string,bytes[]))
func (_ContractICS26Router *ContractICS26RouterSession) GetCounterparty(clientId string) (IICS02ClientMsgsCounterpartyInfo, error) {
	return _ContractICS26Router.Contract.GetCounterparty(&_ContractICS26Router.CallOpts, clientId)
}

// GetCounterparty is a free data retrieval call binding the contract method 0xb0777bfa.
//
// Solidity: function getCounterparty(string clientId) view returns((string,bytes[]))
func (_ContractICS26Router *ContractICS26RouterCallerSession) GetCounterparty(clientId string) (IICS02ClientMsgsCounterpartyInfo, error) {
	return _ContractICS26Router.Contract.GetCounterparty(&_ContractICS26Router.CallOpts, clientId)
}

// GetIBCApp is a free data retrieval call binding the contract method 0x2447af29.
//
// Solidity: function getIBCApp(string portId) view returns(address)
func (_ContractICS26Router *ContractICS26RouterCaller) GetIBCApp(opts *bind.CallOpts, portId string) (common.Address, error) {
	var out []interface{}
	err := _ContractICS26Router.contract.Call(opts, &out, "getIBCApp", portId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetIBCApp is a free data retrieval call binding the contract method 0x2447af29.
//
// Solidity: function getIBCApp(string portId) view returns(address)
func (_ContractICS26Router *ContractICS26RouterSession) GetIBCApp(portId string) (common.Address, error) {
	return _ContractICS26Router.Contract.GetIBCApp(&_ContractICS26Router.CallOpts, portId)
}

// GetIBCApp is a free data retrieval call binding the contract method 0x2447af29.
//
// Solidity: function getIBCApp(string portId) view returns(address)
func (_ContractICS26Router *ContractICS26RouterCallerSession) GetIBCApp(portId string) (common.Address, error) {
	return _ContractICS26Router.Contract.GetIBCApp(&_ContractICS26Router.CallOpts, portId)
}

// GetNextClientSeq is a free data retrieval call binding the contract method 0x27f146f3.
//
// Solidity: function getNextClientSeq() view returns(uint256)
func (_ContractICS26Router *ContractICS26RouterCaller) GetNextClientSeq(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractICS26Router.contract.Call(opts, &out, "getNextClientSeq")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextClientSeq is a free data retrieval call binding the contract method 0x27f146f3.
//
// Solidity: function getNextClientSeq() view returns(uint256)
func (_ContractICS26Router *ContractICS26RouterSession) GetNextClientSeq() (*big.Int, error) {
	return _ContractICS26Router.Contract.GetNextClientSeq(&_ContractICS26Router.CallOpts)
}

// GetNextClientSeq is a free data retrieval call binding the contract method 0x27f146f3.
//
// Solidity: function getNextClientSeq() view returns(uint256)
func (_ContractICS26Router *ContractICS26RouterCallerSession) GetNextClientSeq() (*big.Int, error) {
	return _ContractICS26Router.Contract.GetNextClientSeq(&_ContractICS26Router.CallOpts)
}

// IsConsumingScheduledOp is a free data retrieval call binding the contract method 0x8fb36037.
//
// Solidity: function isConsumingScheduledOp() view returns(bytes4)
func (_ContractICS26Router *ContractICS26RouterCaller) IsConsumingScheduledOp(opts *bind.CallOpts) ([4]byte, error) {
	var out []interface{}
	err := _ContractICS26Router.contract.Call(opts, &out, "isConsumingScheduledOp")

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// IsConsumingScheduledOp is a free data retrieval call binding the contract method 0x8fb36037.
//
// Solidity: function isConsumingScheduledOp() view returns(bytes4)
func (_ContractICS26Router *ContractICS26RouterSession) IsConsumingScheduledOp() ([4]byte, error) {
	return _ContractICS26Router.Contract.IsConsumingScheduledOp(&_ContractICS26Router.CallOpts)
}

// IsConsumingScheduledOp is a free data retrieval call binding the contract method 0x8fb36037.
//
// Solidity: function isConsumingScheduledOp() view returns(bytes4)
func (_ContractICS26Router *ContractICS26RouterCallerSession) IsConsumingScheduledOp() ([4]byte, error) {
	return _ContractICS26Router.Contract.IsConsumingScheduledOp(&_ContractICS26Router.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ContractICS26Router *ContractICS26RouterCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ContractICS26Router.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ContractICS26Router *ContractICS26RouterSession) ProxiableUUID() ([32]byte, error) {
	return _ContractICS26Router.Contract.ProxiableUUID(&_ContractICS26Router.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ContractICS26Router *ContractICS26RouterCallerSession) ProxiableUUID() ([32]byte, error) {
	return _ContractICS26Router.Contract.ProxiableUUID(&_ContractICS26Router.CallOpts)
}

// AckPacket is a paid mutator transaction binding the contract method 0x1bca011a.
//
// Solidity: function ackPacket(((uint64,string,string,uint64,(string,string,string,string,bytes)[]),bytes,bytes,(uint64,uint64)) msg_) returns()
func (_ContractICS26Router *ContractICS26RouterTransactor) AckPacket(opts *bind.TransactOpts, msg_ IICS26RouterMsgsMsgAckPacket) (*types.Transaction, error) {
	return _ContractICS26Router.contract.Transact(opts, "ackPacket", msg_)
}

// AckPacket is a paid mutator transaction binding the contract method 0x1bca011a.
//
// Solidity: function ackPacket(((uint64,string,string,uint64,(string,string,string,string,bytes)[]),bytes,bytes,(uint64,uint64)) msg_) returns()
func (_ContractICS26Router *ContractICS26RouterSession) AckPacket(msg_ IICS26RouterMsgsMsgAckPacket) (*types.Transaction, error) {
	return _ContractICS26Router.Contract.AckPacket(&_ContractICS26Router.TransactOpts, msg_)
}

// AckPacket is a paid mutator transaction binding the contract method 0x1bca011a.
//
// Solidity: function ackPacket(((uint64,string,string,uint64,(string,string,string,string,bytes)[]),bytes,bytes,(uint64,uint64)) msg_) returns()
func (_ContractICS26Router *ContractICS26RouterTransactorSession) AckPacket(msg_ IICS26RouterMsgsMsgAckPacket) (*types.Transaction, error) {
	return _ContractICS26Router.Contract.AckPacket(&_ContractICS26Router.TransactOpts, msg_)
}

// AddClient is a paid mutator transaction binding the contract method 0x1ec43e23.
//
// Solidity: function addClient(string clientId, (string,bytes[]) counterpartyInfo, address client) returns(string)
func (_ContractICS26Router *ContractICS26RouterTransactor) AddClient(opts *bind.TransactOpts, clientId string, counterpartyInfo IICS02ClientMsgsCounterpartyInfo, client common.Address) (*types.Transaction, error) {
	return _ContractICS26Router.contract.Transact(opts, "addClient", clientId, counterpartyInfo, client)
}

// AddClient is a paid mutator transaction binding the contract method 0x1ec43e23.
//
// Solidity: function addClient(string clientId, (string,bytes[]) counterpartyInfo, address client) returns(string)
func (_ContractICS26Router *ContractICS26RouterSession) AddClient(clientId string, counterpartyInfo IICS02ClientMsgsCounterpartyInfo, client common.Address) (*types.Transaction, error) {
	return _ContractICS26Router.Contract.AddClient(&_ContractICS26Router.TransactOpts, clientId, counterpartyInfo, client)
}

// AddClient is a paid mutator transaction binding the contract method 0x1ec43e23.
//
// Solidity: function addClient(string clientId, (string,bytes[]) counterpartyInfo, address client) returns(string)
func (_ContractICS26Router *ContractICS26RouterTransactorSession) AddClient(clientId string, counterpartyInfo IICS02ClientMsgsCounterpartyInfo, client common.Address) (*types.Transaction, error) {
	return _ContractICS26Router.Contract.AddClient(&_ContractICS26Router.TransactOpts, clientId, counterpartyInfo, client)
}

// AddClient0 is a paid mutator transaction binding the contract method 0xe3cb36a0.
//
// Solidity: function addClient((string,bytes[]) counterpartyInfo, address client) returns(string)
func (_ContractICS26Router *ContractICS26RouterTransactor) AddClient0(opts *bind.TransactOpts, counterpartyInfo IICS02ClientMsgsCounterpartyInfo, client common.Address) (*types.Transaction, error) {
	return _ContractICS26Router.contract.Transact(opts, "addClient0", counterpartyInfo, client)
}

// AddClient0 is a paid mutator transaction binding the contract method 0xe3cb36a0.
//
// Solidity: function addClient((string,bytes[]) counterpartyInfo, address client) returns(string)
func (_ContractICS26Router *ContractICS26RouterSession) AddClient0(counterpartyInfo IICS02ClientMsgsCounterpartyInfo, client common.Address) (*types.Transaction, error) {
	return _ContractICS26Router.Contract.AddClient0(&_ContractICS26Router.TransactOpts, counterpartyInfo, client)
}

// AddClient0 is a paid mutator transaction binding the contract method 0xe3cb36a0.
//
// Solidity: function addClient((string,bytes[]) counterpartyInfo, address client) returns(string)
func (_ContractICS26Router *ContractICS26RouterTransactorSession) AddClient0(counterpartyInfo IICS02ClientMsgsCounterpartyInfo, client common.Address) (*types.Transaction, error) {
	return _ContractICS26Router.Contract.AddClient0(&_ContractICS26Router.TransactOpts, counterpartyInfo, client)
}

// AddIBCApp is a paid mutator transaction binding the contract method 0x4b720d5b.
//
// Solidity: function addIBCApp(address app) returns()
func (_ContractICS26Router *ContractICS26RouterTransactor) AddIBCApp(opts *bind.TransactOpts, app common.Address) (*types.Transaction, error) {
	return _ContractICS26Router.contract.Transact(opts, "addIBCApp", app)
}

// AddIBCApp is a paid mutator transaction binding the contract method 0x4b720d5b.
//
// Solidity: function addIBCApp(address app) returns()
func (_ContractICS26Router *ContractICS26RouterSession) AddIBCApp(app common.Address) (*types.Transaction, error) {
	return _ContractICS26Router.Contract.AddIBCApp(&_ContractICS26Router.TransactOpts, app)
}

// AddIBCApp is a paid mutator transaction binding the contract method 0x4b720d5b.
//
// Solidity: function addIBCApp(address app) returns()
func (_ContractICS26Router *ContractICS26RouterTransactorSession) AddIBCApp(app common.Address) (*types.Transaction, error) {
	return _ContractICS26Router.Contract.AddIBCApp(&_ContractICS26Router.TransactOpts, app)
}

// AddIBCApp0 is a paid mutator transaction binding the contract method 0x5f516889.
//
// Solidity: function addIBCApp(string portId, address app) returns()
func (_ContractICS26Router *ContractICS26RouterTransactor) AddIBCApp0(opts *bind.TransactOpts, portId string, app common.Address) (*types.Transaction, error) {
	return _ContractICS26Router.contract.Transact(opts, "addIBCApp0", portId, app)
}

// AddIBCApp0 is a paid mutator transaction binding the contract method 0x5f516889.
//
// Solidity: function addIBCApp(string portId, address app) returns()
func (_ContractICS26Router *ContractICS26RouterSession) AddIBCApp0(portId string, app common.Address) (*types.Transaction, error) {
	return _ContractICS26Router.Contract.AddIBCApp0(&_ContractICS26Router.TransactOpts, portId, app)
}

// AddIBCApp0 is a paid mutator transaction binding the contract method 0x5f516889.
//
// Solidity: function addIBCApp(string portId, address app) returns()
func (_ContractICS26Router *ContractICS26RouterTransactorSession) AddIBCApp0(portId string, app common.Address) (*types.Transaction, error) {
	return _ContractICS26Router.Contract.AddIBCApp0(&_ContractICS26Router.TransactOpts, portId, app)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address authority) returns()
func (_ContractICS26Router *ContractICS26RouterTransactor) Initialize(opts *bind.TransactOpts, authority common.Address) (*types.Transaction, error) {
	return _ContractICS26Router.contract.Transact(opts, "initialize", authority)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address authority) returns()
func (_ContractICS26Router *ContractICS26RouterSession) Initialize(authority common.Address) (*types.Transaction, error) {
	return _ContractICS26Router.Contract.Initialize(&_ContractICS26Router.TransactOpts, authority)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address authority) returns()
func (_ContractICS26Router *ContractICS26RouterTransactorSession) Initialize(authority common.Address) (*types.Transaction, error) {
	return _ContractICS26Router.Contract.Initialize(&_ContractICS26Router.TransactOpts, authority)
}

// InitializeV2 is a paid mutator transaction binding the contract method 0x29b6eca9.
//
// Solidity: function initializeV2(address authority) returns()
func (_ContractICS26Router *ContractICS26RouterTransactor) InitializeV2(opts *bind.TransactOpts, authority common.Address) (*types.Transaction, error) {
	return _ContractICS26Router.contract.Transact(opts, "initializeV2", authority)
}

// InitializeV2 is a paid mutator transaction binding the contract method 0x29b6eca9.
//
// Solidity: function initializeV2(address authority) returns()
func (_ContractICS26Router *ContractICS26RouterSession) InitializeV2(authority common.Address) (*types.Transaction, error) {
	return _ContractICS26Router.Contract.InitializeV2(&_ContractICS26Router.TransactOpts, authority)
}

// InitializeV2 is a paid mutator transaction binding the contract method 0x29b6eca9.
//
// Solidity: function initializeV2(address authority) returns()
func (_ContractICS26Router *ContractICS26RouterTransactorSession) InitializeV2(authority common.Address) (*types.Transaction, error) {
	return _ContractICS26Router.Contract.InitializeV2(&_ContractICS26Router.TransactOpts, authority)
}

// MigrateClient is a paid mutator transaction binding the contract method 0xcce0b265.
//
// Solidity: function migrateClient(string clientId, (string,bytes[]) counterpartyInfo, address client) returns()
func (_ContractICS26Router *ContractICS26RouterTransactor) MigrateClient(opts *bind.TransactOpts, clientId string, counterpartyInfo IICS02ClientMsgsCounterpartyInfo, client common.Address) (*types.Transaction, error) {
	return _ContractICS26Router.contract.Transact(opts, "migrateClient", clientId, counterpartyInfo, client)
}

// MigrateClient is a paid mutator transaction binding the contract method 0xcce0b265.
//
// Solidity: function migrateClient(string clientId, (string,bytes[]) counterpartyInfo, address client) returns()
func (_ContractICS26Router *ContractICS26RouterSession) MigrateClient(clientId string, counterpartyInfo IICS02ClientMsgsCounterpartyInfo, client common.Address) (*types.Transaction, error) {
	return _ContractICS26Router.Contract.MigrateClient(&_ContractICS26Router.TransactOpts, clientId, counterpartyInfo, client)
}

// MigrateClient is a paid mutator transaction binding the contract method 0xcce0b265.
//
// Solidity: function migrateClient(string clientId, (string,bytes[]) counterpartyInfo, address client) returns()
func (_ContractICS26Router *ContractICS26RouterTransactorSession) MigrateClient(clientId string, counterpartyInfo IICS02ClientMsgsCounterpartyInfo, client common.Address) (*types.Transaction, error) {
	return _ContractICS26Router.Contract.MigrateClient(&_ContractICS26Router.TransactOpts, clientId, counterpartyInfo, client)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_ContractICS26Router *ContractICS26RouterTransactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _ContractICS26Router.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_ContractICS26Router *ContractICS26RouterSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _ContractICS26Router.Contract.Multicall(&_ContractICS26Router.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_ContractICS26Router *ContractICS26RouterTransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _ContractICS26Router.Contract.Multicall(&_ContractICS26Router.TransactOpts, data)
}

// RecvPacket is a paid mutator transaction binding the contract method 0x5ebd10ca.
//
// Solidity: function recvPacket(((uint64,string,string,uint64,(string,string,string,string,bytes)[]),bytes,(uint64,uint64)) msg_) returns()
func (_ContractICS26Router *ContractICS26RouterTransactor) RecvPacket(opts *bind.TransactOpts, msg_ IICS26RouterMsgsMsgRecvPacket) (*types.Transaction, error) {
	return _ContractICS26Router.contract.Transact(opts, "recvPacket", msg_)
}

// RecvPacket is a paid mutator transaction binding the contract method 0x5ebd10ca.
//
// Solidity: function recvPacket(((uint64,string,string,uint64,(string,string,string,string,bytes)[]),bytes,(uint64,uint64)) msg_) returns()
func (_ContractICS26Router *ContractICS26RouterSession) RecvPacket(msg_ IICS26RouterMsgsMsgRecvPacket) (*types.Transaction, error) {
	return _ContractICS26Router.Contract.RecvPacket(&_ContractICS26Router.TransactOpts, msg_)
}

// RecvPacket is a paid mutator transaction binding the contract method 0x5ebd10ca.
//
// Solidity: function recvPacket(((uint64,string,string,uint64,(string,string,string,string,bytes)[]),bytes,(uint64,uint64)) msg_) returns()
func (_ContractICS26Router *ContractICS26RouterTransactorSession) RecvPacket(msg_ IICS26RouterMsgsMsgRecvPacket) (*types.Transaction, error) {
	return _ContractICS26Router.Contract.RecvPacket(&_ContractICS26Router.TransactOpts, msg_)
}

// SendPacket is a paid mutator transaction binding the contract method 0x4d6e7ce3.
//
// Solidity: function sendPacket((string,uint64,(string,string,string,string,bytes)) msg_) returns(uint64)
func (_ContractICS26Router *ContractICS26RouterTransactor) SendPacket(opts *bind.TransactOpts, msg_ IICS26RouterMsgsMsgSendPacket) (*types.Transaction, error) {
	return _ContractICS26Router.contract.Transact(opts, "sendPacket", msg_)
}

// SendPacket is a paid mutator transaction binding the contract method 0x4d6e7ce3.
//
// Solidity: function sendPacket((string,uint64,(string,string,string,string,bytes)) msg_) returns(uint64)
func (_ContractICS26Router *ContractICS26RouterSession) SendPacket(msg_ IICS26RouterMsgsMsgSendPacket) (*types.Transaction, error) {
	return _ContractICS26Router.Contract.SendPacket(&_ContractICS26Router.TransactOpts, msg_)
}

// SendPacket is a paid mutator transaction binding the contract method 0x4d6e7ce3.
//
// Solidity: function sendPacket((string,uint64,(string,string,string,string,bytes)) msg_) returns(uint64)
func (_ContractICS26Router *ContractICS26RouterTransactorSession) SendPacket(msg_ IICS26RouterMsgsMsgSendPacket) (*types.Transaction, error) {
	return _ContractICS26Router.Contract.SendPacket(&_ContractICS26Router.TransactOpts, msg_)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address newAuthority) returns()
func (_ContractICS26Router *ContractICS26RouterTransactor) SetAuthority(opts *bind.TransactOpts, newAuthority common.Address) (*types.Transaction, error) {
	return _ContractICS26Router.contract.Transact(opts, "setAuthority", newAuthority)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address newAuthority) returns()
func (_ContractICS26Router *ContractICS26RouterSession) SetAuthority(newAuthority common.Address) (*types.Transaction, error) {
	return _ContractICS26Router.Contract.SetAuthority(&_ContractICS26Router.TransactOpts, newAuthority)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address newAuthority) returns()
func (_ContractICS26Router *ContractICS26RouterTransactorSession) SetAuthority(newAuthority common.Address) (*types.Transaction, error) {
	return _ContractICS26Router.Contract.SetAuthority(&_ContractICS26Router.TransactOpts, newAuthority)
}

// SubmitMisbehaviour is a paid mutator transaction binding the contract method 0x9e2e5c83.
//
// Solidity: function submitMisbehaviour(string clientId, bytes misbehaviourMsg) returns()
func (_ContractICS26Router *ContractICS26RouterTransactor) SubmitMisbehaviour(opts *bind.TransactOpts, clientId string, misbehaviourMsg []byte) (*types.Transaction, error) {
	return _ContractICS26Router.contract.Transact(opts, "submitMisbehaviour", clientId, misbehaviourMsg)
}

// SubmitMisbehaviour is a paid mutator transaction binding the contract method 0x9e2e5c83.
//
// Solidity: function submitMisbehaviour(string clientId, bytes misbehaviourMsg) returns()
func (_ContractICS26Router *ContractICS26RouterSession) SubmitMisbehaviour(clientId string, misbehaviourMsg []byte) (*types.Transaction, error) {
	return _ContractICS26Router.Contract.SubmitMisbehaviour(&_ContractICS26Router.TransactOpts, clientId, misbehaviourMsg)
}

// SubmitMisbehaviour is a paid mutator transaction binding the contract method 0x9e2e5c83.
//
// Solidity: function submitMisbehaviour(string clientId, bytes misbehaviourMsg) returns()
func (_ContractICS26Router *ContractICS26RouterTransactorSession) SubmitMisbehaviour(clientId string, misbehaviourMsg []byte) (*types.Transaction, error) {
	return _ContractICS26Router.Contract.SubmitMisbehaviour(&_ContractICS26Router.TransactOpts, clientId, misbehaviourMsg)
}

// TimeoutPacket is a paid mutator transaction binding the contract method 0xb98c330a.
//
// Solidity: function timeoutPacket(((uint64,string,string,uint64,(string,string,string,string,bytes)[]),bytes,(uint64,uint64)) msg_) returns()
func (_ContractICS26Router *ContractICS26RouterTransactor) TimeoutPacket(opts *bind.TransactOpts, msg_ IICS26RouterMsgsMsgTimeoutPacket) (*types.Transaction, error) {
	return _ContractICS26Router.contract.Transact(opts, "timeoutPacket", msg_)
}

// TimeoutPacket is a paid mutator transaction binding the contract method 0xb98c330a.
//
// Solidity: function timeoutPacket(((uint64,string,string,uint64,(string,string,string,string,bytes)[]),bytes,(uint64,uint64)) msg_) returns()
func (_ContractICS26Router *ContractICS26RouterSession) TimeoutPacket(msg_ IICS26RouterMsgsMsgTimeoutPacket) (*types.Transaction, error) {
	return _ContractICS26Router.Contract.TimeoutPacket(&_ContractICS26Router.TransactOpts, msg_)
}

// TimeoutPacket is a paid mutator transaction binding the contract method 0xb98c330a.
//
// Solidity: function timeoutPacket(((uint64,string,string,uint64,(string,string,string,string,bytes)[]),bytes,(uint64,uint64)) msg_) returns()
func (_ContractICS26Router *ContractICS26RouterTransactorSession) TimeoutPacket(msg_ IICS26RouterMsgsMsgTimeoutPacket) (*types.Transaction, error) {
	return _ContractICS26Router.Contract.TimeoutPacket(&_ContractICS26Router.TransactOpts, msg_)
}

// UpdateClient is a paid mutator transaction binding the contract method 0x6fbf8079.
//
// Solidity: function updateClient(string clientId, bytes updateMsg) returns(uint8)
func (_ContractICS26Router *ContractICS26RouterTransactor) UpdateClient(opts *bind.TransactOpts, clientId string, updateMsg []byte) (*types.Transaction, error) {
	return _ContractICS26Router.contract.Transact(opts, "updateClient", clientId, updateMsg)
}

// UpdateClient is a paid mutator transaction binding the contract method 0x6fbf8079.
//
// Solidity: function updateClient(string clientId, bytes updateMsg) returns(uint8)
func (_ContractICS26Router *ContractICS26RouterSession) UpdateClient(clientId string, updateMsg []byte) (*types.Transaction, error) {
	return _ContractICS26Router.Contract.UpdateClient(&_ContractICS26Router.TransactOpts, clientId, updateMsg)
}

// UpdateClient is a paid mutator transaction binding the contract method 0x6fbf8079.
//
// Solidity: function updateClient(string clientId, bytes updateMsg) returns(uint8)
func (_ContractICS26Router *ContractICS26RouterTransactorSession) UpdateClient(clientId string, updateMsg []byte) (*types.Transaction, error) {
	return _ContractICS26Router.Contract.UpdateClient(&_ContractICS26Router.TransactOpts, clientId, updateMsg)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ContractICS26Router *ContractICS26RouterTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ContractICS26Router.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ContractICS26Router *ContractICS26RouterSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ContractICS26Router.Contract.UpgradeToAndCall(&_ContractICS26Router.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ContractICS26Router *ContractICS26RouterTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ContractICS26Router.Contract.UpgradeToAndCall(&_ContractICS26Router.TransactOpts, newImplementation, data)
}

// ContractICS26RouterAckPacketIterator is returned from FilterAckPacket and is used to iterate over the raw logs and unpacked data for AckPacket events raised by the ContractICS26Router contract.
type ContractICS26RouterAckPacketIterator struct {
	Event *ContractICS26RouterAckPacket // Event containing the contract specifics and raw log

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
func (it *ContractICS26RouterAckPacketIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractICS26RouterAckPacket)
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
		it.Event = new(ContractICS26RouterAckPacket)
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
func (it *ContractICS26RouterAckPacketIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractICS26RouterAckPacketIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractICS26RouterAckPacket represents a AckPacket event raised by the ContractICS26Router contract.
type ContractICS26RouterAckPacket struct {
	ClientId        common.Hash
	Sequence        *big.Int
	Packet          IICS26RouterMsgsPacket
	Acknowledgement []byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterAckPacket is a free log retrieval operation binding the contract event 0xf9bab74bcdb634f4d3dd064cc42a13df056598e1c0336905d2f5750fbfb08b7b.
//
// Solidity: event AckPacket(string indexed clientId, uint256 indexed sequence, (uint64,string,string,uint64,(string,string,string,string,bytes)[]) packet, bytes acknowledgement)
func (_ContractICS26Router *ContractICS26RouterFilterer) FilterAckPacket(opts *bind.FilterOpts, clientId []string, sequence []*big.Int) (*ContractICS26RouterAckPacketIterator, error) {

	var clientIdRule []interface{}
	for _, clientIdItem := range clientId {
		clientIdRule = append(clientIdRule, clientIdItem)
	}
	var sequenceRule []interface{}
	for _, sequenceItem := range sequence {
		sequenceRule = append(sequenceRule, sequenceItem)
	}

	logs, sub, err := _ContractICS26Router.contract.FilterLogs(opts, "AckPacket", clientIdRule, sequenceRule)
	if err != nil {
		return nil, err
	}
	return &ContractICS26RouterAckPacketIterator{contract: _ContractICS26Router.contract, event: "AckPacket", logs: logs, sub: sub}, nil
}

// WatchAckPacket is a free log subscription operation binding the contract event 0xf9bab74bcdb634f4d3dd064cc42a13df056598e1c0336905d2f5750fbfb08b7b.
//
// Solidity: event AckPacket(string indexed clientId, uint256 indexed sequence, (uint64,string,string,uint64,(string,string,string,string,bytes)[]) packet, bytes acknowledgement)
func (_ContractICS26Router *ContractICS26RouterFilterer) WatchAckPacket(opts *bind.WatchOpts, sink chan<- *ContractICS26RouterAckPacket, clientId []string, sequence []*big.Int) (event.Subscription, error) {

	var clientIdRule []interface{}
	for _, clientIdItem := range clientId {
		clientIdRule = append(clientIdRule, clientIdItem)
	}
	var sequenceRule []interface{}
	for _, sequenceItem := range sequence {
		sequenceRule = append(sequenceRule, sequenceItem)
	}

	logs, sub, err := _ContractICS26Router.contract.WatchLogs(opts, "AckPacket", clientIdRule, sequenceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractICS26RouterAckPacket)
				if err := _ContractICS26Router.contract.UnpackLog(event, "AckPacket", log); err != nil {
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

// ParseAckPacket is a log parse operation binding the contract event 0xf9bab74bcdb634f4d3dd064cc42a13df056598e1c0336905d2f5750fbfb08b7b.
//
// Solidity: event AckPacket(string indexed clientId, uint256 indexed sequence, (uint64,string,string,uint64,(string,string,string,string,bytes)[]) packet, bytes acknowledgement)
func (_ContractICS26Router *ContractICS26RouterFilterer) ParseAckPacket(log types.Log) (*ContractICS26RouterAckPacket, error) {
	event := new(ContractICS26RouterAckPacket)
	if err := _ContractICS26Router.contract.UnpackLog(event, "AckPacket", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractICS26RouterAuthorityUpdatedIterator is returned from FilterAuthorityUpdated and is used to iterate over the raw logs and unpacked data for AuthorityUpdated events raised by the ContractICS26Router contract.
type ContractICS26RouterAuthorityUpdatedIterator struct {
	Event *ContractICS26RouterAuthorityUpdated // Event containing the contract specifics and raw log

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
func (it *ContractICS26RouterAuthorityUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractICS26RouterAuthorityUpdated)
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
		it.Event = new(ContractICS26RouterAuthorityUpdated)
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
func (it *ContractICS26RouterAuthorityUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractICS26RouterAuthorityUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractICS26RouterAuthorityUpdated represents a AuthorityUpdated event raised by the ContractICS26Router contract.
type ContractICS26RouterAuthorityUpdated struct {
	Authority common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAuthorityUpdated is a free log retrieval operation binding the contract event 0x2f658b440c35314f52658ea8a740e05b284cdc84dc9ae01e891f21b8933e7cad.
//
// Solidity: event AuthorityUpdated(address authority)
func (_ContractICS26Router *ContractICS26RouterFilterer) FilterAuthorityUpdated(opts *bind.FilterOpts) (*ContractICS26RouterAuthorityUpdatedIterator, error) {

	logs, sub, err := _ContractICS26Router.contract.FilterLogs(opts, "AuthorityUpdated")
	if err != nil {
		return nil, err
	}
	return &ContractICS26RouterAuthorityUpdatedIterator{contract: _ContractICS26Router.contract, event: "AuthorityUpdated", logs: logs, sub: sub}, nil
}

// WatchAuthorityUpdated is a free log subscription operation binding the contract event 0x2f658b440c35314f52658ea8a740e05b284cdc84dc9ae01e891f21b8933e7cad.
//
// Solidity: event AuthorityUpdated(address authority)
func (_ContractICS26Router *ContractICS26RouterFilterer) WatchAuthorityUpdated(opts *bind.WatchOpts, sink chan<- *ContractICS26RouterAuthorityUpdated) (event.Subscription, error) {

	logs, sub, err := _ContractICS26Router.contract.WatchLogs(opts, "AuthorityUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractICS26RouterAuthorityUpdated)
				if err := _ContractICS26Router.contract.UnpackLog(event, "AuthorityUpdated", log); err != nil {
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

// ParseAuthorityUpdated is a log parse operation binding the contract event 0x2f658b440c35314f52658ea8a740e05b284cdc84dc9ae01e891f21b8933e7cad.
//
// Solidity: event AuthorityUpdated(address authority)
func (_ContractICS26Router *ContractICS26RouterFilterer) ParseAuthorityUpdated(log types.Log) (*ContractICS26RouterAuthorityUpdated, error) {
	event := new(ContractICS26RouterAuthorityUpdated)
	if err := _ContractICS26Router.contract.UnpackLog(event, "AuthorityUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractICS26RouterIBCAppAddedIterator is returned from FilterIBCAppAdded and is used to iterate over the raw logs and unpacked data for IBCAppAdded events raised by the ContractICS26Router contract.
type ContractICS26RouterIBCAppAddedIterator struct {
	Event *ContractICS26RouterIBCAppAdded // Event containing the contract specifics and raw log

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
func (it *ContractICS26RouterIBCAppAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractICS26RouterIBCAppAdded)
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
		it.Event = new(ContractICS26RouterIBCAppAdded)
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
func (it *ContractICS26RouterIBCAppAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractICS26RouterIBCAppAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractICS26RouterIBCAppAdded represents a IBCAppAdded event raised by the ContractICS26Router contract.
type ContractICS26RouterIBCAppAdded struct {
	PortId string
	App    common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterIBCAppAdded is a free log retrieval operation binding the contract event 0xa6ec8e860960e638347460dc632fbe0175c51a5ca130e336138bbe26ff304499.
//
// Solidity: event IBCAppAdded(string portId, address app)
func (_ContractICS26Router *ContractICS26RouterFilterer) FilterIBCAppAdded(opts *bind.FilterOpts) (*ContractICS26RouterIBCAppAddedIterator, error) {

	logs, sub, err := _ContractICS26Router.contract.FilterLogs(opts, "IBCAppAdded")
	if err != nil {
		return nil, err
	}
	return &ContractICS26RouterIBCAppAddedIterator{contract: _ContractICS26Router.contract, event: "IBCAppAdded", logs: logs, sub: sub}, nil
}

// WatchIBCAppAdded is a free log subscription operation binding the contract event 0xa6ec8e860960e638347460dc632fbe0175c51a5ca130e336138bbe26ff304499.
//
// Solidity: event IBCAppAdded(string portId, address app)
func (_ContractICS26Router *ContractICS26RouterFilterer) WatchIBCAppAdded(opts *bind.WatchOpts, sink chan<- *ContractICS26RouterIBCAppAdded) (event.Subscription, error) {

	logs, sub, err := _ContractICS26Router.contract.WatchLogs(opts, "IBCAppAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractICS26RouterIBCAppAdded)
				if err := _ContractICS26Router.contract.UnpackLog(event, "IBCAppAdded", log); err != nil {
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

// ParseIBCAppAdded is a log parse operation binding the contract event 0xa6ec8e860960e638347460dc632fbe0175c51a5ca130e336138bbe26ff304499.
//
// Solidity: event IBCAppAdded(string portId, address app)
func (_ContractICS26Router *ContractICS26RouterFilterer) ParseIBCAppAdded(log types.Log) (*ContractICS26RouterIBCAppAdded, error) {
	event := new(ContractICS26RouterIBCAppAdded)
	if err := _ContractICS26Router.contract.UnpackLog(event, "IBCAppAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractICS26RouterIBCAppRecvPacketCallbackErrorIterator is returned from FilterIBCAppRecvPacketCallbackError and is used to iterate over the raw logs and unpacked data for IBCAppRecvPacketCallbackError events raised by the ContractICS26Router contract.
type ContractICS26RouterIBCAppRecvPacketCallbackErrorIterator struct {
	Event *ContractICS26RouterIBCAppRecvPacketCallbackError // Event containing the contract specifics and raw log

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
func (it *ContractICS26RouterIBCAppRecvPacketCallbackErrorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractICS26RouterIBCAppRecvPacketCallbackError)
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
		it.Event = new(ContractICS26RouterIBCAppRecvPacketCallbackError)
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
func (it *ContractICS26RouterIBCAppRecvPacketCallbackErrorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractICS26RouterIBCAppRecvPacketCallbackErrorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractICS26RouterIBCAppRecvPacketCallbackError represents a IBCAppRecvPacketCallbackError event raised by the ContractICS26Router contract.
type ContractICS26RouterIBCAppRecvPacketCallbackError struct {
	Reason []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterIBCAppRecvPacketCallbackError is a free log retrieval operation binding the contract event 0xb9edb487876e8be10f54e377c1a815a54ad92a6db1c9561dfe8fad2f0d1da84f.
//
// Solidity: event IBCAppRecvPacketCallbackError(bytes reason)
func (_ContractICS26Router *ContractICS26RouterFilterer) FilterIBCAppRecvPacketCallbackError(opts *bind.FilterOpts) (*ContractICS26RouterIBCAppRecvPacketCallbackErrorIterator, error) {

	logs, sub, err := _ContractICS26Router.contract.FilterLogs(opts, "IBCAppRecvPacketCallbackError")
	if err != nil {
		return nil, err
	}
	return &ContractICS26RouterIBCAppRecvPacketCallbackErrorIterator{contract: _ContractICS26Router.contract, event: "IBCAppRecvPacketCallbackError", logs: logs, sub: sub}, nil
}

// WatchIBCAppRecvPacketCallbackError is a free log subscription operation binding the contract event 0xb9edb487876e8be10f54e377c1a815a54ad92a6db1c9561dfe8fad2f0d1da84f.
//
// Solidity: event IBCAppRecvPacketCallbackError(bytes reason)
func (_ContractICS26Router *ContractICS26RouterFilterer) WatchIBCAppRecvPacketCallbackError(opts *bind.WatchOpts, sink chan<- *ContractICS26RouterIBCAppRecvPacketCallbackError) (event.Subscription, error) {

	logs, sub, err := _ContractICS26Router.contract.WatchLogs(opts, "IBCAppRecvPacketCallbackError")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractICS26RouterIBCAppRecvPacketCallbackError)
				if err := _ContractICS26Router.contract.UnpackLog(event, "IBCAppRecvPacketCallbackError", log); err != nil {
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

// ParseIBCAppRecvPacketCallbackError is a log parse operation binding the contract event 0xb9edb487876e8be10f54e377c1a815a54ad92a6db1c9561dfe8fad2f0d1da84f.
//
// Solidity: event IBCAppRecvPacketCallbackError(bytes reason)
func (_ContractICS26Router *ContractICS26RouterFilterer) ParseIBCAppRecvPacketCallbackError(log types.Log) (*ContractICS26RouterIBCAppRecvPacketCallbackError, error) {
	event := new(ContractICS26RouterIBCAppRecvPacketCallbackError)
	if err := _ContractICS26Router.contract.UnpackLog(event, "IBCAppRecvPacketCallbackError", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractICS26RouterICS02ClientAddedIterator is returned from FilterICS02ClientAdded and is used to iterate over the raw logs and unpacked data for ICS02ClientAdded events raised by the ContractICS26Router contract.
type ContractICS26RouterICS02ClientAddedIterator struct {
	Event *ContractICS26RouterICS02ClientAdded // Event containing the contract specifics and raw log

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
func (it *ContractICS26RouterICS02ClientAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractICS26RouterICS02ClientAdded)
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
		it.Event = new(ContractICS26RouterICS02ClientAdded)
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
func (it *ContractICS26RouterICS02ClientAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractICS26RouterICS02ClientAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractICS26RouterICS02ClientAdded represents a ICS02ClientAdded event raised by the ContractICS26Router contract.
type ContractICS26RouterICS02ClientAdded struct {
	ClientId         string
	CounterpartyInfo IICS02ClientMsgsCounterpartyInfo
	Client           common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterICS02ClientAdded is a free log retrieval operation binding the contract event 0x0ecded31ecd211a73abf0fb3bc09150bbe321a05550fbe29ea0f16b6e25fbfa8.
//
// Solidity: event ICS02ClientAdded(string clientId, (string,bytes[]) counterpartyInfo, address client)
func (_ContractICS26Router *ContractICS26RouterFilterer) FilterICS02ClientAdded(opts *bind.FilterOpts) (*ContractICS26RouterICS02ClientAddedIterator, error) {

	logs, sub, err := _ContractICS26Router.contract.FilterLogs(opts, "ICS02ClientAdded")
	if err != nil {
		return nil, err
	}
	return &ContractICS26RouterICS02ClientAddedIterator{contract: _ContractICS26Router.contract, event: "ICS02ClientAdded", logs: logs, sub: sub}, nil
}

// WatchICS02ClientAdded is a free log subscription operation binding the contract event 0x0ecded31ecd211a73abf0fb3bc09150bbe321a05550fbe29ea0f16b6e25fbfa8.
//
// Solidity: event ICS02ClientAdded(string clientId, (string,bytes[]) counterpartyInfo, address client)
func (_ContractICS26Router *ContractICS26RouterFilterer) WatchICS02ClientAdded(opts *bind.WatchOpts, sink chan<- *ContractICS26RouterICS02ClientAdded) (event.Subscription, error) {

	logs, sub, err := _ContractICS26Router.contract.WatchLogs(opts, "ICS02ClientAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractICS26RouterICS02ClientAdded)
				if err := _ContractICS26Router.contract.UnpackLog(event, "ICS02ClientAdded", log); err != nil {
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

// ParseICS02ClientAdded is a log parse operation binding the contract event 0x0ecded31ecd211a73abf0fb3bc09150bbe321a05550fbe29ea0f16b6e25fbfa8.
//
// Solidity: event ICS02ClientAdded(string clientId, (string,bytes[]) counterpartyInfo, address client)
func (_ContractICS26Router *ContractICS26RouterFilterer) ParseICS02ClientAdded(log types.Log) (*ContractICS26RouterICS02ClientAdded, error) {
	event := new(ContractICS26RouterICS02ClientAdded)
	if err := _ContractICS26Router.contract.UnpackLog(event, "ICS02ClientAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractICS26RouterICS02ClientMigratedIterator is returned from FilterICS02ClientMigrated and is used to iterate over the raw logs and unpacked data for ICS02ClientMigrated events raised by the ContractICS26Router contract.
type ContractICS26RouterICS02ClientMigratedIterator struct {
	Event *ContractICS26RouterICS02ClientMigrated // Event containing the contract specifics and raw log

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
func (it *ContractICS26RouterICS02ClientMigratedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractICS26RouterICS02ClientMigrated)
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
		it.Event = new(ContractICS26RouterICS02ClientMigrated)
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
func (it *ContractICS26RouterICS02ClientMigratedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractICS26RouterICS02ClientMigratedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractICS26RouterICS02ClientMigrated represents a ICS02ClientMigrated event raised by the ContractICS26Router contract.
type ContractICS26RouterICS02ClientMigrated struct {
	ClientId         string
	CounterpartyInfo IICS02ClientMsgsCounterpartyInfo
	Client           common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterICS02ClientMigrated is a free log retrieval operation binding the contract event 0x23c2e29d6ae84e79fa116b8afd6e28ddc1de7f473d3edb407fbd08093c3ed6bf.
//
// Solidity: event ICS02ClientMigrated(string clientId, (string,bytes[]) counterpartyInfo, address client)
func (_ContractICS26Router *ContractICS26RouterFilterer) FilterICS02ClientMigrated(opts *bind.FilterOpts) (*ContractICS26RouterICS02ClientMigratedIterator, error) {

	logs, sub, err := _ContractICS26Router.contract.FilterLogs(opts, "ICS02ClientMigrated")
	if err != nil {
		return nil, err
	}
	return &ContractICS26RouterICS02ClientMigratedIterator{contract: _ContractICS26Router.contract, event: "ICS02ClientMigrated", logs: logs, sub: sub}, nil
}

// WatchICS02ClientMigrated is a free log subscription operation binding the contract event 0x23c2e29d6ae84e79fa116b8afd6e28ddc1de7f473d3edb407fbd08093c3ed6bf.
//
// Solidity: event ICS02ClientMigrated(string clientId, (string,bytes[]) counterpartyInfo, address client)
func (_ContractICS26Router *ContractICS26RouterFilterer) WatchICS02ClientMigrated(opts *bind.WatchOpts, sink chan<- *ContractICS26RouterICS02ClientMigrated) (event.Subscription, error) {

	logs, sub, err := _ContractICS26Router.contract.WatchLogs(opts, "ICS02ClientMigrated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractICS26RouterICS02ClientMigrated)
				if err := _ContractICS26Router.contract.UnpackLog(event, "ICS02ClientMigrated", log); err != nil {
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

// ParseICS02ClientMigrated is a log parse operation binding the contract event 0x23c2e29d6ae84e79fa116b8afd6e28ddc1de7f473d3edb407fbd08093c3ed6bf.
//
// Solidity: event ICS02ClientMigrated(string clientId, (string,bytes[]) counterpartyInfo, address client)
func (_ContractICS26Router *ContractICS26RouterFilterer) ParseICS02ClientMigrated(log types.Log) (*ContractICS26RouterICS02ClientMigrated, error) {
	event := new(ContractICS26RouterICS02ClientMigrated)
	if err := _ContractICS26Router.contract.UnpackLog(event, "ICS02ClientMigrated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractICS26RouterICS02ClientUpdatedIterator is returned from FilterICS02ClientUpdated and is used to iterate over the raw logs and unpacked data for ICS02ClientUpdated events raised by the ContractICS26Router contract.
type ContractICS26RouterICS02ClientUpdatedIterator struct {
	Event *ContractICS26RouterICS02ClientUpdated // Event containing the contract specifics and raw log

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
func (it *ContractICS26RouterICS02ClientUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractICS26RouterICS02ClientUpdated)
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
		it.Event = new(ContractICS26RouterICS02ClientUpdated)
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
func (it *ContractICS26RouterICS02ClientUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractICS26RouterICS02ClientUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractICS26RouterICS02ClientUpdated represents a ICS02ClientUpdated event raised by the ContractICS26Router contract.
type ContractICS26RouterICS02ClientUpdated struct {
	ClientId string
	Result   uint8
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterICS02ClientUpdated is a free log retrieval operation binding the contract event 0x87bbef2779889a19f0435ddca81fda94132c06ffddb0ea73def256307a293aef.
//
// Solidity: event ICS02ClientUpdated(string clientId, uint8 result)
func (_ContractICS26Router *ContractICS26RouterFilterer) FilterICS02ClientUpdated(opts *bind.FilterOpts) (*ContractICS26RouterICS02ClientUpdatedIterator, error) {

	logs, sub, err := _ContractICS26Router.contract.FilterLogs(opts, "ICS02ClientUpdated")
	if err != nil {
		return nil, err
	}
	return &ContractICS26RouterICS02ClientUpdatedIterator{contract: _ContractICS26Router.contract, event: "ICS02ClientUpdated", logs: logs, sub: sub}, nil
}

// WatchICS02ClientUpdated is a free log subscription operation binding the contract event 0x87bbef2779889a19f0435ddca81fda94132c06ffddb0ea73def256307a293aef.
//
// Solidity: event ICS02ClientUpdated(string clientId, uint8 result)
func (_ContractICS26Router *ContractICS26RouterFilterer) WatchICS02ClientUpdated(opts *bind.WatchOpts, sink chan<- *ContractICS26RouterICS02ClientUpdated) (event.Subscription, error) {

	logs, sub, err := _ContractICS26Router.contract.WatchLogs(opts, "ICS02ClientUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractICS26RouterICS02ClientUpdated)
				if err := _ContractICS26Router.contract.UnpackLog(event, "ICS02ClientUpdated", log); err != nil {
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

// ParseICS02ClientUpdated is a log parse operation binding the contract event 0x87bbef2779889a19f0435ddca81fda94132c06ffddb0ea73def256307a293aef.
//
// Solidity: event ICS02ClientUpdated(string clientId, uint8 result)
func (_ContractICS26Router *ContractICS26RouterFilterer) ParseICS02ClientUpdated(log types.Log) (*ContractICS26RouterICS02ClientUpdated, error) {
	event := new(ContractICS26RouterICS02ClientUpdated)
	if err := _ContractICS26Router.contract.UnpackLog(event, "ICS02ClientUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractICS26RouterICS02MisbehaviourSubmittedIterator is returned from FilterICS02MisbehaviourSubmitted and is used to iterate over the raw logs and unpacked data for ICS02MisbehaviourSubmitted events raised by the ContractICS26Router contract.
type ContractICS26RouterICS02MisbehaviourSubmittedIterator struct {
	Event *ContractICS26RouterICS02MisbehaviourSubmitted // Event containing the contract specifics and raw log

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
func (it *ContractICS26RouterICS02MisbehaviourSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractICS26RouterICS02MisbehaviourSubmitted)
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
		it.Event = new(ContractICS26RouterICS02MisbehaviourSubmitted)
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
func (it *ContractICS26RouterICS02MisbehaviourSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractICS26RouterICS02MisbehaviourSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractICS26RouterICS02MisbehaviourSubmitted represents a ICS02MisbehaviourSubmitted event raised by the ContractICS26Router contract.
type ContractICS26RouterICS02MisbehaviourSubmitted struct {
	ClientId string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterICS02MisbehaviourSubmitted is a free log retrieval operation binding the contract event 0xa263f0a976b2937a51fd2e416491cf0ca724d5499fa870715929dfde4ee4a430.
//
// Solidity: event ICS02MisbehaviourSubmitted(string clientId)
func (_ContractICS26Router *ContractICS26RouterFilterer) FilterICS02MisbehaviourSubmitted(opts *bind.FilterOpts) (*ContractICS26RouterICS02MisbehaviourSubmittedIterator, error) {

	logs, sub, err := _ContractICS26Router.contract.FilterLogs(opts, "ICS02MisbehaviourSubmitted")
	if err != nil {
		return nil, err
	}
	return &ContractICS26RouterICS02MisbehaviourSubmittedIterator{contract: _ContractICS26Router.contract, event: "ICS02MisbehaviourSubmitted", logs: logs, sub: sub}, nil
}

// WatchICS02MisbehaviourSubmitted is a free log subscription operation binding the contract event 0xa263f0a976b2937a51fd2e416491cf0ca724d5499fa870715929dfde4ee4a430.
//
// Solidity: event ICS02MisbehaviourSubmitted(string clientId)
func (_ContractICS26Router *ContractICS26RouterFilterer) WatchICS02MisbehaviourSubmitted(opts *bind.WatchOpts, sink chan<- *ContractICS26RouterICS02MisbehaviourSubmitted) (event.Subscription, error) {

	logs, sub, err := _ContractICS26Router.contract.WatchLogs(opts, "ICS02MisbehaviourSubmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractICS26RouterICS02MisbehaviourSubmitted)
				if err := _ContractICS26Router.contract.UnpackLog(event, "ICS02MisbehaviourSubmitted", log); err != nil {
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

// ParseICS02MisbehaviourSubmitted is a log parse operation binding the contract event 0xa263f0a976b2937a51fd2e416491cf0ca724d5499fa870715929dfde4ee4a430.
//
// Solidity: event ICS02MisbehaviourSubmitted(string clientId)
func (_ContractICS26Router *ContractICS26RouterFilterer) ParseICS02MisbehaviourSubmitted(log types.Log) (*ContractICS26RouterICS02MisbehaviourSubmitted, error) {
	event := new(ContractICS26RouterICS02MisbehaviourSubmitted)
	if err := _ContractICS26Router.contract.UnpackLog(event, "ICS02MisbehaviourSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractICS26RouterInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContractICS26Router contract.
type ContractICS26RouterInitializedIterator struct {
	Event *ContractICS26RouterInitialized // Event containing the contract specifics and raw log

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
func (it *ContractICS26RouterInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractICS26RouterInitialized)
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
		it.Event = new(ContractICS26RouterInitialized)
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
func (it *ContractICS26RouterInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractICS26RouterInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractICS26RouterInitialized represents a Initialized event raised by the ContractICS26Router contract.
type ContractICS26RouterInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ContractICS26Router *ContractICS26RouterFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractICS26RouterInitializedIterator, error) {

	logs, sub, err := _ContractICS26Router.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractICS26RouterInitializedIterator{contract: _ContractICS26Router.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ContractICS26Router *ContractICS26RouterFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractICS26RouterInitialized) (event.Subscription, error) {

	logs, sub, err := _ContractICS26Router.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractICS26RouterInitialized)
				if err := _ContractICS26Router.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ContractICS26Router *ContractICS26RouterFilterer) ParseInitialized(log types.Log) (*ContractICS26RouterInitialized, error) {
	event := new(ContractICS26RouterInitialized)
	if err := _ContractICS26Router.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractICS26RouterNoopIterator is returned from FilterNoop and is used to iterate over the raw logs and unpacked data for Noop events raised by the ContractICS26Router contract.
type ContractICS26RouterNoopIterator struct {
	Event *ContractICS26RouterNoop // Event containing the contract specifics and raw log

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
func (it *ContractICS26RouterNoopIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractICS26RouterNoop)
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
		it.Event = new(ContractICS26RouterNoop)
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
func (it *ContractICS26RouterNoopIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractICS26RouterNoopIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractICS26RouterNoop represents a Noop event raised by the ContractICS26Router contract.
type ContractICS26RouterNoop struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterNoop is a free log retrieval operation binding the contract event 0xd08bf58b0e4eec5bfc697a4fdbb6839057fbf4dd06f1b1ce07445c0e5a654caf.
//
// Solidity: event Noop()
func (_ContractICS26Router *ContractICS26RouterFilterer) FilterNoop(opts *bind.FilterOpts) (*ContractICS26RouterNoopIterator, error) {

	logs, sub, err := _ContractICS26Router.contract.FilterLogs(opts, "Noop")
	if err != nil {
		return nil, err
	}
	return &ContractICS26RouterNoopIterator{contract: _ContractICS26Router.contract, event: "Noop", logs: logs, sub: sub}, nil
}

// WatchNoop is a free log subscription operation binding the contract event 0xd08bf58b0e4eec5bfc697a4fdbb6839057fbf4dd06f1b1ce07445c0e5a654caf.
//
// Solidity: event Noop()
func (_ContractICS26Router *ContractICS26RouterFilterer) WatchNoop(opts *bind.WatchOpts, sink chan<- *ContractICS26RouterNoop) (event.Subscription, error) {

	logs, sub, err := _ContractICS26Router.contract.WatchLogs(opts, "Noop")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractICS26RouterNoop)
				if err := _ContractICS26Router.contract.UnpackLog(event, "Noop", log); err != nil {
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

// ParseNoop is a log parse operation binding the contract event 0xd08bf58b0e4eec5bfc697a4fdbb6839057fbf4dd06f1b1ce07445c0e5a654caf.
//
// Solidity: event Noop()
func (_ContractICS26Router *ContractICS26RouterFilterer) ParseNoop(log types.Log) (*ContractICS26RouterNoop, error) {
	event := new(ContractICS26RouterNoop)
	if err := _ContractICS26Router.contract.UnpackLog(event, "Noop", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractICS26RouterSendPacketIterator is returned from FilterSendPacket and is used to iterate over the raw logs and unpacked data for SendPacket events raised by the ContractICS26Router contract.
type ContractICS26RouterSendPacketIterator struct {
	Event *ContractICS26RouterSendPacket // Event containing the contract specifics and raw log

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
func (it *ContractICS26RouterSendPacketIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractICS26RouterSendPacket)
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
		it.Event = new(ContractICS26RouterSendPacket)
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
func (it *ContractICS26RouterSendPacketIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractICS26RouterSendPacketIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractICS26RouterSendPacket represents a SendPacket event raised by the ContractICS26Router contract.
type ContractICS26RouterSendPacket struct {
	ClientId common.Hash
	Sequence *big.Int
	Packet   IICS26RouterMsgsPacket
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSendPacket is a free log retrieval operation binding the contract event 0xab3a4458a269be61dfa43faa33aa7b1f5d570716f83ad078bc2ba5dab039abae.
//
// Solidity: event SendPacket(string indexed clientId, uint256 indexed sequence, (uint64,string,string,uint64,(string,string,string,string,bytes)[]) packet)
func (_ContractICS26Router *ContractICS26RouterFilterer) FilterSendPacket(opts *bind.FilterOpts, clientId []string, sequence []*big.Int) (*ContractICS26RouterSendPacketIterator, error) {

	var clientIdRule []interface{}
	for _, clientIdItem := range clientId {
		clientIdRule = append(clientIdRule, clientIdItem)
	}
	var sequenceRule []interface{}
	for _, sequenceItem := range sequence {
		sequenceRule = append(sequenceRule, sequenceItem)
	}

	logs, sub, err := _ContractICS26Router.contract.FilterLogs(opts, "SendPacket", clientIdRule, sequenceRule)
	if err != nil {
		return nil, err
	}
	return &ContractICS26RouterSendPacketIterator{contract: _ContractICS26Router.contract, event: "SendPacket", logs: logs, sub: sub}, nil
}

// WatchSendPacket is a free log subscription operation binding the contract event 0xab3a4458a269be61dfa43faa33aa7b1f5d570716f83ad078bc2ba5dab039abae.
//
// Solidity: event SendPacket(string indexed clientId, uint256 indexed sequence, (uint64,string,string,uint64,(string,string,string,string,bytes)[]) packet)
func (_ContractICS26Router *ContractICS26RouterFilterer) WatchSendPacket(opts *bind.WatchOpts, sink chan<- *ContractICS26RouterSendPacket, clientId []string, sequence []*big.Int) (event.Subscription, error) {

	var clientIdRule []interface{}
	for _, clientIdItem := range clientId {
		clientIdRule = append(clientIdRule, clientIdItem)
	}
	var sequenceRule []interface{}
	for _, sequenceItem := range sequence {
		sequenceRule = append(sequenceRule, sequenceItem)
	}

	logs, sub, err := _ContractICS26Router.contract.WatchLogs(opts, "SendPacket", clientIdRule, sequenceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractICS26RouterSendPacket)
				if err := _ContractICS26Router.contract.UnpackLog(event, "SendPacket", log); err != nil {
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

// ParseSendPacket is a log parse operation binding the contract event 0xab3a4458a269be61dfa43faa33aa7b1f5d570716f83ad078bc2ba5dab039abae.
//
// Solidity: event SendPacket(string indexed clientId, uint256 indexed sequence, (uint64,string,string,uint64,(string,string,string,string,bytes)[]) packet)
func (_ContractICS26Router *ContractICS26RouterFilterer) ParseSendPacket(log types.Log) (*ContractICS26RouterSendPacket, error) {
	event := new(ContractICS26RouterSendPacket)
	if err := _ContractICS26Router.contract.UnpackLog(event, "SendPacket", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractICS26RouterTimeoutPacketIterator is returned from FilterTimeoutPacket and is used to iterate over the raw logs and unpacked data for TimeoutPacket events raised by the ContractICS26Router contract.
type ContractICS26RouterTimeoutPacketIterator struct {
	Event *ContractICS26RouterTimeoutPacket // Event containing the contract specifics and raw log

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
func (it *ContractICS26RouterTimeoutPacketIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractICS26RouterTimeoutPacket)
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
		it.Event = new(ContractICS26RouterTimeoutPacket)
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
func (it *ContractICS26RouterTimeoutPacketIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractICS26RouterTimeoutPacketIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractICS26RouterTimeoutPacket represents a TimeoutPacket event raised by the ContractICS26Router contract.
type ContractICS26RouterTimeoutPacket struct {
	ClientId common.Hash
	Sequence *big.Int
	Packet   IICS26RouterMsgsPacket
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTimeoutPacket is a free log retrieval operation binding the contract event 0x01e5ed58494819ef3f6480dd08e433b7c08ed75c7abdf2c22c6f04b71340a168.
//
// Solidity: event TimeoutPacket(string indexed clientId, uint256 indexed sequence, (uint64,string,string,uint64,(string,string,string,string,bytes)[]) packet)
func (_ContractICS26Router *ContractICS26RouterFilterer) FilterTimeoutPacket(opts *bind.FilterOpts, clientId []string, sequence []*big.Int) (*ContractICS26RouterTimeoutPacketIterator, error) {

	var clientIdRule []interface{}
	for _, clientIdItem := range clientId {
		clientIdRule = append(clientIdRule, clientIdItem)
	}
	var sequenceRule []interface{}
	for _, sequenceItem := range sequence {
		sequenceRule = append(sequenceRule, sequenceItem)
	}

	logs, sub, err := _ContractICS26Router.contract.FilterLogs(opts, "TimeoutPacket", clientIdRule, sequenceRule)
	if err != nil {
		return nil, err
	}
	return &ContractICS26RouterTimeoutPacketIterator{contract: _ContractICS26Router.contract, event: "TimeoutPacket", logs: logs, sub: sub}, nil
}

// WatchTimeoutPacket is a free log subscription operation binding the contract event 0x01e5ed58494819ef3f6480dd08e433b7c08ed75c7abdf2c22c6f04b71340a168.
//
// Solidity: event TimeoutPacket(string indexed clientId, uint256 indexed sequence, (uint64,string,string,uint64,(string,string,string,string,bytes)[]) packet)
func (_ContractICS26Router *ContractICS26RouterFilterer) WatchTimeoutPacket(opts *bind.WatchOpts, sink chan<- *ContractICS26RouterTimeoutPacket, clientId []string, sequence []*big.Int) (event.Subscription, error) {

	var clientIdRule []interface{}
	for _, clientIdItem := range clientId {
		clientIdRule = append(clientIdRule, clientIdItem)
	}
	var sequenceRule []interface{}
	for _, sequenceItem := range sequence {
		sequenceRule = append(sequenceRule, sequenceItem)
	}

	logs, sub, err := _ContractICS26Router.contract.WatchLogs(opts, "TimeoutPacket", clientIdRule, sequenceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractICS26RouterTimeoutPacket)
				if err := _ContractICS26Router.contract.UnpackLog(event, "TimeoutPacket", log); err != nil {
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

// ParseTimeoutPacket is a log parse operation binding the contract event 0x01e5ed58494819ef3f6480dd08e433b7c08ed75c7abdf2c22c6f04b71340a168.
//
// Solidity: event TimeoutPacket(string indexed clientId, uint256 indexed sequence, (uint64,string,string,uint64,(string,string,string,string,bytes)[]) packet)
func (_ContractICS26Router *ContractICS26RouterFilterer) ParseTimeoutPacket(log types.Log) (*ContractICS26RouterTimeoutPacket, error) {
	event := new(ContractICS26RouterTimeoutPacket)
	if err := _ContractICS26Router.contract.UnpackLog(event, "TimeoutPacket", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractICS26RouterUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the ContractICS26Router contract.
type ContractICS26RouterUpgradedIterator struct {
	Event *ContractICS26RouterUpgraded // Event containing the contract specifics and raw log

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
func (it *ContractICS26RouterUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractICS26RouterUpgraded)
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
		it.Event = new(ContractICS26RouterUpgraded)
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
func (it *ContractICS26RouterUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractICS26RouterUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractICS26RouterUpgraded represents a Upgraded event raised by the ContractICS26Router contract.
type ContractICS26RouterUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ContractICS26Router *ContractICS26RouterFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*ContractICS26RouterUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ContractICS26Router.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &ContractICS26RouterUpgradedIterator{contract: _ContractICS26Router.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ContractICS26Router *ContractICS26RouterFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *ContractICS26RouterUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ContractICS26Router.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractICS26RouterUpgraded)
				if err := _ContractICS26Router.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_ContractICS26Router *ContractICS26RouterFilterer) ParseUpgraded(log types.Log) (*ContractICS26RouterUpgraded, error) {
	event := new(ContractICS26RouterUpgraded)
	if err := _ContractICS26Router.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractICS26RouterWriteAcknowledgementIterator is returned from FilterWriteAcknowledgement and is used to iterate over the raw logs and unpacked data for WriteAcknowledgement events raised by the ContractICS26Router contract.
type ContractICS26RouterWriteAcknowledgementIterator struct {
	Event *ContractICS26RouterWriteAcknowledgement // Event containing the contract specifics and raw log

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
func (it *ContractICS26RouterWriteAcknowledgementIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractICS26RouterWriteAcknowledgement)
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
		it.Event = new(ContractICS26RouterWriteAcknowledgement)
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
func (it *ContractICS26RouterWriteAcknowledgementIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractICS26RouterWriteAcknowledgementIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractICS26RouterWriteAcknowledgement represents a WriteAcknowledgement event raised by the ContractICS26Router contract.
type ContractICS26RouterWriteAcknowledgement struct {
	ClientId         common.Hash
	Sequence         *big.Int
	Packet           IICS26RouterMsgsPacket
	Acknowledgements [][]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterWriteAcknowledgement is a free log retrieval operation binding the contract event 0x76765590e2b799b0506100f8a6610cfecab2c71e8e1f8aa981b099aff0dfdb74.
//
// Solidity: event WriteAcknowledgement(string indexed clientId, uint256 indexed sequence, (uint64,string,string,uint64,(string,string,string,string,bytes)[]) packet, bytes[] acknowledgements)
func (_ContractICS26Router *ContractICS26RouterFilterer) FilterWriteAcknowledgement(opts *bind.FilterOpts, clientId []string, sequence []*big.Int) (*ContractICS26RouterWriteAcknowledgementIterator, error) {

	var clientIdRule []interface{}
	for _, clientIdItem := range clientId {
		clientIdRule = append(clientIdRule, clientIdItem)
	}
	var sequenceRule []interface{}
	for _, sequenceItem := range sequence {
		sequenceRule = append(sequenceRule, sequenceItem)
	}

	logs, sub, err := _ContractICS26Router.contract.FilterLogs(opts, "WriteAcknowledgement", clientIdRule, sequenceRule)
	if err != nil {
		return nil, err
	}
	return &ContractICS26RouterWriteAcknowledgementIterator{contract: _ContractICS26Router.contract, event: "WriteAcknowledgement", logs: logs, sub: sub}, nil
}

// WatchWriteAcknowledgement is a free log subscription operation binding the contract event 0x76765590e2b799b0506100f8a6610cfecab2c71e8e1f8aa981b099aff0dfdb74.
//
// Solidity: event WriteAcknowledgement(string indexed clientId, uint256 indexed sequence, (uint64,string,string,uint64,(string,string,string,string,bytes)[]) packet, bytes[] acknowledgements)
func (_ContractICS26Router *ContractICS26RouterFilterer) WatchWriteAcknowledgement(opts *bind.WatchOpts, sink chan<- *ContractICS26RouterWriteAcknowledgement, clientId []string, sequence []*big.Int) (event.Subscription, error) {

	var clientIdRule []interface{}
	for _, clientIdItem := range clientId {
		clientIdRule = append(clientIdRule, clientIdItem)
	}
	var sequenceRule []interface{}
	for _, sequenceItem := range sequence {
		sequenceRule = append(sequenceRule, sequenceItem)
	}

	logs, sub, err := _ContractICS26Router.contract.WatchLogs(opts, "WriteAcknowledgement", clientIdRule, sequenceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractICS26RouterWriteAcknowledgement)
				if err := _ContractICS26Router.contract.UnpackLog(event, "WriteAcknowledgement", log); err != nil {
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

// ParseWriteAcknowledgement is a log parse operation binding the contract event 0x76765590e2b799b0506100f8a6610cfecab2c71e8e1f8aa981b099aff0dfdb74.
//
// Solidity: event WriteAcknowledgement(string indexed clientId, uint256 indexed sequence, (uint64,string,string,uint64,(string,string,string,string,bytes)[]) packet, bytes[] acknowledgements)
func (_ContractICS26Router *ContractICS26RouterFilterer) ParseWriteAcknowledgement(log types.Log) (*ContractICS26RouterWriteAcknowledgement, error) {
	event := new(ContractICS26RouterWriteAcknowledgement)
	if err := _ContractICS26Router.contract.UnpackLog(event, "WriteAcknowledgement", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
