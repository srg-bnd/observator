package poller

import "log"

type Poller struct {
}

func NewPoller() *Poller {
	return &Poller{}
}

func (r *Poller) Poll() {
	log.Println("Poll!")
}
