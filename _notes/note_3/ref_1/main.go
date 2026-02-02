package main

import "fmt"

// за пределами функции name уже не живет
func changeName(name *string) {
	fmt.Printf("changeName value = %v, address = %p\n", name, &name)

	// name <- ссылка
	// &name - значение по ссыьке
	newName := "Alice"
	fmt.Printf("changeName value = %v, address = %p\n", newName, &newName)
	name = &newName // <- ссылке name присвоили ссылку на новое имя
	fmt.Printf("changeName value = %v, address = %p\n", name, &name)

	*name = newName
	// Получаем доступ по адресу на который указывает эта ссылка
	// *name = "WTF" // dereference - доступ к переменной name.
	fmt.Printf("changeName value = %v, address = %p\n", name, &name)
}

func main() {
	name := "John"
	fmt.Printf("main value = %v, address = %p\n", name, &name)
	changeName(&name)
	fmt.Printf("main value = %v, address = %p\n", name, &name)

}

// func changeName(name *string) {
// 	rname := name
// 	newName := "Alice"
// 	name = &newName
// 	*rname = newName
// }

// func main() {
// 	name := "John"
// 	fmt.Printf("main value = %v, address = %p\n", name, &name)
// 	changeName(&name)
// 	fmt.Printf("main value = %v, address = %p\n", name, &name)

// }
