package main

import (
	"fmt"
	"sync"
)

func main() {
	// запись из неинициализированного канала - deadlock
	// test_1()

	// запись в открытый канал - блокировка до прихода читателя
	// test_2()

	// нормальная работа
	// test_3()

	//  чтение из закрытого канала
	// test_4()

	// test_5()

	fmt.Println("done")
}

func test_1() {
	var ch chan int

	ch <- 1
}

func test_2() {
	ch := make(chan int)

	ch <- 1
}

func test_3() {
	ch := make(chan int)

	go func() {
		ch <- 1
	}()

	v := <-ch

	fmt.Println(v)
}

func writer() chan int {
	ch := make(chan int)

	go func() {
		for i := range 5 {
			ch <- i + 1
		}

		close(ch)
	}()

	return ch
}

func test_4() {
	ch := writer()

	for {
		v, ok := <-ch
		if !ok {
			break
		}

		fmt.Println("v =", v)
	}
}

func writer2() <-chan int {
	ch := make(chan int)

	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()

		for i := range 5 {
			ch <- i + 1
		}
	}()

	go func() {
		defer wg.Done()

		for i := range 5 {
			ch <- i + 11
		}
	}()

	go func() {
		wg.Wait()

		close(ch)
	}()

	return ch
}

func test_5() {
	ch := writer2()

	for {
		v, ok := <-ch
		if !ok {
			break
		}

		fmt.Println("v =", v)
	}
}
