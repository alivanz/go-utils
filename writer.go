package utils

import (
	"bytes"
	"encoding/binary"
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

func CompactUint(i uint64) []byte {
	b := make([]byte, 9)
	binary.LittleEndian.PutUint64(b[1:], i)
	if i <= 252 {
		return b[1:2]
	} else if i <= 0xffff {
		b[0] = 0xfd
		return b[:3]
	} else if i <= 0xffffffff {
		b[0] = 0xfe
		return b[:5]
	} else {
		b[0] = 0xff
		return b[:9]
	}
}
func (w *bwriter) WriteCompactUint(i uint64) error {
	if _, err := w.Write(CompactUint(i)); err != nil {
		return err
	}
	return nil
}
func (w *bwriter) WriteCompact(bin []byte) error {
	buf := bytes.NewBuffer(nil)
	buf.Write(CompactUint(uint64(len(bin))))
	buf.Write(bin)
	_, err := w.Write(buf.Bytes())
	return err
}

func DERLength(l uint64) []byte {
	if l < 0x80 {
		return []byte{byte(l)}
	}
	buf := bytes.NewBuffer(nil)
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, l)
	var x int
	for x = 0; x < 8; x = x + 1 {
		if b[x] != 0 {
			break
		}
	}
	buf.WriteByte(byte(0x88 - x))
	buf.Write(b[x:])
	return buf.Bytes()
}
func (w *bwriter) WriteDERLength(l uint64) error {
	if _, err := w.Write(DERLength(l)); err != nil {
		return err
	}
	return nil
}
func (w *bwriter) WriteDER(bin []byte) error {
	buf := bytes.NewBuffer(nil)
	buf.Write(DERLength(uint64(len(bin))))
	buf.Write(bin)
	_, err := w.Write(buf.Bytes())
	return err
}
