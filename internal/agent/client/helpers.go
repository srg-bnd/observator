package client

import (
	"bytes"
	"compress/flate"
	"compress/gzip"
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
)

// Returns a new http client (resty)
func newHTTPClient(baseURL string) *resty.Client {
	// Retriable errors
	repetitionIntervals := [3]int{1, 3, 5}

	return resty.New().
		SetBaseURL("http://"+baseURL).
		SetRetryCount(len(repetitionIntervals)).
		SetRetryAfter(func(c *resty.Client, r *resty.Response) (time.Duration, error) {
			var currentRepetitionInterval int

			return func() (time.Duration, error) {
				delay := time.Duration(repetitionIntervals[currentRepetitionInterval]) * time.Second
				currentRepetitionInterval++
				return delay, nil
			}()
		}).
		SetRetryMaxWaitTime(5*time.Second).
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept-Encoding", "gzip").
		SetHeader("Content-Encoding", "gzip")
}

// Compress data
func compress(data []byte) ([]byte, error) {
	var b bytes.Buffer

	w := gzip.NewWriter(&b)

	_, err := w.Write(data)
	if err != nil {
		return nil, fmt.Errorf("failed write data to compress temporary buffer: %v", err)
	}

	err = w.Close()
	if err != nil {
		return nil, fmt.Errorf("failed compress data: %v", err)
	}

	return b.Bytes(), nil
}

// Decompress data
func decompress(compressedData []byte) ([]byte, error) {
	r := flate.NewReader(bytes.NewReader(compressedData))
	defer r.Close()

	var b bytes.Buffer
	_, err := b.ReadFrom(r)
	if err != nil {
		return nil, fmt.Errorf("failed decompress data: %v", err)
	}

	return b.Bytes(), nil
}
