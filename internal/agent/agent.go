package agent

import (
	"time"

	"github.com/srg-bnd/observator/internal/storage"
)

type Agent struct {
	storage storage.Repositories
}

func NewAgent(storage storage.Repositories) *Agent {
	return &Agent{
		storage: storage,
	}
}

func (a *Agent) Start() error {
	for {
		time.Sleep(2 * time.Second)
	}
}
