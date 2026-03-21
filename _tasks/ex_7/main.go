package main

import (
	"fmt"
	"sync"
)

func main() {
	// task1()
	task2()
}

func task1() {
	wg := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			fmt.Println(i)
		}()
	}

	wg.Wait()
}

func task2() {
	wg := &sync.WaitGroup{}
	mu := sync.Mutex{}

	counter := 0
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			mu.Lock()
			defer mu.Unlock()
			counter++
		}()
	}
	wg.Wait()

	fmt.Println(counter)
}
