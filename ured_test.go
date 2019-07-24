package ured

import (
	"testing"

	"github.com/michelia/ulog"
)

func TestNew(t *testing.T) {
	c := Config{
		Addr: "127.0.0.1:6379",
	}
	New(ulog.NewConsole(), c)
}
