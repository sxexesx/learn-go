package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
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
	// если в этом месте будет не инициализированна мапа, тогда будет ошибка
	names := make(map[string]int64, 0)

	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	for _, u := range users {
		wg.Add(1)
		go func() {
			defer wg.Done()

			name, err := fetch(ctx, u)
			if err != nil {
				sync.OnceFunc(func() {
					cancel()
					e = err
				})()
			}

			mu.Lock()
			defer mu.Unlock()
			names[name] = names[name] + 1
		}()
	}
	wg.Wait()

	if e != nil {
		return nil, e
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
