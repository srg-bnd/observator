// Handlers for server
package handlers

import (
	"net/http"
	"strconv"

	"github.com/srg-bnd/observator/internal/server/models"
	"github.com/srg-bnd/observator/internal/server/services"
	"github.com/srg-bnd/observator/internal/storage"
)

type Handler struct {
	service *services.Service
	storage storage.Repositories
}

func NewHandler(storage storage.Repositories) *Handler {
	return &Handler{
		service: services.NewService(storage),
		storage: storage,
	}
}

func (h *Handler) ShowMetricHandler(w http.ResponseWriter, r *http.Request) {
	switch r.PathValue("metricType") {
	case "counter":
		value, err := h.storage.GetCounter(r.PathValue("metricName"))
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.Header().Set("content-type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(strconv.FormatInt(value, 10)))
	case "gauge":
		value, err := h.storage.GetGauge(r.PathValue("metricName"))
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.Header().Set("content-type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(strconv.FormatFloat(value, 'f', -1, 64)))
	default:
		{
			w.WriteHeader(http.StatusNotFound)
		}
	}
}

func (h *Handler) IndexHandler(w http.ResponseWriter, r *http.Request) {
	body := "<html><head></head><body>"

	for _, metric := range models.TrackedMetrics["counter"] {
		value, _ := h.storage.GetCounter(metric)
		body += metric + ": "
		body += strconv.FormatInt(value, 10)
		body += "\n"
	}

	for _, metric := range models.TrackedMetrics["gauge"] {
		body += metric + ": "
		value, _ := h.storage.GetGauge(metric)
		body += strconv.FormatFloat(value, 'f', -1, 64)
		body += "\n</body></html>"
	}

	w.Header().Set("content-type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(body))
}
