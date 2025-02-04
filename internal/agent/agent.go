package agent

import "time"

func Start() error {
	for {
		time.Sleep(2 * time.Second)
	}
}
