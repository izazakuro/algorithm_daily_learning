package main

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)

	dp[0] = 0

	for i := 1; i <= amount; i++ {
		dp[i] = 2 << 31
	}

	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			dp[j] = min(dp[j-coins[i]]+1, dp[j])
		}
	}

	if dp[amount] == 2<<31 {
		return -1
	}

	return dp[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
