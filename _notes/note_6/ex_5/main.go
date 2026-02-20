package main

import "fmt"

func main() {
	// тип переменной x - это функция func() int
	var x = func() int { return 1 }
	x = nil

	test(x)
}

func test(x interface{}) {

	switch x.(type) {
	case int:
		fmt.Println("int", x)

	case string:
		fmt.Println("string", x)

	case nil:
		fmt.Println("nil", x)

	// правильный ответ, но будет ошибка компиляции, т.к. x() - вызов функции,
	// а x типа интерфейс
	case func() int:
		// убираем ошибку компиляции
		f := x.(func() int)
		// fmt.Println("func", x())
		// но всё равно будет nil pointer exception, т.к. функция равна nil
		fmt.Println("func", f())

	default:
		fmt.Println("unknown type")
	}
}
