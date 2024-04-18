package main

import (
	"fmt"
	"sync"
)

func main() {
	nums := [5]int{1, 2, 3, 4, 5}

	for i, item := range nums {
		fmt.Println(i, item)
	}

	var wg sync.WaitGroup
	wg.Add(2)
	ch := make(chan int, 2)
	go sell(ch, &wg)
	go buy(ch, &wg)
	wg.Wait()
}

func sell(ch chan int, wg *sync.WaitGroup) {
	ch <- 1
	ch <- 2
	fmt.Println("send all data")
	close(ch)
	wg.Done()
}

func buy(ch chan int, wg *sync.WaitGroup) {
	fmt.Println("receive data")
	// fmt.Println(<-ch)

	for data := range ch {
		fmt.Println("received channel data-", data)
	}

	wg.Done()
}
