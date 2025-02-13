package collector

import (
	"math/rand"
	"runtime"
	"slices"
)

type Collector struct {
	trackedMetrics *TrackedMetrics
	memStats       runtime.MemStats
}

func NewCollector(trackedMetrics *TrackedMetrics) *Collector {
	return &Collector{
		trackedMetrics: trackedMetrics,
	}
}

type TrackedMetrics struct {
	Counter []string
	Gauge   []string
}

func NewTrackedMetrics(counter []string, gauge []string) *TrackedMetrics {
	return &TrackedMetrics{
		Counter: counter,
		Gauge:   gauge,
	}
}

type CounterMetric struct {
	Name  string
	Value int64
}

func NewCounterMetric(name string, value int64) *CounterMetric {
	return &CounterMetric{
		Name:  name,
		Value: value,
	}
}

type GaugeMetric struct {
	Name  string
	Value float64
}

func NewGaugeMetric(name string, value float64) *GaugeMetric {
	return &GaugeMetric{
		Name:  name,
		Value: value,
	}
}

type Metrics struct {
	Counter []CounterMetric
	Gauge   []GaugeMetric
}

func NewMetrics(counter []CounterMetric, gauge []GaugeMetric) *Metrics {
	return &Metrics{
		Counter: counter,
		Gauge:   gauge,
	}
}

