package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p *Person) changeName(name string) {
	// жизненный цикл копии p завершается, когда завершается функция
	// p = &Person{
	// 	name: name,
	// }

	// I Исправление через dereference
	// *p = Person{
	// 	name: name,
	//  age: age,
	// }

	// II Исправление
	// p.name = name
}

func main() {
	p := Person{
		name: "John",
		age:  12,
	}

	p.changeName("Alice")
	fmt.Println(p)
}
