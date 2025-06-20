// Index Handler
package handlers

import (
	"context"
	"net/http"

	"github.com/srg-bnd/observator/internal/server/models"
)

// Index repository
type IndexRepository interface {
	AllCounterMetrics(context.Context) (map[string]int64, error)
	AllGaugeMetrics(context.Context) (map[string]float64, error)
}

type MetricsByMType struct {
	counter []models.Metrics
	gauge   []models.Metrics
}

// Index handler
type IndexHandler struct {
	repository IndexRepository
}

// Returns new index handler
func NewIndexHandler(repository IndexRepository) *IndexHandler {
	return &IndexHandler{
		repository: repository,
	}
}

// Processes the request
func (h *IndexHandler) Handler(w http.ResponseWriter, r *http.Request) {
	setContentType(w, HTMLFormat)

	metricsByMType, err := h.getAllMetricsByMType(r.Context())
	if err != nil {
		handleError(w, err)
		return
	}

	if err != h.represent(w, metricsByMType) {
		handleError(w, err)
		return
	}
}

// Returns metrics from repository
func (h *IndexHandler) getAllMetricsByMType(ctx context.Context) (*MetricsByMType, error) {
	metricsByMType := MetricsByMType{
		counter: make([]models.Metrics, 0),
		gauge:   make([]models.Metrics, 0),
	}

	if allCounterMetrics, err := h.repository.AllCounterMetrics(ctx); err != nil {
		return nil, err
	} else {
		for ID, delta := range allCounterMetrics {
			metricsByMType.counter = append(metricsByMType.counter, models.Metrics{MType: models.CounterMType, ID: ID, Delta: &delta})
		}
	}

	if allGaugeMetrics, err := h.repository.AllGaugeMetrics(ctx); err != nil {
		return nil, err
	} else {
		for ID, value := range allGaugeMetrics {
			metricsByMType.gauge = append(metricsByMType.gauge, models.Metrics{MType: models.GaugeMType, ID: ID, Value: &value})
		}
	}

	return &metricsByMType, nil
}

// Generates a response
func (h *IndexHandler) represent(w http.ResponseWriter, metricsByMType *MetricsByMType) error {
	var body string

	body += htmlFormatForMetrics(metricsByMType.counter)
	body += htmlFormatForMetrics(metricsByMType.gauge)
	if body == "" {
		body = "<div>Empty</div>"
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(body))

	return nil
}

// Helpers

func htmlFormatForMetrics(metrics []models.Metrics) string {
	var body string

	if len(metrics) == 0 {
		return ""
	}

	body += "<div><h2>" + metrics[0].MType + " metrics:</h2><ul>"
	for _, metric := range metrics {
		body += "<li>" + metric.ID + ": " + metric.GetValueAsString() + "</li>"
	}
	body += "</ul></div"

	return body
}
