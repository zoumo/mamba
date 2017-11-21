package types

import (
	"time"
)

// go:generate flag-gen -i github.com/zoumo/mamba/cmd/flag-gen/types -o github.com/zoumo/mamba

// FlagTypes ...
type FlagTypes struct {
	a bool
	// b []bool not support now
	c time.Duration
	d float32
	e float64
	f int
	// g []int not support now
	j int32
	k int64
	l string
	m []string
	n uint
	// o []uint not support now
	r uint32
	s uint64
}
