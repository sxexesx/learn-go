package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := stage1()
	c2 := stage2(c1)
	for v := range c2 {
		fmt.Println(v)
	}
	fmt.Println("done")
}

func stage1() <-chan int {
	c := make(chan int)

	go func() {
		for i := range 10 {
			c <- i
		}
		close(c)
	}()

	return c
}

func stage2(c <-chan int) <-chan int {
	result := make(chan int)

	go func() {
		for v := range c {
			time.Sleep(500 * time.Millisecond)
			result <- v * 2
		}
		close(result)
	}()
	return result
}
