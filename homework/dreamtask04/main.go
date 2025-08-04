package main

import "fmt"

func main() {
	fmt.Println(longestCommonPrefix1([]string{"flower", "flow", "flight"}))
}

/*
*
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

示例 1：

输入：strs = ["flower","flow","flight"]
输出："fl"
示例 2：

输入：strs = ["dog","racecar","car"]
输出：""
解释：输入不存在公共前缀。

*/

func longestCommonPrefix(str []string) string {
	if len(str) == 0 {
		return ""
	}
	prefix := str[0]
	for i := 1; i < len(str); i++ {
		for j := 0; j < len(prefix); j++ {
			if j >= len(str[i]) || str[i][j] != prefix[j] {
				prefix = prefix[0:j]
				break
			}
		}
	}
	return prefix
}

func longestCommonPrefix1(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		prefix = commonPrefix(prefix, strs[i])
	}
	return prefix
}

func commonPrefix(prefix string, s string) string {
	if len(prefix) > len(s) {
		prefix = prefix[0:len(s)]
	}
	tem := prefix
	for i := 0; i < len(prefix); i++ {
		if i >= len(s) || prefix[i] != s[i] {
			tem = prefix[0:i]
			break
		}
	}
	return tem

}
