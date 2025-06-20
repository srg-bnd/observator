// Collects metrics
package collector

type Metrics struct {
	Counter map[string]int64
	Gauge   map[string]float64
}
