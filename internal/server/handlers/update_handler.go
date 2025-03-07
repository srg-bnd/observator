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

const (
	typeError  = "typeError"
	valueError = "valueError"
	nameError  = "nameError"
)

// POST /update/{metricType}/{metricName}/{metricValue}
func (h *Handler) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/plain; charset=utf-8")

	metrics, err := parseAndValidateMetricsForUpdate(h, r)
	if err != nil {
		handleErrorForUpdate(w, err)
		return
	}

	if err != processForUpdate(h, r, metrics) {
		handleErrorForUpdate(w, err)
		return
	}

	if err != representForUpdate(h, w, r) {
		handleErrorForUpdate(w, err)
		return
	}
}

// POST /update
func (h *Handler) UpdateAsJSONHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	metrics, err := parseAndValidateMetricsForUpdate(h, r)
	if err != nil {
		handleErrorForUpdate(w, err)
		return
	}

	if err != processForUpdate(h, r, metrics) {
		handleErrorForUpdate(w, err)
		return
	}

	if err != representForUpdate(h, w, r) {
		handleErrorForUpdate(w, err)
		return
	}
}

// Helpers

func parseAndValidateMetricsForUpdate(_ *Handler, r *http.Request) (*models.Metrics, error) {
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

func processForUpdate(h *Handler, _ *http.Request, metrics *models.Metrics) error {
	if err := h.service.MetricUpdateService(metrics); err != nil {
		return errors.New(serverError)
	}

	return nil
}

func representForUpdate(_ *Handler, w http.ResponseWriter, _ *http.Request) error {
	w.WriteHeader(http.StatusOK)

	return nil
}

func handleErrorForUpdate(w http.ResponseWriter, err error) {
	switch err.Error() {
	case typeError, valueError:
		w.WriteHeader(http.StatusBadRequest)
	case nameError:
		w.WriteHeader(http.StatusNotFound)
	default:
		handleError(w, err)
	}
}