func (c *Collector) GetMetrics() (*Metrics, error) {
	counterMetrics := make([]CounterMetric, 0)
	gaugeMetrics := make([]GaugeMetric, 0)

	runtime.ReadMemStats(&c.memStats)

	if slices.Contains(c.trackedMetrics.Counter, "PollCount") {
		counterMetrics = append(counterMetrics, *NewCounterMetric("PollCount", c.pollCountCounterMetric()))
	}

	if slices.Contains(c.trackedMetrics.Gauge, "Alloc") {
		gaugeMetrics = append(gaugeMetrics, *NewGaugeMetric("Alloc", c.allocGaugeMetric()))
	}

	if slices.Contains(c.trackedMetrics.Gauge, "BuckHashSys") {
		gaugeMetrics = append(gaugeMetrics, *NewGaugeMetric("BuckHashSys", c.buckHashSysGaugeMetric()))
	}

	if slices.Contains(c.trackedMetrics.Gauge, "Frees") {
		gaugeMetrics = append(gaugeMetrics, *NewGaugeMetric("Frees", c.freesGaugeMetric()))
	}

	if slices.Contains(c.trackedMetrics.Gauge, "GCCPUFraction") {
		gaugeMetrics = append(gaugeMetrics, *NewGaugeMetric("GCCPUFraction", c.gCCPUFractionGaugeMetric()))
	}

	if slices.Contains(c.trackedMetrics.Gauge, "GCSys") {
		gaugeMetrics = append(gaugeMetrics, *NewGaugeMetric("GCSys", c.gCSysGaugeMetric()))
	}

	if slices.Contains(c.trackedMetrics.Gauge, "HeapAlloc") {
		gaugeMetrics = append(gaugeMetrics, *NewGaugeMetric("HeapAlloc", c.heapAllocGaugeMetric()))
	}

	if slices.Contains(c.trackedMetrics.Gauge, "HeapIdle") {
		gaugeMetrics = append(gaugeMetrics, *NewGaugeMetric("HeapIdle", c.heapIdleGaugeMetric()))
	}

	if slices.Contains(c.trackedMetrics.Gauge, "HeapInuse") {
		gaugeMetrics = append(gaugeMetrics, *NewGaugeMetric("HeapInuse", c.heapInuseGaugeMetric()))
	}

	if slices.Contains(c.trackedMetrics.Gauge, "HeapObjects") {
		gaugeMetrics = append(gaugeMetrics, *NewGaugeMetric("HeapObjects", c.heapObjectsGaugeMetric()))
	}

	if slices.Contains(c.trackedMetrics.Gauge, "HeapReleased") {
		gaugeMetrics = append(gaugeMetrics, *NewGaugeMetric("HeapReleased", c.heapReleasedGaugeMetric()))
	}

	if slices.Contains(c.trackedMetrics.Gauge, "HeapSys") {
		gaugeMetrics = append(gaugeMetrics, *NewGaugeMetric("HeapSys", c.heapSysGaugeMetric()))
	}

	if slices.Contains(c.trackedMetrics.Gauge, "LastGC") {
		gaugeMetrics = append(gaugeMetrics, *NewGaugeMetric("LastGC", c.lastGCGaugeMetric()))
	}

	if slices.Contains(c.trackedMetrics.Gauge, "Lookups") {
		gaugeMetrics = append(gaugeMetrics, *NewGaugeMetric("Lookups", c.lookupsGaugeMetric()))
	}

	if slices.Contains(c.trackedMetrics.Gauge, "MCacheInuse") {
		gaugeMetrics = append(gaugeMetrics, *NewGaugeMetric("MCacheInuse", c.mCacheInuseGaugeMetric()))
	}

	if slices.Contains(c.trackedMetrics.Gauge, "MCacheSys") {
		gaugeMetrics = append(gaugeMetrics, *NewGaugeMetric("MCacheSys", c.mCacheSysGaugeMetric()))
	}

	if slices.Contains(c.trackedMetrics.Gauge, "MSpanInuse") {
		gaugeMetrics = append(gaugeMetrics, *NewGaugeMetric("MSpanInuse", c.mSpanInuseGaugeMetric()))
	}

	if slices.Contains(c.trackedMetrics.Gauge, "MSpanSys") {
		gaugeMetrics = append(gaugeMetrics, *NewGaugeMetric("MSpanSys", c.mSpanSysGaugeMetric()))
	}

	if slices.Contains(c.trackedMetrics.Gauge, "Mallocs") {
		gaugeMetrics = append(gaugeMetrics, *NewGaugeMetric("Mallocs", c.mallocsGaugeMetric()))
	}

	if slices.Contains(c.trackedMetrics.Gauge, "NextGC") {
		gaugeMetrics = append(gaugeMetrics, *NewGaugeMetric("NextGC", c.nextGCGaugeMetric()))
	}

	if slices.Contains(c.trackedMetrics.Gauge, "NumForcedGC") {
		gaugeMetrics = append(gaugeMetrics, *NewGaugeMetric("NumForcedGC", c.numForcedGCGaugeMetric()))
	}

	if slices.Contains(c.trackedMetrics.Gauge, "NumGC") {
		gaugeMetrics = append(gaugeMetrics, *NewGaugeMetric("NumGC", c.numGCGaugeMetric()))
	}

	if slices.Contains(c.trackedMetrics.Gauge, "OtherSys") {
		gaugeMetrics = append(gaugeMetrics, *NewGaugeMetric("OtherSys", c.otherSysGaugeMetric()))
	}

	if slices.Contains(c.trackedMetrics.Gauge, "PauseTotalNs") {
		gaugeMetrics = append(gaugeMetrics, *NewGaugeMetric("PauseTotalNs", c.pauseTotalNsGaugeMetric()))
	}

	if slices.Contains(c.trackedMetrics.Gauge, "StackInuse") {
		gaugeMetrics = append(gaugeMetrics, *NewGaugeMetric("StackInuse", c.stackInuseGaugeMetric()))
	}

	if slices.Contains(c.trackedMetrics.Gauge, "StackSys") {
		gaugeMetrics = append(gaugeMetrics, *NewGaugeMetric("StackSys", c.stackSysGaugeMetric()))
	}

	if slices.Contains(c.trackedMetrics.Gauge, "Sys") {
		gaugeMetrics = append(gaugeMetrics, *NewGaugeMetric("Sys", c.sysGaugeMetric()))
	}

	if slices.Contains(c.trackedMetrics.Gauge, "TotalAlloc") {
		gaugeMetrics = append(gaugeMetrics, *NewGaugeMetric("TotalAlloc", c.totalAllocGaugeMetric()))
	}

	if slices.Contains(c.trackedMetrics.Gauge, "RandomValue") {
		gaugeMetrics = append(gaugeMetrics, *NewGaugeMetric("RandomValue", c.randomValueGaugeMetric()))
	}

	return NewMetrics(counterMetrics, gaugeMetrics), nil
}

// Counter Metrics

func (c *Collector) pollCountCounterMetric() int64 {
	return 1
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
