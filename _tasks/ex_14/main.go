package main

import "fmt"

type c chan c

func main() {
	var c = make(c, 1)
	c <- c

	for i := 0; i < 100; i++ {
		select {
		case <-c:
		case <-c:
			c <- c
		default:
			fmt.Println(i)
			return
		}
	}
}

// В каждой итерации мы можем попасть на любой из кейсов. Если попадаем на второй кейс,
// тогда значение вычитывается и сразу же записывается новое значение

// Когда мы попадаем на первый case, тогда следующей итерацией будет default и программа закончится.
