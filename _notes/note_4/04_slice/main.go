package main

import "fmt"

// Пример показывает, что если емкость массива не меняется, тогда элемент в памяти может быть перезаписан
// Решается проблема добавлением max'a - элемента от которого расчитывать капасити
func main() {
	data := []string{"1", "2", "3", "4", "5"}
	addExtra(data[0:3:3]) // <---
	fmt.Println(data)
}

func addExtra(data []string) {
	data = append(data, "-------------------")
	fmt.Println(data)
}
