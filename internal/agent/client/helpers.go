package client

import (
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
