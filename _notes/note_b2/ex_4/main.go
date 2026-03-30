package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 5)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	go func() {
		for {
			v, ok := <-ch
			if !ok {
				return
			}

			fmt.Printf("Got value %d\n", v)
			time.Sleep(time.Duration(v) * time.Second)
		}
	}()
}
