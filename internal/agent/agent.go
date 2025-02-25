// Agent
package agent

import (
	"time"

	"github.com/srg-bnd/observator/internal/agent/client"
	"github.com/srg-bnd/observator/internal/agent/poller"
	"github.com/srg-bnd/observator/internal/agent/reporter"
	"github.com/srg-bnd/observator/internal/storage"
)

type Agent struct {
	poller   *poller.Poller
	reporter *reporter.Reporter
}

func NewAgent(storage storage.Repositories, baseURL string) *Agent {
	return &Agent{
		poller:   poller.NewPoller(storage),
		reporter: reporter.NewReporter(storage, client.NewClient(baseURL)),
	}
}

func (a *Agent) Start(pollInterval int, reportInterval int) error {
	go a.poller.Start(time.Duration(pollInterval) * time.Second)
	a.reporter.Start(time.Duration(reportInterval) * time.Second)

	return nil
}
