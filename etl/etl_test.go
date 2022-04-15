package etl

import (
	"path/filepath"
	"testing"
)

func TestLoadMessageFromJson(t *testing.T) {
	want := Message{Body: "Test message"}
	filename := filepath.Join("../inputs", "message_one.json")
	msg := ReadData(filename)
	if msg.Body != want.Body {
		t.Fatalf(`{Received %s, wanted %s}`, msg.Body, want.Body)
	}
}
