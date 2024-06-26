package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go goOne(ch1)
	go goTwo(ch2)

	select {
	case val1 := <-ch1:
		fmt.Println(val1)
	case val2 := <-ch2:
		fmt.Println(val2)
	default:
		fmt.Println("default case")
	}

	time.Sleep(time.Second * 1)
}

func goOne(ch chan string) {
	ch <- "channel one"
}

func goTwo(ch chan string) {
	ch <- "channel two"
}
