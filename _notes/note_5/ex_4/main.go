package main

import "fmt"

func main() {
	a := []int{} // len = 0, cap = 0, a = []

	for i := range 3 {
		a = append(a, i+1)
	}
	// 0 len = 1, cap = 1, a = [1]
	// 1 len = 2, cap = 2, a = [1,2]
	// 2 len = 3, cap = 4, a = [1,2,3]

	b := append(a, 4) // len = 4, cap = 4, b = [1,2,3,4]
	c := append(b, 5) // len = 5, cap = 8, c = [1,2,3,4,5],0,0,0

	c[1] = 0

	fmt.Println("a = ", a) // [1,2,3]
	fmt.Println("b = ", b) // [1,2,3,4]
	fmt.Println("c = ", c) // [1,0,3,4,5],0,0,0

	d := a[0:4]
	fmt.Println(d)
}
