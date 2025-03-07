// Storage (Metrics)
package storage

type Repositories interface {
	SetGauge(string, float64) error
	GetGauge(string) (float64, error)
	SetCounter(string, int64) error
	GetCounter(string) (int64, error)
}
