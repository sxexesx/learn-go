package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	test_1()
}

func test_1() {
	wg := new(sync.WaitGroup)
	ch := make(chan int)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go producer(i, wg, ch)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	var acc int
	for x := range ch {
		acc += x
	}
	fmt.Println("Result", acc)
}

func producer(i int, wg *sync.WaitGroup, ch chan<- int) {
	defer wg.Done()

	val := i * 2
	ch <- val
	time.Sleep(time.Duration(val) * time.Second)
	fmt.Printf("Горутина %d выполнилась\n", i)
}

// func test_2() {

// }

// ```golang
// func producer(ch chan<- string){}

// var ch chan string
// for i:= 0; i < 10; i++ {
//     go producer(ch)
// }

// for x := range ch {
//     // обрабатываем данные
// }

// 2. Раздача данных для обработки нескольким воркерам

// func consumer(ch chan<- string){}

// var ch chan string
// for i:= 0; i < 10; i++ {
//     go consumer(ch)
// }
// for {
//     ch <- time.Now().Format("")
// }
