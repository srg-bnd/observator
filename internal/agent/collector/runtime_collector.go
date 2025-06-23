// Collects metrics
package collector

import (
	"math/rand"
	"runtime"
)

var TrackedRuntimeMetrics = map[string][]string{
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
}

type RuntimeCollector struct {
	memStats runtime.MemStats

	pollCount int64
}

// Returns new collector
func NewRuntimeCollector() *RuntimeCollector {
	return &RuntimeCollector{
		pollCount: 0,
	}
}

// Returns current values for metrics
func (c *RuntimeCollector) GetMetrics() (*Metrics, error) {
	runtime.ReadMemStats(&c.memStats)

	return &Metrics{
		Counter: map[string]int64{
			"PollCount": c.pollCountCounterMetric(),
		},
		Gauge: map[string]float64{
			"Alloc":         c.allocGaugeMetric(),
			"BuckHashSys":   c.buckHashSysGaugeMetric(),
			"Frees":         c.freesGaugeMetric(),
			"GCCPUFraction": c.gCCPUFractionGaugeMetric(),
			"GCSys":         c.gCSysGaugeMetric(),
			"HeapAlloc":     c.heapAllocGaugeMetric(),
			"HeapIdle":      c.heapIdleGaugeMetric(),
			"HeapInuse":     c.heapInuseGaugeMetric(),
			"HeapObjects":   c.heapObjectsGaugeMetric(),
			"HeapReleased":  c.heapReleasedGaugeMetric(),
			"HeapSys":       c.heapSysGaugeMetric(),
			"LastGC":        c.lastGCGaugeMetric(),
			"Lookups":       c.lookupsGaugeMetric(),
			"MCacheInuse":   c.mCacheInuseGaugeMetric(),
			"MCacheSys":     c.mCacheSysGaugeMetric(),
			"MSpanInuse":    c.mSpanInuseGaugeMetric(),
			"MSpanSys":      c.mSpanSysGaugeMetric(),
			"Mallocs":       c.mallocsGaugeMetric(),
			"NextGC":        c.nextGCGaugeMetric(),
			"NumForcedGC":   c.numForcedGCGaugeMetric(),
			"NumGC":         c.numGCGaugeMetric(),
			"OtherSys":      c.otherSysGaugeMetric(),
			"PauseTotalNs":  c.pauseTotalNsGaugeMetric(),
			"StackInuse":    c.stackInuseGaugeMetric(),
			"StackSys":      c.stackSysGaugeMetric(),
			"Sys":           c.sysGaugeMetric(),
			"TotalAlloc":    c.totalAllocGaugeMetric(),
			"RandomValue":   c.randomValueGaugeMetric(),
		},
	}, nil
}

// Counter Metrics

func (c *RuntimeCollector) pollCountCounterMetric() int64 {
	c.pollCount++
	return c.pollCount
}

// Gauge Metrics

func (c *RuntimeCollector) allocGaugeMetric() float64 {
	return float64(c.memStats.Alloc)
}

func (c *RuntimeCollector) buckHashSysGaugeMetric() float64 {
	return float64(c.memStats.BuckHashSys)
}

func (c *RuntimeCollector) freesGaugeMetric() float64 {
	return float64(c.memStats.Frees)
}

func (c *RuntimeCollector) gCCPUFractionGaugeMetric() float64 {
	return float64(c.memStats.GCCPUFraction)
}

func (c *RuntimeCollector) gCSysGaugeMetric() float64 {
	return float64(c.memStats.GCSys)
}

func (c *RuntimeCollector) heapAllocGaugeMetric() float64 {
	return float64(c.memStats.HeapAlloc)
}

func (c *RuntimeCollector) heapIdleGaugeMetric() float64 {
	return float64(c.memStats.HeapIdle)
}

func (c *RuntimeCollector) heapInuseGaugeMetric() float64 {
	return float64(c.memStats.HeapInuse)
}

func (c *RuntimeCollector) heapObjectsGaugeMetric() float64 {
	return float64(c.memStats.HeapObjects)
}

func (c *RuntimeCollector) heapReleasedGaugeMetric() float64 {
	return float64(c.memStats.HeapReleased)
}

func (c *RuntimeCollector) heapSysGaugeMetric() float64 {
	return float64(c.memStats.HeapSys)
}

func (c *RuntimeCollector) lastGCGaugeMetric() float64 {
	return float64(c.memStats.LastGC)
}

func (c *RuntimeCollector) lookupsGaugeMetric() float64 {
	return float64(c.memStats.Lookups)
}

func (c *RuntimeCollector) mCacheInuseGaugeMetric() float64 {
	return float64(c.memStats.MCacheInuse)
}

func (c *RuntimeCollector) mCacheSysGaugeMetric() float64 {
	return float64(c.memStats.MCacheSys)
}

func (c *RuntimeCollector) mSpanInuseGaugeMetric() float64 {
	return float64(c.memStats.MSpanInuse)
}

func (c *RuntimeCollector) mSpanSysGaugeMetric() float64 {
	return float64(c.memStats.MSpanSys)
}

func (c *RuntimeCollector) mallocsGaugeMetric() float64 {
	return float64(c.memStats.Mallocs)
}

func (c *RuntimeCollector) nextGCGaugeMetric() float64 {
	return float64(c.memStats.NextGC)
}

func (c *RuntimeCollector) numForcedGCGaugeMetric() float64 {
	return float64(c.memStats.NumForcedGC)
}

func (c *RuntimeCollector) numGCGaugeMetric() float64 {
	return float64(c.memStats.NumGC)
}

func (c *RuntimeCollector) otherSysGaugeMetric() float64 {
	return float64(c.memStats.OtherSys)
}

func (c *RuntimeCollector) pauseTotalNsGaugeMetric() float64 {
	return float64(c.memStats.PauseTotalNs)
}

func (c *RuntimeCollector) stackInuseGaugeMetric() float64 {
	return float64(c.memStats.StackInuse)
}

func (c *RuntimeCollector) stackSysGaugeMetric() float64 {
	return float64(c.memStats.StackSys)
}

func (c *RuntimeCollector) sysGaugeMetric() float64 {
	return float64(c.memStats.Sys)
}

func (c *RuntimeCollector) totalAllocGaugeMetric() float64 {
	return float64(c.memStats.TotalAlloc)
}

func (c *RuntimeCollector) randomValueGaugeMetric() float64 {
	return rand.ExpFloat64()
}
