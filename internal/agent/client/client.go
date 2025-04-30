// HTTP Client
package client

import (
	"bytes"
	"compress/flate"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/srg-bnd/observator/internal/agent/models"
)

// Client
type Client struct {
	client  *resty.Client
	baseURL string
}

// Returns a new client
func NewClient(baseURL string) *Client {
	return &Client{
		client: resty.New(),
		// HACK
		baseURL: "http://" + baseURL,
	}
}

// Sends batch of metrics to the server
func (c *Client) SendMetrics(metrics []models.Metrics) error {
	data, err := json.Marshal(&metrics)
	if err != nil {
		return err
	}

	// Compress data
	compressedData, err := compress(data)
	if err != nil {
		return err
	}

	// Init client
	c.client.
		SetRetryCount(0).
		SetRetryWaitTime(1 * time.Second).
		SetRetryMaxWaitTime(1 * time.Second)

	// Execute a request
	res, err := c.client.R().SetBody(compressedData).
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept-Encoding", "gzip").
		SetHeader("Content-Encoding", "gzip").
		Post(c.baseURL + "/updates")

	if err != nil {
		return err
	}

	// Decompress response
	if strings.Contains(res.Header().Get("Accept-Encoding"), "gzip") {
		decompress(res.Body())
	}

	return nil
}

// [deprecated] Sends a metric to the server
func (c *Client) SendMetric(metrics *models.Metrics) error {
	data, err := json.Marshal(&metrics)
	if err != nil {
		return err
	}

	compressedData, err := compress(data)
	if err != nil {
		return err
	}

	c.client.
		SetRetryCount(0).
		SetRetryWaitTime(1 * time.Second).
		SetRetryMaxWaitTime(1 * time.Second)

	res, err := c.client.R().SetBody(compressedData).
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept-Encoding", "gzip").
		SetHeader("Content-Encoding", "gzip").
		Post(c.baseURL + "/update")

	if err != nil {
		return err
	}

	if strings.Contains(res.Header().Get("Accept-Encoding"), "gzip") {
		decompress(res.Body())
	}

	return nil
}

// Helpers

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
