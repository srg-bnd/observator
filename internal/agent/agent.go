package agent

import (
	"github.com/srg-bnd/observator/internal/agent/poller"
	"github.com/srg-bnd/observator/internal/agent/reporter"
	"github.com/srg-bnd/observator/internal/storage"
)

type Agent struct {
	storage  storage.Repositories
	poller   *poller.Poller
	reporter *reporter.Reporter
}

func NewAgent(storage storage.Repositories) *Agent {
	return &Agent{
		storage:  storage,
		poller:   poller.NewPoller(),
		reporter: reporter.NewReporter(),
	}
}

func (a *Agent) Start() error {
	go a.poller.Start()
	a.reporter.Start()

	return nil
}
