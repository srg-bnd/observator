package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/srg-bnd/observator/internal/storage"
	"github.com/stretchr/testify/assert"
)

func TestShowHandler(t *testing.T) {
	storage := storage.NewMemStorage("", 0, false)

	ts := httptest.NewServer(NewHandler(storage).GetRouter())
	defer ts.Close()

	storage.SetCounter("correctKey", 1)
	storage.SetGauge("correctKey", 1)

	testCases := []struct {
		data   DataRequest
		status int
		want   string
	}{
		{data: DataRequest{path: "/value/counter/correctKey", method: "GET"}, status: http.StatusOK, want: "1"},
		{data: DataRequest{path: "/value/counter/unknownKey", method: "GET"}, status: http.StatusNotFound, want: ""},
		{data: DataRequest{path: "/value/gauge/correctKey", method: "GET"}, status: http.StatusOK, want: "1"},
		{data: DataRequest{path: "/value/gauge/unknownKey", method: "GET"}, status: http.StatusNotFound, want: ""},
	}

	for _, tc := range testCases {
		t.Run(tc.data.path, func(t *testing.T) {
			resp, _ := testRequest(t, ts, tc.data)
			defer resp.Body.Close()
			assert.Equal(t, tc.status, resp.StatusCode)
		})
	}
}

func TestShowAsJSONHandler(t *testing.T) {
	storage := storage.NewMemStorage("", 0, false)

	ts := httptest.NewServer(NewHandler(storage).GetRouter())
	defer ts.Close()

	storage.SetCounter("correctKey", 1)
	storage.SetGauge("correctKey", 1)

	// counterMetrics, _ := json.Marshal(&models.Metrics{ID: "correctKey", MType: "counter"})
	// gaugeMetrics, _ := json.Marshal(&models.Metrics{ID: "correctKey", MType: "gauge"})

	// counter, _ := storage.GetCounter("correctKey")
	// gauge, _ := storage.GetGauge("correctKey")

	// wantCounterMetrics, _ := json.Marshal(&models.Metrics{ID: "correctKey", MType: "counter", Delta: &counter})
	// wantGaugeMetrics, _ := json.Marshal(&models.Metrics{ID: "correctKey", MType: "gauge", Value: &gauge})

	testCases := []struct {
		name   string
		status int
		data   DataRequest
		want   string
	}{
		// {name: "correct counter", data: DataRequest{path: "/value", method: "POST", body: string(counterMetrics), acceptEncoding: "gzip"}, status: http.StatusOK, want: string(wantCounterMetrics)},
		// {name: "correct gauge", data: DataRequest{path: "/value", method: "POST", body: string(gaugeMetrics)}, status: http.StatusOK, want: string(wantGaugeMetrics)},
		{name: "incorrect values", data: DataRequest{path: "/value", method: "POST", body: `{"ID": "unknown", "MType": "unknown"}`, acceptEncoding: ""}, status: http.StatusNotFound, want: ""},
		{name: "empty body", data: DataRequest{path: "/value", method: "POST", body: "", acceptEncoding: ""}, status: http.StatusBadRequest, want: ""},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resp, body := testRequestAsJSON(t, ts, tc.data)
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
