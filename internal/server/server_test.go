package server

import (
	"testing"

	"github.com/srg-bnd/observator/internal/storage"
	"github.com/stretchr/testify/assert"
)

func TestStart(t *testing.T) {
	t.Logf("TODO")
}

func TestGetMux(t *testing.T) {
	t.Logf("TODO")
}

func TestNewServer(t *testing.T) {
	server := NewServer(storage.NewMemStorage())
	assert.IsType(t, server, &Server{})
}
