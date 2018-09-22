package utils

import "C"
import (
	"encoding/binary"
)

type NonceGenerator interface {
	Nonce() []byte
}

type Nonce64 uint64
type Nonce96 struct {
	Nonce0 uint64
	Nonce1 uint32
}
type Nonce128 struct {
	Nonce0 uint64
	Nonce1 uint64
}
type Nonce256 struct {
	Nonce0 uint64
	Nonce1 uint64
	Nonce2 uint64
	Nonce3 uint64
}

func (nonce *Nonce64) Nonce() []byte {
	bin := make([]byte, 8)
	binary.LittleEndian.PutUint64(bin, uint64(*nonce))
	(*nonce)++
	return bin
}
func (nonce *Nonce96) Nonce() []byte {
	bin := make([]byte, 12)
	binary.LittleEndian.PutUint64(bin, nonce.Nonce0)
	binary.LittleEndian.PutUint32(bin[8:], nonce.Nonce1)
	nonce.Nonce0++
	if nonce.Nonce0 == 0 {
		nonce.Nonce1++
	}
	return bin
}
func (nonce *Nonce128) Nonce() []byte {
	bin := make([]byte, 16)
	binary.LittleEndian.PutUint64(bin, nonce.Nonce0)
	binary.LittleEndian.PutUint64(bin[8:], nonce.Nonce1)
	nonce.Nonce0++
	if nonce.Nonce0 == 0 {
		nonce.Nonce1++
	}
	return bin
}
func (nonce *Nonce256) Nonce() []byte {
	bin := make([]byte, 32)
	binary.LittleEndian.PutUint64(bin, nonce.Nonce0)
	binary.LittleEndian.PutUint64(bin[8:], nonce.Nonce1)
	binary.LittleEndian.PutUint64(bin[16:], nonce.Nonce2)
	binary.LittleEndian.PutUint64(bin[24:], nonce.Nonce3)
	nonce.Nonce0++
	if nonce.Nonce0 != 0 {
		goto ret
	}
	nonce.Nonce1++
	if nonce.Nonce1 != 0 {
		goto ret
	}
	nonce.Nonce2++
	if nonce.Nonce2 != 0 {
		goto ret
	}
	nonce.Nonce3++
ret:
	return bin
}
