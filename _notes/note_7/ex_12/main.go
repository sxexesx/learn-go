package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		defer fmt.Println("writer exited")
		defer close(ch)

		for i := range 3 {
			select {
			case ch <- i:
				return
			case <-ctx.Done():
				fmt.Println("context done")
				return
			}
		}

		// close(ch)
	}()

	go func() {
		defer fmt.Println("reader exited")
		for i := range ch {
			fmt.Println("i = ", i)
			cancel()
			return
		}
	}()
	time.Sleep(3 * time.Second)

}
