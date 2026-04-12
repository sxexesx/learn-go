package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	res, err := predicableFunc()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Func done with ", res, " seconds")

}

func unpredictableFunc() int {
	n := rand.Intn(100)
	time.Sleep(time.Duration(n) * time.Second)
	return n
}

const (
	pause = time.Duration(10) * time.Second
)

func predicableFunc() (int, error) {
	c := make(chan int)
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	go func() {
		c <- unpredictableFunc()
		close(c)
	}()

	select {
	case v := <-c:
		return v, nil
	case <-ctx.Done():
		return 0, errors.New("Deadline is exceeded")
	}
}
