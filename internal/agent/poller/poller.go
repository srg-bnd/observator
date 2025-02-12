// Collects metrics
package poller

import (
	"fmt"
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
	for {
		time.Sleep(GetPollInterval())
		ShowMetrics()
	}
}

func ShowMetrics() {
	log.Println("=== Poller started ===")
	fmt.Println("Gauge:")
	fmt.Println("- Alloc:", 1)
	fmt.Println("- BuckHashSys:", 1)
	fmt.Println("- Frees:", 1)
	fmt.Println("- GCCPUFraction:", 1)
	fmt.Println("- GCSys:", 1)

	fmt.Println("- HeapAlloc:", 1)
	fmt.Println("- HeapIdle:", 1)
	fmt.Println("- HeapInuse:", 1)
	fmt.Println("- HeapObjects:", 1)
	fmt.Println("- HeapReleased:", 1)
	fmt.Println("- HeapSys:", 1)

	fmt.Println("- LastGC:", 1)
	fmt.Println("- Lookups:", 1)
	fmt.Println("- MCacheInuse:", 1)
	fmt.Println("- MCacheSys:", 1)
	fmt.Println("- MSpanInuse:", 1)
	fmt.Println("- MSpanSys:", 1)

	fmt.Println("- Mallocs:", 1)
	fmt.Println("- NextGC:", 1)
	fmt.Println("- NumForcedGC:", 1)
	fmt.Println("- NumGC:", 1)
	fmt.Println("- NumForcedGC:", 1)
	fmt.Println("- NumGC:", 1)

	fmt.Println("- OtherSys:", 1)
	fmt.Println("- PauseTotalNs:", 1)
	fmt.Println("- StackInuse:", 1)
	fmt.Println("- StackSys:", 1)
	fmt.Println("- Sys:", 1)
	fmt.Println("- TotalAlloc:", 1)

	fmt.Println("- RandomValue:", 1) // custom

	fmt.Println("Counter:")
	fmt.Println("- PollCount:", 1) // custom

	log.Println("=== Poller stopped ===")
}

func GetPollInterval() time.Duration {
	return defaultPollInterval
}
