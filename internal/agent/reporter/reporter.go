// Sends metrics to the server
package reporter

import (
	"log"
	"time"

	"github.com/srg-bnd/observator/internal/agent/client"
	"github.com/srg-bnd/observator/internal/agent/collector"
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
		client:  client.NewClient(storage),
	}
}

func (r *Reporter) Start() {
	for {
		time.Sleep(GetReportInterval())
		log.Println("=> Reporter [started]")
		err := r.client.SendMetrics(trackedMetrics())
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

func trackedMetrics() *collector.TrackedMetrics {
	return collector.NewTrackedMetrics(
		[]string{
			"PollCount",
		},
		[]string{
			"Alloc",
			"BuckHashSys",
			"Frees",
			"GCCPUFraction",
			"GCSys",
			"HeapAlloc",
			"HeapIdle",
			"HeapInuse",
			"HeapObjects",
			"HeapReleased",
			"HeapSys",
			"LastGC",
			"Lookups",
			"MCacheInuse",
			"MCacheSys",
			"MSpanInuse",
			"MSpanSys",
			"Mallocs",
			"NextGC",
			"NumForcedGC",
			"NumGC",
			"OtherSys",
			"PauseTotalNs",
			"StackInuse",
			"StackSys",
			"Sys",
			"TotalAlloc",
			"RandomValue",
		},
	)
}
