package utils

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// SSE A SSE reader
type SSE struct {
	reader *bufio.Reader
}

// SSEEvent represent SSE request sent by a server
type SSEEvent struct {
	Event string
	Data  []byte
}

// NewSSE Create SSE Reader
func NewSSE(r io.Reader) *SSE {
	return &SSE{
		reader: bufio.NewReader(r),
	}
}

// ReadEvent read event from stream
func (sse *SSE) ReadEvent() (*SSEEvent, error) {
	// event
	eventName, err := sse.sseMustMatchKey("event")
	if err != nil {
		return nil, err
	}
	// data
	data, err := sse.sseMustMatchKey("data")
	if err != nil {
		return nil, err
	}
	// skip line
	sse.reader.ReadLine()
	// finally
	return &SSEEvent{
		Event: eventName,
		Data:  []byte(data),
	}, nil
}

func (sse *SSE) sseMustMatchKey(key string) (string, error) {
	k, v, err := sse.sseReadData()
	if err != nil {
		return "", err
	}
	if k != key {
		return "", fmt.Errorf("key not match")
	}
	return v, nil
}
func (sse *SSE) sseReadData() (k string, v string, err error) {
	bline, _, errx := sse.reader.ReadLine()
	if err != nil {
		err = errx
		return
	}
	line := string(bline[:len(bline)])
	return splitKeyValue(line, ": ")
}
func splitKeyValue(line string, delim string) (k string, v string, err error) {
	spl := strings.SplitN(line, delim, 2)
	if len(spl) != 2 {
		err = fmt.Errorf("wrong format")
		return
	}
	k, v = spl[0], spl[1]
	return
}
