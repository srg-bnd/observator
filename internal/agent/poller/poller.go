package poller

import (
	"log"
	"time"
)

const (
	defaultPollInterval = 2 * time.Second
)

type Poller struct {
}

func NewPoller() *Poller {
	return &Poller{}
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
