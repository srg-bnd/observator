package services

import (
	"testing"

	"github.com/srg-bnd/observator/internal/storage"
	"github.com/stretchr/testify/assert"
)

func TestNewService(t *testing.T) {
	storage := storage.NewMemStorage("", 0, false)
	service := NewService(storage, nil)
	assert.IsType(t, service, &Service{})
}

func TestMetricsUpdateService(t *testing.T) {
	t.Logf("TODO")
}

func TestValueSendingService(t *testing.T) {
	t.Logf("TODO")
}
