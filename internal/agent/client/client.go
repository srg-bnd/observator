// Client for working with the server
package client

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/srg-bnd/observator/internal/agent/models"
	"github.com/srg-bnd/observator/internal/server/logger"
	"github.com/srg-bnd/observator/internal/shared/services"
	"go.uber.org/zap"
)

var ErrBadHashSum = errors.New("bad hash")

// TODO: uses `ChecksumBehaviour` instead of `services.Checksum`
type ChecksumBehaviour interface {
	Sum(string) (string, error)
}

// Client
type Client struct {
	httpClient      *resty.Client
	checksumService *services.Checksum
}

// Returns a new client
func NewClient(baseURL string, checksumService *services.Checksum) *Client {
	return &Client{
		httpClient:      newHttpClient(baseURL),
		checksumService: checksumService,
	}
}

// Sends batch of metrics
func (c *Client) SendMetrics(metrics []models.Metrics) error {
	data, err := json.Marshal(&metrics)
	if err != nil {
		return err
	}

	compressedData, err := compress(data)
	if err != nil {
		return err
	}

	request := c.httpClient.R().SetBody(compressedData)
	request, err = c.withChecksum(request, data)
	if err != nil {
		return err
	}

	response, err := request.Post("/updates")
	if err != nil {
		return err
	}

	if strings.Contains(response.Header().Get("Accept-Encoding"), "gzip") {
		// TODO: use the results
		decompress(response.Body())
	}

	return nil
}

// Sets a checksum if need
func (c *Client) withChecksum(request *resty.Request, data []byte) (*resty.Request, error) {
	if c.checksumService != nil {
		return request, nil
	}

	hash, err := c.checksumService.Sum(string(data))
	if err != nil {
		logger.Log.Warn(ErrBadHashSum.Error(), zap.Error(ErrBadHashSum))
		return request, nil
	}

	return request.SetHeader("HashSHA256", string(hash)), nil
}
