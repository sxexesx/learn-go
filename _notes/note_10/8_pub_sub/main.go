package main

import (
	"fmt"
	"sync"
	"time"
)

type Broker struct {
	mu   sync.RWMutex
	subs []chan string
}

func NewBroker() *Broker {
	return &Broker{
		subs: make([]chan string, 0),
	}
}

func (b *Broker) Subscribe() <-chan string {
	c := make(chan string, 5)

	b.mu.Lock()
	b.subs = append(b.subs, c)
	b.mu.Unlock()

	return c
}

// публикация всем подписчикам
func (b *Broker) Publish(msg string) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	for _, ch := range b.subs {
		ch <- msg
	}
}

func (b *Broker) Close() {
	b.mu.Lock()
	defer b.mu.Unlock()

	for _, ch := range b.subs {
		close(ch)
	}
}

func main() {
	broker := NewBroker()

	sub1 := broker.Subscribe()
	sub2 := broker.Subscribe()

	go func() {
		for msg := range sub1 {
			fmt.Println("subscriber 1: ", msg)
		}
	}()

	go func() {
		for msg := range sub2 {
			fmt.Println("subscriber 2: ", msg)
		}
	}()

	broker.Publish("order created")
	broker.Publish("payment receive")

	time.Sleep(time.Second)

	broker.Close()
}
