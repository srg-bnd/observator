// Polls the metrics & saves them to the storage
package poller

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/srg-bnd/observator/internal/agent/collector"
	"github.com/srg-bnd/observator/internal/agent/services"
	"github.com/srg-bnd/observator/internal/storage"
)

type Collector interface {
	GetRuntimeMetrics() (*collector.Metrics, error)
	GetGopsutilMetrics() (*collector.Metrics, error)
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

var (
	ErrGetRuntimeMetrics  = errors.New("get runtime metrics")
	ErrGetGopsutilMetrics = errors.New("get gopsutil metrics")
	ErrUpdateMetrics      = errors.New("update metrics")
)

// Starts the poller
func (p *Poller) Start(ctx context.Context, pollInterval time.Duration) error {
	ticker := time.NewTicker(pollInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-ticker.C:
			go p.pollRuntimeMetrics(ctx)
			go p.pollGopsutilMetrics(ctx)
		}
	}
}

func (p *Poller) pollRuntimeMetrics(ctx context.Context) {
	metrics, err := p.collector.GetRuntimeMetrics()
	if err != nil {
		log.Println(fmt.Errorf("%f%f", ErrGetRuntimeMetrics, err))
	}

	if err = p.metricService.Update(ctx, metrics); err != nil {
		log.Println(fmt.Errorf("%f%f", ErrUpdateMetrics, err))
	}

}

func (p *Poller) pollGopsutilMetrics(ctx context.Context) {
	metrics, err := p.collector.GetGopsutilMetrics()
	if err != nil {
		log.Println(fmt.Errorf("%f%f", ErrGetRuntimeMetrics, err))
	}

	if err = p.metricService.Update(ctx, metrics); err != nil {
		log.Println(fmt.Errorf("%f%f", ErrUpdateMetrics, err))
	}
}
