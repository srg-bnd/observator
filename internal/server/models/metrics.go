// Models (Metrics)
package models

import (
	"strconv"
)

const (
	CounterMType = "counter"
	GaugeMType   = "gauge"
)

var MetricsMTypes = []string{CounterMType, GaugeMType}

type Metrics struct {
	ID    string   `json:"id"`
	MType string   `json:"type"`
	Delta *int64   `json:"delta,omitempty"` // counter
	Value *float64 `json:"value,omitempty"` // gauge
}

func NewMetrics() *Metrics {
	return &Metrics{}
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

func (m *Metrics) GetValueAsString() string {
	if m.MType == CounterMType {
		if m.Delta == nil {
			return ""
		}
		return strconv.FormatInt(*m.Delta, 10)

	} else {
		if m.Value == nil {
			return ""
		}

		return strconv.FormatFloat(*m.Value, 'f', -1, 64)
	}
}

// Helpers

func (m *Metrics) IsCounterMType() bool {
	return m.MType == CounterMType
}

func (m *Metrics) IsGaugeMType() bool {
	return m.MType == GaugeMType
}
