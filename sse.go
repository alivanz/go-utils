package utils

import (
	"bufio"
	"errors"
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

// SSEListener is a listener for SSE.Listen
type SSEListener func(event *SSEEvent)

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
	b, err := sse.reader.ReadByte()
	if err != nil {
		return nil, err
	}
	if b != '\n' {
		return nil, errors.New("wrong format")
	}
	// finally
	return &SSEEvent{
		Event: eventName,
		Data:  []byte(data),
	}, nil
}

// Listen listen SSE events until EOF
func (sse *SSE) Listen(listener SSEListener) error {
	for {
		event, err := sse.ReadEvent()
		if err == io.EOF {
			return nil
		} else if err != nil {
			return err
		}
		go listener(event)
	}
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
	bline, errx := sse.reader.ReadSlice('\n')
	if errx != nil {
		err = errx
		return
	}
	line := string(bline[:len(bline)-1])
	// line := string(bline[:len(bline)])
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
