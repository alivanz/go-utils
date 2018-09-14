package utils

import (
	"bytes"
	"testing"
)

func TestBytesOrder(t *testing.T) {
	vec := []byte{0, 1, 2, 3, 4}
	ref := []byte{4, 3, 2, 1, 0}
	if !bytes.Equal(ReverseBytes(vec), ref) {
		t.Fail()
	}
}
