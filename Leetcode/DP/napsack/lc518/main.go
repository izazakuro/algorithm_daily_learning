package main

// 518. 兑换零钱

func change(amount int, coins []int) int {
	dp := make([]int, amount+1)

	dp[0] = 1

	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			dp[j] += dp[j-coins[i]]
		}
	}
	return dp[amount]
}

// 1. dp含义：0～第i枚钱币的种类，可凑成j的总数
// 2. 递推公式：dp[j] += dp[j-coins[i]]
// 3. 遍历顺序：先物品再背包是组合数， 先背包再物品是排列数