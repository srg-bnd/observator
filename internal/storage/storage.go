// Storage for metrics
package storage

import "context"

type Repositories interface {
	SetGauge(context.Context, string, float64) error
	GetGauge(context.Context, string) (float64, error)
	SetCounter(context.Context, string, int64) error
	GetCounter(context.Context, string) (int64, error)
	SetBatchOfMetrics(context.Context, map[string]int64, map[string]float64) error
}

type (
	gauge   float64
	counter int64
)
