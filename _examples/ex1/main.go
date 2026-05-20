package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for i := 0; i < 6; i++ {
		go networkRequest(i)
	}
}

func networkRequest(threadID int) {
	timeout := rand.Intn(20)
	fmt.Println("[", threadID, "]", " Waiting for ", timeout, " seconds.")
	time.Sleep(time.Duration(timeout) * time.Second)
}
