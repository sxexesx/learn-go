package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"golang.org/x/tools/go/analysis/passes/defers"
)

type Data struct {
	threadID int
	data     int
}

// var bufferPool = sync.Pool{
// 	New: func() any {
// 		return 10
// 	},
// }

func main() {
	wg := sync.WaitGroup{}
	wg.Add(3)

	data := make(chan Data)
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Duration(10)*time.Second)
	defer cancel()

	for i := 0; i < 3; i++ {
		go networkRequest(&wg, i, data)
	}

	go func() {
		wg.Wait()
		close(data)
	}()

	for v := range data {
		fmt.Println("[", v.threadID, "] ", "Finish waiting for ", v.data, " seconds. ", " Buffer pool value: ", bufferPool.Get().(int))
	}

}

func networkRequest(wg *sync.WaitGroup, threadID int, data chan Data) {
	defer wg.Done()

	// a := bufferPool.Get().(int)
	timeout := rand.Intn(30)

	// fmt.Println("[", threadID, "] ", "Executing sync.Pool value: ", a, " Random timout: ", timeout)

	// if timeout > a {
	// 	timeout = a
	// } else {
	// 	fmt.Println("[", threadID, "] ", "Putting value into sync.Pool ", timeout)
	// 	bufferPool.Put(timeout)
	// }

	fmt.Println("[", threadID, "] ", "Waiting for ", timeout, " seconds.")
	time.Sleep(time.Duration(timeout) * time.Second)

	data <- Data{
		threadID: threadID,
		data:     timeout,
	}
}
