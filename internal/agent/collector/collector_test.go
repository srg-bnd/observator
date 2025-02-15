package collector

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

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

func TestNewCollector(t *testing.T) {
	collector := NewCollector()
	assert.IsType(t, collector, &Collector{})
}

func TestGetMetrics(t *testing.T) {
	t.Logf("TODO")
}

// Counter Metrics

func TestPollCountCounterMetric(t *testing.T) {
	c := NewCollector()
	assert.Equal(t, c.pollCountCounterMetric(), int64(1))
}

// Gauge Metrics

func TestAllocGaugeMetric(t *testing.T) {
	c := NewCollector()
	runtime.ReadMemStats(&c.memStats)

	assert.Equal(t, c.allocGaugeMetric(), float64(c.memStats.Alloc))
}

func TestBuckHashSysGaugeMetric(t *testing.T) {
	c := NewCollector()
	runtime.ReadMemStats(&c.memStats)

	assert.Equal(t, c.buckHashSysGaugeMetric(), float64(c.memStats.BuckHashSys))
}

func TestFreesGaugeMetric(t *testing.T) {
	c := NewCollector()
	runtime.ReadMemStats(&c.memStats)

	assert.Equal(t, c.freesGaugeMetric(), float64(c.memStats.Frees))
}

func TestGCCPUFractionGaugeMetric(t *testing.T) {
	c := NewCollector()
	runtime.ReadMemStats(&c.memStats)

	assert.Equal(t, c.gCCPUFractionGaugeMetric(), float64(c.memStats.GCCPUFraction))
}

func TestGCSysGaugeMetric(t *testing.T) {
	c := NewCollector()
	runtime.ReadMemStats(&c.memStats)

	assert.Equal(t, c.gCSysGaugeMetric(), float64(c.memStats.GCSys))
}

func TestHeapAllocGaugeMetric(t *testing.T) {
	c := NewCollector()
	runtime.ReadMemStats(&c.memStats)

	assert.Equal(t, c.heapAllocGaugeMetric(), float64(c.memStats.HeapAlloc))
}

func TestHeapIdleGaugeMetric(t *testing.T) {
	c := NewCollector()
	runtime.ReadMemStats(&c.memStats)

	assert.Equal(t, c.heapIdleGaugeMetric(), float64(c.memStats.HeapIdle))
}

func TestHeapInuseGaugeMetric(t *testing.T) {
	c := NewCollector()
	runtime.ReadMemStats(&c.memStats)

	assert.Equal(t, c.heapInuseGaugeMetric(), float64(c.memStats.HeapInuse))
}

func TestHeapObjectsGaugeMetric(t *testing.T) {
	c := NewCollector()
	runtime.ReadMemStats(&c.memStats)

	assert.Equal(t, c.heapObjectsGaugeMetric(), float64(c.memStats.HeapObjects))
}

func TestHeapReleasedGaugeMetric(t *testing.T) {
	c := NewCollector()
	runtime.ReadMemStats(&c.memStats)

	assert.Equal(t, c.heapReleasedGaugeMetric(), float64(c.memStats.HeapReleased))
}

func TestHeapSysGaugeMetric(t *testing.T) {
	c := NewCollector()
	runtime.ReadMemStats(&c.memStats)

	assert.Equal(t, c.heapSysGaugeMetric(), float64(c.memStats.HeapSys))
}

func TestLastGCGaugeMetric(t *testing.T) {
	c := NewCollector()
	runtime.ReadMemStats(&c.memStats)

	assert.Equal(t, c.lastGCGaugeMetric(), float64(c.memStats.LastGC))
}

func TestLookupsGaugeMetric(t *testing.T) {
	c := NewCollector()
	runtime.ReadMemStats(&c.memStats)

	assert.Equal(t, c.lookupsGaugeMetric(), float64(c.memStats.Lookups))
}

func TestMCacheInuseGaugeMetric(t *testing.T) {
	c := NewCollector()
	runtime.ReadMemStats(&c.memStats)

	assert.Equal(t, c.mCacheInuseGaugeMetric(), float64(c.memStats.MCacheInuse))
}

func TestMCacheSysGaugeMetric(t *testing.T) {
	c := NewCollector()
	runtime.ReadMemStats(&c.memStats)

	assert.Equal(t, c.mCacheSysGaugeMetric(), float64(c.memStats.MCacheSys))
}

func TestMSpanInuseGaugeMetric(t *testing.T) {
	c := NewCollector()
	runtime.ReadMemStats(&c.memStats)

	assert.Equal(t, c.mSpanInuseGaugeMetric(), float64(c.memStats.MSpanInuse))
}

func TestMSpanSysGaugeMetric(t *testing.T) {
	c := NewCollector()
	runtime.ReadMemStats(&c.memStats)

	assert.Equal(t, c.mSpanSysGaugeMetric(), float64(c.memStats.MSpanSys))
}

func TestMallocsGaugeMetric(t *testing.T) {
	c := NewCollector()
	runtime.ReadMemStats(&c.memStats)

	assert.Equal(t, c.mallocsGaugeMetric(), float64(c.memStats.Mallocs))
}

func TestNextGCGaugeMetric(t *testing.T) {
	c := NewCollector()
	runtime.ReadMemStats(&c.memStats)

	assert.Equal(t, c.nextGCGaugeMetric(), float64(c.memStats.NextGC))
}

func TestNumForcedGCGaugeMetric(t *testing.T) {
	c := NewCollector()
	runtime.ReadMemStats(&c.memStats)

	assert.Equal(t, c.numForcedGCGaugeMetric(), float64(c.memStats.NumForcedGC))
}

func TestNumGCGaugeMetric(t *testing.T) {
	c := NewCollector()
	runtime.ReadMemStats(&c.memStats)

	assert.Equal(t, c.numGCGaugeMetric(), float64(c.memStats.NumGC))
}

func TestOtherSysGaugeMetric(t *testing.T) {
	c := NewCollector()
	runtime.ReadMemStats(&c.memStats)

	assert.Equal(t, c.otherSysGaugeMetric(), float64(c.memStats.OtherSys))
}

func TestPauseTotalNsGaugeMetric(t *testing.T) {
	c := NewCollector()
	runtime.ReadMemStats(&c.memStats)

	assert.Equal(t, c.pauseTotalNsGaugeMetric(), float64(c.memStats.PauseTotalNs))
}

func TestStackInuseGaugeMetric(t *testing.T) {
	c := NewCollector()
	runtime.ReadMemStats(&c.memStats)

	assert.Equal(t, c.stackInuseGaugeMetric(), float64(c.memStats.StackInuse))
}

func TestStackSysGaugeMetric(t *testing.T) {
	c := NewCollector()
	runtime.ReadMemStats(&c.memStats)

	assert.Equal(t, c.stackSysGaugeMetric(), float64(c.memStats.StackSys))
}

func TestSysGaugeMetric(t *testing.T) {
	c := NewCollector()
	runtime.ReadMemStats(&c.memStats)

	assert.Equal(t, c.sysGaugeMetric(), float64(c.memStats.Sys))
}

func TestTotalAllocGaugeMetric(t *testing.T) {
	c := NewCollector()
	runtime.ReadMemStats(&c.memStats)

	assert.Equal(t, c.totalAllocGaugeMetric(), float64(c.memStats.TotalAlloc))
}

func TestRandomValueGaugeMetric(t *testing.T) {
	c := NewCollector()
	runtime.ReadMemStats(&c.memStats)

	last := c.randomValueGaugeMetric()
	if assert.IsType(t, last, float64(last)) {
		assert.NotEqual(t, last, c.randomValueGaugeMetric())
	}
}
