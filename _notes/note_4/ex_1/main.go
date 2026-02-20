package main

import "fmt"

func main() {
	greetings := "привет как дела"

	fmt.Println(len(greetings)) // 13 * 2 + 2 = 28

	fmt.Println("%v %b %c \n", greetings[1], greetings[1], greetings[1])
	// v = сырое
	// b = binary
	// c = symbol
}
