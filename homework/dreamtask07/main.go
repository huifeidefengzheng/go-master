package main

import (
	"fmt"
	"sort"
)

/*
*
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。
示例 1：

输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
输出：[[1,6],[8,10],[15,18]]
解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
示例 2：

输入：intervals = [[1,4],[4,5]]
输出：[[1,5]]
解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。
*/
func mer(inte [][]int) [][]int {
	if len(inte) == 0 {
		return [][]int{}
	}
	if len(inte) == 1 {
		return inte
	}
	// 对二维数组进行排序 按照二维数组的每个元素的起始位置进行排序
	sort.Slice(inte, func(i, j int) bool {
		return inte[i][0] < inte[j][0]
	})
	fmt.Println(inte)
	// 创建一个二维切片
	res := make([][]int, 0)
	for i := 0; i < len(inte); i++ {

		if len(res) != 0 && res[len(res)-1][1] >= inte[i][0] {
			res[len(res)-1][1] = max(res[len(res)-1][1], inte[i][1])
		} else {
			// 将aa[0]和inte[i][1]添加到res内
			res = append(res, inte[i])
		}
	}
	return res
}
func main() {
	inte := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	//inte := [][]int{{1, 4}, {4, 5}}
	//inte := [][]int{{1, 4}, {2, 3}}
	//inte := [][]int{{1, 4}, {0, 2}, {3, 5}}
	fmt.Println(mer(inte))
}
