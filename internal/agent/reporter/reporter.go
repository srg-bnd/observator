// Sends metrics to the server
package reporter

import (
	"context"
	"errors"
	"fmt"
	"log"
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

var (
	ErrReportMetrics = errors.New("report metrics")
)

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
			go r.reportMetrics(ctx, collector.TrackedGopsutilMetrics)
			go r.reportMetrics(ctx, collector.TrackedRuntimeMetrics)
		}
	}
}

func (r *Reporter) reportMetrics(ctx context.Context, trackedRuntimeMetrics map[string][]string) {
	if err := r.metricService.Send(ctx, trackedRuntimeMetrics); err != nil {
		log.Println(fmt.Errorf("%f%f", ErrReportMetrics, err))
	}
}
