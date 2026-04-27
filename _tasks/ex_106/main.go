package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type myMutex struct {
	locker int64
}

func (m *myMutex) Lock() {
	for {
		if atomic.CompareAndSwapInt64(&m.locker, 0, 1) {
			return
		}
	}
}

func (m *myMutex) Unlock() {
	atomic.StoreInt64(&m.locker, 0)
}

func main() {
	wg := &sync.WaitGroup{}
	mu := myMutex{}

	wg.Add(1000)
	c := 0
	for range 1000 {
		go func() {
			defer wg.Done()

			mu.Lock()
			c++
			mu.Unlock()
		}()
	}
	wg.Wait()

	fmt.Println(c)
}
