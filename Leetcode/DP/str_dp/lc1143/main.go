package main

// 1143. 最长公共子序列

func longestCommonSubsequence(text1 string, text2 string) int {
	n := len(text1)
	m := len(text2)

	dp := make([][]int, n+1)

	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
	}

	res := 0

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
			res = max(res, dp[i][j])
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

// dp含义：text1从0～i-1 与 text2从0～j-1中存在的公共子序列的最大长度
// 递推： 可以从三个方向得出dp[i][j]，如果text1[i] == text2[j] 
// 则等于没有这两个字符的最大长度+1: dp[i][j] = dp[i-1][j-1] + 1
// 另外情况下等于左或右的最大长度 else : dp[i][j] = max(dp[i-1][j], dp[i][j-1])