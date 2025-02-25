package handlers

import (
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/go-chi/chi"
	"github.com/srg-bnd/observator/internal/server/services"
	"github.com/srg-bnd/observator/internal/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewHandler(t *testing.T) {
	storage := storage.NewMemStorage()
	handler := NewHandler(storage)
	assert.IsType(t, handler, &Handler{})
}

func TestGetRouter(t *testing.T) {
	type fields struct {
		service      *services.Service
		storage      storage.Repositories
		rootFilePath string
	}
	tests := []struct {
		name   string
		fields fields
		want   chi.Router
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				service:      tt.fields.service,
				storage:      tt.fields.storage,
				rootFilePath: tt.fields.rootFilePath,
			}
			if got := h.GetRouter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Handler.GetRouter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateMetricHandler(t *testing.T) {
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
			assert.Equal(t, tc.status, resp.StatusCode)
		})
	}
}

func TestShowMetricHandler(t *testing.T) {
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
			assert.Equal(t, tc.status, resp.StatusCode)
		})
	}
}

func TestIndexHandler(t *testing.T) {
	storage := storage.NewMemStorage()
	handler := NewHandler(storage)
	handler.rootFilePath = "../../../" + handler.rootFilePath

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

// Helpers

func testRequest(t *testing.T, ts *httptest.Server, method, path string) (*http.Response, string) {
	req, err := http.NewRequest(method, ts.URL+path, nil)
	require.NoError(t, err)

	resp, err := ts.Client().Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	return resp, string(respBody)
}
