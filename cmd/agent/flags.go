// Flags
package main

import "flag"

var AgentFlags = struct {
	pollInterval   int
	reportInterval int
	serverAddr     string
}{}

func parseFlags() {
	flag.StringVar(&AgentFlags.serverAddr, "a", "localhost:8080", "address and port to run server")
	flag.IntVar(&AgentFlags.reportInterval, "r", 10, "reportInterval – frequency (seconds) of sending values to the server")
	flag.IntVar(&AgentFlags.pollInterval, "p", 2, "pollInterval – frequency (seconds) of metric polling")
	flag.Parse()
}
