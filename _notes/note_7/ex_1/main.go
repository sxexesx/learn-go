package main

import "fmt"

func main() {
	var ch chan int

	fmt.Printf("capacity=%d; length=%d; ch=%v\n", cap(ch), len(ch), ch)

	someVar := 3

	ch = make(chan int, 5)
	ch <- someVar
	fmt.Printf("capacity=%d; length=%d; ch=%v\n", cap(ch), len(ch), ch)

	a := <-ch
	fmt.Printf("capacity=%d; length=%d; ch=%v\n", cap(ch), len(ch), ch)
	println(a)
}
