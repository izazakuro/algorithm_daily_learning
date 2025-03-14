package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {

	var s, t string

	fmt.Scan(&s, &t)

	n := len(s)
	m := len(t)

	dp := make([][]int, n+1)

	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, m+1)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	fmt.Println(dp[n][m])

}
