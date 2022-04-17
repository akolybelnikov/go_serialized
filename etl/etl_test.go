package etl

import (
	"github.com/akolybelnikov/go_serialized/utils"
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

func TestEncryptMessage(t *testing.T) {
	filename := filepath.Join("../inputs", "message_one.json")
	msg := ReadData(filename)
	passphrase := utils.MakePassphrase()
	e := msg.EncodeData(passphrase)
	if len(e.Message) == 0 || len(e.Passphrase) == 0 {
		t.Fatalf("encrypted message has empty fields %v", e)
	}
}

func TestDecryptMessage(t *testing.T) {
	filename := filepath.Join("../inputs", "message_one.json")
	msg := ReadData(filename)
	passphrase := utils.MakePassphrase()
	e := msg.EncodeData(passphrase)
	de := e.DecodeData()
	if msg.Body != de {
		t.Fatalf(`{Received %s, wanted %s}`, de, msg.Body)
	}
}
