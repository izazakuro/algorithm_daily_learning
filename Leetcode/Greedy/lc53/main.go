package main

// 53. 最大子序和

func maxSubArray(nums []int) int {
	n := len(nums)
	res := -10001

	cnt := 0

	for i := 0; i < n; i++ {
		cnt += nums[i]
		if cnt > res {
			res = cnt
		}
		if cnt < 0 {
			cnt = 0
		}
	}

	return res
}

// 贪心策略： 和变为负数后从下一个值开始计算
