// Models (Metrics)
package models

import "strconv"

type Metrics struct {
	ID    string   `json:"id"`
	MType string   `json:"type"`
	Delta *int64   `json:"delta,omitempty"` // counter
	Value *float64 `json:"value,omitempty"` // gauge
}

func NewMetrics() *Metrics {
	return &Metrics{}
}

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
	},
}

func (m *Metrics) SetCounter(value int64) {
	m.Delta = &value
}

func (m *Metrics) GetCounter() int64 {
	return *m.Delta
}

func (m *Metrics) SetGauge(value float64) {
	m.Value = &value
}

func (m *Metrics) GetGauge() float64 {
	return *m.Value
}

func (m *Metrics) GetCounterAsString() string {
	return strconv.FormatInt(*m.Delta, 10)
}

func (m *Metrics) GetGaugeAsString() string {
	return strconv.FormatFloat(*m.Value, 'f', -1, 64)
}
