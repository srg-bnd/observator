// Handlers for server
package services

import (
	"github.com/srg-bnd/observator/internal/server/models"
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

func (service *Service) MetricUpdateService(metric *models.Metrics) error {
	switch metric.MType {
	case "counter":
		service.storage.SetCounter(metric.ID, metric.GetCounter())
	case "gauge":
		service.storage.SetGauge(metric.ID, metric.GetGauge())
	}

	return nil
}
