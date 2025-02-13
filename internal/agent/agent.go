package agent

import (
	"github.com/srg-bnd/observator/internal/agent/poller"
	"github.com/srg-bnd/observator/internal/agent/reporter"
	"github.com/srg-bnd/observator/internal/storage"
)

type Agent struct {
	poller   *poller.Poller
	reporter *reporter.Reporter
}

func NewAgent(storage storage.Repositories) *Agent {
	return &Agent{
		poller:   poller.NewPoller(storage),
		reporter: reporter.NewReporter(storage),
	}
}

func (a *Agent) Start() error {
	go a.poller.Start()
	a.reporter.Start()

	return nil
}
