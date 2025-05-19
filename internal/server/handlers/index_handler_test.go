package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/srg-bnd/observator/internal/storage"
	"github.com/stretchr/testify/assert"
)

func TestIndexHandler(t *testing.T) {
	db := getTempDB()
	defer db.Close()

	storage := storage.NewMemStorage()

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

			NewIndexHandler(storage).Handler(w, r)

			assert.Equal(t, tc.expectedCode, w.Code)
		})
	}
}
