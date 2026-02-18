package main

import "fmt"

func main() {
	a := []int{} // len = 0, cap = 0, a = nil

	a = append(a, []int{1, 2, 3, 4, 5}...) // len = 5, cap = 6, [1,0,3,4,5],0

	b := append(a, 6) // len = 6, cap = 10, b = [1,2,3,4,5,6] <- ссылается на a
	c := append(a, 7) // len = 6, cap = 10, c = [1,2,3,4,5,7] <- ссылается на a, поэтому меняется и слайс b

	c[1] = 0

	fmt.Println("a = ", a) // a = [1,0,3,4,5]
	fmt.Println("b = ", b) // b = [1,0,3,4,5,7]
	fmt.Println("c = ", c) // c = [1,0,3,4,5,7]
}
