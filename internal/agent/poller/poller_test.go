package poller

import (
	"fmt"
	"testing"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/srg-bnd/observator/internal/agent/collector"
	"github.com/srg-bnd/observator/internal/storage"
	"github.com/stretchr/testify/assert"
)

func TestNewPoller(t *testing.T) {
	poller := NewPoller(storage.NewMemStorage())
	assert.IsType(t, poller, &Poller{})
}

func TestStart(t *testing.T) {
	t.Logf("TODO")
}

func TestTrackedRuntimeMetrics(t *testing.T) {
	trackedMetrics := collector.TrackedRuntimeMetrics

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
					"TotalAlloc",
					"RandomValue",
				},
				trackedMetric)
		})
	}
}

func TestTrackedGopsutilMetrics(t *testing.T) {
	trackedMetrics := collector.TrackedGopsutilMetrics

	for _, trackedMetric := range trackedMetrics["counter"] {
		assert.Contains(t, []string{"PollCount"}, trackedMetric)
	}

	countOfCPU, _ := cpu.Counts(true)
	expectedMetrics := []string{
		"TotalMemory",
		"FreeMemory",
		"CPUutilization1",
	}
	for numberCPU := 1; numberCPU < countOfCPU; numberCPU++ {
		expectedMetrics = append(expectedMetrics, fmt.Sprint("CPUutilization", numberCPU+1))
	}

	for _, trackedMetric := range trackedMetrics["gauge"] {
		t.Run(trackedMetric, func(t *testing.T) {
			assert.Contains(t, expectedMetrics, trackedMetric)
		})
	}
}

func Test_pollRuntimeMetrics(t *testing.T) {
	t.Logf("TODO")
}

func Test_pollGopsutilMetrics(t *testing.T) {
	t.Logf("TODO")
}
