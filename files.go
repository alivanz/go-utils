package utils

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func ReadFileLines(filename string) ([]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	r := bufio.NewReader(f)
	out := make([]string, 0)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		text := strings.TrimSuffix(line, "\n")
		out = append(out, text)
	}
	return out, nil
}
