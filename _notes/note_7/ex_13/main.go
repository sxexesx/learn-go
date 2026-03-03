package main

import (
	"context"
	"fmt"
)

func main() {

}

func leakyChannels() {
	ch := make(chan int)

	go func() {
		for i := range 3 {
			ch <- i
		}
		// close(ch)
	}()

	go func() {
		for v := range ch {
			fmt.Println(v)
		}
	}()
}

func leakyFunc() {
	// ctx, cancel := context.WithCancel(context.Background())
	// defer cancel()

	go func() {
		select {
		// case <-ctx.Done():
		}
	}()
}
