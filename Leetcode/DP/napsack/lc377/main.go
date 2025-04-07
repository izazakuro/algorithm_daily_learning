package main

// 377. 组合总和4

func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)

	dp[0] = 1

	for j := 1; j <= target; j++ {
		for i := 0; i < len(nums); i++ {
			if nums[i] <= j {
				dp[j] += dp[j-nums[i]]
			}
		}
	}

	return dp[target]
}
