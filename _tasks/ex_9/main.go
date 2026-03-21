package main

import (
	"fmt"
	"time"
)

func worker() <-chan int {
	ch := make(chan int)

	go func() {
		time.Sleep(1 * time.Second)
		close(ch)
	}()

	return ch
}

func main() {
	start := time.Now()

	// операции запускаются последовательно,
	// блокировки не будет.
	_, _ = <-worker(), <-worker()

	// ответ:
	// 2s + ..µs
	fmt.Println(time.Since(start))
}
