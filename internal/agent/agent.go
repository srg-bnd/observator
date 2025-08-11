// Package agent provides implementation of monitoring agent responsible for data collection and reporting.
//
// The package contains core components for monitoring system:
// - Agent structure representing main monitoring entity
// - Methods for starting and managing monitoring processes
// - Dependencies for data collection and reporting
//
// This package serves as a central component for monitoring operations,
// coordinating data collection and reporting activities.
//
// Usage example:
//
//	container := config.NewContainer()
//	agent := agent.NewAgent(container)
//	agent.Start(60, 300)
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
	"github.com/srg-bnd/observator/internal/shared/compressor"
)

// Agent represents the main monitoring agent structure.
//
// The Agent struct is responsible for managing data collection and reporting processes.
// It consists of two main components:
// - Poller for collecting monitoring data
// - Reporter for processing and sending collected data
//
// This structure serves as the core component for monitoring operations,
// coordinating data collection and reporting activities.
type Agent struct {
	// Poller is responsible for collecting monitoring data from various sources
	poller *poller.Poller
	// Reporter is responsible for processing and sending collected data
	reporter *reporter.Reporter
}

// NewAgent creates a new Agent instance configured with provided settings.
//
// The function initializes all necessary components for the agent to operate:
// - Poller for data collection
// - Reporter for data processing and reporting
// - Client for communication with the server
//
// Parameters:
// - container: configuration container containing all necessary settings and dependencies
//
// The configuration container should provide:
// * Storage for data persistence
// * WorkerPoolReporter for managing reporting workers
// * Server address for client communication
// * ChecksumService for data integrity verification
// * Other required dependencies
//
// Returns a fully initialized Agent instance ready for operation.
func NewAgent(container *config.Container) *Agent {
	return &Agent{
		poller: poller.NewPoller(container.Storage),
		reporter: reporter.NewReporter(
			container.Storage,
			container.WorkerPoolReporter,
			client.NewClient(
				container.ServerAddr,
				container.ChecksumService,
				compressor.NewCompressor(),
			)),
	}
}

// Start launches the agent with specified polling and reporting intervals.
//
// The agent starts working in two modes:
// - Data collection with a specified polling interval
// - Sending collected data with a specified reporting interval
//
// The function creates a context with interrupt signal handling and launches:
// 1. Data collector (poller) in a separate goroutine
// 2. Reporting component (reporter) in the current thread
//
// Parameters:
// - pollInterval: polling interval in seconds
// - reportInterval: reporting interval in seconds
//
// Return values:
// - error: nil in case of successful launch, error in case of failure
func (a *Agent) Start(pollInterval int, reportInterval int) error {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	go a.poller.Start(ctx, time.Duration(pollInterval)*time.Second)
	a.reporter.Start(ctx, time.Duration(reportInterval)*time.Second)

	return nil
}
