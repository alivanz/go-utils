package utils

import (
	"bytes"
	"sync"
	"testing"
)

func TestSSE(t *testing.T) {
	testCase := map[string]string{
		"test1": "123",
		"test2": "[123,456]",
	}
	// write buffer
	buf := bytes.NewBuffer(nil)
	wg := sync.WaitGroup{}
	for k, v := range testCase {
		wg.Add(1)
		sseBufWrite(buf, k, v)
	}
	t.Log(string(buf.Bytes()))
	// parse
	sse := NewSSE(buf)
	if err := sse.Listen(func(event *SSEEvent) {
		defer wg.Done()
		ref, found := testCase[event.Event]
		if !found {
			t.Log("testcase not found")
			t.Log(event)
			t.Fail()
			return
		}
		if string(event.Data) != ref {
			t.Log(event)
			t.Fail()
		}
	}); err != nil {
		t.Log(err)
		t.Fail()
	}
	wg.Wait()
}
func sseBufWrite(buf *bytes.Buffer, event, data string) {
	buf.WriteString("event: " + event + "\n")
	buf.WriteString("data: " + data + "\n")
	buf.WriteString("\n")
}
