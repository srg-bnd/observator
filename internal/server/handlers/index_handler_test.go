package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/srg-bnd/observator/internal/storage"
	"github.com/stretchr/testify/assert"
)

func TestIndexHandler(t *testing.T) {
	storage := storage.NewMemStorage()
	handler := NewHandler(storage)

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
