package utils

import (
	"encoding/binary"
	"fmt"
	"io"
)

type bwriter struct {
	io.Writer
}

func NewBinaryWriter(w io.Writer) BinaryWriter {
	return &bwriter{w}
}

func (w *bwriter) WriteByte(b byte) error {
	x := []byte{b}
	_, err := w.Write(x)
	return err
}
func (w *bwriter) WriteUint16(d uint16) error {
	b := make([]byte, 2)
	binary.LittleEndian.PutUint16(b, d)
	_, err := w.Write(b)
	return err
}
func (w *bwriter) WriteUint32(d uint32) error {
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, d)
	_, err := w.Write(b)
	return err
}
func (w *bwriter) WriteUint64(d uint64) error {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, d)
	_, err := w.Write(b)
	return err
}

func (w *bwriter) WriteCompactUint(i uint64) error {
	if i <= 252 {
		return w.WriteByte(byte(i))
	} else if i <= 0xffff {
		w.WriteByte(0xfd)
		return w.WriteUint16(uint16(i))
	} else if i <= 0xffffffff {
		w.WriteByte(0xfe)
		return w.WriteUint32(uint32(i))
	} else {
		w.WriteByte(0xff)
		return w.WriteUint64(i)
	}
}
func (w *bwriter) WriteCompact(bin []byte) error {
	w.WriteCompactUint(uint64(len(bin)))
	_, err := w.Write(bin)
	return err
}

func (w *bwriter) WriteDERLength(l uint64) error {
	if l < 0x80 {
		return w.WriteByte(byte(l))
	}
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, l)
	for x := 0; x < 8; x = x + 1 {
		if b[x] != 0 {
			w.WriteByte(byte(0x88 - x))
			_, err := w.Write(b[x:])
			return err
		}
	}
	return fmt.Errorf("Wrong algo")
}
func (w *bwriter) WriteDER(bin []byte) error {
	w.WriteDERLength(uint64(len(bin)))
	_, err := w.Write(bin)
	return err
}
