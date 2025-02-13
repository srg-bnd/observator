package collector

import "slices"

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
		gaugeMetrics = append(gaugeMetrics, *NewGaugeMetric("BuckHashSys", allocGaugeMetric()))
	}

	if slices.Contains(c.trackedMetrics.Gauge, "Frees") {
		gaugeMetrics = append(gaugeMetrics, *NewGaugeMetric("Frees", allocGaugeMetric()))
	}

	if slices.Contains(c.trackedMetrics.Gauge, "GCCPUFraction") {
		gaugeMetrics = append(gaugeMetrics, *NewGaugeMetric("GCCPUFraction", allocGaugeMetric()))
	}

	if slices.Contains(c.trackedMetrics.Gauge, "GCSys") {
		gaugeMetrics = append(gaugeMetrics, *NewGaugeMetric("GCSys", allocGaugeMetric()))
	}

	if slices.Contains(c.trackedMetrics.Gauge, "HeapAlloc") {
		gaugeMetrics = append(gaugeMetrics, *NewGaugeMetric("HeapAlloc", allocGaugeMetric()))
	}

	if slices.Contains(c.trackedMetrics.Gauge, "HeapIdle") {
		gaugeMetrics = append(gaugeMetrics, *NewGaugeMetric("HeapIdle", allocGaugeMetric()))
	}

	if slices.Contains(c.trackedMetrics.Gauge, "HeapInuse") {
		gaugeMetrics = append(gaugeMetrics, *NewGaugeMetric("HeapInuse", allocGaugeMetric()))
	}

	if slices.Contains(c.trackedMetrics.Gauge, "HeapObjects") {
		gaugeMetrics = append(gaugeMetrics, *NewGaugeMetric("HeapObjects", allocGaugeMetric()))
	}

	if slices.Contains(c.trackedMetrics.Gauge, "HeapReleased") {
		gaugeMetrics = append(gaugeMetrics, *NewGaugeMetric("HeapReleased", allocGaugeMetric()))
	}

	if slices.Contains(c.trackedMetrics.Gauge, "HeapSys") {
		gaugeMetrics = append(gaugeMetrics, *NewGaugeMetric("HeapSys", allocGaugeMetric()))
	}

	if slices.Contains(c.trackedMetrics.Gauge, "LastGC") {
		gaugeMetrics = append(gaugeMetrics, *NewGaugeMetric("LastGC", allocGaugeMetric()))
	}

	if slices.Contains(c.trackedMetrics.Gauge, "Lookups") {
		gaugeMetrics = append(gaugeMetrics, *NewGaugeMetric("Lookups", allocGaugeMetric()))
	}

	if slices.Contains(c.trackedMetrics.Gauge, "MCacheInuse") {
		gaugeMetrics = append(gaugeMetrics, *NewGaugeMetric("MCacheInuse", allocGaugeMetric()))
	}

	if slices.Contains(c.trackedMetrics.Gauge, "MCacheSys") {
		gaugeMetrics = append(gaugeMetrics, *NewGaugeMetric("MCacheSys", allocGaugeMetric()))
	}

	if slices.Contains(c.trackedMetrics.Gauge, "MSpanInuse") {
		gaugeMetrics = append(gaugeMetrics, *NewGaugeMetric("MSpanInuse", allocGaugeMetric()))
	}

	if slices.Contains(c.trackedMetrics.Gauge, "MSpanSys") {
		gaugeMetrics = append(gaugeMetrics, *NewGaugeMetric("MSpanSys", allocGaugeMetric()))
	}

	if slices.Contains(c.trackedMetrics.Gauge, "Mallocs") {
		gaugeMetrics = append(gaugeMetrics, *NewGaugeMetric("Mallocs", allocGaugeMetric()))
	}

	if slices.Contains(c.trackedMetrics.Gauge, "NextGC") {
		gaugeMetrics = append(gaugeMetrics, *NewGaugeMetric("NextGC", allocGaugeMetric()))
	}

	if slices.Contains(c.trackedMetrics.Gauge, "NumForcedGC") {
		gaugeMetrics = append(gaugeMetrics, *NewGaugeMetric("NumForcedGC", allocGaugeMetric()))
	}

	if slices.Contains(c.trackedMetrics.Gauge, "NumGC") {
		gaugeMetrics = append(gaugeMetrics, *NewGaugeMetric("NumGC", allocGaugeMetric()))
	}

	if slices.Contains(c.trackedMetrics.Gauge, "OtherSys") {
		gaugeMetrics = append(gaugeMetrics, *NewGaugeMetric("OtherSys", allocGaugeMetric()))
	}

	if slices.Contains(c.trackedMetrics.Gauge, "PauseTotalNs") {
		gaugeMetrics = append(gaugeMetrics, *NewGaugeMetric("PauseTotalNs", allocGaugeMetric()))
	}

	if slices.Contains(c.trackedMetrics.Gauge, "StackInuse") {
		gaugeMetrics = append(gaugeMetrics, *NewGaugeMetric("StackInuse", allocGaugeMetric()))
	}

	if slices.Contains(c.trackedMetrics.Gauge, "StackSys") {
		gaugeMetrics = append(gaugeMetrics, *NewGaugeMetric("StackSys", allocGaugeMetric()))
	}

	if slices.Contains(c.trackedMetrics.Gauge, "Sys") {
		gaugeMetrics = append(gaugeMetrics, *NewGaugeMetric("Sys", allocGaugeMetric()))
	}

	if slices.Contains(c.trackedMetrics.Gauge, "TotalAlloc") {
		gaugeMetrics = append(gaugeMetrics, *NewGaugeMetric("TotalAlloc", allocGaugeMetric()))
	}

	if slices.Contains(c.trackedMetrics.Gauge, "RandomValue") {
		gaugeMetrics = append(gaugeMetrics, *NewGaugeMetric("RandomValue", allocGaugeMetric()))
	}

	return NewMetrics(counterMetrics, gaugeMetrics), nil
}

func pollCountCounterMetric() int64 {
	return 0
}

func allocGaugeMetric() float64 {
	return 0.0
}
