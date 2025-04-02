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
