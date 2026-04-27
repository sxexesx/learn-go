package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	const limit = 3
	const jobs = 20

	sem := make(chan struct{}, limit)
	wg := sync.WaitGroup{}

	for i := 0; i < jobs; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			// занимаем слот симафора
			sem <- struct{}{}

			// освобождаем слот
			defer func() {
				<-sem
			}()

			fmt.Printf("start processing job[%d]...\n", i)
			time.Sleep(2 * time.Second)
			fmt.Printf("end processing job[%d]\n", i)
		}()
	}
	wg.Wait()
	fmt.Println("done")
}
