// Show Handlers (Metrics)
package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/srg-bnd/observator/internal/server/models"
)

const (
	notExistError = "notExistError"
)

// GET /value/{metricType}/{metricName}
func (h *Handler) ShowHandler(w http.ResponseWriter, r *http.Request) {
	setContentType(w, TextFormat)

	metrics, err := findMetricsForShow(h, r, TextFormat)
	if err != nil {
		handleErrorForShow(w, err)
		return
	}

	if err != representForShow(w, r, metrics, TextFormat) {
		handleErrorForShow(w, err)
		return
	}
}

// GET /value
func (h *Handler) ShowAsJSONHandler(w http.ResponseWriter, r *http.Request) {
	setContentType(w, JSONFormat)

	metrics, err := findMetricsForShow(h, r, JSONFormat)
	if err != nil {
		handleErrorForShow(w, err)
		return
	}

	if err != representForShow(w, r, metrics, JSONFormat) {
		handleErrorForShow(w, err)
		return
	}
}

// Helpers

func findMetricsForShow(h *Handler, r *http.Request, format string) (*models.Metrics, error) {
	var metrics models.Metrics

	// Build metrics
	if format == JSONFormat {
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
			if format == JSONFormat {
				zero := int64(0)
				metrics.Delta = &zero
			} else {
				return &metrics, errors.New(notExistError)
			}
		} else {
			metrics.Delta = &delta
		}
	case "gauge":
		value, err := h.storage.GetGauge(metrics.ID)
		if err != nil {
			if format == JSONFormat {
				zero := float64(0)
				metrics.Value = &zero
			} else {
				return &metrics, errors.New(notExistError)
			}
		} else {
			metrics.Value = &value
		}
	default:
		return &metrics, errors.New(notExistError)
	}

	return &metrics, nil
}

func representForShow(w http.ResponseWriter, _ *http.Request, metrics *models.Metrics, format string) error {
	var body []byte

	if format == JSONFormat {
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
