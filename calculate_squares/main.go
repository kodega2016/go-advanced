package main

import (
	"fmt"
	"time"
)

func calculateSquare(num int) {
	time.Sleep(1 * time.Second)
	fmt.Println("Square of", num, "is", num*num)
}

func main() {
	start := time.Now()

	for i := 0; i < 10; i++ {
		go calculateSquare(i)
	}

	elapsed := time.Since(start)
	time.Sleep(2 * time.Second)
	fmt.Println("Function took:", elapsed)
}
