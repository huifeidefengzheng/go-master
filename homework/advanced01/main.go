package main

import "fmt"

/*
*
题目 ：编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
*/
func add(num *int) int {
	return *num + 10
}

func main() {
	// 定义一个变量
	var num int = 10
	fmt.Println(add(&num))

}
