package main

import (
	"context"
	"fmt"

	"time"
)

func main() {
	start := time.Now()

	a, err := getFiles(context.TODO(), "1", "2", "3", "4", "5")
	if err != nil {
		panic("error!")
	}

	fmt.Println(a)
	fmt.Println("time = ", time.Since(start))
}

func getFiles(ctx context.Context, names ...string) (result map[string][]byte, err error) {
	if len(names) == 0 {
		return nil, nil
	}

	// wg := &sync.WaitGroup{}
	// ch := make(chan map[string]byte)
	ch := make(chan struct {
		name string
		data []byte
	})
	defer close(ch)

	result = make(map[string][]byte, len(names))
	for _, name := range names {
		// wg.Add(1)
		go func() {
			// defer wg.Done()

			data, err := getFile(ctx, name)
			if err != nil {
				panic("error!")
			}
			ch <- struct {
				name string
				data []byte
			}{
				name: name,
				data: data,
			}

			// ch <- data
			// result[name] = data
		}()
	}

	go func() {
		b := <-ch
		key := b.name
		result[key] = b.data
	}()

	// wg.Wait()

	// select {
	// case ctx.Done():
	// case
	// }

	return result, nil
}

func getFile(ctx context.Context, name string) ([]byte, error) {

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-ticker.C:

	}

	return []byte(name), nil
}
