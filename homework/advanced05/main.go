package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	count int
	lock  sync.Mutex
}

func (c *Counter) Increment() {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.count++
}

func (c *Counter) Get() int {
	c.lock.Lock()
	defer c.lock.Unlock()
	return c.count
}

func main() {
	counter := &Counter{}

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(goroutineID int) {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				counter.Increment()
			}
			fmt.Printf("协程 %d 完成\n", counter.Get())
		}(i)
	}

	wg.Wait()
	fmt.Printf("所有协程完成:%d", counter.Get())

}
