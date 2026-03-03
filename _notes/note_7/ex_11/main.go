package main

import (
	"fmt"
	"net/http"
	"strconv"

	"context"
	// "strings"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second)
	ch := make(chan any)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// configLoaded := 3
	// go example1(configLoaded, ticker, ch)

	configLoadedStr := "3"
	errCh := make(chan any)
	go example2(ctx, configLoadedStr, ticker, errCh, ch)

	select {
	case err := <-errCh:
		fmt.Println("error parsing config ", err)
	case <-ch:
	case <-ctx.Done():
		fmt.Println("Context done")
	}

	fmt.Println("server started...")
	server := http.Server{
		Addr: "localhost:8000",
	}
	server.ListenAndServe()
}

func example1(configLoaded int, ticker *time.Ticker, ch chan any) {
	defer ticker.Stop()

	counter := 0
	for {
		select {
		case <-ticker.C:
			counter++
			fmt.Println("counter increased, new value = ", counter)

			if counter == configLoaded {
				close(ch)

				// Исправление 1: завершаем работу горутины, когда она больше не нужна
				return
			}
		}
	}
}

func example2(ctx context.Context, configLoadedStr string, ticker *time.Ticker, errCh chan any, ch chan any) {
	defer ticker.Stop()

	counter := 0
	for {
		select {
		// Исправление 3:
		case <-ctx.Done():
			return

		case <-ticker.C:
			counter++
			fmt.Println("counter increased, new value = ", counter)

			_, err := strconv.Atoi(configLoadedStr)
			if err != nil {
				errCh <- err
				return
			}
		}
	}
}
