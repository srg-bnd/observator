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
		metrics, err := r.collector.GetMetrics()

		if err != nil {
			log.Println(err)
			return
		}
		ShowMetrics(metrics)
		log.Println("=== Poller stopped ===")
	}
}

func ShowMetrics(metrics *collector.Metrics) {
	fmt.Println("Counter:")
	for _, counterMetric := range *metrics.Counter {
		fmt.Println("-", counterMetric.Name, ":", counterMetric.Value)
	}

	fmt.Println("Gauge:")
	for _, gaugeMetric := range *metrics.Gauge {
		fmt.Println("-", gaugeMetric.Name, ":", gaugeMetric.Value)
	}
}

func GetPollInterval() time.Duration {
	return defaultPollInterval
}
