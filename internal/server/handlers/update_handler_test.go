package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

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
