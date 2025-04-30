package main

import "sort"

// 35. 查找插入位置

func searchInsert(nums []int, target int) int {
	pos := sort.Search(len(nums), func(index int) bool {
		return target <= nums[index]
	})

	return pos
}
