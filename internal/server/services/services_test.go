package services

import (
	"testing"

	"github.com/srg-bnd/observator/internal/server/models"
	"github.com/srg-bnd/observator/internal/storage"
	"github.com/stretchr/testify/assert"
)

func TestNewService(t *testing.T) {
	storage := storage.NewMemStorage()
	service := NewService(storage)
	assert.IsType(t, service, &Service{})
}

func TestNewResult(t *testing.T) {
	result := NewResult(true, "OK")
	assert.IsType(t, result, &Result{})
}

func TestUpdateMetricServiceForCounter(t *testing.T) {
	storage := storage.NewMemStorage()
	service := NewService(storage)
	metric := models.NewMetric()
	metric.Type = "counter"
	metric.Name = "metric"
	metric.SetCounter(1)

	result := service.UpdateMetricService(metric)
	assert.Equal(t, result.Ok, true)
	assert.Equal(t, storage.GetCounter(metric.Name), metric.GetCounter())
}

func TestUpdateMetricServiceForGauge(t *testing.T) {
	storage := storage.NewMemStorage()
	service := NewService(storage)
	metric := models.NewMetric()
	metric.Type = "gauge"
	metric.Name = "metric"
	metric.SetCounter(1.0)

	result := service.UpdateMetricService(metric)
	assert.Equal(t, result.Ok, true)
	assert.Equal(t, storage.GetGauge(metric.Name), metric.GetGauge())
}
