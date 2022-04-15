package utils

import (
	"testing"
)

func TestPassphrase(t *testing.T) {
	passphrase := MakePassphrase()
	if len(passphrase) != 32 {
		t.Fatalf(`passphrase %s has length %d`, passphrase, len(passphrase))
	}
}
