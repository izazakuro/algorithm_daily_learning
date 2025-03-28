package main

func rob(nums []int) int {
	n := len(nums)

	if len(nums) == 1 {
		return nums[0]
	}

	dp := make([]int, n)

	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])

	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}

	return dp[n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
