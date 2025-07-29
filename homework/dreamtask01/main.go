package main

import "fmt"

/*
*
有效的括号

考察：字符串处理、栈的使用

题目：给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效

链接：https://leetcode-cn.com/problems/valid-parentheses/
*/
func main() {

	fmt.Println(isValid1("()[]{}"))

}

func isValid1(str string) bool {
	if len(str) == 0 || len(str)%2 != 0 {
		return false
	}
	var m1 = make(map[string]string)
	m1[")"] = "("
	m1["]"] = "["
	m1["}"] = "{"
	stack := []string{}
	//fmt.Println(stack)
	for i := 0; i < len(str); i++ {
		//fmt.Println(string(str[i]))
		if m1[string(str[i])] != "" {
			if len(stack) == 0 {
				return false
			}
			if m1[string(str[i])] != stack[len(stack)-1] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, string(str[i]))
		}

	}
	return len(stack) == 0
}
