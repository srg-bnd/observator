package client

import (
	"testing"

	"github.com/srg-bnd/observator/internal/storage"
	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	client := NewClient(storage.NewMemStorage())
	assert.IsType(t, client, &Client{})
}

func TestSendMetrics(t *testing.T) {
	t.Logf("TODO")
}

func TestSendMetric(t *testing.T) {
	t.Logf("TODO")
}

func TestGetBaseURL(t *testing.T) {
	assert.Equal(t, getBaseURL(), "http://localhost:8080")
}

func TestConsts(t *testing.T) {
	assert.Equal(t, baseURL, "http://localhost:8080")
}
