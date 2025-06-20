// Collects metrics
package collector

import (
	"math/rand"
	"runtime"

	"github.com/shirou/gopsutil/v4/mem"
)

// Collector
type Collector struct {
	memStats runtime.MemStats

	pollCount int64
}

// Metrics
type Metrics struct {
	Counter    map[string]int64
	Gauge      map[string]float64
	GaugeMatch map[string]float64
}

// Tracked metrics
var TrackedMetrics = map[string][]string{
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

		"TotalMemory",
		"FreeMemory",
	},
}

// Returns new collector
func NewCollector() *Collector {
	return &Collector{
		pollCount: 0,
	}
}

// Returns current runtime metrics
func (c *Collector) GetRuntimeMetrics() (*Metrics, error) {
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

// Returns current gops (gopsutil) metrics
func (c *Collector) GetGopsutilMetrics() (*Metrics, error) {
	virtualMemoryStat, _ := mem.VirtualMemory()

	return &Metrics{
		Gauge: map[string]float64{
			"TotalMemory": float64(virtualMemoryStat.Total),
			"FreeMemory":  float64(virtualMemoryStat.Free),
		}}, nil
}

// Counter Metrics

func (c *Collector) pollCountCounterMetric() int64 {
	c.pollCount++
	return c.pollCount
}

// Gauge Metrics

func (c *Collector) allocGaugeMetric() float64 {
	return float64(c.memStats.Alloc)
}

func (c *Collector) buckHashSysGaugeMetric() float64 {
	return float64(c.memStats.BuckHashSys)
}

func (c *Collector) freesGaugeMetric() float64 {
	return float64(c.memStats.Frees)
}

func (c *Collector) gCCPUFractionGaugeMetric() float64 {
	return float64(c.memStats.GCCPUFraction)
}

func (c *Collector) gCSysGaugeMetric() float64 {
	return float64(c.memStats.GCSys)
}

func (c *Collector) heapAllocGaugeMetric() float64 {
	return float64(c.memStats.HeapAlloc)
}

func (c *Collector) heapIdleGaugeMetric() float64 {
	return float64(c.memStats.HeapIdle)
}

func (c *Collector) heapInuseGaugeMetric() float64 {
	return float64(c.memStats.HeapInuse)
}

func (c *Collector) heapObjectsGaugeMetric() float64 {
	return float64(c.memStats.HeapObjects)
}

func (c *Collector) heapReleasedGaugeMetric() float64 {
	return float64(c.memStats.HeapReleased)
}

func (c *Collector) heapSysGaugeMetric() float64 {
	return float64(c.memStats.HeapSys)
}

func (c *Collector) lastGCGaugeMetric() float64 {
	return float64(c.memStats.LastGC)
}

func (c *Collector) lookupsGaugeMetric() float64 {
	return float64(c.memStats.Lookups)
}

func (c *Collector) mCacheInuseGaugeMetric() float64 {
	return float64(c.memStats.MCacheInuse)
}

func (c *Collector) mCacheSysGaugeMetric() float64 {
	return float64(c.memStats.MCacheSys)
}

func (c *Collector) mSpanInuseGaugeMetric() float64 {
	return float64(c.memStats.MSpanInuse)
}

func (c *Collector) mSpanSysGaugeMetric() float64 {
	return float64(c.memStats.MSpanSys)
}

func (c *Collector) mallocsGaugeMetric() float64 {
	return float64(c.memStats.Mallocs)
}

func (c *Collector) nextGCGaugeMetric() float64 {
	return float64(c.memStats.NextGC)
}

func (c *Collector) numForcedGCGaugeMetric() float64 {
	return float64(c.memStats.NumForcedGC)
}

func (c *Collector) numGCGaugeMetric() float64 {
	return float64(c.memStats.NumGC)
}

func (c *Collector) otherSysGaugeMetric() float64 {
	return float64(c.memStats.OtherSys)
}

func (c *Collector) pauseTotalNsGaugeMetric() float64 {
	return float64(c.memStats.PauseTotalNs)
}

func (c *Collector) stackInuseGaugeMetric() float64 {
	return float64(c.memStats.StackInuse)
}

func (c *Collector) stackSysGaugeMetric() float64 {
	return float64(c.memStats.StackSys)
}

func (c *Collector) sysGaugeMetric() float64 {
	return float64(c.memStats.Sys)
}

func (c *Collector) totalAllocGaugeMetric() float64 {
	return float64(c.memStats.TotalAlloc)
}

func (c *Collector) randomValueGaugeMetric() float64 {
	return rand.ExpFloat64()
}
