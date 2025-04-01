package main

import (
	"sort"
	"strconv"
)

// 最大数
func largestNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		a := strconv.Itoa(nums[i])
		b := strconv.Itoa(nums[j])
		return a+b > b+a
	})
	if nums[0] == 0 {
		return "0"
	}
	res := ""
	for _, x := range nums {
		res += strconv.Itoa(x)
	}
	return res
}

// 贪心策略：找出两个数拼接，按拼接结果进行重新排序，最终返回的数组就是最大数
// 将整型转换为字符串进行拼接比较，更为高效，且因为是按字典序比较大小
// 例如：a = 3 , b = 30 a+b = 330 , b+a = 303 303 < 330
