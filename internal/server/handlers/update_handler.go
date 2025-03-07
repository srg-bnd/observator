// Update Handlers (Metrics)
package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"slices"
	"strconv"
	"strings"

	"github.com/go-chi/chi"
	"github.com/srg-bnd/observator/internal/server/models"
)

const (
	invalidNameError = "invalidNameError"
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

	if err != representForUpdate(h, w, r, metrics) {
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

	if err != representForUpdate(h, w, r, metrics) {
		handleErrorForUpdate(w, err)
		return
	}
}

// Helpers

func parseAndValidateMetricsForUpdate(_ *Handler, r *http.Request) (*models.Metrics, error) {
	var metricsValue string
	metrics := models.Metrics{}

	if strings.Contains(r.Header.Get("content-type"), "application/json") {
		var buf bytes.Buffer
		_, err := buf.ReadFrom(r.Body)
		if err != nil {
			return &metrics, errors.New(invalidDataError)
		}

		// TODO: use `json.NewDecoder(req.Body).Decode`
		if err = json.Unmarshal(buf.Bytes(), &metrics); err != nil {
			return &metrics, errors.New(invalidDataError)
		}
	} else {
		metrics.MType = chi.URLParam(r, "metricType")
		metrics.ID = chi.URLParam(r, "metricName")
	}

	// Check type
	if !slices.Contains(models.MetricsMTypes, metrics.MType) {
		return &metrics, errors.New(invalidDataError)
	}

	// Check name
	if metrics.ID == "" {
		return &metrics, errors.New(invalidNameError)
	}

	// Check value
	if !strings.Contains(r.Header.Get("content-type"), "application/json") {
		metricsValue = chi.URLParam(r, "metricValue")

		switch metrics.MType {
		case "counter":
			value, err := strconv.ParseInt(metricsValue, 10, 64)
			if err != nil {
				return nil, errors.New(invalidDataError)
			}

			metrics.SetCounter(value)
		case "gauge":
			value, err := strconv.ParseFloat(metricsValue, 64)
			if err != nil {
				return nil, errors.New(invalidDataError)
			}

			metrics.SetGauge(value)
		}
	}

	return &metrics, nil
}

func processForUpdate(h *Handler, _ *http.Request, metrics *models.Metrics) error {
	if err := h.service.MetricUpdateService(metrics); err != nil {
		return errors.New(serverError)
	}

	return nil
}

func representForUpdate(_ *Handler, w http.ResponseWriter, r *http.Request, metrics *models.Metrics) error {
	if strings.Contains(r.Header.Get("content-type"), "application/json") {
		return representForShow(w, r, metrics)
	}

	w.WriteHeader(http.StatusOK)
	return nil
}

func handleErrorForUpdate(w http.ResponseWriter, err error) {
	switch err.Error() {
	case invalidNameError:
		w.WriteHeader(http.StatusNotFound)
	default:
		handleError(w, err)
	}
}
