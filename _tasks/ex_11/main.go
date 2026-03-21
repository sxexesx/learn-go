package main

import "fmt"

func spawnMessagesCorrect(n int) chan string {
	ch := make(chan string, 1)

	go func() {
		for i := 0; i < n; i++ {
			ch <- fmt.Sprintf("msg %d", i+1)
		}
		close(ch)
	}()

	return ch
}

func spawnMessages(n int) chan string {
	ch := make(chan string, 1)

	for i := 0; i < n; i++ {
		ch <- fmt.Sprintf("msg %d", i+1)
	}
	return ch
}

func main() {
	n := 10

	// for msg := range spawnMessagesCorrect(n) {
	for msg := range spawnMessages(n) {
		fmt.Println("received:", msg)
	}
}

// ответ:
// произойдет моментальная блокировка, потому что запись выполняется запись
// в переполненный буфферизированный канал
// как исправить? добавить запись в отдельной горутине
