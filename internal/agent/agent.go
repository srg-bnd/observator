package agent

import (
	"time"

	"github.com/srg-bnd/observator/internal/agent/poller"
	"github.com/srg-bnd/observator/internal/agent/reporter"
	"github.com/srg-bnd/observator/internal/storage"
)

const (
	defaultPollInterval   = 2
	defaultReportInterval = 10
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
	startPoll := time.Now()
	startReport := time.Now()

	for {
		if time.Since(startPoll) >= defaultPollInterval*time.Second {
			go a.poller.Poll()
			startPoll = time.Now()
		}

		if time.Since(startReport) >= defaultReportInterval*time.Second {
			go a.reporter.Report()
			startReport = time.Now()
		}
	}
}
