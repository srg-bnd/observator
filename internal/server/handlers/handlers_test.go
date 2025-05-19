package handlers

import (
	"database/sql"
	"testing"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func TestHandleError(t *testing.T) {
	t.Logf("TODO")
}

func TestSetContentType(t *testing.T) {
	t.Logf("TODO")
}

// Helpers

func getTempDB() *sql.DB {
	db, err := sql.Open("pgx", "")
	if err != nil {
		panic(err)
	}

	return db
}
