package utils

import (
	"encoding/binary"
	"io"
)

type breader struct {
	io.Reader
}

func NewBinaryReader(r io.Reader) BinaryReader {
	return breader{r}
}
func (r breader) ReadFull(b []byte) error {
	_, err := io.ReadFull(r, b)
	return err
}
func (r breader) ReadByte() (byte, error) {
	b := make([]byte, 1)
	_, err := io.ReadFull(r, b)
	return b[0], err
}
func (r breader) ReadUint16() (uint16, error) {
	b := make([]byte, 2)
	err := r.ReadFull(b)
	return binary.LittleEndian.Uint16(b), err
}
func (r breader) ReadUint32() (uint32, error) {
	b := make([]byte, 4)
	err := r.ReadFull(b)
	return binary.LittleEndian.Uint32(b), err
}
func (r breader) ReadUint64() (uint64, error) {
	b := make([]byte, 8)
	err := r.ReadFull(b)
	return binary.LittleEndian.Uint64(b), err
}

func (r breader) ReadCompactUint() (uint64, error) {
	b, err := r.ReadByte()
	if err != nil {
		return 0, err
	}
	if b == 0xfd {
		x, err := r.ReadUint16()
		return uint64(x), err
	} else if b == 0xfe {
		x, err := r.ReadUint32()
		return uint64(x), err
	} else if b == 0xff {
		return r.ReadUint64()
	} else {
		return uint64(b), nil
	}
}
func (r breader) ReadCompact() ([]byte, error) {
	l, err := r.ReadCompactUint()
	if err != nil {
		return nil, err
	}
	bin := make([]byte, l)
	err = r.ReadFull(bin)
	return bin, err
}

func (r breader) ReadDERLength() (uint64, error) {
	b, err := r.ReadByte()
	if b < 0x80 {
		return uint64(b), nil
	}
	llen := b - 0x80
	bin := make([]byte, 8)
	err = r.ReadFull(bin[8-llen:])
	return binary.BigEndian.Uint64(bin), err
}
func (r breader) ReadDER() ([]byte, error) {
	l, err := r.ReadDERLength()
	if err != nil {
		return nil, err
	}
	bin := make([]byte, l)
	err = r.ReadFull(bin)
	return bin, err
}
