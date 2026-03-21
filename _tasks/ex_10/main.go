package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			ch <- i
		}
		// close(ch)
	}()

	for n := range ch {
		fmt.Println(n)
	}
}

// ответ:
// Последовательно 100 чисел будут выведены, а потом произойдет deadlock.
// Канал заблокируется до прихода писателя.
// Для исравления необходимо закрыть канал
