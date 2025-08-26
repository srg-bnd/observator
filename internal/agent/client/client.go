// Client for working with the server
package client

import (
	"context"
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

type CompressorBehaviour interface {
	Compress([]byte) ([]byte, error)
	Decompress([]byte) ([]byte, error)
}

// Client
type Client struct {
	httpClient      *resty.Client
	checksumService *services.Checksum
	compressor      CompressorBehaviour
	publicKey       *services.PublicKey
}

// Returns a new client
func NewClient(baseURL string, checksumService *services.Checksum, compress CompressorBehaviour, publicKey *services.PublicKey) *Client {
	return &Client{
		httpClient:      newHTTPClient(baseURL),
		checksumService: checksumService,
		compressor:      compress,
		publicKey:       publicKey,
	}
}

// Sends batch of metrics
func (c *Client) SendMetrics(context context.Context, metrics []models.Metrics) error {
	data, err := json.Marshal(&metrics)
	if err != nil {
		return err
	}

	if c.publicKey != nil {
		data, err = c.publicKey.Encrypt(data)
		if err != nil {
			return err
		}
	}

	compressedData, err := c.compressor.Compress(data)
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
		c.compressor.Decompress(response.Body())
	}

	return nil
}

// Sets a checksum if need
func (c *Client) withChecksum(request *resty.Request, data []byte) (*resty.Request, error) {
	if c.checksumService == nil {
		return request, nil
	}

	hash, err := c.checksumService.Sum(string(data))
	if err != nil {
		logger.Log.Warn(ErrBadHashSum.Error(), zap.Error(ErrBadHashSum))
		return request, nil
	}

	return request.SetHeader("HashSHA256", string(hash)), nil
}
