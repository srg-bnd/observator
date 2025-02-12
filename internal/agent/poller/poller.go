// Collects metrics
package poller

import (
	"fmt"
	"log"
	"time"

	"github.com/srg-bnd/observator/internal/agent/collector"
	"github.com/srg-bnd/observator/internal/storage"
)

const (
	defaultPollInterval = 2 * time.Second
)

type Poller struct {
	storage   storage.Repositories
	collector *collector.Collector
}

func NewPoller(storage storage.Repositories) *Poller {
	return &Poller{
		storage:   storage,
		collector: collector.NewCollector(),
	}
}

func (r *Poller) Start() {
	for {
		time.Sleep(GetPollInterval())
		log.Println("=== Poller started ===")
		metricsByType, err := r.collector.GetMetrics()

		if err != nil {
			log.Println(err)
			return
		}
		ShowMetrics(metricsByType)
		log.Println("=== Poller stopped ===")
	}
}

func ShowMetrics(metricsByType map[string][]string) {
	fmt.Println("Gauge:")

	for typeOfMetric, metrics := range metricsByType {
		fmt.Println(typeOfMetric, ":")
		for _, metric := range metrics {
			fmt.Println("-", metric, ":", 0)
		}
	}
}

func GetPollInterval() time.Duration {
	return defaultPollInterval
}
