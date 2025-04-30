// Mem storage for metrics
package storage

import (
	"errors"
)

// Mem storage
type MemStorage struct {
	gauges   map[string]gauge
	counters map[string]counter
}

// Returns a new mem storage
func NewMemStorage() *MemStorage {
	return &MemStorage{
		gauges:   make(map[string]gauge),
		counters: make(map[string]counter),
	}
}

// Changes gauge by key
func (mStore *MemStorage) SetGauge(key string, value float64) error {
	mStore.gauges[key] = gauge(value)

	return nil
}

// Returns gauge by key
func (mStore *MemStorage) GetGauge(key string) (float64, error) {
	value, ok := mStore.gauges[key]
	if !ok {
		return -1, errors.New("unknown")
	}

	return float64(value), nil
}

// Changes counter by key
func (mStore *MemStorage) SetCounter(key string, value int64) error {
	mStore.counters[key] += counter(value)

	return nil
}

// Returns gauge by counter
func (mStore *MemStorage) GetCounter(key string) (int64, error) {
	value, ok := mStore.counters[key]
	if !ok {
		return -1, errors.New("unknown")
	}

	return int64(value), nil
}

// Updates batch of metrics
func (mStore *MemStorage) SetBatchOfMetrics() error {
	return nil
}
