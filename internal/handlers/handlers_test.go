package handlers

import (
	"testing"

	"github.com/srg-bnd/observator/internal/storage"
	"github.com/stretchr/testify/assert"
)

func TestNewHandler(t *testing.T) {
	storage := storage.NewMemStorage()
	handler := NewHandler(storage)
	assert.IsType(t, handler, &Handler{})
}

func TestUpdateMetricHandler(t *testing.T) {
	t.Logf("TODO")
}

func TestParseAndValidateMetric(t *testing.T) {
	t.Logf("TODO")
}

func TestProcessMetric(t *testing.T) {
	t.Logf("TODO")
}

func TestHandleError(t *testing.T) {
	t.Logf("TODO")
}
