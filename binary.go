package utils

import (
	"io"
)

type BinaryReadWriter interface {
	BinaryReader
	BinaryWriter
}
type BinaryWriter interface {
	io.Writer
	WriteByte(b byte) error
	WriteUint16(d uint16) error
	WriteUint32(d uint32) error
	WriteUint64(d uint64) error
	WriteDERLength(uint64) error
	WriteDER([]byte) error
	WriteCompactUint(uint64) error
	WriteCompact([]byte) error
}
type BinaryReader interface {
	io.Reader
	ReadByte() (byte, error)
	ReadUint16() (uint16, error)
	ReadUint32() (uint32, error)
	ReadUint64() (uint64, error)
	ReadDERLength() (uint64, error)
	ReadDER() ([]byte, error)
	ReadCompactUint() (uint64, error)
	ReadCompact() ([]byte, error)
}
