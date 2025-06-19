package services

import (
	"context"
	"fmt"
	"testing"

	"github.com/srg-bnd/observator/internal/agent/collector"
	"github.com/srg-bnd/observator/internal/agent/models"
	"github.com/srg-bnd/observator/internal/storage"
	"github.com/stretchr/testify/assert"
)

type MockClient struct {
	metrics []models.Metrics
}

func (c *MockClient) SendMetrics(ctx context.Context, metrics []models.Metrics) error {
	c.metrics = metrics
	return nil
}

func TestNewService(t *testing.T) {
	storage := storage.NewMemStorage()
	service := NewMetricService(storage, nil)
	assert.IsType(t, service, &MetricService{})
}

func TestUpdate(t *testing.T) {
	ctx := context.Background()
	repository := storage.NewMemStorage()
	metrics := collector.Metrics{
		Counter: map[string]int64{"firstKey": 1},
		Gauge:   map[string]float64{"firstKey": 1.0},
	}

	service := NewMetricService(repository, nil)
	err := service.Update(ctx, &metrics)
	assert.Nil(t, err)

	for key, value := range metrics.Counter {
		valueFromRepository, _ := repository.GetCounter(ctx, key)
		assert.Equal(t, valueFromRepository, value)
	}

	for key, value := range metrics.Gauge {
		valueFromRepository, _ := repository.GetGauge(ctx, key)
		assert.Equal(t, valueFromRepository, value)
	}
}

func TestSend(t *testing.T) {
	ctx := context.Background()
	repository := storage.NewMemStorage()

	repository.SetCounter(ctx, "one", 1)
	repository.SetCounter(ctx, "two", 2)
	repository.SetGauge(ctx, "one", 1.0)
	repository.SetGauge(ctx, "two", 2.0)

	testCases := []struct {
		name         string
		CounterMType []string
		GaugeMType   []string
	}{
		{
			name:         "tracked metrics not empty",
			CounterMType: []string{"one", "two"},
			GaugeMType:   []string{"one"},
		},
		{
			name: "tracked metrics empty",
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprint("Test ", i+1, ": ", tc.name), func(t *testing.T) {
			client := MockClient{}
			service := NewMetricService(repository, &MockClient{})
			err := service.Send(ctx, map[string][]string{
				models.CounterMType: tc.CounterMType,
				models.GaugeMType:   tc.GaugeMType,
			})
			assert.Nil(t, err)

			for _, metric := range client.metrics {
				if metric.MType == models.CounterMType {
					assert.Contains(t, tc.CounterMType, metric.ID)
				} else {
					assert.Contains(t, tc.GaugeMType, metric.ID)
				}
			}
		})
	}
}
