// Index Handlers (Metrics)
package handlers

import (
	"net/http"

	"github.com/srg-bnd/observator/internal/server/models"
)

// Index repository
type IndexRepository interface {
	GetGauge(string) (float64, error)
	GetCounter(string) (int64, error)
}

// Index handler
type IndexHandler struct {
	repository IndexRepository
}

// Returns new index handler
func NewIndexHandler(repository IndexRepository) *IndexHandler {
	return &IndexHandler{
		repository: repository,
	}
}

// Processes the request
func (h *IndexHandler) Handler(w http.ResponseWriter, r *http.Request) {
	setContentType(w, HTMLFormat)

	metricsByMType, err := h.getMetricsByMType()
	if err != nil {
		handleError(w, err)
		return
	}

	if err != h.represent(w, metricsByMType) {
		handleError(w, err)
		return
	}
}

// Returns metrics from repository
func (h *IndexHandler) getMetricsByMType() (map[string][]models.Metrics, error) {
	metricsByMType := make(map[string][]models.Metrics, 0)
	metricsByMType["counter"] = make([]models.Metrics, 0)
	metricsByMType["gauge"] = make([]models.Metrics, 0)

	for _, ID := range models.TrackedMetrics["counter"] {
		counter, err := h.repository.GetCounter(ID)
		if err == nil {
			metricsByMType["counter"] = append(metricsByMType["counter"], models.Metrics{MType: "counter", ID: ID, Delta: &counter})
		}
	}

	for _, ID := range models.TrackedMetrics["gauge"] {
		gauge, err := h.repository.GetGauge(ID)
		if err == nil {
			metricsByMType["gauge"] = append(metricsByMType["gauge"], models.Metrics{MType: "gauge", ID: ID, Value: &gauge})
		}
	}

	return metricsByMType, nil
}

// Generates a response
func (h *IndexHandler) represent(w http.ResponseWriter, metricsByMType map[string][]models.Metrics) error {
	var body string

	for MType, allMetrics := range metricsByMType {
		body += "<h1>" + MType + ":</h1><ul>"
		for _, metrics := range allMetrics {
			if MType == "gauge" {
				body += "<li>" + metrics.ID + ": " + metrics.GetGaugeAsString() + "</li>"
			} else {
				body += "<li>" + metrics.ID + ": " + metrics.GetCounterAsString() + "</li>"
			}
		}
		body += "</ul>"
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(body))

	return nil
}
