package main

import "fmt"

const MAX = 5

func main() {
	s := generate() // len = 4, cap = 5, s = base = [1, 2, 3, 4],0

	matation(s) // len = 5; cap = 5, [1,2,3,4],-1

	fmt.Printf("s[0:MAX] = %v", s[0:MAX]) // <- [1,2,3,4,-1]
}

func generate() []int {
	out := make([]int, 0, MAX) // len = 0, cap = 5, out = []

	for i := 1; i < MAX; i++ {
		out = append(out, i)
	}
	// 1 len = 1, cap = 1, out = [1]
	// 2 len = 2, cap = 2, out = [1, 2]
	// 3 len = 3, cap = 4, out = [1, 2, 3]
	// 4 len = 4, cap = 4, out = [1, 2, 3, 4]

	return out
}

func matation(s []int) {
	s = append(s, -1)
}
