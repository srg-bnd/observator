package models

/* Metric */

type Metric struct {
	Type         string
	Name         string
	counterValue int64
	gaugeValue   float64
}

func NewMetric() *Metric {
	return &Metric{}
}

func (m *Metric) SetCounter(value int64) {
	m.counterValue = value
}

func (m *Metric) GetCounter() int64 {
	return m.counterValue
}

func (m *Metric) SetGauge(value float64) {
	m.gaugeValue = value
}

func (m *Metric) GetGauge() float64 {
	return m.gaugeValue
}
