package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	mu := http.NewServeMux()
	mu.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("test handler activated")
	})

	t := time.NewTicker(time.Second)
	go func() {
		for c := range t.C {
			a := c
			fmt.Println("tick", a)
		}
	}()

	http.ListenAndServe(":8080", mu)

	fmt.Println("finished")
}
