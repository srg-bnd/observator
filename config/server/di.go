// Dependency Injection Container
package config

import (
	"database/sql"

	"github.com/srg-bnd/observator/internal/storage"
)

type Container struct {
	Storage storage.Repositories
	DB      *sql.DB
}
