package main

import "fmt"

func main() {
	s := make([]int, 0, 5) // len = 0, cap = 5, [],0,0,0,0,0
	s = append(s, 1, 2, 3) // len = 3, cap = 5, [1,2,3],0,0

	subSlice := s[1:3] // len = 2, cap = 4, subSlice = [2,3],0,0; based on s

	subSlice[0] = 99               // [99,3],0,0;
	subSlice = append(subSlice, 4) // len = 3, cap = 4, subSlice = [99,3,4],0; // s = [1,99,3],4,0

	s = append(s, 5, 6, 7) // len = 6, cap = 10, new  s = [1,99,3,5,6,7]

	subSlice[1] = 100 // [99,100,4],0

	fmt.Println("s: ", s)               // [1,99,3,5,6,7]
	fmt.Println("subSlice: ", subSlice) // [99,100,4],0

}
