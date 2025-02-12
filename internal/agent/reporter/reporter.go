package reporter

import (
	"log"
	"time"
)

const (
	defaultReportInterval = 10 * time.Second
)

type Reporter struct {
}

func NewReporter() *Reporter {
	return &Reporter{}
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
