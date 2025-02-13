package collector

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCollector(t *testing.T) {
	trackedMetrics := NewTrackedMetrics([]string{}, []string{})
	collector := NewCollector(trackedMetrics)
	assert.IsType(t, collector, &Collector{})
}

func TestNewTrackedMetrics(t *testing.T) {
	trackedMetrics := NewTrackedMetrics([]string{}, []string{})
	assert.IsType(t, trackedMetrics, &TrackedMetrics{})
}

func TestNewCounterMetric(t *testing.T) {
	counterMetric := NewCounterMetric("metric", 1.0)
	assert.IsType(t, counterMetric, &CounterMetric{})
}

func TestNewGaugeMetric(t *testing.T) {
	gaugeMetric := NewGaugeMetric("metric", 1.0)
	assert.IsType(t, gaugeMetric, &GaugeMetric{})
}

func TestNewMetrics(t *testing.T) {
	metrics := NewMetrics(
		[]CounterMetric{*NewCounterMetric("metric", 1)},
		[]GaugeMetric{*NewGaugeMetric("metric", 1.0)},
	)
	assert.IsType(t, metrics, &Metrics{})
}

func TestGetMetrics(t *testing.T) {
	t.Logf("TODO")
}
