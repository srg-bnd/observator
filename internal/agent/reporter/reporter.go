// Sends metrics to the server
package reporter

import (
	"log"
	"time"

	"github.com/srg-bnd/observator/internal/agent/models"
	"github.com/srg-bnd/observator/internal/agent/services"
	"github.com/srg-bnd/observator/internal/storage"
)

const (
	defaultReportInterval = 10 * time.Second
)

type Reporter struct {
	storage storage.Repositories
	service *services.Service
}

func NewReporter(storage storage.Repositories) *Reporter {
	return &Reporter{
		storage: storage,
		service: services.NewService(storage),
	}
}

func (r *Reporter) Start() {
	for {
		time.Sleep(GetReportInterval())
		log.Println("=> Reporter [started]")
		err := r.service.ValueSendingService(models.TrackedMetrics)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println("=> Reporter [stopped]")
	}
}

func GetReportInterval() time.Duration {
	return defaultReportInterval
}
