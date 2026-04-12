package main

import (
	"fmt"
	// "sync"
	// "time"
)

func main() {
	c := make(chan int)

	for i := 0; i < 100; i++ {
		go func() {
			c <- i
		}()
	}

	for range 100 {
		v := <-c
		fmt.Println(v)
	}
}

/*
counter := make([]int, 1000)

	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			counter[i] = i
			// counter = append(counter, i)
			// time.Sleep(1000)
		}(i)
	}

	wg.Wait()
	fmt.Println(len(counter))
*/
