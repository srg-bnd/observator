package services

import (
	"strconv"

	"github.com/srg-bnd/observator/internal/agent/client"
	"github.com/srg-bnd/observator/internal/agent/collector"
	"github.com/srg-bnd/observator/internal/storage"
)

type Service struct {
	storage storage.Repositories
	client  *client.Client
}

func NewService(storage storage.Repositories) *Service {
	return &Service{
		storage: storage,
		client:  client.NewClient(),
	}
}

func (s *Service) MetricsUpdateService(metrics *collector.Metrics) error {
	for _, metric := range metrics.Counter {
		s.storage.SetCounter(metric.Name, metric.Value)
	}

	for _, metric := range metrics.Gauge {
		s.storage.SetGauge(metric.Name, metric.Value)
	}

	return nil
}

func (s *Service) ValueSendingService(trackedMetrics *collector.TrackedMetrics) error {
	for _, metricName := range trackedMetrics.Counter {
		value, err := s.storage.GetCounter(metricName)
		if err != nil {
			return err
		}
		metricValue := strconv.FormatInt(value, 10)

		s.client.SendMetric("counter", metricName, metricValue)
	}

	for _, metricName := range trackedMetrics.Gauge {
		value, err := s.storage.GetGauge(metricName)
		metricValue := strconv.FormatFloat(value, 'f', -1, 64)
		if err != nil {
			return err
		}

		s.client.SendMetric("gauge", metricName, metricValue)
	}

	return nil
}
