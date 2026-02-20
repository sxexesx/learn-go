package main

import (
	"fmt"
)

type SomeInterface interface {
}

type Interface interface {
	Method(in string) string
}

type InterfaceStruct struct {
	prefix string
}

func (i InterfaceStruct) Method(in string) string {
	return i.prefix + in
}

func main() {
	var si SomeInterface

	fmt.Printf("value %#=v'n", si)
	fmt.Printf("pointer %#=v'n", &si)

	var is Interface = InterfaceStruct{"I say: "}

	fmt.Printf("value %#+v\n", is)
	fmt.Printf("pointer %#+v\n", &is)
	fmt.Printf("value %q\n", is.Method("Hello"))
}
