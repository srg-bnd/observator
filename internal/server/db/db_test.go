package db

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDB(t *testing.T) {
	testDSN := "postgresql://user:password@localhost:5432/dbname"
	db := NewDB(testDSN)

	assert.IsType(t, &sql.DB{}, db)
}
