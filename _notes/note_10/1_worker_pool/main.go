package main

import (
	"fmt"
	"sync"
)

func main() {
	jobs := make(chan int)
	results := make(chan int)

	wg := sync.WaitGroup{}

	limit := 3

	go func() {
		for i := range 10 {
			jobs <- i
		}
		close(jobs)
	}()

	for i := 0; i < limit; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for j := range jobs {
				results <- j * 2
			}
		}()
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for c := range results {
		fmt.Println(c)
	}
}
