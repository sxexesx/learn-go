package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := new(sync.WaitGroup)

	wg.Add(5) // вариант 1
	for i := 0; i < 5; i++ {
		wg.Add(1) // вариант 2
		go longOperation(i, wg)
	}
	wg.Wait()
	fmt.Printf("That's all\n")
}

// func longOperation(i int) {
// 	time.Sleep(time.Second * time.Duration(i))
// 	fmt.Printf("%v passed\n", i)
// }

func longOperation(i int, wg *sync.WaitGroup) {
	defer wg.Done()

	time.Sleep(time.Second * time.Duration(i))
	fmt.Printf("%v passed\n", i)
}
