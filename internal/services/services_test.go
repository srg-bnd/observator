package services

import (
	"testing"

	"github.com/srg-bnd/observator/internal/storage"
	"github.com/stretchr/testify/assert"
)

func TestNewService(t *testing.T) {
	storage := storage.NewMemStorage()
	deque := NewService(storage)
	assert.IsType(t, deque, &Service{})
}

func TestUpdateMetricService(t *testing.T) {
	// TODO
}
