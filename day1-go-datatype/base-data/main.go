package main

import "fmt"

// 十六进制
var a uint8 = 0xF
var b uint8 = 0xf

// 八进制
var c uint8 = 017
var d uint8 = 0o17
var e uint8 = 0o17

// 二进制
var f uint8 = 0b1111
var g uint8 = 0b1111

// 十进制
var h uint8 = 15

var c1 complex64 = 1.10 + 0.1i
var c2 complex64 = 1.10 + 0.1i
var c3 complex64 = complex(1.10, 0.1) // c2与c3是等价的

var x float32 // 使用 var 声明并指定类型
var y float32

var s string = "Hello, world!"
var bytes []byte = []byte(s)

var r1 rune = 'a'
var r2 rune = '世'

var ss string = "abc，你好，世界！"
var runes []rune = []rune(ss)

var str1 string = "Hello\nworld!\n"
var str2 string = `Hello
world!
`

var sg string = "Go语言"
var bytess []byte = []byte(sg)
var runess []rune = []rune(sg)

func main() {
	x = real(c2)
	y = imag(c2)
	fmt.Println("c1==c2: ", c1 == c2)
	fmt.Println("c2==c3: ", c2 == c3)
	fmt.Println("real part of c2:", x)
	fmt.Println("imaginary part of c2:", y)

	fmt.Println("convert \"Hello, world!\" to bytes: ", bytes)
	var s string = string(bytes)
	fmt.Println(s)
	fmt.Println("rune:", runes)
	fmt.Println("str1==str2:", str1 == str2)
	fmt.Println("string sub: ", s[0:7])
	fmt.Println("bytes sub: ", string(bytes[0:7]))
	fmt.Println("runes sub: ", string(runes[0:3]))

	fmt.Println("string length: ", len(sg))
	fmt.Println("bytes length: ", len(bytess))
	fmt.Println("runes length: ", len(runess))
	fmt.Println("=================: ")
	fmt.Println("string sub: ", sg[0:7])
	fmt.Println("bytes sub: ", string(bytess[0:7]))
	fmt.Println("runes sub: ", string(runess[0:3]))

}
