package main

import "fmt"

func main() {
	a := []int{1, 2, 3} //len = 3, cap = 3, a = [1, 2, 3]
	println("cap(a) = ", cap(a))

	b := append(a, 4) // len = 4, cap = 6, b = [1,2,3,4] ,0,0
	c := append(a, 5) // len = 4, cap = 6, c = [1,2,3,5] ,0,0

	c[1] = 0

	fmt.Println("a = ", a) // 1, 2, 3
	fmt.Println("b = ", b) // 1, 2, 3, 4
	fmt.Println("c = ", c) // 1, 0, 3, 5
}
