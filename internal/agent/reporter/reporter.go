// Sends metrics to the server
package reporter

import (
	"log"
	"time"

	"github.com/srg-bnd/observator/internal/agent/client"
	"github.com/srg-bnd/observator/internal/storage"
)

const (
	defaultReportInterval = 10 * time.Second
)

type Reporter struct {
	storage storage.Repositories
	client  *client.Client
}

func NewReporter(storage storage.Repositories) *Reporter {
	return &Reporter{
		storage: storage,
		client:  client.NewClient(),
	}
}

func (r *Reporter) Start() {
	reporterStarted := time.Now()

	for {
		if time.Since(reporterStarted) >= GetReportInterval() {
			log.Println("Reporter!")
			reporterStarted = time.Now()
		}
	}
}

func GetReportInterval() time.Duration {
	return defaultReportInterval
}
