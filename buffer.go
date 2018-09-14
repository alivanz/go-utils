package utils

type WriteBuffer struct {
	data []byte
}

func NewWriteBuffer() *WriteBuffer {
	return &WriteBuffer{[]byte{}}
}
func (b *WriteBuffer) Write(data []byte) (int, error) {
	b.data = append(b.data, data...)
	return len(data), nil
}
func (b *WriteBuffer) Binary() []byte {
	return b.data
}
