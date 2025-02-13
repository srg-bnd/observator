package client

import (
	"log"
	"net/http"
	"strconv"

	"github.com/srg-bnd/observator/internal/agent/collector"
	"github.com/srg-bnd/observator/internal/storage"
)

type Client struct {
	storage storage.Repositories
}

func NewClient(storage storage.Repositories) *Client {
	return &Client{
		storage: storage,
	}
}

func (c *Client) SendMetrics(trackedMetrics *collector.TrackedMetrics) error {
	for _, metricName := range trackedMetrics.Counter {
		metricValue := strconv.FormatInt(c.storage.GetCounter(metricName), 10)

		c.SendMetric("counter", metricName, metricValue)
	}

	for _, metricName := range trackedMetrics.Gauge {
		metricValue := strconv.FormatFloat(c.storage.GetGauge(metricName), 'f', -1, 64)

		c.SendMetric("gauge", metricName, metricValue)
	}

	return nil
}

func (c *Client) SendMetric(metricType string, metricName string, metricValue string) {
	url := "http://localhost:8080/update/" + metricType + "/" + metricName + "/" + string(metricValue)
	_, err := http.Post(url, "text/plain", nil)
	if err != nil {
		log.Println(err)
	}
}
