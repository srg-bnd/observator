// Show Handlers (Metrics)
package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/srg-bnd/observator/internal/server/models"
)

// GET /value/{metricType}/{metricName}
func (h *Handler) ShowHandler(w http.ResponseWriter, r *http.Request) {
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

// GET /value
func (h *Handler) ShowAsJSONHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	var metrics models.Metrics
	var buf bytes.Buffer
	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// TODO: use `json.NewDecoder(req.Body).Decode`
	if err = json.Unmarshal(buf.Bytes(), &metrics); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	switch metrics.MType {
	case "counter":
		value, err := h.storage.GetCounter(metrics.ID)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		metrics.Delta = &value
		resp, err := json.Marshal(metrics)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(resp)
	case "gauge":
		value, err := h.storage.GetGauge(metrics.ID)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		metrics.Value = &value
		resp, err := json.Marshal(metrics)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(resp)
	default:
		{
			w.WriteHeader(http.StatusNotFound)
		}
	}
}
