// Show Handler
package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/srg-bnd/observator/internal/server/models"
)

// Show repository
type ShowRepository interface {
	GetGauge(context.Context, string) (float64, error)
	GetCounter(context.Context, string) (int64, error)
}

// Show handler
type ShowHandler struct {
	repository ShowRepository
}

// Returns new show handler
func NewShowHandler(repository ShowRepository) *ShowHandler {
	return &ShowHandler{
		repository: repository,
	}
}

// Processes the request GET /value/{metricType}/{metricName}
func (h *ShowHandler) Handler(w http.ResponseWriter, r *http.Request) {
	setContentType(w, TextFormat)

	metric, err := h.findMetric(r, TextFormat)
	if err != nil {
		handleError(w, err)
		return
	}

	if err != h.represent(w, metric, TextFormat) {
		handleError(w, err)
		return
	}
}

// Processes the request GET /value (JSON)
func (h *ShowHandler) JSONHandler(w http.ResponseWriter, r *http.Request) {
	setContentType(w, JSONFormat)

	metric, err := h.findMetric(r, JSONFormat)
	if err != nil {
		handleError(w, err)
		return
	}

	if err != h.represent(w, metric, JSONFormat) {
		handleError(w, err)
		return
	}
}

// Returns metric from repository
func (h *ShowHandler) findMetric(r *http.Request, format string) (*models.Metrics, error) {
	metric, err := h.buildMetric(r, format)
	if err != nil {
		return metric, err
	}

	// Sets the value for the metric
	switch {
	case metric.IsCounterMType():
		err := h.setCounterValue(metric, r, format)
		if err != nil {
			return metric, err
		}
	case metric.IsGaugeMType():
		err := h.setGaugeValue(metric, r, format)
		if err != nil {
			return metric, err
		}
	default:
		return metric, errNotFound
	}

	return metric, nil
}

// Generates a response
func (h *ShowHandler) represent(w http.ResponseWriter, metric *models.Metrics, format string) error {
	var body []byte

	if format == JSONFormat {
		data, err := json.Marshal(metric)
		if err != nil {
			return errServer
		}
		body = data
	} else {
		body = []byte(metric.GetValueAsString())
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)

	return nil
}

// Builds a metric
func (h *ShowHandler) buildMetric(r *http.Request, format string) (*models.Metrics, error) {
	var metric models.Metrics

	if format == JSONFormat {
		var buf bytes.Buffer
		_, err := buf.ReadFrom(r.Body)

		if err != nil {
			return &metric, errInvalidData
		}

		// TODO: use `json.NewDecoder(req.Body).Decode`
		if err = json.Unmarshal(buf.Bytes(), &metric); err != nil {
			return &metric, errInvalidData
		}
	} else {
		metric.MType = chi.URLParam(r, "metricType")
		metric.ID = chi.URLParam(r, "metricName")
	}

	return &metric, nil
}

// Sets the counter value for the metric
func (h *ShowHandler) setCounterValue(metric *models.Metrics, r *http.Request, format string) error {
	delta, err := h.repository.GetCounter(r.Context(), metric.ID)
	if err != nil {
		if format == JSONFormat {
			var newDelta int64
			if metric.Delta == nil {
				newDelta = int64(0)
			} else {
				newDelta = int64(*metric.Delta)
			}
			metric.Delta = &newDelta
		} else {
			return errNotFound
		}
	} else {
		metric.Delta = &delta
	}

	return nil
}

// Sets the gauge value for the metric
func (h *ShowHandler) setGaugeValue(metric *models.Metrics, r *http.Request, format string) error {
	value, err := h.repository.GetGauge(r.Context(), metric.ID)
	if err != nil {
		if format == JSONFormat {
			var newValue float64
			if metric.Value == nil {
				newValue = float64(0)
			} else {
				newValue = float64(*metric.Value)
			}
			metric.Value = &newValue
		} else {
			return errNotFound
		}
	} else {
		metric.Value = &value
	}

	return nil
}
