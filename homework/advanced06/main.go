package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	var counter int64
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				atomic.AddInt64(&counter, 1)
				fmt.Printf("协程 %d 完成\n", atomic.LoadInt64(&counter))
			}
		}(i)
	}
	wg.Wait()
	fmt.Printf("所有协程完成:%d", counter)
}
