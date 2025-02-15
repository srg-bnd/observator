// Handlers for server
package handlers

import (
	"log"
	"net/http"
	"strconv"
	"text/template"

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

type IndexData struct {
	Metrics []Metric
}

type Metric struct {
	Name  string
	Value string
}

func (h *Handler) IndexHandler(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles("web/server/index.html")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	metrics := make([]Metric, 0)

	for _, metric := range models.TrackedMetrics["counter"] {
		value, err := h.storage.GetCounter(metric)
		if err == nil {
			metrics = append(metrics, Metric{Name: metric, Value: strconv.FormatInt(value, 10)})
		}
	}

	for _, metric := range models.TrackedMetrics["gauge"] {
		value, err := h.storage.GetGauge(metric)
		if err == nil {
			metrics = append(metrics, Metric{Name: metric, Value: strconv.FormatFloat(value, 'f', -1, 64)})
		}
	}

	w.Header().Set("content-type", "text/html; charset=utf-8")
	err = html.Execute(w, IndexData{Metrics: metrics})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
