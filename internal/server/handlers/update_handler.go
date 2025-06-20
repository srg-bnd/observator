// Update Handler
package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"slices"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/srg-bnd/observator/internal/server/models"
)

// Reader repository
type ReaderRepository interface {
	GetGauge(context.Context, string) (float64, error)
	GetCounter(context.Context, string) (int64, error)
}

// Writer repository
type WriterRepository interface {
	SetGauge(context.Context, string, float64) error
	SetCounter(context.Context, string, int64) error
}

// Update repository
type UpdateRepository interface {
	ReaderRepository
	WriterRepository
}

// Update handler
type UpdateHandler struct {
	repository UpdateRepository
}

// Returns new update handler
func NewUpdateHandler(repository UpdateRepository) *UpdateHandler {
	return &UpdateHandler{
		repository: repository,
	}
}

// POST /update/{metricType}/{metricName}/{metricValue}
func (h *UpdateHandler) Handler(w http.ResponseWriter, r *http.Request) {
	setContentType(w, TextFormat)

	metric, err := h.parseAndValidateMetric(r, TextFormat)
	if err != nil {
		handleError(w, err)
		return
	}

	if err != h.process(metric, r) {
		handleError(w, err)
		return
	}

	if err != h.represent(w, metric, TextFormat) {
		handleError(w, err)
		return
	}
}

// POST /update
func (h *UpdateHandler) JSONHandler(w http.ResponseWriter, r *http.Request) {
	setContentType(w, JSONFormat)

	metric, err := h.parseAndValidateMetric(r, JSONFormat)
	if err != nil {
		handleError(w, err)
		return
	}

	if err != h.process(metric, r) {
		handleError(w, err)
		return
	}

	if err != h.represent(w, metric, JSONFormat) {
		handleError(w, err)
		return
	}
}

// Parses and validates metric
func (h *UpdateHandler) parseAndValidateMetric(r *http.Request, format string) (*models.Metrics, error) {
	var metricValue string
	metric := models.Metrics{}

	if format == JSONFormat {
		var buf bytes.Buffer
		_, err := buf.ReadFrom(r.Body)
		defer r.Body.Close()

		if err != nil {
			return &metric, invalidDataError
		}

		// TODO: use `json.NewDecoder(req.Body).Decode`
		if err = json.Unmarshal(buf.Bytes(), &metric); err != nil {
			return &metric, invalidDataError
		}
	} else {
		metric.MType = chi.URLParam(r, "metricType")
		metric.ID = chi.URLParam(r, "metricName")
	}

	// Check type
	if !slices.Contains(models.MetricsMTypes, metric.MType) {
		return &metric, invalidDataError
	}

	// Check name
	if metric.ID == "" {
		return &metric, invalidNameError
	}

	// Check value
	if format != JSONFormat {
		metricValue = chi.URLParam(r, "metricValue")

		switch metric.MType {
		case models.CounterMType:
			value, err := strconv.ParseInt(metricValue, 10, 64)
			if err != nil {
				return nil, invalidDataError
			}

			metric.SetCounter(value)
		case models.GaugeMType:
			value, err := strconv.ParseFloat(metricValue, 64)
			if err != nil {
				return nil, invalidDataError
			}

			metric.SetGauge(value)
		}
	}

	return &metric, nil
}

// Processes the metric
func (h *UpdateHandler) process(metric *models.Metrics, r *http.Request) error {
	switch metric.MType {
	case models.CounterMType:
		h.repository.SetCounter(r.Context(), metric.ID, metric.GetCounter())
		delta, err := h.repository.GetCounter(r.Context(), metric.ID)
		if err != nil {
			return nil
		}
		metric.Delta = &delta
	case models.GaugeMType:
		h.repository.SetGauge(r.Context(), metric.ID, metric.GetGauge())
		value, err := h.repository.GetGauge(r.Context(), metric.ID)
		if err != nil {
			return nil
		}
		metric.Value = &value
	}

	return nil
}

// Generates a response
func (h *UpdateHandler) represent(w http.ResponseWriter, metric *models.Metrics, contentType string) error {
	if contentType == JSONFormat {
		var body []byte

		data, err := json.Marshal(metric)
		if err != nil {
			return serverError
		}
		body = data

		w.WriteHeader(http.StatusOK)
		w.Write(body)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	return nil
}
