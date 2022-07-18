package tggo

import (
	"os"
	"testing"
)

const (
	token = "5404543726:AAHFy9XubKviQZYrvUwuK1NHKbRsiS33NL8"
)

var (
	tlg *Tggo
)

func TestMain(m *testing.M) {
	tlg = NewTggo(token)
	os.Exit(m.Run())
}
