package subscriber

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"

	contractICS26Router "operator/bindings/ICS26Router"
	tendermintContract "operator/bindings/SP1ICS07Tendermint"
	"operator/services"
	"operator/utils"

	channeltypesv2 "github.com/cosmos/ibc-go/v10/modules/core/04-channel/v2/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gogo/protobuf/proto"
)

// read abi json file once in runtime
var tendermintAbiJson []byte
var initErr error

func init() {
	tendermintAbiJson, initErr = os.ReadFile("../../abi/SP1ICS07Tendermint.json")
	if initErr != nil {
		log.Fatal(initErr)
	}
}

const COMETBFT_SEND_PACKET_EVENT = "tm.event = 'Tx' AND message.action = '/ibc.applications.transfer.v1.MsgTransfer'"
const EVENT_SEND_PACKET_FIELD = "send_packet.encoded_packet_hex"
const EVENT_TX_HASH_FIELD = "tx.hash"

type Subscriber struct {
	txHandler services.TransactionHandler
}

func NewSubscriber(txHandler services.TransactionHandler) *Subscriber {
	return &Subscriber{
		txHandler,
	}
}

func (s *Subscriber) SubscribeCosmos(ctx services.Context) {
	sub, err := ctx.CosmosClient().WSEvents.Subscribe(context.Background(), "", COMETBFT_SEND_PACKET_EVENT)
	if err != nil {
		ctx.Logger.Println(err.Error())
	}

	for {
		select {
		case e := <-sub:
			// handle event
			sendPacketEvent := e.Events[EVENT_SEND_PACKET_FIELD]
			if sendPacketEvent == nil {
				continue
			}

			txHashStr := e.Events[EVENT_TX_HASH_FIELD]
			if txHashStr == nil {
				ctx.Logger.Println(fmt.Errorf("Invalid tx hash"))
				continue
			}
			txHash, err := hex.DecodeString(txHashStr[0])
			if err != nil {
				// TODO handle log here
				ctx.Logger.Println(fmt.Errorf("Failed to decode tx hash: %s", err.Error()))
				continue
			}

			packetEncodedStr := sendPacketEvent[0]
			packetBytes, err := hex.DecodeString(packetEncodedStr)
			if err != nil {
				// TODO handle log here
				ctx.Logger.Println(fmt.Errorf("Failed to decode packet hex: %s", err.Error()))
				continue
			}

			var packet channeltypesv2.Packet
			err = proto.Unmarshal(packetBytes, &packet)
			if err != nil {
				// TODO handle log here
				ctx.Logger.Println(fmt.Errorf("Failed to unmarshal packet: %s", err.Error()))
				continue
			}

			txResp, err := ctx.CosmosClient().Tx(context.Background(), txHash, true)
			if err != nil {
				// TODO handle log here
				ctx.Logger.Println(fmt.Errorf("Failed to fetch tx from tx hash: %s", err.Error()))
				continue
			}

			revisionHeight := int64(ctx.LatestCosmosTimestamp().LatestUpdateHeight)

			// get trusted block
			trustedBlockResp, err := ctx.CosmosClient().Block(context.Background(), &revisionHeight)
			if err != nil {
				// TODO handle log here
				ctx.Logger.Println(fmt.Errorf("Failed to fetch block from block height %v: %s", txResp.Height, err.Error()))
				continue
			}

			// target height are the latest block height
			value, merkleProof, err := utils.ProvePath(ctx, txResp.Proof.Proof.Aunts, uint64(txResp.Height))

			membershipMsg := tendermintContract.ILightClientMsgsMsgVerifyMembership{
				Height: tendermintContract.IICS02ClientMsgsHeight{
					RevisionHeight: uint64(txResp.Height),
					RevisionNumber: 0,
				},
				KvPairs: []tendermintContract.IMembershipMsgsKVPair{
					{
						Path:  txResp.Proof.Proof.Aunts,
						Value: utils.BytesToBytes32(value),
					},
				},
				MerkleProofs: []tendermintContract.IMembershipMsgsMerkleProof{
					*merkleProof,
				},
				// current appHash
				AppHash: utils.BytesToBytes32(txResp.TxResult),
				// trusted consensus from revision height block
				TrustedConsensusState: tendermintContract.IICS07TendermintMsgsConsensusState{
					Timestamp:          big.NewInt(trustedBlockResp.Block.Header.Time.Unix()),
					Root:               utils.BytesToBytes32(trustedBlockResp.Block.Header.ConsensusHash),
					NextValidatorsHash: utils.BytesToBytes32(trustedBlockResp.Block.Header.NextValidatorsHash),
				},
				MembershipType: 1,
			}

			parsedABI, err := abi.JSON(strings.NewReader(string(tendermintAbiJson)))
			if err != nil {
				ctx.Logger.Println(fmt.Errorf("Failed to read abi json file: %s", err.Error()))
			}

			calldata, err := parsedABI.Pack("verifyMembership", membershipMsg)
			if err != nil {
				ctx.Logger.Println(fmt.Errorf("Failed to abi encode verify msg: %s", err.Error()))
			}

			payloads := make([]contractICS26Router.IICS26RouterMsgsPayload, len(packet.Payloads))
			for _, p := range packet.Payloads {
				payloads = append(payloads, contractICS26Router.IICS26RouterMsgsPayload{
					SourcePort: p.SourcePort,
					DestPort:   p.DestinationPort,
					Version:    p.Version,
					Encoding:   p.Encoding,
					Value:      p.Value,
				})
			}

			msgRecvPacket := contractICS26Router.IICS26RouterMsgsMsgRecvPacket{
				Packet: contractICS26Router.IICS26RouterMsgsPacket{
					Sequence:         packet.Sequence,
					SourceClient:     packet.SourceClient,
					DestClient:       packet.DestinationClient,
					TimeoutTimestamp: packet.TimeoutTimestamp,
					Payloads:         payloads,
				},
				MembershipMsg: calldata,
			}

			// try send recv packet msg to ethereum
			s.txHandler.SendTx(ctx, msgRecvPacket)
		}
	}
}

