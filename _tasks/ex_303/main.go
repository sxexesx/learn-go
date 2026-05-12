package main

import (
	"fmt"
	"sync"
)

func main() {
	wp := NewWorkerPool(3)
	f := func() {
		fmt.Println("do sth")
	}
	wp.Submit(f)
}

func wp() {
	ch := make(chan func())
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			ch <- func() {}
		}()
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
}

type WorkerPool struct {
	// wg          sync.WaitGroup{}
	// queue       chan func()
	// results     chan struct()
	// buffer      int
}

// type ElementType struct {
//     id int
//     fun func()
// }

// type Task struct {
//     fun     func()
//     result  chan bool
// }

func NewWorkerPool(numberOfWorkers int) *WorkerPool {
	jobs := make(chan struct{}, numberOfWorkers)

	for i := 0; i < numberOfWorkers; i++ {
		go func() {

			// for {
			task := <-queue
			task.fun()

			task.result <- true
			// }
		}()
	}

	// return &WorkerPool{
	//     queue: queue,
	//     // result
	// }
}

// Submit - добавить таску в воркер пул и возвращает управление (неблокирующая операция).
// Обеспечить соблюдение очереди при запуске задач.
func (wp *WorkerPool) Submit(task func()) error {
	// wp.queue <- task()

	// return nil
}

// SubmitWait - добавить таску в воркер пул и дождаться окончания ее выполнения.
// Если был вызван метод Stop, SubmitWait выходит с ошибкой ErrStopped для задач
// которые ждут выполнения. Задачи которые выполняются должны доработать.
func (wp *WorkerPool) SubmitWait(task func()) error {
	// st := make(ch struct)

	// wp.queue.fun <- task()

	// <-wp.queue.result

	// return nil
}

// Stop - остановить воркер пул, дождаться выполнения только тех тасок, которые выполняются сейчас
func (wp *WorkerPool) Stop() error {
}

// StopWait - остановить воркер пул, дождаться выполнения всех тасок, даже тех, что не начали выполняться, но лежат в очереди
func (wp *WorkerPool) StopWait() error {
}
