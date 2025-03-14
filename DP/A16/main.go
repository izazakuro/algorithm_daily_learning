package main

import "fmt"

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	var N int
	fmt.Scan(&N)

	A := make([]int, N+1)
	B := make([]int, N+1)
	dp := make([]int, N+1)

	for i := 2; i <= N; i++ {
		fmt.Scan(&A[i])
	}

	for i := 3; i <= N; i++ {
		fmt.Scan(&B[i])
	}

	dp[1] = 0
	dp[2] = A[2]

	for i := 3; i <= N; i++ {
		dp[i] = Min(dp[i-1]+A[i], dp[i-2]+B[i])
	}

	fmt.Println(dp[N])
}
