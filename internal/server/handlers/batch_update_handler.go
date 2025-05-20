// Updates Handler
package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/srg-bnd/observator/internal/server/models"
)

// Updates repository
type BatchUpdateRepository interface {
	SetBatchOfMetrics(map[string]int64, map[string]float64) error
}

// Updates handler
type BatchUpdateHandler struct {
	repository BatchUpdateRepository
}

// Returns new update handler
func NewBatchUpdateHandler(repository BatchUpdateRepository) *BatchUpdateHandler {
	return &BatchUpdateHandler{
		repository: repository,
	}
}

// POST /updates
func (h *BatchUpdateHandler) Handler(w http.ResponseWriter, r *http.Request) {
	setContentType(w, JSONFormat)

	// Parse and validate metrics
	metrics := make([]models.Metrics, 0)

	var buf bytes.Buffer
	_, err := buf.ReadFrom(r.Body)

	if err != nil {
		handleError(w, errors.New(invalidDataError))
		return
	}

	if err = json.Unmarshal(buf.Bytes(), &metrics); err != nil {
		handleError(w, errors.New(invalidDataError))
		return
	}

	// Updates metrics in storage
	counterMetrics := make(map[string]int64, 0)
	gaugeMetrics := make(map[string]float64, 0)

	for _, metric := range metrics {
		if metric.MType == "counter" {
			counterMetrics[metric.ID] = *metric.Delta
		} else {
			gaugeMetrics[metric.ID] = *metric.Value
		}
	}

	if err := h.repository.SetBatchOfMetrics(counterMetrics, gaugeMetrics); err != nil {
		handleError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
