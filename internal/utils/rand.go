package utils

import (
	i "github.com/AgentCoop/peppermint/internal"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func RandUint64() uint64 {
	return rand.Uint64()
}
func UniqueId() i.UniqueId {
	return i.UniqueId(rand.Uint64())
}
