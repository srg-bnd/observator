package collector

type Collector struct {
}

type CounterMetric struct {
	Name  string
	Value int64
}

type GaugeMetric struct {
	Name  string
	Value float64
}

func NewCounterMetric(name string, value int64) *CounterMetric {
	return &CounterMetric{
		Name:  name,
		Value: value,
	}
}

func NewGaugeMetric(name string, value float64) *GaugeMetric {
	return &GaugeMetric{
		Name:  name,
		Value: value,
	}
}

type Metrics struct {
	Counter *[]CounterMetric
	Gauge   *[]GaugeMetric
}

func NewMetrics(counter *[]CounterMetric, gauge *[]GaugeMetric) *Metrics {
	return &Metrics{
		Counter: counter,
		Gauge:   gauge,
	}
}

func NewCollector() *Collector {
	return &Collector{}
}

func (c *Collector) GetMetrics() (*Metrics, error) {
	return NewMetrics(
		&[]CounterMetric{
			// Custom
			*NewCounterMetric("PollCount", 0),
		},
		&[]GaugeMetric{
			// Runtime
			*NewGaugeMetric("Alloc", 0.0),
			*NewGaugeMetric("BuckHashSys", 0.0),
			*NewGaugeMetric("Frees", 0.0),
			*NewGaugeMetric("GCCPUFraction", 0.0),
			*NewGaugeMetric("GCSys", 0.0),
			*NewGaugeMetric("HeapAlloc", 0.0),
			*NewGaugeMetric("HeapIdle", 0.0),
			*NewGaugeMetric("HeapInuse", 0.0),
			*NewGaugeMetric("HeapObjects", 0.0),
			*NewGaugeMetric("HeapReleased", 0.0),
			*NewGaugeMetric("HeapSys", 0.0),
			*NewGaugeMetric("LastGC", 0.0),
			*NewGaugeMetric("Lookups", 0.0),
			*NewGaugeMetric("MCacheInuse", 0.0),
			*NewGaugeMetric("MCacheSys", 0.0),
			*NewGaugeMetric("MSpanInuse", 0.0),
			*NewGaugeMetric("MSpanSys", 0.0),
			*NewGaugeMetric("Mallocs", 0.0),
			*NewGaugeMetric("NextGC", 0.0),
			*NewGaugeMetric("NumForcedGC", 0.0),
			*NewGaugeMetric("NumGC", 0.0),
			*NewGaugeMetric("OtherSys", 0.0),
			*NewGaugeMetric("PauseTotalNs", 0.0),
			*NewGaugeMetric("StackInuse", 0.0),
			*NewGaugeMetric("StackSys", 0.0),
			*NewGaugeMetric("Sys", 0.0),
			*NewGaugeMetric("TotalAlloc", 0.0),
			// Custom
			*NewGaugeMetric("RandomValue", 0.0),
		},
	), nil
}
