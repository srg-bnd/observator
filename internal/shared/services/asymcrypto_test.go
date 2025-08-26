package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPublicKey(t *testing.T) {
	assert.IsType(t, &PublicKey{}, NewPublicKey(""))
}

func TestNewPrivateKey(t *testing.T) {
	assert.IsType(t, &PrivateKey{}, NewPrivateKey(""))
}
