// Collects metrics
package poller

import (
	"log"
	"time"

	"github.com/srg-bnd/observator/internal/agent/collector"
	"github.com/srg-bnd/observator/internal/agent/services"
	"github.com/srg-bnd/observator/internal/storage"
)

const (
	defaultPollInterval = 2 * time.Second
)

type Poller struct {
	storage   storage.Repositories
	collector *collector.Collector
	services  *services.Service
}

func NewPoller(storage storage.Repositories) *Poller {
	return &Poller{
		storage:   storage,
		collector: collector.NewCollector(),
		services:  services.NewService(storage),
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

		err = r.services.MetricsUpdateService(metrics)
		if err != nil {
			log.Println(err)
			return
		}

		log.Println("=== Poller stopped ===")
	}
}

func GetPollInterval() time.Duration {
	return defaultPollInterval
}
