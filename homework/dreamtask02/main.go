package main

/*
*
 136. 只出现一次的数字：给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。可以使用 for 循环遍历数组，结合 if 条件判断和 map
    数据结构来解决，例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素。
*/
func main() {
	print(singleNumber([]int{4, 1, 2, 1, 2}))
}

func singleNumber(nums []int) int {
	var res int
	for _, num := range nums {
		//将数字num转成二进制
		//binaryStr := strconv.FormatInt(int64(num), 2)
		//fmt.Println(binaryStr)
		//binaryStrres := strconv.FormatInt(int64(res), 2)
		//fmt.Println(binaryStrres)

		res ^= num
		//fmt.Println(res)
	}
	return res
}

func singleNumber2(nums []int) int {
	var mm = make(map[int]int)
	for _, num := range nums {
		mm[num]++
	}
	var res int
	for k, v := range mm {
		if v == 1 {
			res = k
			break
		}
	}
	return res
}
