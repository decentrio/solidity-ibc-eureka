package services

import "time"

type BatchBuilder struct {
	timestamp time.Time
	packets   []Packet
}

func NewBatchBuidler() *BatchBuilder {
	return &BatchBuilder{
		timestamp: time.Now(),
		packets:   []Packet{},
	}
}

func (b *BatchBuilder) InsertPacket(packet Packet) {
	b.packets = append(b.packets, packet)
}

func (b *BatchBuilder) ClearBatch() {
	b.timestamp = time.Now()
	b.packets = []Packet{}
}

func (b *BatchBuilder) CheckBatch(ctx Context) {
	if len(b.packets) < int(ctx.Config.BatchConfig.BatchSize) && b.timestamp.After(time.Now().Add(ctx.Config.BatchConfig.BatchPeriods)) {
		ctx.BatchPackets <- BatchPackets{
			Packets: b.packets,
		}

		b.ClearBatch()
	} else if len(b.packets) > int(ctx.Config.BatchConfig.BatchSize) {
		ctx.BatchPackets <- BatchPackets{
			Packets: b.packets,
		}

		b.ClearBatch()
	}
}
