package main

import "fmt"

/*
*
1. 两数之和
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
你可以按任意顺序返回答案。
示例 1：
输入：nums = [2,7,11,15], target = 9
输出：[0,1]
*/
func ts(nums []int, target int) []int {
	var res []int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				res = append(res, i, j)
				return res
			}
		}
	}
	return res
}

func ts2(nums []int, target int) []int {
	hastTable := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if v, ok := hastTable[target-nums[i]]; ok {
			return []int{v, i}
		}
		hastTable[nums[i]] = i
	}
	return nil
}

func main() {
	fmt.Println(ts([]int{2, 7, 11, 15}, 9))
	fmt.Println(ts2([]int{2, 7, 11, 15}, 9))
	println("hello world")
}
