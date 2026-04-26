package main

import (
	"container/list"
	"fmt"
)

func main() {
	// ll := list.List{}
	ll := list.New()
	fmt.Printf("front %v\n", ll.Front())

	ll.PushFront(Order{
		Item:  100,
		Sku:   100,
		Title: "iPhone 17",
	})
	fmt.Printf("front %v\n", ll.Front())

	ll.PushFront(Order{
		Item:  200,
		Sku:   200,
		Title: "iPhone 26",
	})
	fmt.Printf("front %v\n", ll.Front())

	ll.PushFront(Order{
		Item:  300,
		Sku:   300,
		Title: "iPhone 56",
	})
	fmt.Printf("front %v\n", ll.Front())

	for i := ll.Front(); i != nil; i = i.Next() {
		fmt.Printf("Element: %+v\n", i)

	}

	// e1 := ll.PushBack(1)
	// _ = ll.PushFront(4)

	// ll.InsertAfter(3, e1)

	// firstElement := ll.Front()
	// fmt.Printf("First element: %v\n", firstElement.Value)
	// fmt.Printf("length %v\n", ll.Len())

	// for e := ll.Front(); e != nil; e = e.Next() {
	// 	// if eNext := e.Next(); eNext == nil {
	// 	// 	fmt.Printf("Element: %v\n", eNext)

	// 	// 	break
	// 	// }
	// 	fmt.Printf("Element: %v\n", e.Value)
	// }

}

type Order struct {
	Item  int
	Sku   int
	Title string
}
