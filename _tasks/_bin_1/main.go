package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

const numRequests = 10000

var count atomic.Int32

func networkRequest(wg *sync.WaitGroup) {
	defer wg.Done()

	time.Sleep(time.Millisecond)
	count.Add(1)
}

func main() {
	start := time.Now()
	wg := sync.WaitGroup{}

	for i := 0; i < numRequests; i++ {
		wg.Add(1)
		go networkRequest(&wg)
	}

	wg.Wait()

	fmt.Println(count.Load())
	fmt.Println("processing time ... ", time.Since(start), "s")
}
