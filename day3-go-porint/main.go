package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var p1 *int
	var p2 *string

	i := 1
	s := "Hello"
	// 基础类型数据，必须使用变量名获取指针，无法直接通过字面量获取指针
	// 因为字面量会在编译期被声明为成常量，不能获取到内存中的指针信息
	p1 = &i
	p2 = &s

	p3 := &p2
	fmt.Println(p1)
	fmt.Println(*p1)
	fmt.Println(p2)
	fmt.Println(p2)
	fmt.Println(**p3)

	var p4 *int
	i1 := 1
	p4 = &i1
	fmt.Println(*p4 == i1)
	*p4 = 2
	fmt.Println(i1)

	a := "Hello, world!"
	fmt.Println(&a)
	fmt.Println(unsafe.Pointer(&a))
	upA := uintptr(unsafe.Pointer(&a))
	upA += 1

	c := (*uint8)(unsafe.Pointer(upA))
	fmt.Println(*c)
}
