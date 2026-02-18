package main

import "fmt"

func getBytes(start, end int) []byte {
	arr := [999999999]byte{}

	slice := arr[start:end]
	return slice
}

func main() {
	// s ссылается на массив огромной capacity и пока мы работаем с этим s в хипе лежит этот массив
	s := getBytes(10, 20)

	println(cap(s))
	fmt.Println(s)
}
