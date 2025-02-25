// HTTP Client
package client

import (
	"log"
	"time"

	"github.com/go-resty/resty/v2"
)

const (
	baseURL = "http://localhost:8080"
)

type Client struct {
	client *resty.Client
}

func NewClient() *Client {
	return &Client{
		client: resty.New(),
	}
}

func (c *Client) SendMetric(metricType string, metricName string, metricValue string) {
	c.client.
		SetRetryCount(0).
		SetRetryWaitTime(1 * time.Second).
		SetRetryMaxWaitTime(1 * time.Second)

	_, err := c.client.R().SetPathParams(map[string]string{
		"metricType":  metricType,
		"metricName":  metricName,
		"metricValue": metricValue,
	}).
		SetHeader("Content-Type", "plain/text").
		Post(getBaseURL() + "/update/{metricType}/{metricName}/{metricValue}")

	if err != nil {
		log.Println(err)
		return
	}
}

func getBaseURL() string {
	return baseURL
}
