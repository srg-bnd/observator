package handlers

import (
	"bytes"
	"compress/gzip"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/go-chi/chi"
	"github.com/srg-bnd/observator/internal/server/services"
	"github.com/srg-bnd/observator/internal/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type DataRequest struct {
	method, path, body, contentType, acceptEncoding, contentEncoding string
}

func TestNewHandler(t *testing.T) {
	storage := storage.NewMemStorage("", 0, false)
	handler := NewHandler(storage)
	assert.IsType(t, handler, &Handler{})
}

func TestGetRouter(t *testing.T) {
	type fields struct {
		service *services.Service
		storage storage.Repositories
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
				service: tt.fields.service,
				storage: tt.fields.storage,
			}
			if got := h.GetRouter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Handler.GetRouter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHandleError(t *testing.T) {
	t.Logf("TODO")
}

func TestSetContentType(t *testing.T) {
	t.Logf("TODO")
}

// Helpers

func testRequest(t *testing.T, ts *httptest.Server, data DataRequest) (*http.Response, string) {
	req, err := http.NewRequest(data.method, ts.URL+data.path, nil)
	require.NoError(t, err)

	if data.acceptEncoding != "" {
		req.Header.Add("Accept-Encoding", data.acceptEncoding)
	} else {
		req.Header.Set("Accept-Encoding", "")
	}

	if data.contentEncoding != "" {
		req.Header.Add("Content-Encoding", data.contentEncoding)
	} else {
		req.Header.Set("Content-Encoding", "")
	}

	resp, err := ts.Client().Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	return resp, string(respBody)
}

func testRequestAsJSON(t *testing.T, ts *httptest.Server, data DataRequest) (*http.Response, string) {
	var body io.Reader

	if data.contentEncoding == "gzip" {
		// Compress if need
		buf := bytes.NewBuffer(nil)
		zb := gzip.NewWriter(buf)
		_, err := zb.Write([]byte(data.body))
		require.NoError(t, err)
		err = zb.Close()
		require.NoError(t, err)
		body = buf
	} else {
		body = strings.NewReader(data.body)
	}

	req, err := http.NewRequest(data.method, ts.URL+data.path, body)
	require.NoError(t, err)

	if data.contentType != "" {
		req.Header.Add("Content-Type", data.contentType)
	} else {
		req.Header.Add("Content-Type", "application/json")
	}

	if data.acceptEncoding != "" {
		req.Header.Add("Accept-Encoding", data.acceptEncoding)
	} else {
		req.Header.Set("Accept-Encoding", "")
	}

	if data.contentEncoding != "" {
		req.Header.Add("Content-Encoding", data.contentEncoding)
	} else {
		req.Header.Set("Content-Encoding", "")
	}

	resp, err := ts.Client().Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	var bodyAsString []byte
	if resp.Header.Get("Content-Encoding") == "gzip" {
		// Decompress if need
		zb, _ := gzip.NewReader(bytes.NewReader(respBody))
		var b bytes.Buffer
		_, err := b.ReadFrom(zb)
		if err != nil {
			require.NoError(t, err)
		}
		bodyAsString = b.Bytes()
	} else {
		bodyAsString = respBody
	}

	return resp, string(bodyAsString)
}
