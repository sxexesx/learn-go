package main

import "fmt"

func main() {
	a := []int{} // len = 0, cap = 0

	a = append(a, []int{1, 2, 3, 4, 5}...) // len = 5, cap = 6, a = [1,2,3,4,5]

	b := append(a, 6) // len = 6, cap = 6, b (= a) = [1,2,3,4,5,6]
	c := append(b, 7) // len = 7, cap = 12, c = [1,2,3,4,5,6,7]

	c[1] = 0

	fmt.Println("a = ", a) // [1,2,3,4,5]
	fmt.Println("b = ", b) // [1,2,3,4,5,6]
	fmt.Println("c = ", c) // [1,0,3,4,5,6,7]

	d := c[0:11] // len = 7, cap = 12, ...
	println(d)   // [1,0,3,4,5,6,7,0,0,0,0,0]
}
