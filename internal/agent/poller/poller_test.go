package poller

import (
	"testing"
	"time"

	"github.com/srg-bnd/observator/internal/storage"
	"github.com/stretchr/testify/assert"
)

func TestNewPoller(t *testing.T) {
	poller := NewPoller(storage.NewMemStorage())
	assert.IsType(t, poller, &Poller{})
}

func TestStart(t *testing.T) {
	t.Logf("TODO")
}

func TestConsts(t *testing.T) {
	assert.Equal(t, defaultPollInterval, 2)
}

func TestGetPollInterval(t *testing.T) {
	assert.Equal(t, GetPollInterval(), 2*time.Second)
}
