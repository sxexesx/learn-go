package main
package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// реализовать processParallel
// прокинуть контекст
func processDataT(ctx context.Context, val int) int {
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	return val * 2
}

func mainT() {
	in := make(chan int)
	out := make(chan int)

	go func() {
		for i := range 10 {
			in <- i
		}
		close(in)
	}()

	now := time.Now()
	processParallel(ctx, in, out, 5)

	for val := range out {
		fmt.Println("v = ", val)
	}

	fmt.Println("program duration", time.Since(now))
}

// операция должна выполняться не более 5 секунд
func processParallelT(ctx context.Context, in <-chan int, out chan<- int, numWorkers int) {
	// ...
}
