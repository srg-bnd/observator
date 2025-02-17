// HTTP Client
package client

import (
	"log"
	"net/http"
)

const (
	baseURL = "http://localhost:8080"
)

type Client struct {
	client *http.Client
}

func NewClient() *Client {
	return &Client{
		client: &http.Client{
			Timeout: 0,
		},
	}
}

func (c *Client) SendMetric(metricType string, metricName string, metricValue string) {
	endpoint := getBaseURL() + "/update/" + metricType + "/" + metricName + "/" + string(metricValue) + "/"
	request, err := http.NewRequest(http.MethodPost, endpoint, http.NoBody)
	if err != nil {
		panic(err)
	}

	request.Header.Set("Content-Type", "text/plain; charset=UTF-8")
	response, err := c.client.Do(request)
	if err != nil {
		log.Println(err)
	} else {
		defer response.Body.Close()
	}
}

func getBaseURL() string {
	return baseURL
}
