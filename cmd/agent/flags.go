// Flags
package main

import (
	"flag"
	"os"
	"strconv"
)

var AgentFlags = struct {
	pollInterval   int
	reportInterval int
	serverAddr     string
}{}

func parseFlags() {
	flag.IntVar(&AgentFlags.pollInterval, "p", 2, "pollInterval – frequency (seconds) of metric polling")
	flag.IntVar(&AgentFlags.reportInterval, "r", 10, "reportInterval – frequency (seconds) of sending values to the server")
	flag.StringVar(&AgentFlags.serverAddr, "a", "localhost:8080", "address and port to run server")

	flag.Parse()

	if envPollInterval := os.Getenv("POLL_INTERVAL"); envPollInterval != "" {
		AgentFlags.pollInterval, _ = strconv.Atoi(envPollInterval)
	}

	if envReportInterval := os.Getenv("REPORT_INTERVAL"); envReportInterval != "" {
		AgentFlags.reportInterval, _ = strconv.Atoi(envReportInterval)
	}

	if envServerAddr := os.Getenv("ADDRESS"); envServerAddr != "" {
		AgentFlags.serverAddr = envServerAddr
	}
}
