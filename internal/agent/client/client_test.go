package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	client := NewClient()
	assert.IsType(t, client, &Client{})
}

func TestSendMetrics(t *testing.T) {
	t.Logf("TODO")
}

func TestSendMetric(t *testing.T) {
	t.Logf("TODO")
}
