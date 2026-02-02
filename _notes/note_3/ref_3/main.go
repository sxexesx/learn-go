package main

import "fmt"

type Address struct {
	city   string
	street string
	house  int
}

// тут ошибка
func (a *Address) setCity(city string) {
	// *a = Address{
	// 	city:   city,
	// 	street: a.city,
	// 	house:  a.house,
	// }
	a = &Address{
		city: city,
	}
}

// правильно
func (a *Address) setStreet(street string) {
	a.street = street
}

func setHouse(addr *Address, house int) {
	addr = &Address{
		house: house,
	}
}

func main() {
	addr := Address{
		city:   "New Your",
		street: "5 Avenu",
		house:  10,
	}

	addr.setCity("London")
	addr.setStreet("Piccadilly")
	setHouse(&addr, 5)

	fmt.Println(addr)
}
