package collector

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVars(t *testing.T) {
	assert.Equal(t, TrackedRuntimeMetrics, map[string][]string{
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

	assert.Equal(t, TrackedGopsutilMetrics, map[string][]string{
		"gauge": {
			"TotalMemory",
			"FreeMemory",
			"CPUutilization1",
		},
	})
}
