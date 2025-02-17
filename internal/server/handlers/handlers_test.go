package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/srg-bnd/observator/internal/storage"
	"github.com/stretchr/testify/assert"
)

func TestNewHandler(t *testing.T) {
	storage := storage.NewMemStorage()
	handler := NewHandler(storage)
	assert.IsType(t, handler, &Handler{})
}

func TestShowMetricHandler(t *testing.T) {
	storage := storage.NewMemStorage()
	handler := NewHandler(storage)

	storage.SetCounter("existKey", 1)
	storage.SetGauge("existKey", 1)

	testCases := []struct {
		method       string
		expectedCode int
		path         string
	}{
		{method: http.MethodGet, expectedCode: http.StatusOK, path: "/value/counter/existKey/"},
		{method: http.MethodGet, expectedCode: http.StatusNotFound, path: "/value/counter/unknownKey"},
		{method: http.MethodGet, expectedCode: http.StatusOK, path: "/value/gauge/existKey"},
		{method: http.MethodGet, expectedCode: http.StatusNotFound, path: "/value/gauge/unknownKey"},
	}

	for _, tc := range testCases {
		t.Run(tc.path, func(t *testing.T) {
			r := httptest.NewRequest(tc.method, tc.path, nil)
			w := httptest.NewRecorder()

			handler.ShowMetricHandler(w, r)

			assert.Equal(t, tc.expectedCode, w.Code)
		})
	}
}

func TestIndexHandler(t *testing.T) {
	storage := storage.NewMemStorage()
	handler := NewHandler(storage)
	handler.indexFilePath = "../../../" + handler.indexFilePath

	testCases := []struct {
		method       string
		expectedCode int
	}{
		{method: http.MethodGet, expectedCode: http.StatusOK},
	}

	for _, tc := range testCases {
		t.Run(tc.method, func(t *testing.T) {
			r := httptest.NewRequest(tc.method, "/", nil)
			w := httptest.NewRecorder()

			handler.IndexHandler(w, r)

			assert.Equal(t, tc.expectedCode, w.Code)
		})
	}
}
