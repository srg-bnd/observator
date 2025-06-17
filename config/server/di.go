// Dependency Injection Container
package config

import (
	"database/sql"

	"github.com/srg-bnd/observator/internal/storage"
)

type Container struct {
	EncryptionKey string
	DB            *sql.DB
	Storage       storage.Repositories
}
