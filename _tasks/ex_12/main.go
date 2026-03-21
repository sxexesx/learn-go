package main

import "fmt"

func main() {
	ch := make(chan int, 1)

	for i := 0; i < 5; i++ {
		select {
		case val := <-ch:
			fmt.Println(val)
		case ch <- i:
		}
	}
}

// 0 -> добавление 0 в канал
// 1 -> чтение из канала
// 2 -> добавление 2 в канал
// 3 -> чтение из канала
// 4 -> добавление в канал

// ответ:
// 0
// 2
