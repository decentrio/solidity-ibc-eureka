package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"operator/utils"
	"os"
	"strings"

	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
	"github.com/cosmos/gogoproto/proto"
	channeltypesv2 "github.com/cosmos/ibc-go/v10/modules/core/04-channel/v2/types"
)

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

func subscribeCosmos(logger *log.Logger, client *rpchttp.HTTP) {
	sub, err := client.WSEvents.Subscribe(context.Background(), "", COMETBFT_SEND_PACKET_EVENT)
	if err != nil {
		logger.Println(err.Error())
	}

	for {
		select {
		case e := <-sub:
			fmt.Println("events: ", e.Events)
			// handle event
			sendPacketEvent := e.Events[channeltypesv2.EventTypeSendPacket]
			if sendPacketEvent == nil {
				continue
			}
			packetHex := e.Events[channeltypesv2.AttributeKeyEncodedPacketHex]
			fmt.Println("packetHex: ", packetHex)

			txHashStr := e.Events[EVENT_TX_HASH_FIELD]
			if txHashStr == nil {
				continue
			}
			txHash, err := mustTxHashBytes(txHashStr[0])
			if err != nil {
				fmt.Println(fmt.Errorf("Failed to decode tx hash: %s", err.Error()))
				continue
			}

			packetEncodedStr := sendPacketEvent[0]
			packetBytes, err := hex.DecodeString(packetEncodedStr)
			if err != nil {
				fmt.Println(fmt.Errorf("Failed to decode packet hex: %s", err.Error()))
				continue
			}

			var packet channeltypesv2.Packet
			err = proto.Unmarshal(packetBytes, &packet)
			if err != nil {
				fmt.Println(fmt.Errorf("Failed to unmarshal packet: %s", err.Error()))
				continue
			}

			txResp, err := client.Tx(context.Background(), txHash, true)
			if err != nil {
				fmt.Println(fmt.Errorf("Failed to fetch tx from tx hash: %s", err.Error()))
				continue
			}

			revisionHeight := int64(txResp.Height)

			// target height are the latest block height
			_, merkleProof, err := utils.ProvePath(client, txResp.Proof.Proof.Aunts, uint64(revisionHeight))

			fmt.Println(merkleProof)

		}
	}
}

func main() {

	tendermintRpcClient, err := rpchttp.New("http://127.0.0.1:26657", "/websocket")
	if err != nil {
		panic(fmt.Errorf("failed to create RPC client: %w", err).Error())
	}
	err = tendermintRpcClient.Start()
	if err != nil {
		panic(fmt.Errorf("err starting rpc  client: %w", err).Error())
	}
	log := log.Default()
	fmt.Println("start indexing...")

	subscribeCosmos(log, tendermintRpcClient)
}

func mustTxHashBytes(txHashHex string) ([]byte, error) {
	h := strings.TrimSpace(txHashHex)
	h = strings.TrimPrefix(h, "0x")
	bz, err := hex.DecodeString(h)
	if err != nil {
		return nil, fmt.Errorf("invalid tx hash hex: %w", err)
	}
	if len(bz) != 32 {
		return nil, fmt.Errorf("tx hash must decode to 32 bytes, got %d", len(bz))
	}
	return bz, nil
}
