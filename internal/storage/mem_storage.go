// Mem storage for metrics
package storage

import (
	"context"
	"errors"
	"sync"
)

// Mem storage
type MemStorage struct {
	gauges   map[string]gauge
	counters map[string]counter
	mtx      sync.RWMutex
}

// Returns a new mem storage
func NewMemStorage() *MemStorage {
	return &MemStorage{
		gauges:   make(map[string]gauge),
		counters: make(map[string]counter),
		mtx:      sync.RWMutex{},
	}
}

// Changes gauge by key
func (mStore *MemStorage) SetGauge(ctx context.Context, key string, value float64) error {
	mStore.mtx.Lock()
	mStore.gauges[key] = gauge(value)
	mStore.mtx.Unlock()

	return nil
}

// Returns gauge by key
func (mStore *MemStorage) GetGauge(ctx context.Context, key string) (float64, error) {
	mStore.mtx.RLock()

	value, ok := mStore.gauges[key]
	if !ok {
		return -1, errors.New("unknown")
	}

	mStore.mtx.RUnlock()

	return float64(value), nil
}

// Changes counter by key
func (mStore *MemStorage) SetCounter(ctx context.Context, key string, value int64) error {
	mStore.mtx.Lock()
	mStore.counters[key] += counter(value)
	mStore.mtx.Unlock()

	return nil
}

// Returns gauge by counter
func (mStore *MemStorage) GetCounter(ctx context.Context, key string) (int64, error) {
	mStore.mtx.RLock()
	value, ok := mStore.counters[key]
	if !ok {
		return -1, errors.New("unknown")
	}
	mStore.mtx.RUnlock()

	return int64(value), nil
}

// Batch update batch of metrics
func (mStore *MemStorage) SetBatchOfMetrics(ctx context.Context, counterMetrics map[string]int64, gaugeMetrics map[string]float64) error {
	for key, value := range counterMetrics {
		mStore.SetCounter(ctx, key, value)
	}

	for key, value := range gaugeMetrics {
		mStore.SetGauge(ctx, key, value)
	}

	return nil
}
