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

func TestMetricUpdateServiceForCounter(t *testing.T) {
	storage := storage.NewMemStorage()
	service := NewService(storage)
	metric := models.NewMetric()
	metric.Type = "counter"
	metric.Name = "metric"
	metric.SetCounter(1)

	err := service.MetricUpdateService(metric)
	assert.Nil(t, err)
	value, _ := storage.GetCounter(metric.Name)
	assert.Equal(t, value, metric.GetCounter())
}

func TestMetricUpdateServiceForGauge(t *testing.T) {
	storage := storage.NewMemStorage()
	service := NewService(storage)
	metric := models.NewMetric()
	metric.Type = "gauge"
	metric.Name = "metric"
	metric.SetCounter(1.0)

	err := service.MetricUpdateService(metric)
	assert.Nil(t, err)
	value, _ := storage.GetGauge(metric.Name)
	assert.Equal(t, value, metric.GetGauge())
}
