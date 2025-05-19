package router

import (
	"bytes"
	"compress/gzip"
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/go-chi/chi"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/srg-bnd/observator/internal/server/models"
	"github.com/srg-bnd/observator/internal/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type DataRequest struct {
	method, path, body, contentType, acceptEncoding, contentEncoding string
}

func TestNewRouter(t *testing.T) {
	db, err := sql.Open("pgx", "temp.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	tests := []struct {
		name string
		want chi.Router
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := NewRouter(storage.NewDBStorage(db), db)
			if got := router; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRouter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShowHandler(t *testing.T) {
	db := getTempDB()
	defer db.Close()
	storage := storage.NewMemStorage()

	ts := httptest.NewServer(NewRouter(storage, db))
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

func TestShowHandlerForJSON(t *testing.T) {
	db := getTempDB()
	defer db.Close()
	storage := storage.NewMemStorage()

	ts := httptest.NewServer(NewRouter(storage, db))
	defer ts.Close()

	storage.SetCounter("correctKey", 1)
	storage.SetGauge("correctKey", 1)

	testCases := []struct {
		name   string
		status int
		data   DataRequest
		want   string
	}{
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

func TestUpdateHandler(t *testing.T) {
	db := getTempDB()
	defer db.Close()
	storage := storage.NewMemStorage()

	ts := httptest.NewServer(NewRouter(storage, db))
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
	storage := storage.NewMemStorage()

	ts := httptest.NewServer(NewRouter(storage, db))
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

func getTempDB() *sql.DB {
	db, err := sql.Open("pgx", "")
	if err != nil {
		panic(err)
	}

	return db
}
