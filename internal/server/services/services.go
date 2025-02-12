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

type Result struct {
	Ok     bool
	Status string
}

func NewResult(ok bool, status string) *Result {
	return &Result{Ok: ok, Status: status}
}

func (service *Service) UpdateMetricService(metric *models.Metric) *Result {
	switch metric.Type {
	case "counter":
		service.storage.SetCounter(metric.Name, metric.GetCounter())
	case "gauge":
		service.storage.SetGauge(metric.Name, metric.GetGauge())
	}

	return NewResult(true, "Ok")
}
