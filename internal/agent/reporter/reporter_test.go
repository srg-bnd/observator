package reporter

import (
	"testing"
	"time"

	"github.com/srg-bnd/observator/internal/storage"
	"github.com/stretchr/testify/assert"
)

func TestNewPoller(t *testing.T) {
	poller := NewReporter(storage.NewMemStorage())
	assert.IsType(t, poller, &Reporter{})
}

func TestStart(t *testing.T) {
	t.Logf("TODO")
}

func TestConsts(t *testing.T) {
	assert.Equal(t, defaultReportInterval, 10)
}

func TestGetReportInterval(t *testing.T) {
	assert.Equal(t, GetReportInterval(), 10*time.Second)
}
