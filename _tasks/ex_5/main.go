package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	predicableTimeWork()
}

func randomTimeWork() {
	time.Sleep(time.Duration(rand.Intn(100)) * time.Second)
}

// решение
func predicableTimeWork() {
	ch := make(chan struct{})

	go func() {
		randomTimeWork()
		close(ch)
	}()

	select {
	case <-ch:
		fmt.Println("3 seconds not exceeded")
	case <-time.After(3 * time.Second):
		fmt.Println("error: 3 seconds exceeded")
	}
}
