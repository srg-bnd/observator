// Index Handlers (Metrics)
package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/srg-bnd/observator/internal/server/models"
)

// GET /
func (h *Handler) IndexHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")
	if len(path) > 2 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	body := "<p>counter:</p><ul>"
	for _, metric := range models.TrackedMetrics["counter"] {
		value, err := h.storage.GetCounter(metric)
		if err == nil {
			body += "<li>" + metric + ": " + strconv.FormatInt(value, 10) + "</li>"
		}
	}
	body += "</ul>"

	body += "<p>gauge:</p><ul>"
	for _, metric := range models.TrackedMetrics["gauge"] {
		value, err := h.storage.GetGauge(metric)
		if err == nil {
			body += "<li>" + metric + ": " + strconv.FormatFloat(value, 'f', -1, 64) + "</li>"
		}
	}
	body += "</ul>"

	w.Header().Set("content-type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(body))
}
