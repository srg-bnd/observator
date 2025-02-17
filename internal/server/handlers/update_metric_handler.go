// UpdateMetricHandler for server
package handlers

import (
	"errors"
	"net/http"
	"slices"
	"strconv"
	"strings"

	"github.com/srg-bnd/observator/internal/server/models"
)

func (h *Handler) UpdateMetricHandler(w http.ResponseWriter, r *http.Request) {
	metric, err := h.parseAndValidateMetric(r)
	if err != nil {
		h.handleError(w, err)
		return
	}

	h.processMetric(w, metric)
}

func (h *Handler) parseAndValidateMetric(r *http.Request) (*models.Metric, error) {
	metric := models.Metric{}

	path := strings.Split(r.URL.Path, "/")
	if len(path) < 4 {
		return nil, errors.New("typeError")
	}
	metricType := path[2]
	metricName := path[3]
	metricValue := path[4]

	// Check type
	if !slices.Contains([]string{"counter", "gauge"}, metricType) {
		return nil, errors.New("typeError")
	}
	metric.Type = metricType

	// Check name
	if metricName == "" {
		return nil, errors.New("nameError")
	}
	metric.Name = metricName

	// Check value
	switch metric.Type {
	case "counter":
		value, err := strconv.ParseInt(metricValue, 10, 64)
		if err != nil {
			return nil, errors.New("valueError")
		}

		metric.SetCounter(value)
	case "gauge":
		value, err := strconv.ParseFloat(metricValue, 64)
		if err != nil {
			return nil, errors.New("valueError")
		}

		metric.SetGauge(value)
	}

	return &metric, nil
}

func (h *Handler) processMetric(w http.ResponseWriter, metric *models.Metric) {
	err := h.service.MetricUpdateService(metric)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.Header().Set("content-type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
	}
}

func (h *Handler) handleError(w http.ResponseWriter, err error) {
	switch err.Error() {
	case "typeError", "valueError":
		w.WriteHeader(http.StatusBadRequest)
	case "nameError":
		w.WriteHeader(http.StatusNotFound)
	}
}
