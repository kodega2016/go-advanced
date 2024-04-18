package main

import "fmt"

func main() {
	ch := make(chan int, 5)
	ch <- 100
	ch <- 200

	close(ch)

	for val := range ch {
		fmt.Println(val)
	}
}
