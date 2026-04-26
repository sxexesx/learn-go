package main

import (
	"fmt"
	"time"
)

func main() {
	jobs := make(chan struct{})

	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(time.Second)
			jobs <- struct{}{}
		}
		close(jobs)
	}()

	for {
		select {
		case v, ok := <-jobs:
			if !ok {
				fmt.Println("done")
				return
			}
			fmt.Printf("%+v\n", v)

		case <-time.After(1500 * time.Millisecond):
			fmt.Printf("timeout")
			return
		}
	}
}
