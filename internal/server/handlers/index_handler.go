// Index Handlers (Metrics)
package handlers

import (
	"net/http"

	"github.com/srg-bnd/observator/internal/server/models"
)

// GET /
func (h *Handler) IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html; charset=utf-8")

	metricsByMType, err := getMetricsByMTypeForIndex(h, r)
	if err != nil {
		handleErrorForIndex(w, err)
		return
	}

	if err != representMetricsByMTypeForIndex(w, r, metricsByMType) {
		handleErrorForIndex(w, err)
		return
	}
}

// Helpers

func getMetricsByMTypeForIndex(h *Handler, _ *http.Request) (map[string][]models.Metrics, error) {
	metricsByMType := make(map[string][]models.Metrics, 0)
	metricsByMType["counter"] = make([]models.Metrics, 0)
	metricsByMType["gauge"] = make([]models.Metrics, 0)

	for _, ID := range models.TrackedMetrics["counter"] {
		counter, err := h.storage.GetCounter(ID)
		if err == nil {
			metricsByMType["counter"] = append(metricsByMType["counter"], models.Metrics{MType: "counter", ID: ID, Delta: &counter})
		}
	}

	for _, ID := range models.TrackedMetrics["gauge"] {
		gauge, err := h.storage.GetGauge(ID)
		if err == nil {
			metricsByMType["counter"] = append(metricsByMType["counter"], models.Metrics{MType: "counter", ID: ID, Value: &gauge})
		}
	}

	return metricsByMType, nil
}

func representMetricsByMTypeForIndex(w http.ResponseWriter, _ *http.Request, metricsByMType map[string][]models.Metrics) error {
	var body string

	for MType, allMetrics := range metricsByMType {
		body += "<h1>" + MType + ":</h1><ul>"
		for _, metrics := range allMetrics {
			body += "<li>" + metrics.ID + ": " + metrics.GetCounterAsString() + "</li>"
		}
		body += "</ul>"
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(body))

	return nil
}

func handleErrorForIndex(w http.ResponseWriter, err error) {
	// TODO
}
