// Handlers for server
package services

import (
	"slices"
	"strconv"

	"github.com/srg-bnd/observator/internal/storage"
)

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

var (
	MemStorage *storage.MemStorage
)

// Services

func UpdateMetricService(metricType, metricName, metricValue string) *Data {
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
			MemStorage.SetCounter(metricName, value)
		}
	case "gauge":
		// Check and set value
		value, err := strconv.ParseFloat(metricValue, 64)
		if err != nil {
			return errorData("valueError")
		} else {
			MemStorage.SetGauge(metricName, value)
		}
	}

	return successData("Ok")
}
