package main

// 376. 摆动序列

func wiggleMaxLength(nums []int) int {
	n := len(nums)
	pre := 0
	cur := 0
	result := 1

	for i := 0; i < n-1; i++ {
		cur = nums[i+1] - nums[i]
		if (pre >= 0 && cur < 0) || (pre <= 0 && cur > 0) {
			result++
			pre = cur
		}
	}

	return result
}

// 贪心策略：删除单调区间中间的元素
// 尾部因为是单调区终点，因此必定存在一个摆动
// 其中，出现平坡的情况也要考虑，在出现摆动时才进行pre和cur的更新
