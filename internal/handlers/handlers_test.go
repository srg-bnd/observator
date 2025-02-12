package handlers

import (
	"testing"

	"github.com/srg-bnd/observator/internal/storage"
	"github.com/stretchr/testify/assert"
)

func TestNewHttpHandler(t *testing.T) {
	storage := storage.NewMemStorage()
	deque := NewHttpHandler(storage)
	assert.IsType(t, deque, &HttpHandler{})
}

func TestUpdateMetricHandler(t *testing.T) {
	t.Logf("TODO")
}
