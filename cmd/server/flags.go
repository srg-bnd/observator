// Flags
package main

import "flag"

var flagHostAddr string

func parseFlags() {
	flag.StringVar(&flagHostAddr, "a", ":8080", "address and port to run server")
	flag.Parse()
}
