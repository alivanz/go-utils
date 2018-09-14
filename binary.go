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
	WriteUint32(d uint32) error
	WriteUint64(d uint64) error
	WriteDERLength(uint64) error
	WriteDER([]byte) error
}
type BinaryReader interface {
	io.Reader
	ReadByte() (byte, error)
	ReadUint32() (uint32, error)
	ReadUint64() (uint64, error)
	ReadDERLength() (uint64, error)
	ReadDER() ([]byte, error)
}
