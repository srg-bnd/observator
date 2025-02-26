// Flags
package main

import (
	"flag"
	"os"
	"strconv"
)

type AppConfigs struct {
	pollInterval   int
	reportInterval int
	serverAddr     string
}

var appConfigs = AppConfigs{}

func parseFlags() {
	flag.IntVar(&appConfigs.pollInterval, "p", 2, "pollInterval – frequency (seconds) of metric polling")
	flag.IntVar(&appConfigs.reportInterval, "r", 10, "reportInterval – frequency (seconds) of sending values to the server")
	flag.StringVar(&appConfigs.serverAddr, "a", "localhost:8080", "address and port to run server")

	flag.Parse()

	if envPollInterval := os.Getenv("POLL_INTERVAL"); envPollInterval != "" {
		appConfigs.pollInterval, _ = strconv.Atoi(envPollInterval)
	}

	if envReportInterval := os.Getenv("REPORT_INTERVAL"); envReportInterval != "" {
		appConfigs.reportInterval, _ = strconv.Atoi(envReportInterval)
	}

	if envServerAddr := os.Getenv("ADDRESS"); envServerAddr != "" {
		appConfigs.serverAddr = envServerAddr
	}
}
