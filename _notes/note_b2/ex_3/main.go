package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := new(sync.WaitGroup)
	ch := make(chan int)

	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			val := i * 2
			ch <- val

			time.Sleep(time.Duration(val) * time.Second)
			fmt.Printf("Горутина %d выполнилась\n", i)
		}()
	}

	// ВАЖНО что закрываем канал в отдельной горутине, когда дождались всеx остальныx
	go func() {
		wg.Wait()

		close(ch)
	}()

	var acc int
	for x := range ch {
		acc += x
	}
	fmt.Println("Result", acc)
}
