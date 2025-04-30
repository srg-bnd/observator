// Services for server
package services

import (
	"github.com/srg-bnd/observator/internal/server/models"
	"github.com/srg-bnd/observator/internal/storage"
)

// Service
type Service struct {
	storage storage.Repositories
}

// Returns a new service
func NewService(storage storage.Repositories) *Service {
	return &Service{
		storage: storage,
	}
}

// Metric update service
func (service *Service) MetricUpdateService(metric *models.Metrics) error {
	switch metric.MType {
	case "counter":
		service.storage.SetCounter(metric.ID, metric.GetCounter())
		counter, _ := service.storage.GetCounter(metric.ID)
		metric.Delta = &counter
	case "gauge":
		service.storage.SetGauge(metric.ID, metric.GetGauge())
		gauge, _ := service.storage.GetGauge(metric.ID)
		metric.Value = &gauge
	}

	return nil
}

// Batch metric update service
func (service *Service) BatchMetricUpdateService(metrics []*models.Metrics) error {
	if err := service.storage.SetBatchOfMetrics(); err != nil {
		return err
	}

	return nil
}
