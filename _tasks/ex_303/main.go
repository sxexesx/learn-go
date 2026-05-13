package main

import (
	"errors"
	"fmt"
	"sync"
)

var ErrStopped = errors.New("ErrStopped")

func main() {
	wp := NewWorkerPool(3)
	f := func() {
		fmt.Println("do sth")
	}
	wp.Submit(f)
}

type WorkerPool struct {
	mu      sync.Mutex
	cond    *sync.Cond
	wg      sync.WaitGroup
	queue   []Task
	stopped bool
	drain   bool

	// stop chan struct{}
}

type Task struct {
	t    func()
	done chan error
}

func NewWorkerPool(numberOfWorkers int) *WorkerPool {
	if numberOfWorkers <= 0 {
		numberOfWorkers = 1
	}

	wp := &WorkerPool{}

	wp.wg.Add(numberOfWorkers)
	for i := 0; i < numberOfWorkers; i++ {
		go wp.worker()
	}

	return wp
}

func (wp *WorkerPool) worker() {
	defer wp.wg.Done()

	for {
		wp.mu.Lock()

		for len(wp.queue) == 0 && !wp.stopped {
			wp.cond.Wait()
		}

		if wp.stopped && len(wp.queue) == 0 {
			wp.mu.Unlock()
		}

		task := wp.queue[0]
		wp.queue = wp.queue[1:]

		wp.mu.Unlock()

		task.t()

		if task.done != nil {
			task.done <- nil
		}
	}
}

// Submit - добавить таску в воркер пул и возвращает управление (неблокирующая операция).
// Обеспечить соблюдение очереди при запуске задач.
func (wp *WorkerPool) Submit(task func()) error {
	wp.mu.Lock()
	defer wp.mu.Unlock()

	if wp.stopped {
		return ErrStopped
	}

	wp.queue = append(wp.queue, Task{t: task})

	wp.cond.Signal()
	return nil
}

// SubmitWait - добавить таску в воркер пул и дождаться окончания ее выполнения.
// Если был вызван метод Stop, SubmitWait выходит с ошибкой ErrStopped для задач
// которые ждут выполнения. Задачи которые выполняются должны доработать.
func (wp *WorkerPool) SubmitWait(task func()) error {
	done := make(chan error, 1)

	wp.mu.Lock()

	if wp.stopped {
		wp.mu.Unlock()
		return ErrStopped
	}

	wp.queue = append(wp.queue, Task{
		t:    task,
		done: done,
	})

	wp.cond.Signal()
	wp.mu.Unlock()

	return <-done
}

// Stop - остановить воркер пул, дождаться выполнения только тех тасок, которые выполняются сейчас
func (wp *WorkerPool) Stop() error {
	wp.mu.Lock()

	if wp.stopped {
		wp.mu.Unlock()
		wp.wg.Wait()

		return nil
	}
	wp.stopped = true
	wp.drain = false

	for _, task := range wp.queue {
		if task.done != nil {
			task.done <- ErrStopped
		}
	}

	wp.queue = nil
	wp.cond.Broadcast()

	wp.mu.Unlock()
	wp.wg.Wait()
	return nil
}

// StopWait - остановить воркер пул, дождаться выполнения всех тасок, даже тех, что не начали выполняться, но лежат в очереди
func (wp *WorkerPool) StopWait() error {
	wp.mu.Lock()

	if !wp.stopped {
		wp.stopped = true
		wp.drain = true

		wp.cond.Broadcast()
	}

	wp.wg.Wait()
	return nil
}
