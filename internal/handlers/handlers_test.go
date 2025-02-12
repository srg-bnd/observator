package handlers

import (
	"testing"

	"github.com/srg-bnd/observator/internal/storage"
	"github.com/stretchr/testify/assert"
)

func TestNewHTTPHandler(t *testing.T) {
	storage := storage.NewMemStorage()
	deque := NewHTTPHandler(storage)
	assert.IsType(t, deque, &HTTPHandler{})
}

func TestUpdateMetricHandler(t *testing.T) {
	t.Logf("TODO")
}
