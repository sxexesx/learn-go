package main

import "fmt"

func main() {
	//////////////
	// пример 1 //
	//////////////

	a := []int{1, 2, 3}
	fmt.Println(a)

	//////////////
	// пример 2 //
	//////////////

	s := make([]int, 3, 4) // длина 3, емкость 4
	for i := range 3 {
		s[i] = i + 1
	}

	changeSlice(s)
	fmt.Printf("%v, len = %d, capacity = %d\n", s, len(s), cap(s))

	newS := s[0:4]
	fmt.Printf("%v, len = %d, capacity = %d\n", newS, len(newS), cap(newS))

	//////////////
	// пример 3 //
	//////////////

	m := make(map[string]int)
	m["Bob"] = 25
	changeMap(m)

	fmt.Println(m)

	//////////////
	// пример 4 //
	//////////////

	ch := make(chan string)

	go func() {
		ch <- "John"
	}()
}

func modifySlice(a []int) {
	a[1] = 10
}

func changeSlice(s []int) {
	s = append(s, 4)
}

func changeMap(m map[string]int) {
	m["Bob"] = 52
}

func readFromChannel(ch chan string) {
	val := <-ch
	fmt.Println(val)
}
