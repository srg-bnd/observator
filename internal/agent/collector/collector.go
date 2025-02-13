package collector

import (
	"math/rand"
	"runtime"
	"slices"
)

type Collector struct {
	trackedMetrics *TrackedMetrics
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

	if slices.Contains(c.trackedMetrics.Counter, "PollCount") {
		counterMetrics = append(counterMetrics, *NewCounterMetric("PollCount", pollCountCounterMetric()))
	}

	if slices.Contains(c.trackedMetrics.Gauge, "Alloc") {
		gaugeMetrics = append(gaugeMetrics, *NewGaugeMetric("Alloc", allocGaugeMetric()))
	}

	if slices.Contains(c.trackedMetrics.Gauge, "BuckHashSys") {
		gaugeMetrics = append(gaugeMetrics, *NewGaugeMetric("BuckHashSys", buckHashSysGaugeMetric()))
	}

	if slices.Contains(c.trackedMetrics.Gauge, "Frees") {
		gaugeMetrics = append(gaugeMetrics, *NewGaugeMetric("Frees", freesGaugeMetric()))
	}

	if slices.Contains(c.trackedMetrics.Gauge, "GCCPUFraction") {
		gaugeMetrics = append(gaugeMetrics, *NewGaugeMetric("GCCPUFraction", gCCPUFractionGaugeMetric()))
	}

	if slices.Contains(c.trackedMetrics.Gauge, "GCSys") {
		gaugeMetrics = append(gaugeMetrics, *NewGaugeMetric("GCSys", gCSysGaugeMetric()))
	}

	if slices.Contains(c.trackedMetrics.Gauge, "HeapAlloc") {
		gaugeMetrics = append(gaugeMetrics, *NewGaugeMetric("HeapAlloc", heapAllocGaugeMetric()))
	}

	if slices.Contains(c.trackedMetrics.Gauge, "HeapIdle") {
		gaugeMetrics = append(gaugeMetrics, *NewGaugeMetric("HeapIdle", heapIdleGaugeMetric()))
	}

	if slices.Contains(c.trackedMetrics.Gauge, "HeapInuse") {
		gaugeMetrics = append(gaugeMetrics, *NewGaugeMetric("HeapInuse", heapInuseGaugeMetric()))
	}

	if slices.Contains(c.trackedMetrics.Gauge, "HeapObjects") {
		gaugeMetrics = append(gaugeMetrics, *NewGaugeMetric("HeapObjects", heapObjectsGaugeMetric()))
	}

	if slices.Contains(c.trackedMetrics.Gauge, "HeapReleased") {
		gaugeMetrics = append(gaugeMetrics, *NewGaugeMetric("HeapReleased", heapReleasedGaugeMetric()))
	}

	if slices.Contains(c.trackedMetrics.Gauge, "HeapSys") {
		gaugeMetrics = append(gaugeMetrics, *NewGaugeMetric("HeapSys", heapSysGaugeMetric()))
	}

	if slices.Contains(c.trackedMetrics.Gauge, "LastGC") {
		gaugeMetrics = append(gaugeMetrics, *NewGaugeMetric("LastGC", lastGCGaugeMetric()))
	}

	if slices.Contains(c.trackedMetrics.Gauge, "Lookups") {
		gaugeMetrics = append(gaugeMetrics, *NewGaugeMetric("Lookups", lookupsGaugeMetric()))
	}

	if slices.Contains(c.trackedMetrics.Gauge, "MCacheInuse") {
		gaugeMetrics = append(gaugeMetrics, *NewGaugeMetric("MCacheInuse", mCacheInuseGaugeMetric()))
	}

	if slices.Contains(c.trackedMetrics.Gauge, "MCacheSys") {
		gaugeMetrics = append(gaugeMetrics, *NewGaugeMetric("MCacheSys", mCacheSysGaugeMetric()))
	}

	if slices.Contains(c.trackedMetrics.Gauge, "MSpanInuse") {
		gaugeMetrics = append(gaugeMetrics, *NewGaugeMetric("MSpanInuse", mSpanInuseGaugeMetric()))
	}

	if slices.Contains(c.trackedMetrics.Gauge, "MSpanSys") {
		gaugeMetrics = append(gaugeMetrics, *NewGaugeMetric("MSpanSys", mSpanSysGaugeMetric()))
	}

	if slices.Contains(c.trackedMetrics.Gauge, "Mallocs") {
		gaugeMetrics = append(gaugeMetrics, *NewGaugeMetric("Mallocs", mallocsGaugeMetric()))
	}

	if slices.Contains(c.trackedMetrics.Gauge, "NextGC") {
		gaugeMetrics = append(gaugeMetrics, *NewGaugeMetric("NextGC", nextGCGaugeMetric()))
	}

	if slices.Contains(c.trackedMetrics.Gauge, "NumForcedGC") {
		gaugeMetrics = append(gaugeMetrics, *NewGaugeMetric("NumForcedGC", numForcedGCGaugeMetric()))
	}

	if slices.Contains(c.trackedMetrics.Gauge, "NumGC") {
		gaugeMetrics = append(gaugeMetrics, *NewGaugeMetric("NumGC", numGCGaugeMetric()))
	}

	if slices.Contains(c.trackedMetrics.Gauge, "OtherSys") {
		gaugeMetrics = append(gaugeMetrics, *NewGaugeMetric("OtherSys", otherSysGaugeMetric()))
	}

	if slices.Contains(c.trackedMetrics.Gauge, "PauseTotalNs") {
		gaugeMetrics = append(gaugeMetrics, *NewGaugeMetric("PauseTotalNs", pauseTotalNsGaugeMetric()))
	}

	if slices.Contains(c.trackedMetrics.Gauge, "StackInuse") {
		gaugeMetrics = append(gaugeMetrics, *NewGaugeMetric("StackInuse", stackInuseGaugeMetric()))
	}

	if slices.Contains(c.trackedMetrics.Gauge, "StackSys") {
		gaugeMetrics = append(gaugeMetrics, *NewGaugeMetric("StackSys", stackSysGaugeMetric()))
	}

	if slices.Contains(c.trackedMetrics.Gauge, "Sys") {
		gaugeMetrics = append(gaugeMetrics, *NewGaugeMetric("Sys", sysGaugeMetric()))
	}

	if slices.Contains(c.trackedMetrics.Gauge, "TotalAlloc") {
		gaugeMetrics = append(gaugeMetrics, *NewGaugeMetric("TotalAlloc", totalAllocGaugeMetric()))
	}

	if slices.Contains(c.trackedMetrics.Gauge, "RandomValue") {
		gaugeMetrics = append(gaugeMetrics, *NewGaugeMetric("RandomValue", randomValueGaugeMetric()))
	}

	return NewMetrics(counterMetrics, gaugeMetrics), nil
}

