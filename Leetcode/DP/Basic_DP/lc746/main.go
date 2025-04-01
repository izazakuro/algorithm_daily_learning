package main

// 爬楼梯最小花费
func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 0
	for i := 2; i <= n; i++ {
		dp[i] = min(cost[i-1]+dp[i-1], cost[i-2]+dp[i-2])
	}

	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// dp[i] = min(dp[i-1] + cost[i-1] , dp[i-2] + cost[i-2]
// dp数组含义：到达台阶i所需的最小cost
// 初始化从0，1开始，不跳即不消耗体力，所以都初始化为0
