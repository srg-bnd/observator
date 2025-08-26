package middleware

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDecryptor(t *testing.T) {
	assert.IsType(t, &Decryptor{}, NewDecryptor(nil))
}
