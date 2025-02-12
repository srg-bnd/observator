package poller

import (
	"log"
	"time"

	"github.com/srg-bnd/observator/internal/storage"
)

const (
	defaultPollInterval = 2 * time.Second
)

type Poller struct {
	storage storage.Repositories
}

func NewPoller(storage storage.Repositories) *Poller {
	return &Poller{
		storage: storage,
	}
}

func (r *Poller) Start() {
	pollerStarted := time.Now()

	for {
		if time.Since(pollerStarted) >= GetPollInterval() {
			log.Println("Poller!")
			pollerStarted = time.Now()
		}
	}
}

func GetPollInterval() time.Duration {
	return defaultPollInterval
}
