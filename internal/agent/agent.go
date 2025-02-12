package agent

import (
	"time"

	"github.com/srg-bnd/observator/internal/agent/poller"
	"github.com/srg-bnd/observator/internal/agent/reporter"
	"github.com/srg-bnd/observator/internal/storage"
)

const (
	defaultPollInterval   = 2 * time.Second
	defaultReportInterval = 10 * time.Second
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
	pollerStarted := time.Now()
	reporterStarted := time.Now()

	for {
		if time.Since(pollerStarted) >= GetPollInterval() {
			go a.poller.Start()
			pollerStarted = time.Now()
		}

		if time.Since(reporterStarted) >= GetReportInterval() {
			go a.reporter.Start()
			reporterStarted = time.Now()
		}
	}
}

func GetPollInterval() time.Duration {
	return defaultPollInterval
}

func GetReportInterval() time.Duration {
	return defaultReportInterval
}