// SubscribeEth subscribes to Ethereum events from the ICS26Router contract
func (s *Subscriber) SubscribeEth(ctx services.Context) {
	// Get the ICS26Router contract address from environment variable
	ics26RouterAddr := os.Getenv("ICS26_ROUTER_ADDRESS")
	if ics26RouterAddr == "" {
		ctx.Logger.Println("ICS26_ROUTER_ADDRESS environment variable is required")
		return
	}

	// Parse the contract address
	contractAddr := common.HexToAddress(ics26RouterAddr)

	// Create a new ICS26Router filterer instance (only for event subscriptions)
	filterer, err := contractICS26Router.NewContractICS26RouterFilterer(contractAddr, ctx.EthClient())
	if err != nil {
		ctx.Logger.Printf("Failed to create ICS26Router filterer instance: %v", err)
		return
	}

	// Create event channels for each event type
	sendPacketCh := make(chan *contractICS26Router.ContractICS26RouterSendPacket)
	writeAckCh := make(chan *contractICS26Router.ContractICS26RouterWriteAcknowledgement)
	ackPacketCh := make(chan *contractICS26Router.ContractICS26RouterAckPacket)
	timeoutPacketCh := make(chan *contractICS26Router.ContractICS26RouterTimeoutPacket)

	// Set up watch options (nil for all events, no filtering by clientId or sequence)
	watchOpts := &bind.WatchOpts{Context: context.Background()}

	// Subscribe to SendPacket events
	sendPacketSub, err := filterer.WatchSendPacket(watchOpts, sendPacketCh, nil, nil)
	if err != nil {
		ctx.Logger.Printf("Failed to subscribe to SendPacket events: %v", err)
		return
	}
	defer sendPacketSub.Unsubscribe()

	// Subscribe to WriteAcknowledgement events
	writeAckSub, err := filterer.WatchWriteAcknowledgement(watchOpts, writeAckCh, nil, nil)
	if err != nil {
		ctx.Logger.Printf("Failed to subscribe to WriteAcknowledgement events: %v", err)
		return
	}
	defer writeAckSub.Unsubscribe()

	// Subscribe to AckPacket events
	ackPacketSub, err := filterer.WatchAckPacket(watchOpts, ackPacketCh, nil, nil)
	if err != nil {
		ctx.Logger.Printf("Failed to subscribe to AckPacket events: %v", err)
		return
	}
	defer ackPacketSub.Unsubscribe()

	// Subscribe to TimeoutPacket events
	timeoutPacketSub, err := filterer.WatchTimeoutPacket(watchOpts, timeoutPacketCh, nil, nil)
	if err != nil {
		ctx.Logger.Printf("Failed to subscribe to TimeoutPacket events: %v", err)
		return
	}
	defer timeoutPacketSub.Unsubscribe()

	ctx.Logger.Println("Successfully subscribed to ICS26Router events")

	// Event loop to handle incoming events
	for {
		select {
		case ev := <-sendPacketCh:
			// TODO: handle SendPacket event
			// This event is emitted when a packet is sent from Ethereum to Cosmos
			ctx.Logger.Printf("SendPacket event received: clientId=%x, sequence=%s", ev.ClientId, ev.Sequence.String())

		case ev := <-writeAckCh:
			// TODO: handle WriteAcknowledgement event
			// This event is emitted when an acknowledgement is written on Ethereum
			ctx.Logger.Printf("WriteAcknowledgement event received: clientId=%x, sequence=%s", ev.ClientId, ev.Sequence.String())

		case ev := <-ackPacketCh:
			// TODO: handle AckPacket event
			// This event is emitted when a packet acknowledgement is received on Ethereum
			ctx.Logger.Printf("AckPacket event received: clientId=%x, sequence=%s", ev.ClientId, ev.Sequence.String())

		case ev := <-timeoutPacketCh:
			// TODO: handle TimeoutPacket event
			// This event is emitted when a packet times out
			ctx.Logger.Printf("TimeoutPacket event received: clientId=%x, sequence=%s", ev.ClientId, ev.Sequence.String())

		case err := <-sendPacketSub.Err():
			ctx.Logger.Printf("SendPacket subscription error: %v", err)
			return

		case err := <-writeAckSub.Err():
			ctx.Logger.Printf("WriteAcknowledgement subscription error: %v", err)
			return

		case err := <-ackPacketSub.Err():
			ctx.Logger.Printf("AckPacket subscription error: %v", err)
			return

		case err := <-timeoutPacketSub.Err():
			ctx.Logger.Printf("TimeoutPacket subscription error: %v", err)
			return
		}
	}
}
