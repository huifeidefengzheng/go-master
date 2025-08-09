package main

import (
	"fmt"
	"sync"
)

func sendChannel(ch chan<- int) {
	for i := 0; i < 100; i++ {
		ch <- i
	}
	close(ch)
}

func receiveChannel(ch <-chan int) {
	for i := range ch {
		fmt.Println(i)
	}
}
func main() {

	ch := make(chan int, 10)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		sendChannel(ch)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		receiveChannel(ch)
	}()

	wg.Wait()
}
