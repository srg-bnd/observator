// Dependency Injection Container
package config

import (
	"github.com/srg-bnd/observator/internal/storage"
)

type Container struct {
	Storage    storage.Repositories
	ServerAddr string
}
