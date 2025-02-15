package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/srg-bnd/observator/internal/storage"
	"github.com/stretchr/testify/assert"
)

func TestUpdateMetricHandler(t *testing.T) {
	storage := storage.NewMemStorage()
	handler := NewHandler(storage)

	testCases := []struct {
		method       string
		expectedCode int
		path         string
	}{
		{method: http.MethodPost, expectedCode: http.StatusOK, path: "/update/counter/key/1"},
		{method: http.MethodPost, expectedCode: http.StatusOK, path: "/update/gauge/key/1"},
	}

	for _, tc := range testCases {
		t.Run(tc.path, func(t *testing.T) {
			r := httptest.NewRequest(tc.method, tc.path, nil)
			w := httptest.NewRecorder()

			handler.UpdateMetricHandler(w, r)

			assert.Equal(t, tc.expectedCode, w.Code)
		})
	}
}

func TestParseAndValidateMetric(t *testing.T) {
	t.Logf("TODO")
}

func TestProcessMetric(t *testing.T) {
	t.Logf("TODO")
}

func TestHandleError(t *testing.T) {
	t.Logf("TODO")
}
