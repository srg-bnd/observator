package reporter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPoller(t *testing.T) {
	poller := NewReporter()
	assert.IsType(t, poller, &Reporter{})
}

func TestReport(t *testing.T) {
	t.Logf("TODO")
}
