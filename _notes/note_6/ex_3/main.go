package main

import "fmt"

type ABC interface {
	A()
	B()
	C()
}

type abc struct{}

func (a abc) A() {}
func (a abc) B() {}
func (a abc) C() {}

type ab struct{}

func (a ab) A() {}
func (a ab) B() {}

func main() {
	// var a interface{}

	// переменная a имеет конкретный тип, поэтмоу type assertion не срабатывает
	// a := abc{}
	// a1 := a.(ABC)

	// fmt.Println(a1)

	// нельзя привести, тк ab{} не реализовывает интерфейс ab
	// упадет в runtime, т.к.
	var b interface{}
	b = ab{}
	b1 := b.(ABC)

	fmt.Println(b1)
}
