package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	client := NewClient(":8080")
	assert.IsType(t, client, &Client{})
}

func TestSendMetric(t *testing.T) {
	t.Logf("TODO")
}
