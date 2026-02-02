package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// a := 1
	// println("before", &a)
	// foo()
	// println("after", &a)

	// указатель
	a := false
	b := true
	c := &a
	d := &b
	fmt.Printf("value = %v, address = %p\n", a, &a)
	fmt.Printf("value = %v, address = %p\n", b, &b)
	fmt.Printf("value = %v, address = %p\n", c, &c)
	println(unsafe.Sizeof(c))
	fmt.Printf("value = %v, address = %p\n", d, &d)
	println(unsafe.Sizeof(d))
}

//go:noinline
func foo() {
	bara()
	baza()
}

//go:noinline
// func bar() {
// 	arr := [1500]byte{}
// 	println(&arr)
//}

//go:noinline
func bara() *int {
	a := 100
	return &a
}

//go:noinline
func baza() int {
	c := 100
	return c
}
