package util

import (
	"math/rand"
	"time"
)

func RandomBytes(bit int) []byte {
	bs := make([]byte, bit)

	n, err := rand.New(rand.NewSource(time.Now().UnixNano())).Read(bs)
	if err != nil {
		panic(err)
	}
	return bs[:n]
}