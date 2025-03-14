package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	var n, m int

	fmt.Scan(&n, &m)

	w := make([]int, n)
	v := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&w[i], &v[i])
	}

	dp := make([][]int, n)

	for i := 0; i < n; i++ {
		dp[i] = make([]int, m+1)
	}
	for i := 0; i < n; i++ {
		dp[i][0] = 0
	}

	for i := 1; i < m+1; i++ {
		if w[0] <= i {
			dp[0][i] = v[0]
		}
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m+1; j++ {
			if j < w[i] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-w[i]]+v[i])
			}
		}
	}
	answer := 0
	for i := 0; i < m+1; i++ {
		answer = max(answer, dp[n-1][i])
	}

	fmt.Println(answer)
}
