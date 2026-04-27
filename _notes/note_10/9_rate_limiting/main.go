package main

import (
	"fmt"
	"time"
)

func main() {
	jobs := make(chan int)

	go func() {
		defer close(jobs)

		for i := 0; i < 10; i++ {
			jobs <- i
		}
	}()

	ticker := time.Tick(1 * time.Second)

	for j := range jobs {
		<-ticker
		fmt.Println("process job ", j, time.Now().Format("15:04:05.000"))
	}
	fmt.Println("done")
}
