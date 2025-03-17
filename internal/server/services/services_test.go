package services

import (
	"testing"

	"github.com/srg-bnd/observator/internal/storage"
	"github.com/stretchr/testify/assert"
)

func TestNewService(t *testing.T) {
	storage := storage.NewMemStorage("", 0, false)
	service := NewService(storage)
	assert.IsType(t, service, &Service{})
}
