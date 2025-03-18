package poller

import (
	"testing"

	"github.com/srg-bnd/observator/internal/agent/collector"
	"github.com/srg-bnd/observator/internal/storage"
	"github.com/stretchr/testify/assert"
)

func TestNewPoller(t *testing.T) {
	poller := NewPoller(storage.NewMemStorage("", 0, false))
	assert.IsType(t, poller, &Poller{})
}

func TestStart(t *testing.T) {
	t.Logf("TODO")
}

func TestTrackedMetrics(t *testing.T) {
	trackedMetrics := collector.TrackedMetrics

	for _, trackedMetric := range trackedMetrics["counter"] {
		assert.Contains(t, []string{"PollCount"}, trackedMetric)
	}

	for _, trackedMetric := range trackedMetrics["gauge"] {
		t.Run(trackedMetric, func(t *testing.T) {
			assert.Contains(
				t,
				[]string{
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
				trackedMetric)
		})
	}
}
