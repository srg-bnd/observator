// Polls the metrics & saves them to the storage
package poller

import (
	"time"

	"github.com/srg-bnd/observator/internal/agent/collector"
	"github.com/srg-bnd/observator/internal/agent/services"
	"github.com/srg-bnd/observator/internal/storage"
)

// Poller
type Poller struct {
	storage   storage.Repositories
	collector *collector.Collector
	services  *services.Service
}

// Returns a new poller
func NewPoller(storage storage.Repositories) *Poller {
	return &Poller{
		storage:   storage,
		collector: collector.NewCollector(),
		services:  services.NewService(storage, nil),
	}
}

// Starts the poller
func (r *Poller) Start(pollInterval time.Duration) error {
	ticker := time.NewTicker(pollInterval * time.Second)
	defer ticker.Stop()

	for {
		<-ticker.C

		metrics, err := r.collector.GetMetrics()
		if err != nil {
			return err
		}

		err = r.services.MetricsUpdateService(metrics)
		if err != nil {
			return err
		}
	}
}
