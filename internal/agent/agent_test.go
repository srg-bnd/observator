package agent

import (
	"testing"

	config "github.com/srg-bnd/observator/config/agent"
	"github.com/srg-bnd/observator/internal/storage"
	"github.com/stretchr/testify/assert"
)

func TestNewAgent(t *testing.T) {
	agent := NewAgent(&config.Container{Storage: storage.NewMemStorage()})
	assert.IsType(t, agent, &Agent{})
}

func TestStart(t *testing.T) {
	t.Logf("TODO")
}
