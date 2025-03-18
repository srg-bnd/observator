// Services
package services

import (
	"github.com/srg-bnd/observator/internal/agent/client"
	"github.com/srg-bnd/observator/internal/agent/collector"
	"github.com/srg-bnd/observator/internal/agent/models"
	"github.com/srg-bnd/observator/internal/storage"
)

type Service struct {
	storage storage.Repositories
	client  *client.Client
}

func NewService(storage storage.Repositories, client *client.Client) *Service {
	return &Service{
		storage: storage,
		client:  client,
	}
}

func (s *Service) MetricsUpdateService(metrics *collector.Metrics) error {
	for metricName, metricValue := range metrics.Counter {
		s.storage.SetCounter(metricName, metricValue)
	}

	for metricName, metricValue := range metrics.Gauge {
		s.storage.SetGauge(metricName, metricValue)
	}

	return nil
}

func (s *Service) ValueSendingService(trackedMetrics map[string][]string) error {
	for _, metricName := range trackedMetrics["counter"] {
		value, err := s.storage.GetCounter(metricName)
		if err != nil {
			return err
		}

		s.client.SendMetric(&models.Metrics{ID: metricName, MType: "counter", Delta: &value})
	}

	for _, metricName := range trackedMetrics["gauge"] {
		value, err := s.storage.GetGauge(metricName)
		if err != nil {
			return err
		}

		s.client.SendMetric(&models.Metrics{ID: metricName, MType: "gauge", Value: &value})
	}

	return nil
}
