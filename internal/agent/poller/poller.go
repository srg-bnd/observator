// Polls the metrics & saves them to the storage
package poller

import (
	"context"
	"time"

	"github.com/srg-bnd/observator/internal/agent/collector"
	"github.com/srg-bnd/observator/internal/agent/services"
	"github.com/srg-bnd/observator/internal/storage"
)

type Collector interface {
	GetMetrics() (*collector.Metrics, error)
}

type MetricService interface {
	Update(context.Context, *collector.Metrics) error
}

type Poller struct {
	collector     Collector
	metricService MetricService
}

// Returns a new poller
func NewPoller(repository storage.Repositories) *Poller {
	return &Poller{
		collector:     collector.NewCollector(),
		metricService: services.NewMetricService(repository, nil),
	}
}

// Starts the poller
func (r *Poller) Start(ctx context.Context, pollInterval time.Duration) error {
	ticker := time.NewTicker(pollInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-ticker.C:
			metrics, err := r.collector.GetMetrics()
			if err != nil {
				return err
			}

			if err = r.metricService.Update(ctx, metrics); err != nil {
				return err
			}
		}
	}
}
