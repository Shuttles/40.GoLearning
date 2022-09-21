package workerpool

import (
	"errors"
	"fmt"
	"sync"
)

const (
	defaultCapacity = 100
	maxCapacity     = 10000
)

type Task func()

type Pool struct {
	capacity int

	active chan struct{}
	tasks chan Task

	quit chan struct{}
	wg sync.WaitGroup
}

var ErrWorkerPoolFreed =  errors.New("workerpool freed")

func New(capacity int) *Pool {
	if capacity < 0 {
		capacity = defaultCapacity
	}
	if capacity > maxCapacity {
		capacity = maxCapacity
	}

	p := &Pool{
		capacity: capacity,
		active: make(chan struct{}, capacity),
		tasks: make(chan Task),
		quit: make(chan struct{}),
	}

	fmt.Printf("workerpool start\n")

	go p.run()

	return p
}

func (p *Pool) run() {
	idx := 0
	for {
		select {
		case <-p.quit:
			return
		case p.active <- struct{}{}:
			idx++
			p.newWorker(idx)
		}
	}
}

func (p *Pool) newWorker(idx int) {
	p.wg.Add(1)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("worker[%03d]: recover panic[%s] and exit\n", idx, err)
				<-p.active
			}
			p.wg.Done()
		}()

		fmt.Printf("worker[%03d] start\n", idx)

		for {
			select {
			case <-p.quit:
				fmt.Printf("worker[%03d]: exit\n", idx)
				<-p.active
				return
			case t := <-p.tasks:
				fmt.Printf("worker[%03d] receive a task\n", idx)
				t()
			}
		}
	}()
}

func (p *Pool) Schedule(t Task) error {
	select {
	case <-p.quit:
		return ErrWorkerPoolFreed
	case p.tasks <- t:
		return nil
	}
}

func (p *Pool) Free() {
	close(p.quit)
	p.wg.Wait()
	fmt.Printf("workerpool freed\n")
}