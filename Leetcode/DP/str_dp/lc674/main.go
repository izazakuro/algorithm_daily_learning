package main

import "fmt"

// 674. 最长连续递增序列

// 双指针解法
func findLengthOfLCIS(nums []int) int {
	n := len(nums)

	if n == 1 {
		return 1
	}

	l, r := 0, 1

	res := 0

	for r < n {
		if nums[r] <= nums[r-1] {
			res = max(res, r-l)
			l = r
			r++
		} else {
			res = max(res, r-l+1)
			r++
		}
	}

	return res
}

// dp解法

func findLengthOfLCIS_dp(nums []int) int {
	n := len(nums)

	dp := make([]int, n)

	dp[0] = 1

	res := 1

	for i := 1; i < n; i++ {
		dp[i] = 1
		if nums[i] > nums[i-1] {
			dp[i] += dp[i-1]
		}
		res = max(dp[i], res)
	}

	fmt.Println(dp)

	return res

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
