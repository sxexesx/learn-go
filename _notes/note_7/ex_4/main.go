package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)

	for i := 0; i < 10; i++ {
		go func() {
			fmt.Printf("Goroutine %v\n", i)
		}()
		runtime.Gosched()
	}

	time.Sleep(time.Second * 1)
}
