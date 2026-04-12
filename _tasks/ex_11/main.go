package main

import "fmt"

func main() {
	n := 10

	for msg := range spawnMessagesCorrect2(n) {
		fmt.Println("received:", msg)
	}
}

func spawnMessagesCorrect1(n int) chan string {
	ch := make(chan string, 1)

	go func() { // <----
		for i := 0; i < n; i++ {
			ch <- fmt.Sprintf("msg %d", i+1)
		}
		close(ch) // <----
	}()

	return ch
}

func spawnMessagesCorrect2(n int) chan string {
	ch := make(chan string, n) // <----

	for i := 0; i < n; i++ {
		ch <- fmt.Sprintf("msg %d", i+1)
	}
	close(ch)

	return ch
}
