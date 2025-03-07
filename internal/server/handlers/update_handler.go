// Update Handlers (Metrics)
package handlers

import (
	"errors"
	"net/http"
	"slices"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/srg-bnd/observator/internal/server/models"
)

// POST /update/{metricType}/{metricName}/{metricValue}
func (h *Handler) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	metrics, err := h.parseAndValidateMetricsForUpdate(r)
	if err != nil {
		h.handleErrorForUpdate(w, err)
		return
	}

	h.processMetricsForUpdate(w, metrics)
}

// POST /update
func (h *Handler) UpdateAsJSONHandler(w http.ResponseWriter, r *http.Request) {
	metrics, err := h.parseAndValidateMetricsForUpdate(r)
	if err != nil {
		h.handleErrorForUpdate(w, err)
		return
	}

	h.processMetricsForUpdate(w, metrics)
}

// Helpers

func (h *Handler) parseAndValidateMetricsForUpdate(r *http.Request) (*models.Metrics, error) {
	metrics := models.Metrics{}

	metricType := chi.URLParam(r, "metricType")
	metricName := chi.URLParam(r, "metricName")
	metricValue := chi.URLParam(r, "metricValue")

	// Check type
	if !slices.Contains([]string{"counter", "gauge"}, metricType) {
		return nil, errors.New("typeError")
	}
	metrics.MType = metricType

	// Check name
	if metricName == "" {
		return nil, errors.New("nameError")
	}
	metrics.ID = metricName

	// Check value
	switch metrics.MType {
	case "counter":
		value, err := strconv.ParseInt(metricValue, 10, 64)
		if err != nil {
			return nil, errors.New("valueError")
		}

		metrics.SetCounter(value)
	case "gauge":
		value, err := strconv.ParseFloat(metricValue, 64)
		if err != nil {
			return nil, errors.New("valueError")
		}

		metrics.SetGauge(value)
	}

	return &metrics, nil
}

func (h *Handler) processMetricsForUpdate(w http.ResponseWriter, metrics *models.Metrics) {
	err := h.service.MetricUpdateService(metrics)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.Header().Set("content-type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
	}
}

func (h *Handler) handleErrorForUpdate(w http.ResponseWriter, err error) {
	switch err.Error() {
	case "typeError", "valueError":
		w.WriteHeader(http.StatusBadRequest)
	case "nameError":
		w.WriteHeader(http.StatusNotFound)
	}
}
