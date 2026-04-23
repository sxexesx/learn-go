package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	storage := make([]int, 0)
	buffer := make([]int, 0)
	c := make(chan int, 10)
	rw := sync.RWMutex{}

	go func() {
		var counter int

		for {
			time.Sleep(1 * time.Second)
			counter++
			if counter > 10 {
				counter = 1
			}

			rw.Lock()
			buffer = append(buffer, counter)
			rw.Unlock()

			c <- counter
			fmt.Println("writing buffer... ", counter)
		}
	}()

	go func() {
		for {
			if len(c) == 10 {
				fmt.Println("flushing buffer...")

				tmp := make([]int, 0, 10)
				for i := 0; i < 10; i++ {
					tmp = append(tmp, <-c)
				}

				rw.Lock()
				storage = append(storage, tmp...)
				buffer = buffer[:0]
				rw.Unlock()
			}

			time.Sleep(100 * time.Millisecond)
		}
	}()

	for {
		time.Sleep(15 * time.Second)

		rw.RLock()
		fmt.Println("read storage after 15 sec")
		fmt.Println("Buffer: ", buffer)
		fmt.Println("Storage: ", storage)
		println("-------------------------")
		rw.RUnlock()
	}
}
