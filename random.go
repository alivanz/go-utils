package utils

import (
	"io"
	"math/rand"
	"time"
)

var (
	HexChars = []byte("0123456789abcdef")
	Alphabet = []byte("abcdefghijklmnopqrstuvwxyz")
)

func RandomChars(chars []byte, length int) []byte {
	out := make([]byte, length)
	l := len(chars)
	for i := range out {
		out[i] = chars[rand.Int()%l]
	}
	return out
}
func RandomBytes(length int) []byte {
	out := make([]byte, length)
	io.ReadFull(rand.New(rand.NewSource(time.Now().UnixNano())), out)
	return out
}
