package collector

import (
	"github.com/shirou/gopsutil/v4/mem"
)

var TrackedGopsutilMetrics = map[string][]string{
	"gauge": {
		"TotalMemory",
		"FreeMemory",
	},
}

type GopsutilCollector struct {
}

// Returns new collector
func NewGopsutilCollector() *GopsutilCollector {
	return &GopsutilCollector{}
}

// Returns current values for metrics
func (c *GopsutilCollector) GetMetrics() (*Metrics, error) {
	virtualMemoryStat, _ := mem.VirtualMemory()

	return &Metrics{
		Gauge: map[string]float64{
			"TotalMemory": float64(virtualMemoryStat.Total),
			"FreeMemory":  float64(virtualMemoryStat.Free),
		}}, nil
}
