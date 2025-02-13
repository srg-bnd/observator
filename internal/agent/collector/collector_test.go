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
	counterMetric := NewCounterMetric("metric", 1)
	assert.IsType(t, counterMetric, &CounterMetric{})
}

func TestNewGaugeMetric(t *testing.T) {
	gaugeMetric := NewGaugeMetric("metric", 1)
	assert.IsType(t, gaugeMetric, &GaugeMetric{})
}

func TestNewMetrics(t *testing.T) {
	metrics := NewMetrics(
		[]CounterMetric{*NewCounterMetric("metric", 1)},
		[]GaugeMetric{*NewGaugeMetric("metric", 1)},
	)
	assert.IsType(t, metrics, &Metrics{})
}

func TestGetMetrics(t *testing.T) {
	t.Logf("TODO")
}

// Counter Metrics

func TestPollCountCounterMetric(t *testing.T) {
	t.Logf("TODO")
}

// Gauge Metrics

func TestAllocGaugeMetric(t *testing.T) {
	t.Logf("TODO")
}

func TestBuckHashSysGaugeMetric(t *testing.T) {
	t.Logf("TODO")
}

func TestFreesGaugeMetric(t *testing.T) {
	t.Logf("TODO")
}

func TestGCCPUFractionGaugeMetric(t *testing.T) {
	t.Logf("TODO")
}

func TestGCSysGaugeMetric(t *testing.T) {
	t.Logf("TODO")
}

func TestHeapAllocGaugeMetric(t *testing.T) {
	t.Logf("TODO")
}

func TestheapIdleGaugeMetric(t *testing.T) {
	t.Logf("TODO")
}

func TestHeapInuseGaugeMetric(t *testing.T) {
	t.Logf("TODO")
}

func TestHeapObjectsGaugeMetric(t *testing.T) {
	t.Logf("TODO")
}

func TestHeapReleasedGaugeMetric(t *testing.T) {
	t.Logf("TODO")
}

func TestHeapSysGaugeMetric(t *testing.T) {
	t.Logf("TODO")
}

func TestLastGCGaugeMetric(t *testing.T) {
	t.Logf("TODO")
}

func TestLookupsGaugeMetric(t *testing.T) {
	t.Logf("TODO")
}

func TestMCacheInuseGaugeMetric(t *testing.T) {
	t.Logf("TODO")
}

func TestMCacheSysGaugeMetric(t *testing.T) {
	t.Logf("TODO")
}

func TestMSpanInuseGaugeMetric(t *testing.T) {
	t.Logf("TODO")
}

func TestMSpanSysGaugeMetric(t *testing.T) {
	t.Logf("TODO")
}

func TestMallocsGaugeMetric(t *testing.T) {
	t.Logf("TODO")
}

func TestNextGCGaugeMetric(t *testing.T) {
	t.Logf("TODO")
}

func TestNumForcedGCGaugeMetric(t *testing.T) {
	t.Logf("TODO")
}

func TestNumGCGaugeMetric(t *testing.T) {
	t.Logf("TODO")
}

func TestOtherSysGaugeMetric(t *testing.T) {
	t.Logf("TODO")
}

func TestPauseTotalNsGaugeMetric(t *testing.T) {
	t.Logf("TODO")
}

func TestStackInuseGaugeMetric(t *testing.T) {
	t.Logf("TODO")
}

func TestStackSysGaugeMetric(t *testing.T) {
	t.Logf("TODO")
}

func TestSysGaugeMetric(t *testing.T) {
	t.Logf("TODO")
}

func TestTotalAllocGaugeMetric(t *testing.T) {
	t.Logf("TODO")
}

func TestRandomValueGaugeMetric(t *testing.T) {
	t.Logf("TODO")
}
