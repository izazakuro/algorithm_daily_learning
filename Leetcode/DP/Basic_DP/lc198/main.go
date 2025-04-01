package main

// 打家劫舍
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

// dp数组含义：前i个房间能获取的最大价值
// 递推公式：取决于max(dp[i-1], dp[i-2]+cost[i]) 偷这个房间和不偷这个房间
// 初始化：第0个房间能获取的最大价值就是cost[0] 取决于前两个房间的状态所以得初始化dp[0],[1]
// dp[1] = 前两个房间中的最大价值
