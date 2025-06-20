// Services
package services

import (
	"context"
	"log"

	"github.com/srg-bnd/observator/internal/agent/collector"
	"github.com/srg-bnd/observator/internal/agent/models"
)

type Client interface {
	SendMetrics(context.Context, []models.Metrics) error
}

type Repository interface {
	SetGauge(context.Context, string, float64) error
	SetCounter(context.Context, string, int64) error
	GetGauge(context.Context, string) (float64, error)
	GetCounter(context.Context, string) (int64, error)
}

type MetricService struct {
	repository Repository
	client     Client
}

// Returns a new service
func NewMetricService(repository Repository, client Client) *MetricService {
	return &MetricService{
		repository: repository,
		client:     client,
	}
}

// Updates metrics in the repository
func (s *MetricService) Update(ctx context.Context, metrics *collector.Metrics) error {
	for metricName, metricValue := range metrics.Counter {
		s.repository.SetCounter(ctx, metricName, metricValue)
	}

	for metricName, metricValue := range metrics.Gauge {
		s.repository.SetGauge(ctx, metricName, metricValue)
	}

	return nil
}

// Sends metrics to the server
func (s *MetricService) Send(ctx context.Context, trackedMetrics map[string][]string) error {
	metrics := make([]models.Metrics, 0, len(trackedMetrics))

	for _, metricName := range trackedMetrics[models.CounterMType] {
		if value, err := s.repository.GetCounter(ctx, metricName); err == nil {
			metrics = append(metrics, models.Metrics{ID: metricName, MType: models.CounterMType, Delta: &value})
		}
	}

	for _, metricName := range trackedMetrics[models.GaugeMType] {
		if value, err := s.repository.GetGauge(ctx, metricName); err == nil {
			metrics = append(metrics, models.Metrics{ID: metricName, MType: models.GaugeMType, Value: &value})
		}
	}

	if len(metrics) == 0 {
		return nil
	}

	err := s.client.SendMetrics(ctx, metrics)
	if err != nil {
		// TODO: uses normal logger
		log.Println("send metrics:", err)
	}

	return nil
}
