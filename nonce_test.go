package utils

import (
	"testing"
)

func TestNonce(t *testing.T) {
	var nonce64 Nonce64
	var nonce96 Nonce96
	var nonce128 Nonce128
	var nonce256 Nonce256
	nonce96.Nonce0 = 0xffffffffffffffff
	nonce128.Nonce0 = 0xffffffffffffffff
	nonce256.Nonce0 = 0xffffffffffffffff
	nonce256.Nonce1 = 0xffffffffffffffff
	nonce256.Nonce2 = 0xffffffffffffffff
	for i := 0; i < 5; i++ {
		t.Log(nonce64.Nonce())
	}
	for i := 0; i < 5; i++ {
		t.Log(nonce96.Nonce())
	}
	for i := 0; i < 5; i++ {
		t.Log(nonce128.Nonce())
	}
	for i := 0; i < 5; i++ {
		t.Log(nonce256.Nonce())
	}
}
