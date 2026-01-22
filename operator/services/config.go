package services

import "time"

const DEFAULT_INTERVAL = 5
const DEFAULT_TIME_PERIODS = time.Second * 20 //default 20 seconds

type IntervalType uint8

const (
	blockHeight IntervalType = iota
	timestamp
)

type IntervalConig struct {
	blockHeight uint8
	blockTime   time.Duration
}

type Config struct {
	KeyPath        string
	IntervalParams IntervalConig
	IntervalType   IntervalType
}

func NewConfig(KeyPath string, params IntervalConig, intervalType IntervalType) Config {
	return Config{
		IntervalParams: params,
		IntervalType:   intervalType,
	}
}

func DefaultConfig() Config {
	return Config{
		IntervalParams: IntervalConig{
			blockHeight: DEFAULT_INTERVAL,
			blockTime:   DEFAULT_TIME_PERIODS,
		},
		IntervalType: timestamp,
	}
}
