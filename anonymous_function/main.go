package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		fmt.Println("Hello, World!")
	}()

	func() {
		fmt.Println("this is anonymous call")
	}() // anonymous function

	time.Sleep(time.Second * 1)
}
