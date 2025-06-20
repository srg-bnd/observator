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
	ErrReportGopsutilMetrics = errors.New("report gopsutil metrics")
	ErrReportRuntimeMetrics  = errors.New("report runtime metrics")
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
			go r.reportRuntimeMetrics(ctx)
			go r.reportGopsutilMetrics(ctx)
		}
	}
}

func (r *Reporter) reportGopsutilMetrics(ctx context.Context) {
	if err := r.metricService.Send(ctx, collector.TrackedGopsutilMetrics); err != nil {
		log.Println(fmt.Errorf("%f%f", ErrReportRuntimeMetrics, err))
	}
}

func (r *Reporter) reportRuntimeMetrics(ctx context.Context) {
	if err := r.metricService.Send(ctx, collector.TrackedRuntimeMetrics); err != nil {
		log.Println(fmt.Errorf("%f%f", ErrReportRuntimeMetrics, err))
	}
}
