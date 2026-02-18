package main

import "fmt"

type SomeError struct{}

func (s SomeError) Error() string {
	return "some error"
}

func foo() error {
	var result *SomeError

	return result
}

func main() {
	result := foo()

	if result != nil {
		fmt.Println("Error occured!", result)
		return
	}
}
