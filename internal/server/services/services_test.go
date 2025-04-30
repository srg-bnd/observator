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
	metrics := models.NewMetrics()
	metrics.MType = "counter"
	metrics.ID = "metric"
	metrics.SetCounter(1)

	err := service.MetricUpdateService(metrics)
	assert.Nil(t, err)
	value, _ := storage.GetCounter(metrics.ID)
	assert.Equal(t, value, metrics.GetCounter())
}

func TestMetricUpdateServiceForGauge(t *testing.T) {
	storage := storage.NewMemStorage()
	service := NewService(storage)
	metrics := models.NewMetrics()
	metrics.MType = "gauge"
	metrics.ID = "metric"
	metrics.SetGauge(1.0)

	err := service.MetricUpdateService(metrics)
	assert.Nil(t, err)
	value, _ := storage.GetGauge(metrics.ID)
	assert.Equal(t, value, metrics.GetGauge())
}
