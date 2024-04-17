package main

import (
	"fmt"
	"sync"
	"time"
)

func calculateSqaure(n int, wg *sync.WaitGroup) {
	fmt.Println(n * n)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(10)
	start := time.Now()
	for i := 0; i < 10; i++ {
		go calculateSqaure(i, &wg)
	}
	elapsed := time.Since(start)
	wg.Wait()
	fmt.Printf("Time taken: %s\n", elapsed)
}
