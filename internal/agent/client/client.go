// HTTP Client
package client

import (
	"encoding/json"
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

	c.client.
		SetRetryCount(0).
		SetRetryWaitTime(1 * time.Second).
		SetRetryMaxWaitTime(1 * time.Second)

	_, err = c.client.R().SetBody(data).
		SetHeader("Content-Type", "plain/text").
		Post(c.baseURL + "/update/{metricType}/{metricName}/{metricValue}")

	if err != nil {
		// log.Println(err)
		return
	}
}
