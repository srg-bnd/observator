package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/srg-bnd/observator/internal/server/models"
	"github.com/srg-bnd/observator/internal/storage"
	"github.com/stretchr/testify/assert"
)

func TestShowHandler(t *testing.T) {
	storage := storage.NewMemStorage()

	ts := httptest.NewServer(NewHandler(storage).GetRouter())
	defer ts.Close()

	storage.SetCounter("correctKey", 1)
	storage.SetGauge("correctKey", 1)

	testCases := []struct {
		path   string
		method string
		status int
		want   string
	}{
		{path: "/value/counter/correctKey", method: "GET", status: http.StatusOK, want: "1"},
		{path: "/value/counter/unknownKey", method: "GET", status: http.StatusNotFound, want: ""},
		{path: "/value/gauge/correctKey", method: "GET", status: http.StatusOK, want: "1"},
		{path: "/value/gauge/unknownKey", method: "GET", status: http.StatusNotFound, want: ""},
	}

	for _, tc := range testCases {
		t.Run(tc.path, func(t *testing.T) {
			resp, _ := testRequest(t, ts, tc.method, tc.path)
			defer resp.Body.Close()
			assert.Equal(t, tc.status, resp.StatusCode)
		})
	}
}

func TestShowAsJSONHandler(t *testing.T) {
	storage := storage.NewMemStorage()

	ts := httptest.NewServer(NewHandler(storage).GetRouter())
	defer ts.Close()

	storage.SetCounter("correctKey", 1)
	storage.SetGauge("correctKey", 1)

	counterMetrics, _ := json.Marshal(&models.Metrics{ID: "correctKey", MType: "counter"})
	gaugeMetrics, _ := json.Marshal(&models.Metrics{ID: "correctKey", MType: "gauge"})

	counter, _ := storage.GetCounter("correctKey")
	gauge, _ := storage.GetGauge("correctKey")

	wantCounterMetrics, _ := json.Marshal(&models.Metrics{ID: "correctKey", MType: "counter", Delta: &counter})
	wantGaugeMetrics, _ := json.Marshal(&models.Metrics{ID: "correctKey", MType: "gauge", Value: &gauge})

	testCases := []struct {
		name        string
		path        string
		method      string
		contentType string
		status      int
		body        string
		want        string
	}{
		{name: "correct counter", path: "/value", method: "POST", status: http.StatusOK, body: string(counterMetrics), want: string(wantCounterMetrics)},
		{name: "correct gauge", path: "/value", method: "POST", status: http.StatusOK, body: string(gaugeMetrics), want: string(wantGaugeMetrics)},
		{name: "incorrect values", path: "/value", method: "POST", status: http.StatusNotFound, body: `{"ID": "unknown", "MType": "unknown"}`, want: ""},
		{name: "empty body", path: "/value", method: "POST", status: http.StatusBadRequest, body: "", want: ""},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resp, body := testRequestAsJSON(t, ts, tc.method, tc.path, tc.body, tc.contentType)
			defer resp.Body.Close()
			assert.Equal(t, tc.status, resp.StatusCode)
			assert.Equal(t, tc.want, body)
		})
	}
}

func TestFindMetricsForShow(t *testing.T) {
	t.Logf("TODO")
}

func TestRepresentMetricsForShow(t *testing.T) {
	t.Logf("TODO")
}

func TestHandleErrorForShow(t *testing.T) {
	t.Logf("TODO")
}
