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
	metricService MetricService
	repository    storage.Repositories
	workerPool    int
}

type Job struct {
	trackedMetrics map[string][]string
}

var ErrReportMetrics = errors.New("report metrics")

// Returns a new reporter
func NewReporter(repository storage.Repositories, workerPool int, client *client.Client) *Reporter {
	return &Reporter{
		repository:    repository,
		workerPool:    workerPool,
		metricService: services.NewMetricService(repository, client),
	}
}

// Starts the reporter
func (r *Reporter) Start(ctx context.Context, reportInterval time.Duration) error {
	ticker := time.NewTicker(reportInterval)
	defer ticker.Stop()

	jobs := make(chan Job, r.workerPool)
	defer close(jobs)

	// Runs workers
	for range r.workerPool {
		go r.worker(jobs, ctx)
	}

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-ticker.C:
			// Sets jobs
			jobs <- Job{trackedMetrics: collector.TrackedRuntimeMetrics}
			jobs <- Job{trackedMetrics: collector.TrackedGopsutilMetrics}
		}
	}
}

func (r *Reporter) reportMetrics(ctx context.Context, trackedMetrics map[string][]string) {
	if err := r.metricService.Send(ctx, trackedMetrics); err != nil {
		log.Println(fmt.Errorf("%f%f", ErrReportMetrics, err))
	}
}

func (r *Reporter) worker(jobs <-chan Job, ctx context.Context) {
	for job := range jobs {
		r.reportMetrics(ctx, job.trackedMetrics)
	}
}
