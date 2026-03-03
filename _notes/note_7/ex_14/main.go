package main

import (
	"fmt"
	"runtime"
	"time"
)

func leakyFunc() {
	go func() {
		select {}
	}()
}

func main() {
	ticker := time.NewTicker(time.Second)
	var m runtime.MemStats

	for {
		select {
		case <-ticker.C:
			runtime.ReadMemStats(&m)
			for range 10000 {
				leakyFunc()
			}

			fmt.Println("num goroutines: ", runtime.NumGoroutine())
			fmt.Println("total heap: ", m.HeapAlloc/1024/1024, "Mb")
			fmt.Println("total stack: ", m.StackInuse/1024/1024, "Mb")
		}
	}
}
