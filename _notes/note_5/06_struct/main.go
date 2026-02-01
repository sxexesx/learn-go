package main

import (
	"fmt"
	"strings"
)

// Пример разницы с вызовом по ссылки и по значению.
// В первом варианте (с Point) setX изменит исходный объект point
// Во втором - нет

func main() {
	point := Point{1, 1}

	point.SetX(100)
	fmt.Println(point)

	p := Person{
		Name:    "Igor",
		Surname: "Starshinov",
	}
	p.UpPerson()

	fmt.Printf("%#+v", p)
}

type Point struct {
	X int
	Y int
}

func (p *Point) SetX(v int) {
	p.X = v
}

type Person struct {
	Name    string
	Surname string
}

func (p Person) UpPerson() {
	p.Name = strings.ToUpper(p.Name)
	p.Surname = strings.ToUpper(p.Surname)
}
