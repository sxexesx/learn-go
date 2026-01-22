package main

import "fmt"

func main() {
	var a [7]string
	// a = append(a, "!")
	var b [7]string
	copy(b[:], a[:])

	a[6] = "!"
	a[0] = "Hello"
	a[1] = "World"

	fmt.Println(a[0], a[1])
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(len(a))
	fmt.Println(cap(a))
}
