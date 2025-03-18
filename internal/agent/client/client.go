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

type Client struct {
	client  *resty.Client
	baseURL string
}

func NewClient(baseURL string) *Client {
	return &Client{
		client: resty.New(),
		// HACK
		baseURL: "http://" + baseURL,
	}
}

func (c *Client) SendMetric(metrics *models.Metrics) {
	data, err := json.Marshal(&metrics)
	if err != nil {
		// log.Println(err)
		return
	}

	compressedData, err := compress(data)
	if err != nil {
		// log.Println(err)
		return
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
		// log.Println(err)
		return
	}

	if strings.Contains(res.Header().Get("Accept-Encoding"), "gzip") {
		decompress(res.Body())
	}
}

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
