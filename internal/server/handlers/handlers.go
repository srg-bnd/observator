// Handlers for server
package handlers

import (
	"net/http"
	"strconv"
	"strings"
	"text/template"

	"github.com/go-chi/chi"
	"github.com/srg-bnd/observator/internal/server/models"
	"github.com/srg-bnd/observator/internal/server/services"
	"github.com/srg-bnd/observator/internal/storage"
)

type Handler struct {
	service      *services.Service
	storage      storage.Repositories
	rootFilePath string
}

func NewHandler(storage storage.Repositories) *Handler {
	return &Handler{
		service:      services.NewService(storage),
		storage:      storage,
		rootFilePath: "web/",
	}
}

// Init router
func (h *Handler) GetRouter() chi.Router {
	r := chi.NewRouter()

	r.Post("/update/{metricType}/{metricName}/{metricValue}", h.UpdateMetricHandler)
	r.Get("/value/{metricType}/{metricName}", h.ShowMetricHandler)
	r.Get("/", h.IndexHandler)

	return r
}

func (h *Handler) UpdateMetricHandler(w http.ResponseWriter, r *http.Request) {
	metric, err := h.parseAndValidateMetric(r)
	if err != nil {
		h.handleError(w, err)
		return
	}

	h.processMetric(w, metric)
}

func (h *Handler) ShowMetricHandler(w http.ResponseWriter, r *http.Request) {
	metricType := chi.URLParam(r, "metricType")
	metricName := chi.URLParam(r, "metricName")

	switch metricType {
	case "counter":
		value, err := h.storage.GetCounter(metricName)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.Header().Set("content-type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(strconv.FormatInt(value, 10)))
	case "gauge":
		value, err := h.storage.GetGauge(metricName)
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

	path := strings.Split(r.URL.Path, "/")
	if len(path) > 2 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	html, err := template.ParseFiles(h.rootFilePath + "server/index.html")
	if err != nil {
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
