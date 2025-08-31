package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	client := NewClient(":8080", nil, nil, nil)
	assert.IsType(t, client, &Client{})
}

func TestSendMetrics(t *testing.T) {
	t.Logf("TODO")
}
