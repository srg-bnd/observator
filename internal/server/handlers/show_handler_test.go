package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/srg-bnd/observator/internal/server/models"
	"github.com/srg-bnd/observator/internal/storage"
	"github.com/stretchr/testify/assert"
)

func TestShowHandler(t *testing.T) {
	db := getTempDB()
	defer db.Close()

	storage := storage.NewMemStorage()
	storage.SetCounter(context.Background(), "counter", 1)
	handler := NewShowHandler(storage)

	testCases := []struct {
		method       string
		expectedCode int
		metric       models.Metrics
	}{
		{method: http.MethodGet, expectedCode: http.StatusOK, metric: models.Metrics{ID: "counter", MType: "counter"}},
	}

	for _, tc := range testCases {
		t.Run("GET /value (JSON)", func(t *testing.T) {
			jsonData, err := json.Marshal(tc.metric)
			if err != nil {
				t.Fatal(err)
			}

			r := httptest.NewRequest(tc.method, "/value", bytes.NewBuffer(jsonData))
			w := httptest.NewRecorder()

			handler.JSONHandler(w, r)
			assert.Equal(t, tc.expectedCode, w.Code)
		})

		t.Run("GET /value/{metricType}/{metricName}", func(t *testing.T) {
			params := url.Values{}
			params.Add("metricType", tc.metric.MType)
			params.Add("metricName", tc.metric.ID)

			path := "/value?" + params.Encode()

			r := httptest.NewRequest(tc.method, path, nil)
			w := httptest.NewRecorder()

			handler.Handler(w, r)
			t.Logf("TODO: check chi.URLParam")
			// assert.Equal(t, tc.expectedCode, w.Code)
		})
	}
}
