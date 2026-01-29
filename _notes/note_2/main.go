package main

import (
	"fmt"
	"net/http"
)

func main() {
	go122()
}

func go122() {
	// values := []int{1, 2, 3, 4, 5}
	// for _, val := range values {
	// 	go func() {
	// 		fmt.Printf("%d ", val) // ошибка: замыкание на переменной цикла, все горутины используют одно и то же значение val
	// 	}()
	// }

	// -----------
	for i := range 10 {
		fmt.Println(10 - i)
	}

	// -----------
	mux := http.NewServeMux()

	mux.HandleFunc("/item/", getItem)

	http.ListenAndServe(":8080", mux)
}

func getItem(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path
	fmt.Printf("path = %s", id)
}
