package scanner

import (
	"sync"

	"github.com/positive-builder2/scannerN/types"
)

type Pool struct {
	Jobs    chan types.Job
	Results chan types.Result
	Wg      sync.WaitGroup
}

func (p *Pool) worker() {
	for job := range p.Jobs {
		res := ScanPort(job)
		if res.Open {
			p.Results <- res
		}
		p.Wg.Done()
	}
}

func NewPool(workers int) *Pool {
	p := &Pool{
		Jobs:    make(chan types.Job, 1000),
		Results: make(chan types.Result, 1000),
	}

	for i := 0; i < workers; i++ {
		go p.worker()
	}
	return p
}
