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


func Rand_BytesArray(len int) []byte {
	out := make([]byte, len)
	_, err := rand.Read(out)
	if err != nil {
		panic("utils: rand.Read failed")
	}
	return out
}