package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	go buy(ch)
	go sell(ch)
	time.Sleep(1 * time.Second)
}

func buy(ch chan string) {
	fmt.Println("Waiting for a data")
	fmt.Println("Data received:", <-ch)
}

func sell(ch chan string) {
	ch <- "Samsung Galaxy S10"
	fmt.Println("Data sent")
}
