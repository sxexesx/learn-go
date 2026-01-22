package main

import "fmt"

func main() {
	m := map[int]string{
		1: "one",
		2: "two",
		3: "three",
	}
	var i int = 0
	// i := 0

	inc(&i, m)
	fmt.Printf("%#+v\n", i)
	fmt.Printf("%#+v\n", m)

}

func inc(i *int, m map[int]string) {
	*i++

	for k := range m {
		m[k] = "zero"
	}
}
