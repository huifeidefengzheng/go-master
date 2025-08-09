package main

import (
	"fmt"
	"sync"
)

func main() {
	// 创建一个缓冲大小为10的通道
	ch := make(chan int, 10)
	var wg sync.WaitGroup

	// 添加生产者和消费者协程到等待组
	wg.Add(2)

	// 生产者协程：向通道发送100个整数
	go func() {
		defer wg.Done()
		for i := 1; i <= 100; i++ {
			ch <- i
			fmt.Printf("生产者发送: %d\n", i)
		}
		close(ch) // 发送完所有数据后关闭通道
		fmt.Println("生产者完成")
	}()

	// 消费者协程：从通道接收并打印整数
	go func() {
		defer wg.Done()
		for num := range ch {
			fmt.Printf("消费者接收: %d\n", num)
		}
		fmt.Println("消费者完成")
	}()

	// 等待所有协程完成
	wg.Wait()
	fmt.Println("程序执行完毕")
}
