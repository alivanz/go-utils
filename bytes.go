package utils

func ReverseBytes(b []byte) []byte {
	l := len(b)
	out := make([]byte, l)
	for i, x := range b {
		out[l-i-1] = x
	}
	return out
}
