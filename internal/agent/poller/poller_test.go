package poller

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPoller(t *testing.T) {
	poller := NewPoller()
	assert.IsType(t, poller, &Poller{})
}

func TestPoll(t *testing.T) {
	t.Logf("TODO")
}
