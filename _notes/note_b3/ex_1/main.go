package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	errgroup "golang.org/x/sync/errgroup"
)

type User struct {
	Name string
}

func main() {
	names := []User{
		{"Ann"},
		{"Bob"},
		{"Cindy"},
		{"Bob"},
	}

	ctx := context.Background()

	start := time.Now()
	res, err := process(ctx, names)
	if err != nil {
		fmt.Println("an err occured", err.Error())
	}

	fmt.Println("time:", time.Since(start))
	fmt.Println(res)
}

func process(ctx context.Context, users []User) (map[string]int64, error) {
	names := make(map[string]int64, 0)
	mu := sync.Mutex{}

	egroup, ectx := errgroup.WithContext(ctx)
	egroup.SetLimit(100)

	for _, u := range users {
		egroup.Go(
			func() error {
				name, err := fetch(ectx, u)
				if err != nil {
					return err
				}

				mu.Lock()
				defer mu.Unlock()
				names[name] = names[name] + 1

				return nil
			})
	}

	if err := egroup.Wait(); err != nil {
		return nil, err
	}

	return names, nil
}

func fetch(ctx context.Context, user User) (string, error) {
	if user.Name == "Ann" {
		return "", errors.New("error")
	}

	ch := make(chan any)

	go func() {
		time.Sleep(10 * time.Millisecond)
		close(ch)
	}()

	select {
	case <-ch:
		return user.Name, nil
	case <-ctx.Done():
		return "", errors.New("context canceled")
	}
}
