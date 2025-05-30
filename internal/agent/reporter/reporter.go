// Sends metrics to the server
package reporter

import (
	"time"

	"github.com/srg-bnd/observator/internal/agent/client"
	"github.com/srg-bnd/observator/internal/agent/collector"
	"github.com/srg-bnd/observator/internal/agent/services"
	"github.com/srg-bnd/observator/internal/storage"
)

// Reporter
type Reporter struct {
	storage storage.Repositories
	service *services.Service
}

// Returns a new reporter
func NewReporter(storage storage.Repositories, client *client.Client) *Reporter {
	return &Reporter{
		storage: storage,
		service: services.NewService(storage, client),
	}
}

// Starts the reporter
func (r *Reporter) Start(reportInterval time.Duration) error {
	ticker := time.NewTicker(reportInterval)
	defer ticker.Stop()

	for {
		<-ticker.C

		err := r.service.ValueSendingService(collector.TrackedMetrics)
		if err != nil {
			return err
		}
	}
}
