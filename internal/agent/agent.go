// Agent
package agent

import (
	"context"
	"os"
	"os/signal"
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
		poller: poller.NewPoller(container.Storage),
		reporter: reporter.NewReporter(
			container.Storage,
			client.NewClient(container.ServerAddr, container.ChecksumService)),
	}
}

// Starts processes poller and reporter
func (a *Agent) Start(pollInterval int, reportInterval int) error {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	go a.poller.Start(ctx, time.Duration(pollInterval)*time.Second)
	a.reporter.Start(ctx, time.Duration(reportInterval)*time.Second)

	return nil
}
