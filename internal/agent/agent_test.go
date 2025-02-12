package agent

import (
	"testing"
	"time"

	"github.com/srg-bnd/observator/internal/storage"
	"github.com/stretchr/testify/assert"
)

func TestNewAgent(t *testing.T) {
	agent := NewAgent(storage.NewMemStorage())
	assert.IsType(t, agent, &Agent{})
}

func TestStart(t *testing.T) {
	t.Logf("TODO")
}

func TestConsts(t *testing.T) {
	assert.Equal(t, defaultPollInterval, 2)
	assert.Equal(t, defaultReportInterval, 10)
}

func TestGetPollInterval(t *testing.T) {
	assert.Equal(t, GetPollInterval(), 2*time.Second)
}

func TestGetReportInterval(t *testing.T) {
	assert.Equal(t, GetReportInterval(), 10*time.Second)
}
