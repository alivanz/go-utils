package utils

import (
	"bytes"
	"io"
	"os/exec"
	"strings"
)

func ExecReadLines(name string, params []string) ([]string, error) {
	out, err := exec.Command(name, params...).Output()
	if err != nil {
		return nil, err
	}
	r := bytes.NewBuffer(out)
	lines := make([]string, 0)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			return lines, nil
		} else if err != nil {
			return nil, err
		}
		line = strings.TrimSuffix(line, "\n")
		line = strings.TrimSuffix(line, "\x10")
		lines = append(lines, line)
	}
}
