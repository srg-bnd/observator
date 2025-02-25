// Sends metrics to the server
package reporter

import (
	"log"
	"time"

	"github.com/srg-bnd/observator/internal/agent/client"
	"github.com/srg-bnd/observator/internal/agent/collector"
	"github.com/srg-bnd/observator/internal/agent/services"
	"github.com/srg-bnd/observator/internal/storage"
)

type Reporter struct {
	storage storage.Repositories
	service *services.Service
}

func NewReporter(storage storage.Repositories, client *client.Client) *Reporter {
	return &Reporter{
		storage: storage,
		service: services.NewService(storage, client),
	}
}

func (r *Reporter) Start(reportInterval time.Duration) {
	for {
		time.Sleep(reportInterval)
		log.Println("=> Reporter [started]")
		err := r.service.ValueSendingService(collector.TrackedMetrics)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println("=> Reporter [stopped]")
	}
}
