package main

import (
	"fmt"
	"sync"
)

func main() {
	jobs := make(chan int, 0)

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()

		for i := range 20 {
			jobs <- i
		}

		close(jobs)
	}()

	go func() {
		for j := range jobs {
			fmt.Println("worker [1]: ", j)
		}
	}()

	for j := range jobs {
		fmt.Println("worker [2]: ", j)
	}
	wg.Wait()
}
