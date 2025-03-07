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

func TestUpdateHandler(t *testing.T) {
	storage := storage.NewMemStorage()

	ts := httptest.NewServer(NewHandler(storage).GetRouter())
	defer ts.Close()

	testCases := []struct {
		path   string
		method string
		status int
		want   string
	}{
		{path: "/update/counter/correctKey/1", method: "POST", status: http.StatusOK},
		{path: "/update/gauge/correctKey/1", method: "POST", status: http.StatusOK},
	}

	for _, tc := range testCases {
		t.Run(tc.path, func(t *testing.T) {
			resp, _ := testRequest(t, ts, tc.method, tc.path)
			defer resp.Body.Close()
			assert.Equal(t, tc.status, resp.StatusCode)
		})
	}
}

func TestUpdateAsJSONHandler(t *testing.T) {
	storage := storage.NewMemStorage()

	ts := httptest.NewServer(NewHandler(storage).GetRouter())
	defer ts.Close()

	counter := int64(1)
	gauge := float64(1.0)
	existCounter := int64(2)

	counterMetrics, _ := json.Marshal(&models.Metrics{ID: "correctKey", MType: "counter", Delta: &counter})
	gaugeMetrics, _ := json.Marshal(&models.Metrics{ID: "correctKey", MType: "gauge", Value: &gauge})
	existCounterMetrics, _ := json.Marshal(&models.Metrics{ID: "existKey", MType: "counter", Delta: &counter})
	storage.SetCounter("existKey", counter)
	wantExistCounterMetrics, _ := json.Marshal(&models.Metrics{ID: "existKey", MType: "counter", Delta: &existCounter})

	testCases := []struct {
		name   string
		path   string
		method string
		status int
		body   string
		want   string
	}{
		{name: "correct counter", path: "/update", method: "POST", status: http.StatusOK, body: string(counterMetrics), want: string(counterMetrics)},
		{name: "exist counter", path: "/update", method: "POST", status: http.StatusOK, body: string(existCounterMetrics), want: string(wantExistCounterMetrics)},
		{name: "correct gauge", path: "/update", method: "POST", status: http.StatusOK, body: string(gaugeMetrics), want: string(gaugeMetrics)},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resp, body := testRequestAsJSON(t, ts, tc.method, tc.path, tc.body)
			defer resp.Body.Close()
			assert.Equal(t, tc.status, resp.StatusCode)
			assert.Equal(t, tc.want, body)
		})
	}
}

func TestParseAndValidateMetricsForUpdate(t *testing.T) {
	t.Logf("TODO")
}

func TestProcessForUpdate(t *testing.T) {
	t.Logf("TODO")
}

func TestRepresentForUpdate(t *testing.T) {
	t.Logf("TODO")
}

func TestHandleErrorForUpdate(t *testing.T) {
	t.Logf("TODO")
}
