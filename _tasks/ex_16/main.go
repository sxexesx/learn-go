package main

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"time"
)

type Request struct {
	Payload string
}

type Client interface {
	SendRequest(ctx context.Context, request Request) error
	WithLimiter(ctx context.Context, requests []Request)
}

type client struct{}

func (c client) SendRequest(ctx context.Context, request Request) error {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("sending request", request.Payload)
	return nil
}

// 1 - ограничение на количество коннектов
// 2 - одновременно работающих горутин
// 3 - ограничение RPS

var maxConnects = 10

func (c client) WithLimiter1(ctx context.Context, ch chan Request) {
	wg := sync.WaitGroup{}
	wg.Add(maxConnects)

	for range maxConnects {

		go func() {
			defer wg.Done()

			for req := range ch {
				c.SendRequest(ctx, req)
			}
		}()
	}

	wg.Wait()
}

var maxGoroutines = 10000

func (c client) WithLimiter2(ctx context.Context, reqs []Request) {
	tokens := make(chan struct{}, maxGoroutines)

	go func() {
		for range maxGoroutines {
			tokens <- struct{}{}
		}
	}()

	for _, req := range reqs {
		<-tokens
		go func() {
			defer func() {
				tokens <- struct{}{}
			}()

			c.SendRequest(ctx, req)
		}()
	}

	for range maxGoroutines {
		<-tokens
	}
}

var rps = 100

func (c client) WithLimiter3(ctx context.Context, reqs []Request) {
	ticker := time.NewTicker(time.Second / time.Duration(rps))
	wg := sync.WaitGroup{}

	wg.Add(len(reqs))
	for _, req := range reqs {
		<-ticker.C
		go func() {
			defer wg.Done()

			c.SendRequest(ctx, req)
		}()
	}

	wg.Wait()
}

func main() {
	ctx := context.Background()
	c := client{}

	requests := make([]Request, 1000)
	for i := 0; i < 1000; i++ {
		requests[i] = Request{
			Payload: strconv.Itoa(i),
		}
	}
	c.WithLimiter1(ctx, generate(requests))
	// c.WithLimiter2(ctx, generate(requests))
}

func generate(reqs []Request) chan Request {
	ch := make(chan Request)

	go func() {
		for _, v := range reqs {
			ch <- v
		}
		close(ch)
	}()
	return ch
}
