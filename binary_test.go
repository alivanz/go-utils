package utils

import (
	"bytes"
	"testing"
)

func TestDERLength(t *testing.T) {
	// test vector
	vector := []uint64{2, 7, 49, 48, 70, 255, 437, 946, 606, 551, 557432421, 960369495, 132093869, 379953793, 195865593, 768670073}
	// writer
	buf := NewWriteBuffer()
	w := NewBinaryWriter(buf)
	for _, x := range vector {
		w.WriteDERLength(x)
	}
	bin := buf.Binary()
	// Reader
	r := NewBinaryReader(bytes.NewBuffer(bin))
	for _, x := range vector {
		u, _ := r.ReadDERLength()
		if u != x {
			t.Fail()
		}
	}
}

func TestUint(t *testing.T) {
	// test vector
	vector32 := []uint32{2, 7, 49, 48, 70, 255, 437, 946, 606, 551}
	vector64 := []uint64{557432421, 960369495, 132093869, 379953793, 195865593, 768670073}
	// writer
	buf := NewWriteBuffer()
	w := NewBinaryWriter(buf)
	for _, x32 := range vector32 {
		w.WriteUint32(x32)
	}
	for _, x64 := range vector64 {
		w.WriteUint64(x64)
	}
	bin := buf.Binary()
	// Reader
	r := NewBinaryReader(bytes.NewBuffer(bin))
	for _, x32 := range vector32 {
		u, _ := r.ReadUint32()
		if u != x32 {
			t.Fail()
		}
	}
	for _, x64 := range vector64 {
		u, _ := r.ReadUint64()
		if u != x64 {
			t.Fail()
		}
	}
}

func TestDERBytes(t *testing.T) {
	vector := [][]byte{
		[]byte("PfW7hstIlU"),
		[]byte("8ubC094AjN"),
		[]byte("5HbmHcVFA8"),
		[]byte("VvtoL3YTOQ"),
		[]byte("Cq6yxiAhi4"),
		[]byte("OU3JvPxVjk"),
		[]byte("sppszV5rRz"),
		[]byte("lekbWJTkMT"),
		[]byte("gpKHFMqtJb"),
		[]byte("PYQWHWRKqc"),
	}
	// writer
	buf := NewWriteBuffer()
	w := NewBinaryWriter(buf)
	for _, x := range vector {
		w.WriteDER(x)
	}
	bin := buf.Binary()
	// Reader
	r := NewBinaryReader(bytes.NewBuffer(bin))
	for _, x := range vector {
		u, _ := r.ReadDER()
		if !bytes.Equal(u, x) {
			t.Fail()
		}
	}
}
