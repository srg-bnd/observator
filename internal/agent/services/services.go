package services

import (
	"fmt"

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
	ShowMetrics(metrics)

	return nil
}

func ShowMetrics(metrics *collector.Metrics) {
	fmt.Println("Counter:")
	for _, counterMetric := range metrics.Counter {
		fmt.Println("-", counterMetric.Name, ":", counterMetric.Value)
	}

	fmt.Println("Gauge:")
	for _, gaugeMetric := range metrics.Gauge {
		fmt.Println("-", gaugeMetric.Name, ":", gaugeMetric.Value)
	}
}
