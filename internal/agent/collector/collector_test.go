package collector

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewcollector(t *testing.T) {
	collector := NewCollector()
	assert.IsType(t, collector, &Collector{})
}

func TestGetMetrics(t *testing.T) {
	t.Logf("TODO")
}
