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

func TestVars(t *testing.T) {
	assert.Equal(t, TrackedMetrics, map[string][]string{
		"counter": {
			"PollCount",
		},
		"gauge": {
			"Alloc",
			"BuckHashSys",
			"Frees",
			"GCCPUFraction",
			"GCSys",
			"HeapAlloc",
			"HeapIdle",
			"HeapInuse",
			"HeapObjects",
			"HeapReleased",
			"HeapSys",
			"LastGC",
			"Lookups",
			"MCacheInuse",
			"MCacheSys",
			"MSpanInuse",
			"MSpanSys",
			"Mallocs",
			"NextGC",
			"NumForcedGC",
			"NumGC",
			"OtherSys",
			"PauseTotalNs",
			"StackInuse",
			"StackSys",
			"Sys",
			"TotalAlloc",
			"RandomValue",
		},
	})
}
