package utils

import "io"

func Pump(src io.Reader, dst io.Writer) {
	var err error
	var n int
	buf := make([]byte, 1024)
	for {
		n, err = src.Read(buf)
		if err != nil {
			return
		}
		_, err = dst.Write(buf[:n])
		if err != nil {
			return
		}
	}
}

func PumpClose(src io.ReadCloser, dst io.WriteCloser) {
	defer src.Close()
	defer dst.Close()
	Pump(src, dst)
}
