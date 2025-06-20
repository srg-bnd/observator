package collector

import (
	"strconv"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/mem"
)

var TrackedGopsutilMetrics = map[string][]string{
	"gauge": {
		"TotalMemory",
		"FreeMemory",
		"CPUutilization1",
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
	percentageOfCPU, _ := cpu.Percent(0, true)

	gauges := make(map[string]float64, 0)
	for numberCPU, percentage := range percentageOfCPU {
		gauges["CPUutilization"+strconv.Itoa(numberCPU+1)] = percentage
	}

	gauges["TotalMemory"] = float64(virtualMemoryStat.Total)
	gauges["FreeMemory"] = float64(virtualMemoryStat.Free)

	return &Metrics{Gauge: gauges}, nil
}
