package main

import "fmt"

type Address struct {
	city   string
	street string
	house  int
}

// тут ошибка
func (a *Address) setCity(city string) {
	*a = Address{
		city:   city,
		street: a.street,
		house:  a.house,
	}
	// a = &Address{
	// 	city: city,
	// }
}

// правильно
func (a *Address) setStreet(street string) {
	a.street = street
}

// тут ошибка
func setHouse(addr *Address, house int) {
	// addr = &Address{
	// 	house: house,
	// }
	fmt.Printf("set house address = %p\n", addr)
	addr.house = house
}

func main() {
	addr := Address{
		city:   "New York",
		street: "5 Avenu",
		house:  10,
	}

	addr.setCity("London")
	addr.setStreet("Piccadilly")

	fmt.Printf("main address = %p\n", &addr)
	setHouse(&addr, 5)

	fmt.Println(addr)
}
