// Handlers for server
package services

import (
	"slices"
	"strconv"

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

type Data struct {
	Ok     bool
	Status string
}

func successData(status string) *Data {
	return &Data{Ok: true, Status: status}
}

func errorData(status string) *Data {
	return &Data{Ok: false, Status: status}
}

// Services

func (service *Service) UpdateMetricService(metricType, metricName, metricValue string) *Data {
	// Check type
	if !slices.Contains([]string{"counter", "gauge"}, metricType) {
		return errorData("typeError")
	}

	// Check name
	if metricName == "" {
		return errorData("nameError")
	}

	switch metricType {
	case "counter":
		// Check and set value
		value, err := strconv.ParseInt(metricValue, 10, 64)
		if err != nil {
			return errorData("valueError")
		} else {
			service.storage.SetCounter(metricName, value)
		}
	case "gauge":
		// Check and set value
		value, err := strconv.ParseFloat(metricValue, 64)
		if err != nil {
			return errorData("valueError")
		} else {
			service.storage.SetGauge(metricName, value)
		}
	}

	return successData("Ok")
}
