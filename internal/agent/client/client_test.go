package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	client := NewClient()
	assert.IsType(t, client, &Client{})
}

func TestSendMetric(t *testing.T) {
	t.Logf("TODO")
}

func TestGetBaseURL(t *testing.T) {
	assert.Equal(t, getBaseURL(), ":8080")
}

func TestConsts(t *testing.T) {
	assert.Equal(t, baseURL, ":8080")
}
