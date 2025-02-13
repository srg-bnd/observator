package agent

import (
	"testing"

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
