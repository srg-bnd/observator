// Dependency Injection Container
package config

import (
	"database/sql"

	"github.com/srg-bnd/observator/internal/shared/services"
	"github.com/srg-bnd/observator/internal/storage"
)

type Container struct {
	PrivateKey      *services.PrivateKey
	ChecksumService *services.Checksum
	DB              *sql.DB
	Storage         storage.Repositories
}
