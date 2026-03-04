package services

import (
	channeltypesv2 "github.com/cosmos/ibc-go/v10/modules/core/04-channel/v2/types"
)

type Packet struct {
	Packet *channeltypesv2.Packet
}

type BatchPackets struct {
	Packets []Packet
}
