// Show Handlers (Metrics)
package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	"github.com/srg-bnd/observator/internal/server/models"
)

const (
	invalidDataError = "invalidDataError"
	notExistError    = "notExistError"
)

// GET /value/{metricType}/{metricName}
func (h *Handler) ShowHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/plain; charset=utf-8")

	metrics, err := findMetricsForShow(h, r)
	if err != nil {
		handleErrorForShow(w, err)
		return
	}

	if err != representForShow(w, r, metrics) {
		handleErrorForShow(w, err)
		return
	}
}

// GET /value
func (h *Handler) ShowAsJSONHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	metrics, err := findMetricsForShow(h, r)
	if err != nil {
		handleErrorForShow(w, err)
		return
	}

	if err != representForShow(w, r, metrics) {
		handleErrorForShow(w, err)
		return
	}
}

// Helpers

func findMetricsForShow(h *Handler, r *http.Request) (*models.Metrics, error) {
	var metrics models.Metrics

	// Build metrics
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

	// Find metrics
	switch metrics.MType {
	case "counter":
		delta, err := h.storage.GetCounter(metrics.ID)
		if err != nil {
			return &metrics, errors.New(notExistError)
		}
		metrics.Delta = &delta
	case "gauge":
		value, err := h.storage.GetGauge(metrics.ID)
		if err != nil {
			return &metrics, errors.New(notExistError)
		}
		metrics.Value = &value
	default:
		return &metrics, errors.New(notExistError)
	}

	return &metrics, nil
}

func representForShow(w http.ResponseWriter, r *http.Request, metrics *models.Metrics) error {
	var body []byte

	if strings.Contains(r.Header.Get("content-type"), "application/json") {
		data, err := json.Marshal(metrics)
		if err != nil {
			return errors.New(serverError)
		}
		body = data
	} else {
		switch metrics.MType {
		case "counter":
			body = []byte(metrics.GetCounterAsString())
		case "gauge":
			body = []byte(metrics.GetGaugeAsString())
		}
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)

	return nil
}

func handleErrorForShow(w http.ResponseWriter, err error) {
	switch err.Error() {
	case invalidDataError:
		w.WriteHeader(http.StatusBadRequest)
	case notExistError:
		w.WriteHeader(http.StatusNotFound)
	default:
		handleError(w, err)
	}
}
