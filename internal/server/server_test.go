package server

import (
	"database/sql"
	"testing"

	"github.com/srg-bnd/observator/internal/storage"
	"github.com/stretchr/testify/assert"
	_ "modernc.org/sqlite"
)

func TestStart(t *testing.T) {
	t.Logf("TODO")
}

func TestGetMux(t *testing.T) {
	t.Logf("TODO")
}

func TestNewServer(t *testing.T) {
	db, err := sql.Open("sqlite", "temp.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	server := NewServer(storage.NewMemStorage("", 0, false), db)
	assert.IsType(t, server, &Server{})
}
