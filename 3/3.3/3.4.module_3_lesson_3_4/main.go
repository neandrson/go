package pool

import (
	"errors"
	"sync"
)

type PoolTask interface {
	Execute() error
	OnFailure(error)
}

type MyPool struct {
	numberWorkers int
	tasks         chan PoolTask
	wg            sync.WaitGroup
	isActive      bool
	mu            sync.RWMutex
}

func (p *MyPool) Start() {
	p.mu.RLock()
	if p.isActive {
		defer p.mu.RUnlock()
		return
	}
	p.mu.RUnlock()

	p.wg.Add(p.numberWorkers)
	for i := 0; i < p.numberWorkers; i++ {
		go func() {
			defer p.wg.Done()
			for w := range p.tasks {
				err := w.Execute()
				if err != nil {
					w.OnFailure(err)
				}
			}
		}()
	}

	p.mu.Lock()
	p.isActive = true
	p.mu.Unlock()
}

func (p *MyPool) Stop() {
	p.mu.Lock()
	if !p.isActive {
		defer p.mu.Unlock()
		return
	}
	p.isActive = false
	p.mu.Unlock()

	close(p.tasks)
	p.wg.Wait()
}

func (p *MyPool) AddWork(task PoolTask) {
	p.mu.RLock()
	if !p.isActive {
		defer p.mu.RUnlock()
		return
	}
	p.mu.RUnlock()

	for {

		p.mu.RLock()
		if !p.isActive {
			defer p.mu.RUnlock()
			return
		}
		select {
		case p.tasks <- task:
			return
		default:
			continue
		}
		p.mu.RUnlock()
	}
}

func NewWorkerPool(numWorkers int, channelSize int) (*MyPool, error) {
	if numWorkers <= 0 {
		return nil, errors.New("non-positive number of workers")
	}
	if channelSize < 0 {
		return nil, errors.New("negative size of the channel")
	}

	var p *MyPool
	if channelSize == 0 {
		p = &MyPool{
			tasks:         make(chan PoolTask),
			numberWorkers: numWorkers,
		}
	} else {
		p = &MyPool{
			tasks:         make(chan PoolTask, channelSize),
			numberWorkers: numWorkers,
		}
	}

	return p, nil
}
