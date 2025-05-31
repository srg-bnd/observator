// Services
package services

import (
	"context"

	"github.com/srg-bnd/observator/internal/agent/client"
	"github.com/srg-bnd/observator/internal/agent/collector"
	"github.com/srg-bnd/observator/internal/agent/models"
	"github.com/srg-bnd/observator/internal/storage"
)

// Service
type Service struct {
	storage storage.Repositories
	client  *client.Client
}

// Returns a new service
func NewService(storage storage.Repositories, client *client.Client) *Service {
	return &Service{
		storage: storage,
		client:  client,
	}
}

// Updates metrics in the storage
func (s *Service) MetricsUpdateService(ctx context.Context, metrics *collector.Metrics) error {
	for metricName, metricValue := range metrics.Counter {
		s.storage.SetCounter(ctx, metricName, metricValue)
	}

	for metricName, metricValue := range metrics.Gauge {
		s.storage.SetGauge(ctx, metricName, metricValue)
	}

	return nil
}

// Collects metrics and sends them to the server
func (s *Service) ValueSendingService(ctx context.Context, trackedMetrics map[string][]string) error {
	metrics := make([]models.Metrics, 0, len(trackedMetrics))

	for _, metricName := range trackedMetrics["counter"] {
		if value, err := s.storage.GetCounter(ctx, metricName); err == nil {
			metrics = append(metrics, models.Metrics{ID: metricName, MType: "counter", Delta: &value})
		}
		// TODO: send to logs if error
	}

	for _, metricName := range trackedMetrics["gauge"] {
		if value, err := s.storage.GetGauge(ctx, metricName); err == nil {
			metrics = append(metrics, models.Metrics{ID: metricName, MType: "gauge", Value: &value})
		}
		// TODO: send to logs if error

	}

	if len(metrics) > 0 {
		s.client.SendMetrics(metrics)
	}

	return nil
}
