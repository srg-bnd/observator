package server

import (
	"database/sql"
	"testing"

	_ "github.com/jackc/pgx/v5/stdlib"
	config "github.com/srg-bnd/observator/config/server"
	"github.com/srg-bnd/observator/internal/server/router"
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
	db, err := sql.Open("pgx", "temp.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	server := NewServer(router.NewRouter(&config.Container{DB: db, Storage: storage.NewMemStorage()}))
	assert.IsType(t, server, &Server{})
}
