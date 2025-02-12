package reporter

import "log"

type Reporter struct {
}

func NewReporter() *Reporter {
	return &Reporter{}
}

func (r *Reporter) Report() {
	log.Println("Report!")
}
