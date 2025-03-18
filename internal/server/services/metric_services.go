// Services (Metrics)
package services

import (
	"github.com/srg-bnd/observator/internal/server/models"
)

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
