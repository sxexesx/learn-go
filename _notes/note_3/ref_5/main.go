package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	//////////////
	// пример 1 //
	//////////////
	p := getPerson()
	fmt.Println(p)

	//////////////
	// пример 2 //
	//////////////
	w := getSlice()
	fmt.Printf("main w = %p\n", &w)

	//////////////
	// пример 3 //
	//////////////
	a, b := named()
	fmt.Println(a, b)
	fmt.Printf("named. pointers. a = %p, b = %p\n", &a, &b)
	a, b = unnamed()
	fmt.Println(a, b)
}

func changeName(p *Person, name string) {
	p.name = name
}

func getPerson() (p Person) {
	// var p Person

	defer changeName(&p, "Alice")

	p = Person{
		name: "Bob",
		age:  52,
	}
	return p
}

func changeSlice(s []int) {
	s[1] = 10
	fmt.Printf("defer s = %p\n", &s)
}

func getSlice() []int {
	s := []int{1, 2, 3}
	fmt.Printf("getSlice. pointer = %p\n", &s)

	defer changeSlice(s)

	fmt.Printf("getSlice s = %p\n", &s)
	return s
}

func named() (a, b int) {
	a, b = 1, 2
	fmt.Printf("named. pointers. a = %p, b = %p\n", &a, &b)

	defer func() {
		a = 10
		b = 20
	}()

	return a, b
}

func unnamed() (int, int) {
	a, b := 1, 2
	fmt.Printf("named. pointers. a = %p, b = %p\n", &a, &b)

	defer func() {
		a = 10
		b = 20
	}()

	return a, b
}
