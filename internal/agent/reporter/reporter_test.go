package reporter

import (
	"testing"

	"github.com/srg-bnd/observator/internal/agent/client"
	"github.com/srg-bnd/observator/internal/storage"
	"github.com/stretchr/testify/assert"
)

func TestNewPoller(t *testing.T) {
	reporter := NewReporter(storage.NewMemStorage(), client.NewClient(""))
	assert.IsType(t, reporter, &Reporter{})
}

func TestStart(t *testing.T) {
	t.Logf("TODO")
}
