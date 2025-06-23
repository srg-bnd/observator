package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMetrics(t *testing.T) {
	metrics := NewMetrics()
	assert.IsType(t, metrics, &Metrics{})
}

func TestSetCounter(t *testing.T) {
	metrics := NewMetrics()
	metrics.SetCounter(1)

	assert.Equal(t, *metrics.Delta, int64(1))
}

func TestGetCounter(t *testing.T) {
	metrics := NewMetrics()
	metrics.SetCounter(1)

	assert.Equal(t, metrics.GetCounter(), int64(1))
}

func TestSetGauge(t *testing.T) {
	metrics := NewMetrics()
	metrics.SetGauge(1)

	assert.Equal(t, *metrics.Value, float64(1))
}

func TestGetGauge(t *testing.T) {
	metrics := NewMetrics()
	metrics.SetGauge(1)

	assert.Equal(t, metrics.GetGauge(), float64(1))
}
