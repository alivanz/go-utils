package utils

import "io"

type readwriter struct {
	*breader
	*bwriter
}

func NewBinaryReadWriter(x io.ReadWriter) BinaryReadWriter {
	return readwriter{&breader{x}, &bwriter{x}}
}
