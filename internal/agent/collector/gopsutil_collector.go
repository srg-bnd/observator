package collector

import (
	"fmt"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/mem"
)

var TrackedGopsutilMetrics = map[string][]string{
	"gauge": {
		"TotalMemory",
		"FreeMemory",
		// NOTE: CPUutilization from 1 to ...
		"CPUutilization1",
	},
}

type GopsutilCollector struct {
}

// Returns new collector
func NewGopsutilCollector() *GopsutilCollector {
	// Dynamic metrics
	countOfCPU, _ := cpu.Counts(true)
	for numberCPU := 1; numberCPU < countOfCPU; numberCPU++ {
		TrackedGopsutilMetrics["gauge"] = append(TrackedGopsutilMetrics["gauge"], fmt.Sprint("CPUutilization", numberCPU+1))
	}

	return &GopsutilCollector{}
}

// Returns current values for metrics
func (c *GopsutilCollector) GetMetrics() (*Metrics, error) {
	virtualMemoryStat, _ := mem.VirtualMemory()
	percentageOfCPU, _ := cpu.Percent(0, true)

	gauges := make(map[string]float64, 0)
	for numberCPU, percentage := range percentageOfCPU {
		gauges[fmt.Sprint("CPUutilization", numberCPU+1)] = percentage
	}

	gauges["TotalMemory"] = float64(virtualMemoryStat.Total)
	gauges["FreeMemory"] = float64(virtualMemoryStat.Free)

	return &Metrics{Gauge: gauges}, nil
}
