// Services (Metrics)
package services

import (
	"github.com/srg-bnd/observator/internal/server/models"
)

func (service *Service) MetricUpdateService(metric *models.Metrics) error {
	switch metric.MType {
	case "counter":
		service.storage.SetCounter(metric.ID, metric.GetCounter())
	case "gauge":
		service.storage.SetGauge(metric.ID, metric.GetGauge())
	}

	return nil
}
