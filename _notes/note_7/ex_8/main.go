package main

import (
	"fmt"
	"time"
)

func main() {
	reader(doubler(writer()))
}

// генерирует числа от 1 до 10
func writer() <-chan int {
	ch := make(chan int)

	go func() {
		for i := range 10 {
			ch <- i + 1
		}
		close(ch)
	}()

	return ch
}

// умножает числа на 2, имитируя работу c разницей 500 мс
func doubler(ch <-chan int) <-chan int {
	cch := make(chan int)

	go func() {
		for a := range ch {
			time.Sleep(time.Duration(500) * time.Millisecond)
			cch <- a * 2
		}
		close(cch)
	}()

	return cch
}

// читает и выводит на экран
func reader(ch <-chan int) {
	for i := range ch {
		fmt.Println(i)
	}
}
