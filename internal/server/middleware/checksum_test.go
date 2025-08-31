package middleware

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewChecksum(t *testing.T) {
	assert.IsType(t, &Checksum{}, NewChecksum(nil))
}
