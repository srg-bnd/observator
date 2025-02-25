package server

import (
	"reflect"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/srg-bnd/observator/internal/server/handlers"
	"github.com/srg-bnd/observator/internal/storage"
	"github.com/stretchr/testify/assert"
)

func TestStart(t *testing.T) {
	t.Logf("TODO")
}

func TestGetRouter(t *testing.T) {
	t.Logf("TODO")
}

func TestNewServer(t *testing.T) {
	server := NewServer(storage.NewMemStorage())
	assert.IsType(t, server, &Server{})
}

func TestGetHost(t *testing.T) {
	got, _ := getHost()
	want := ":8080"

	if got != want {
		t.Errorf(`Incorrect 'getHost()' method behavior, got "%v", want "%v"`, got, want)
	}
}

func TestServer_GetRouter(t *testing.T) {
	type fields struct {
		handler *handlers.Handler
	}
	tests := []struct {
		name   string
		fields fields
		want   chi.Router
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := &Server{
				handler: tt.fields.handler,
			}
			if got := server.GetRouter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.GetRouter() = %v, want %v", got, tt.want)
			}
		})
	}
}
