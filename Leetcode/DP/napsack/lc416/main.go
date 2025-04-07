package main

// 416. 分割子和集 01背包

func canPartition(nums []int) bool {

	n := len(nums)

	sum := 0

	for i := 0; i < n; i++ {
		sum += nums[i]
	}

	if sum%2 != 0 {
		return false
	}

	target := sum / 2

	dp := make([]int, target+1)

	for i := 0; i < n; i++ {
		for j := target; j >= nums[i]; j-- {
			dp[j] = max(dp[j], dp[j-nums[i]]+nums[i])
		}
	}

	return dp[target] == target

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
