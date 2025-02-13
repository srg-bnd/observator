// Storage for metrics
package storage

import "errors"

type (
	gauge   float64
	counter int64
)

type Repositories interface {
	SetGauge(string, float64) error
	GetGauge(string) (float64, error)
	SetCounter(string, int64) error
	GetCounter(string) (int64, error)
}

type MemStorage struct {
	gauges   map[string]gauge
	counters map[string]counter
}

// Create MemStorage instance
func NewMemStorage() *MemStorage {
	return &MemStorage{
		gauges:   make(map[string]gauge),
		counters: make(map[string]counter),
	}
}

// Change gauge by key
func (mStore *MemStorage) SetGauge(key string, value float64) error {
	mStore.gauges[key] = gauge(value)
	return nil
}

// Return gauge by key
func (mStore *MemStorage) GetGauge(key string) (float64, error) {
	value, ok := mStore.gauges[key]
	if !ok {
		return -1, errors.New("unknown")
	}

	return float64(value), nil
}

// Change counter by key
func (mStore *MemStorage) SetCounter(key string, value int64) error {
	mStore.counters[key] += counter(value)
	return nil
}

// Return gauge by counter
func (mStore *MemStorage) GetCounter(key string) (int64, error) {
	value, ok := mStore.counters[key]
	if !ok {
		return -1, errors.New("unknown")
	}

	return int64(value), nil
}
