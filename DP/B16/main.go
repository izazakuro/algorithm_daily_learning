package main

import "fmt"

func absInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	var N int
	fmt.Scan(&N)

	arr := make([]int, N+1)
	dp := make([]int, N+1)

	for i := 1; i <= N; i++ {
		fmt.Scan(&arr[i])
	}
	dp[1] = 0
	dp[2] = absInt(arr[1] - arr[2])

	for i := 3; i <= N; i++ {
		dp[i] = min(dp[i-1]+absInt(arr[i-1]-arr[i]), dp[i-2]+absInt(arr[i-2]-arr[i]))
	}

	fmt.Println(dp[N])

}
