package main

import (
	"fmt"
	"net/http"
	"strconv"

	// "strings"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second)
	ch := make(chan any)

	// configLoaded := 3
	// go example1(configLoaded, ticker, ch)

	configLoadedStr := "3a"
	errCh := make(chan any)
	go example2(configLoadedStr, ticker, errCh, ch)

	select {
	case err := <-errCh:
		fmt.Println("error parsing config ", err)
	case <-ch:
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

func example2(configLoadedStr string, ticker *time.Ticker, errCh chan any, ch chan any) {
	defer ticker.Stop()

	counter := 0
	for {
		select {
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
