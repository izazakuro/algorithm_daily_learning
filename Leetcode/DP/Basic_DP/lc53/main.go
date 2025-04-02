package main

// 53. 最大子序和

func maxSubArray(nums []int) int {
	n := len(nums)
	dp := make([]int, n)

	dp[0] = nums[0]

	res := dp[0]

	for i := 1; i < n; i++ {
		dp[i] = max(nums[i]+dp[i-1], nums[i])
		if res < dp[i] {
			res = dp[i]
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 1. DP数组含义：下标为i时的最大和
// 2. 递推：max(前i-1的最大和+nums[i], 重新从num[i]开始计数）
