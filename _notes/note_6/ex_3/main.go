package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func main() {
	// Сбор данных из нескольких источников
	// test_1()

	// Раздача данных для обработки несколькими воркерами
	test_2()
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

func test_2() {
	ch := make(chan int)

	wg := new(sync.WaitGroup)
	for range 10 {
		wg.Add(1)
		go consumer(wg, ch)
	}

	for i := 0; i < 10; i++ {
		x := rand.IntN(10)
		ch <- x
	}

	wg.Wait()
	close(ch)

	fmt.Println("All done!")
}

func consumer(wg *sync.WaitGroup, ch <-chan int) {
	defer wg.Done()

	a := <-ch
	time.Sleep(time.Duration(a) * time.Second)
	fmt.Printf("Consumer got value %d\n", a)
}
