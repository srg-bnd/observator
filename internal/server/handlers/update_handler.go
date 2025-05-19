// Update Handler
package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"slices"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/srg-bnd/observator/internal/server/models"
)

// Show repository
type UpdateRepository interface {
	SetGauge(string, float64) error
	GetGauge(string) (float64, error)
	SetCounter(string, int64) error
	GetCounter(string) (int64, error)
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

	if err != h.process(metric) {
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

	if err != h.process(metric) {
		handleError(w, err)
		return
	}

	if err != h.represent(w, metric, JSONFormat) {
		handleError(w, err)
		return
	}
}

// Parses and validate metric
func (h *UpdateHandler) parseAndValidateMetric(r *http.Request, format string) (*models.Metrics, error) {
	var metricValue string
	metric := models.Metrics{}

	if format == JSONFormat {
		var buf bytes.Buffer
		_, err := buf.ReadFrom(r.Body)

		if err != nil {
			return &metric, errors.New(invalidDataError)
		}

		// TODO: use `json.NewDecoder(req.Body).Decode`
		if err = json.Unmarshal(buf.Bytes(), &metric); err != nil {
			return &metric, errors.New(invalidDataError)
		}
	} else {
		metric.MType = chi.URLParam(r, "metricType")
		metric.ID = chi.URLParam(r, "metricName")
	}

	// Check type
	if !slices.Contains(models.MetricsMTypes, metric.MType) {
		return &metric, errors.New(invalidDataError)
	}

	// Check name
	if metric.ID == "" {
		return &metric, errors.New(invalidNameError)
	}

	// Check value
	if format != JSONFormat {
		metricValue = chi.URLParam(r, "metricValue")

		switch metric.MType {
		case "counter":
			value, err := strconv.ParseInt(metricValue, 10, 64)
			if err != nil {
				return nil, errors.New(invalidDataError)
			}

			metric.SetCounter(value)
		case "gauge":
			value, err := strconv.ParseFloat(metricValue, 64)
			if err != nil {
				return nil, errors.New(invalidDataError)
			}

			metric.SetGauge(value)
		}
	}

	return &metric, nil
}

func (h *UpdateHandler) process(metric *models.Metrics) error {
	switch metric.MType {
	case "counter":
		h.repository.SetCounter(metric.ID, metric.GetCounter())
		counter, _ := h.repository.GetCounter(metric.ID)
		metric.Delta = &counter
	case "gauge":
		h.repository.SetGauge(metric.ID, metric.GetGauge())
		gauge, _ := h.repository.GetGauge(metric.ID)
		metric.Value = &gauge
	}

	return nil
}

// Generates a response
func (h *UpdateHandler) represent(w http.ResponseWriter, metric *models.Metrics, contentType string) error {
	if contentType == JSONFormat {
		var body []byte

		data, err := json.Marshal(metric)
		if err != nil {
			return errors.New(serverError)
		}
		body = data

		w.Write(body)
	}

	w.WriteHeader(http.StatusOK)
	return nil
}
