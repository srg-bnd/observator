package services

import (
	"testing"

	"github.com/srg-bnd/observator/internal/storage"
	"github.com/stretchr/testify/assert"
)

func TestNewService(t *testing.T) {
	storage := storage.NewMemStorage()
	service := NewService(storage)
	assert.IsType(t, service, &Service{})
}

func TestMetricsUpdateService(t *testing.T) {
	t.Logf("TODO")
}
