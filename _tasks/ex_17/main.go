package main

import (
	"fmt"
)

// type Request struct {
// 	Payload string
// }

// type Client interface {
// 	SendRequest(ctx context.Context, request Request) error
// 	WithLimiter(ctx context.Context, requests []Request)
// }

// type client struct{}

func main() {
	// ctx := context.Background()
	// c := client{}

	println("start")

	requests := make([]string, 0, 4)
	for i := 0; i < 4; i++ {
		requests = append(requests, fmt.Sprintf("%d", i*i))
	}
	fmt.Printf("%+v", requests)

}
