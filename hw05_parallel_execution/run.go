package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

type errCounter struct {
	sync.RWMutex
	threshold int
	count     int
}

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	tasksCh := make(chan Task, len(tasks))
	for i := range tasks {
		tasksCh <- tasks[i]
	}
	close(tasksCh)

	ec := errCounter{threshold: m}
	wg := sync.WaitGroup{}

	for i := 0; i < n; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for {
				if ec.ThresholdExceeded() {
					return
				}

				task, ok := <-tasksCh
				if !ok {
					return
				}

				err := task()
				if err != nil {
					ec.Inc()
				}
			}
		}()
	}

	wg.Wait()

	if ec.ThresholdExceeded() {
		return ErrErrorsLimitExceeded
	}

	return nil
}

func (ec *errCounter) Inc() {
	ec.Lock()
	defer ec.Unlock()

	ec.count++
}

func (ec *errCounter) ThresholdExceeded() bool {
	ec.RLock()
	defer ec.RUnlock()

	if ec.threshold < 0 {
		return true
	}

	return ec.count >= ec.threshold
}
