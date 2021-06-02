package utils

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func RandUint64() uint64 {
	return rand.Uint64()
}
