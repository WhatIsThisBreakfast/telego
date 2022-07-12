package telego

import (
	"os"
	"testing"
)

const (
	token = "5404543726:AAHFy9XubKviQZYrvUwuK1NHKbRsiS33NL8"
)

var (
	tlg *Telego
)

func TestMain(m *testing.M) {
	tlg = NewTelego(token)
	os.Exit(m.Run())
}
