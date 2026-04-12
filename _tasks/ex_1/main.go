package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	predicableFunc()
}

func unpredictableFunc() int {
	n := rand.Intn(100)
	time.Sleep(time.Duration(n) * time.Second)
	return n
}

const (
	pause = 10 * time.Second
)

func predicableFunc() {
	c := make(chan int)

	go func() {
		c <- unpredictableFunc()
		close(c)
	}()

	select {
	case v := <-c:
		fmt.Println("Func done with ", v, " seconds")
	case <-time.After(pause):
		fmt.Println("Deadline is exceeded")
	}
}
