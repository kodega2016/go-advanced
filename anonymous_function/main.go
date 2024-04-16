package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		fmt.Println("Hello, World!")
	}()

	time.Sleep(time.Second * 1)
}
