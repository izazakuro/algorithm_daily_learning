package main

import "fmt"

func main() {
	var N, S int

	fmt.Scan(&N, &S)

	arr := make([]int, N+1)

	for i := 1; i <= N; i++ {
		fmt.Scan(&arr[i])
	}

	dp := make([][]bool, S+1)

	for i := 0; i <= S; i++ {
		dp[i] = make([]bool, S+1)
	}

	dp[0][0] = true

	for i := 1; i <= S; i++ {
		dp[0][i] = false
	}

	for i := 1; i <= N; i++ {
		for j := 0; j <= S; j++ {
			if j < arr[i] {
				if dp[i-1][j] == true {
					dp[i][j] = true
				} else {
					dp[i][j] = false
				}
			} else if j >= arr[i] {
				if dp[i-1][j] == true || dp[i-1][j-arr[i]] == true {
					dp[i][j] = true
				} else {
					dp[i][j] = false
				}
			}

		}
	}

	if dp[N][S] {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}
