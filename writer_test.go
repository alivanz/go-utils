package utils

import (
	"bytes"
	"testing"
)

func TestCompact(t *testing.T) {
	vec := []uint64{0, 1, 2, 3, 2344, 1237768, 31273128736, 13276378163682}
	for _, v := range vec {
		buf := bytes.NewBuffer(nil)
		w := NewBinaryWriter(buf)
		w.WriteCompactUint(v)
		if !bytes.Equal(buf.Bytes(), CompactUint(v)) {
			t.Fail()
		}
	}
}
func TestDER(t *testing.T) {
	vec := []uint64{0, 1, 2, 3, 2344, 1237768, 31273128736, 13276378163682}
	for _, v := range vec {
		buf := bytes.NewBuffer(nil)
		w := NewBinaryWriter(buf)
		w.WriteDERLength(v)
		if !bytes.Equal(buf.Bytes(), DERLength(v)) {
			t.Fail()
		}
	}
}
