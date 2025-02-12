package reporter

import "log"

type Reporter struct {
}

func NewReporter() *Reporter {
	return &Reporter{}
}

func (r *Reporter) Start() {
	log.Println("Reporter!")
}
