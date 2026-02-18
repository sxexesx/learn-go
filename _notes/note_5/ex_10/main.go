package main

import "sync"

func main() {
	m := make(map[int]int)

	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}

	wg.Add(100)

	for i := range 100 {
		go func() {
			defer wg.Done()

			mu.Lock()
			defer mu.Unlock()

			m[i] = i
		}()
	}

	wg.Wait()
}
