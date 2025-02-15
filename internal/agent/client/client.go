// HTTP Client
package client

import (
	"log"
	"net/http"
)

const (
	baseURL = ":8080"
)

type Client struct {
}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) SendMetric(metricType string, metricName string, metricValue string) {
	url := getBaseURL() + "/update/" + metricType + "/" + metricName + "/" + string(metricValue)
	response, err := http.Post(url, "text/plain", nil)
	if err != nil {
		log.Println(err)
	}

	response.Body.Close()
}

func getBaseURL() string {
	return baseURL
}
