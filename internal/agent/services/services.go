package services

import (
	"github.com/srg-bnd/observator/internal/agent/collector"
	"github.com/srg-bnd/observator/internal/storage"
)

type Service struct {
	storage storage.Repositories
}

func NewService(storage storage.Repositories) *Service {
	return &Service{
		storage: storage,
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
