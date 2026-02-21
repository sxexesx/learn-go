package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	// money := 0
	var money atomic.Int32
	var donation atomic.Int32

	mutex := &sync.Mutex{}

	go func() {
		for {
			mutex.Lock()

			m := money.Load()
			d := donation.Load()
			mutex.Unlock()

			if m != d {
				fmt.Println(" error !!! money = ", m, " donation = ", d)
				break
			}
			// fmt.Println(m, " ", d)
		}
	}()

	wg := &sync.WaitGroup{}

	wg.Add(1000)
	for range 1000 {
		go func() {
			defer wg.Done()

			// money++
			mutex.Lock()
			money.Add(1)
			donation.Add(1)
			mutex.Unlock()
		}()
	}
	wg.Wait()

	// fmt.Println(money)
	fmt.Println(money.Load())
	fmt.Println(donation.Load())
}
