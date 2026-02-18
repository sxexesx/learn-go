package main

import "fmt"

type User struct{}

func main() {
	m := make(map[User]int)

	u := User{}

	// m[u] = 2

	fmt.Println(&m[u]) // нельзя взять из-за эвакуации данных
	fmt.Println(m[&u]) // не тот тип

}
