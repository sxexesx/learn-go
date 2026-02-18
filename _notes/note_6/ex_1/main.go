package main

import "fmt"

type SomeStruct struct{}

func foo() interface{} {

	var result *SomeStruct // ссылка на тип <- nil

	return result
	// информация о типе есть
	// а значение будет nil
}

func main() {
	res := foo()

	if res != nil {
		fmt.Println("res != nil, res = ", res)
		// значение равно nil
	}

}
