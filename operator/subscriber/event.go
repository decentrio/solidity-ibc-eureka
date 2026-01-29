package subscriber

import (
	"context"
	"operator/services"

	"github.com/cosmos/ibc-go/v10/modules/apps/transfer/internal/events"
)

const COMETBFT_SEND_PACKET_EVENT = "tm.event = 'Tx' AND message.action = '/ibc.core.channel.v1.MsgSendPacket'"

type Subscriber struct {
}

func (s *Subscriber) SubscribeCosmos(ctx services.Context) {
	sub, err := ctx.CosmosClient().WSEvents.Subscribe(context.Background(), "", COMETBFT_SEND_PACKET_EVENT)
	if err != nil {
		ctx.Logger.Println(err.Error())
	}

	for {
		select {
		case ev := <-sub:
			// handle event
			ev.Events[]
		}
	}
}

