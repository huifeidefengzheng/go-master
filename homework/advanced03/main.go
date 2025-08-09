package main

import (
	"fmt"
	"sync"
)

// 编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
func main() {
	var wg sync.WaitGroup

	// 添加两个协程到等待组
	wg.Add(2)

	// 启动协程打印奇数
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i += 2 {
			fmt.Printf("奇数: %d\n", i)
		}
	}()

	// 启动协程打印偶数
	go func() {
		defer wg.Done()
		for i := 2; i <= 10; i += 2 {
			fmt.Printf("偶数: %d\n", i)
		}
	}()

	// 等待所有协程完成
	wg.Wait()
	fmt.Println("所有数字打印完成")
}
