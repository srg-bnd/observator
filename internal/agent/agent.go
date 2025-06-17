// Agent
package agent

import (
	"time"

	config "github.com/srg-bnd/observator/config/agent"
	"github.com/srg-bnd/observator/internal/agent/client"
	"github.com/srg-bnd/observator/internal/agent/poller"
	"github.com/srg-bnd/observator/internal/agent/reporter"
)

// Agent
type Agent struct {
	poller   *poller.Poller
	reporter *reporter.Reporter
}

// Returns a new agent
func NewAgent(container *config.Container) *Agent {
	return &Agent{
		poller:   poller.NewPoller(container.Storage),
		reporter: reporter.NewReporter(container.Storage, client.NewClient(container.ServerAddr)),
	}
}

// Starts processes poller and reporter
func (a *Agent) Start(pollInterval int, reportInterval int) error {
	go a.poller.Start(time.Duration(pollInterval) * time.Second)
	a.reporter.Start(time.Duration(reportInterval) * time.Second)

	return nil
}
