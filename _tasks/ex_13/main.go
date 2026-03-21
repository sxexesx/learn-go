package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	tt, err := predictableFunc(context.Background())
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Function completed in ", tt, " seconds")
}

func predictableFunc(ctx context.Context) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	ch := make(chan int)
	go func() {
		ch <- unpredictableFunc()
		close(ch)
	}()

	select {
	case a := <-ch:
		return a, nil
	case <-ctx.Done():
		return 0, errors.New("time exceeded")
	}
}

func unpredictableFunc() int {
	n := rand.Intn(15)
	time.Sleep(time.Duration(n) * time.Second)

	return n
}
