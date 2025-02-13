package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMetric(t *testing.T) {
	metric := NewMetric()
	assert.IsType(t, metric, &Metric{})
}

func TestSetCounter(t *testing.T) {
	metric := NewMetric()
	metric.SetCounter(1)

	assert.Equal(t, metric.counterValue, int64(1))
}

func TestGetCounter(t *testing.T) {
	metric := NewMetric()
	metric.SetCounter(1)

	assert.Equal(t, metric.GetCounter(), int64(1))
}

func TestSetGauge(t *testing.T) {
	metric := NewMetric()
	metric.SetGauge(1)

	assert.Equal(t, metric.gaugeValue, float64(1))
}

func TestGetGauge(t *testing.T) {
	metric := NewMetric()
	metric.SetGauge(1)

	assert.Equal(t, metric.GetGauge(), float64(1))
}
