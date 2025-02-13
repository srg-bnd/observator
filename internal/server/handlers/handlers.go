// Handlers for server
package handlers

import (
	"errors"
	"net/http"
	"slices"
	"strconv"

	"github.com/srg-bnd/observator/internal/server/models"
	"github.com/srg-bnd/observator/internal/server/services"
	"github.com/srg-bnd/observator/internal/storage"
)

type Handler struct {
	service *services.Service
}

func NewHandler(storage storage.Repositories) *Handler {
	return &Handler{
		service: services.NewService(storage),
	}
}

/* UpdateMetricHandler */

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

	// Check type
	if !slices.Contains([]string{"counter", "gauge"}, r.PathValue("metricType")) {
		return nil, errors.New("typeError")
	}
	metric.Type = r.PathValue("metricType")

	// Check name
	if r.PathValue("metricName") == "" {
		return nil, errors.New("nameError")
	}
	metric.Name = r.PathValue("metricName")

	// Check value
	switch metric.Type {
	case "counter":
		value, err := strconv.ParseInt(r.PathValue("metricValue"), 10, 64)
		if err != nil {
			return nil, errors.New("valueError")
		}

		metric.SetCounter(value)
	case "gauge":
		value, err := strconv.ParseFloat(r.PathValue("metricValue"), 64)
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
