package main

// 474. 一和零

func findMaxForm(strs []string, m int, n int) int {

	dp := make([][]int, m+1)

	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	for _, str := range strs {
		x, y := 0, 0
		for _, c := range str {
			if c == '0' {
				x++
			}
			if c == '1' {
				y++
			}

		}
		for i := m; i >= x; i-- {
			for j := n; j >= y; j-- {
				dp[i][j] = max(dp[i-x][j-y]+1, dp[i][j])
			}
		}
	}
	return dp[m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 1. dp含义，组成i个0 j个1 最大有几个物品（字符串）
// 2. 递推：dp[i][j] = max(dp[i-x][j-y] + 1 装这个物品 + 1 容量m-x,n-y , dp[i][j])
