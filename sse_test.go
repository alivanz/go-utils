package utils

import (
	"bytes"
	"testing"
)

func TestSSE(t *testing.T) {
	testCase := map[string]string{
		"test1": "123",
		"test2": "[123,456]",
	}
	// write buffer
	buf := bytes.NewBuffer(nil)
	for k, v := range testCase {
		sseBufWrite(buf, k, v)
	}
	t.Log(string(buf.Bytes()))
	// parse
	sse := NewSSE(buf)
	for k, v := range testCase {
		seeTestMustMatch(t, sse, k, v)
	}
}
func sseBufWrite(buf *bytes.Buffer, event, data string) {
	buf.WriteString("event: " + event + "\n")
	buf.WriteString("data: " + data + "\n")
	buf.WriteString("\n")
}
func seeTestMustMatch(t *testing.T, sse *SSE, event, data string) {
	e, err := sse.ReadEvent()
	if err != nil {
		panic(err)
	}
	switch true {
	case e.Event == event:
	case string(e.Data) == data:
	default:
		t.Log(e)
		t.Fail()
	}
}
