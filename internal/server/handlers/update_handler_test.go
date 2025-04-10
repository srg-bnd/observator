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
	db := getTempDB()
	defer db.Close()
	storage := storage.NewMemStorage("", 0, false)

	ts := httptest.NewServer(NewHandler(storage, db).GetRouter())
	defer ts.Close()

	testCases := []struct {
		data   DataRequest
		status int
		want   string
	}{
		{data: DataRequest{path: "/update/counter/correctKey/1", method: "POST"}, status: http.StatusOK},
		{data: DataRequest{path: "/update/gauge/correctKey/1", method: "POST"}, status: http.StatusOK},
	}

	for _, tc := range testCases {
		t.Run(tc.data.path, func(t *testing.T) {
			resp, _ := testRequest(t, ts, tc.data)
			defer resp.Body.Close()
			assert.Equal(t, tc.status, resp.StatusCode)
		})
	}
}

func TestUpdateAsJSONHandler(t *testing.T) {
	db := getTempDB()
	defer db.Close()
	storage := storage.NewMemStorage("", 0, false)

	ts := httptest.NewServer(NewHandler(storage, db).GetRouter())
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
		data   DataRequest
		status int
		want   string
	}{
		{name: "correct counter", data: DataRequest{path: "/update", method: "POST", body: string(counterMetrics)}, status: http.StatusOK, want: string(counterMetrics)},
		{name: "exist counter", data: DataRequest{path: "/update", method: "POST", body: string(existCounterMetrics)}, status: http.StatusOK, want: string(wantExistCounterMetrics)},
		{name: "correct gauge", data: DataRequest{path: "/update", method: "POST", body: string(gaugeMetrics)}, status: http.StatusOK, want: string(gaugeMetrics)},
		{name: "correct with plain_text", data: DataRequest{path: "/update", method: "POST", body: string(gaugeMetrics), contentType: "plain/text"}, status: http.StatusOK, want: string(gaugeMetrics)},
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
