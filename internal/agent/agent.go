package agent

import (
	"time"

	"github.com/srg-bnd/observator/internal/storage"
)

var (
	MemStorage *storage.MemStorage
)

func Start() error {
	for {
		time.Sleep(2 * time.Second)
	}
}
