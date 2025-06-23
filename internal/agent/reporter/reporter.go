// Sends metrics to the server
package reporter

import (
	"context"
	"time"

	"github.com/srg-bnd/observator/internal/agent/client"
	"github.com/srg-bnd/observator/internal/agent/collector"
	"github.com/srg-bnd/observator/internal/agent/services"
	"github.com/srg-bnd/observator/internal/storage"
)

type MetricService interface {
	Send(context.Context, map[string][]string) error
}

// Reporter
type Reporter struct {
	repository    storage.Repositories
	metricService MetricService
}

// Returns a new reporter
func NewReporter(repository storage.Repositories, client *client.Client) *Reporter {
	return &Reporter{
		repository:    repository,
		metricService: services.NewMetricService(repository, client),
	}
}

// Starts the reporter
func (r *Reporter) Start(ctx context.Context, reportInterval time.Duration) error {
	ticker := time.NewTicker(reportInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-ticker.C:
			if err := r.metricService.Send(ctx, collector.TrackedMetrics); err != nil {
				return err
			}
		}
	}
}
