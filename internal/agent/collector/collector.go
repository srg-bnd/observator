package collector

type Collector struct {
}

func NewCollector() *Collector {
	return &Collector{}
}

func (c *Collector) GetMetrics() (map[string][]string, error) {
	metrics := make(map[string][]string, 0)

	metrics["counter"] = []string{
		"PollCount",
	}
	metrics["gauge"] = []string{
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
		"NumForcedGC",
		"NumGC",

		"OtherSys",
		"PauseTotalNs",
		"StackInuse",
		"StackSys",
		"Sys",
		"TotalAlloc",

		"RandomValue",
	}

	return metrics, nil
}
