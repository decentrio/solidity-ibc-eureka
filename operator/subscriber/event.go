package subscriber

import (
	"context"
	"encoding/hex"
	"operator/services"

	channeltypes "github.com/cosmos/ibc-go/v10/modules/core/04-channel/v2/types"
)

const COMETBFT_SEND_PACKET_EVENT = "tm.event = 'Tx' AND message.action = '/ibc.core.channel.v2.MsgSendPacket'"

type Subscriber struct {
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
			sendPacketEvent := e.Events[channeltypes.EventTypeSendPacket]
			if sendPacketEvent == nil {
				continue
			}
			packetHex := sendPacketEvent[channeltypes.AttributeKeyEncodedPacketHex]
			packet, err := hex.DecodeString(packetHex)
		}
	}
}
