package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

var maxWaitingSeconds int32 = 5

func randomWait() int {
	workSeconds := rand.IntN(5 + 1)
	time.Sleep(time.Duration(workSeconds) * time.Second)
	return workSeconds
}

func main() {
	now := time.Now()
	totalWorkSeconds := 0

	ch := make(chan int)

	for range 100 {
		go func() {
			res := randomWait()
			ch <- res
		}()
	}

	for range 100 {
		totalWorkSeconds += <-ch
	}

	fmt.Printf("main: %v\n", time.Since(now).Seconds())
	fmt.Printf("total: %d\n", totalWorkSeconds)
}
