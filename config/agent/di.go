// Dependency Injection Container
package config

import (
	"github.com/srg-bnd/observator/internal/shared/services"
	"github.com/srg-bnd/observator/internal/storage"
)

type Container struct {
	PublicKey          *services.PublicKey
	ChecksumService    *services.Checksum
	Storage            storage.Repositories
	ServerAddr         string
	WorkerPoolReporter int
}
