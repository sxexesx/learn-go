package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	arr := make([]chan int, 0)
	results := make(chan string)

	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	for i := 0; i < 3; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			c := make(chan int, 10)

			for range 10 {
				c <- rand.Intn(10)
			}

			mu.Lock()
			defer mu.Unlock()
			arr = append(arr, c)

			close(c)
		}()
	}

	wg.Wait()
	fmt.Println("len ", len(arr))

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for v := range arr[i] {
				results <- fmt.Sprintf("arr [%d]: %d", i, v)
			}
		}()
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for v := range results {
		fmt.Println(v)
	}
}