// Counter Metrics

func pollCountCounterMetric() int64 {
	return 1
}

// Gauge Metrics

func allocGaugeMetric() float64 {
	return float64(runtime.MemStats.Alloc)
}

func buckHashSysGaugeMetric() float64 {
	return float64(runtime.MemStats.BuckHashSys)
}

func freesGaugeMetric() float64 {
	return float64(runtime.MemStats.Frees)
}

func gCCPUFractionGaugeMetric() float64 {
	return float64(runtime.MemStats.GCCPUFraction)
}

func gCSysGaugeMetric() float64 {
	return float64(runtime.MemStats.GCSys)
}

func heapAllocGaugeMetric() float64 {
	return float64(runtime.MemStats.HeapAlloc)
}

func heapIdleGaugeMetric() float64 {
	return float64(runtime.MemStats.HeapIdle)
}

func heapInuseGaugeMetric() float64 {
	return float64(runtime.MemStats.HeapInuse)
}

func heapObjectsGaugeMetric() float64 {
	return float64(runtime.MemStats.HeapObjects)
}

func heapReleasedGaugeMetric() float64 {
	return float64(runtime.MemStats.HeapReleased)
}

func heapSysGaugeMetric() float64 {
	return float64(runtime.MemStats.HeapSys)
}

func lastGCGaugeMetric() float64 {
	return float64(runtime.MemStats.LastGC)
}

func lookupsGaugeMetric() float64 {
	return float64(runtime.MemStats.Lookups)
}

func mCacheInuseGaugeMetric() float64 {
	return float64(runtime.MemStats.MCacheInuse)
}

func mCacheSysGaugeMetric() float64 {
	return float64(runtime.MemStats.MCacheSys)
}

func mSpanInuseGaugeMetric() float64 {
	return float64(runtime.MemStats.MSpanInuse)
}

func mSpanSysGaugeMetric() float64 {
	return float64(runtime.MemStats.MSpanSys)
}

func mallocsGaugeMetric() float64 {
	return float64(runtime.MemStats.Mallocs)
}

func nextGCGaugeMetric() float64 {
	return float64(runtime.MemStats.NextGC)
}

func numForcedGCGaugeMetric() float64 {
	return float64(runtime.MemStats.NumForcedGC)
}

func numGCGaugeMetric() float64 {
	return float64(runtime.MemStats.NumGC)
}

func otherSysGaugeMetric() float64 {
	return float64(runtime.MemStats.OtherSys)
}

func pauseTotalNsGaugeMetric() float64 {
	return float64(runtime.MemStats.PauseTotalNs)
}

func stackInuseGaugeMetric() float64 {
	return float64(runtime.MemStats.StackInuse)
}

func stackSysGaugeMetric() float64 {
	return float64(runtime.MemStats.StackSys)
}

func sysGaugeMetric() float64 {
	return float64(runtime.MemStats.Sys)
}

func totalAllocGaugeMetric() float64 {
	return float64(runtime.MemStats.TotalAlloc)
}

func randomValueGaugeMetric() float64 {
	return rand.ExpFloat64()
}
