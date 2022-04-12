package testes

import "

const (
	// scaling factor (in powers of 2)
	defaultMaxScalingFactor = 6
	maxScalingFactor        = 10
)

const (
	maxArgChoices int     = 10
	one           uint64  = 1<<64 - 1
	halfBool      int64   = 1<<63 - 1
	halfHalfBool  int64   = 1<<62 - 1
	ratio         float64 = float64(one) / float64(halfBool)
	halfRatio     float64 = float64(halfBool) / float64(halfHalfBool)
)

const (
	b8  = 1<<8 - 1
	b16 = 1<<16 - 1
	b32 = 1<<32 - 1
	b64 = 1<<64 - 1
	u8  = b64 >> 56
	u16 = b64 >> 48
)

const arune rune = 'A'

var (
	global Any

	Log = errorlogger.New()
	log = Log
)
