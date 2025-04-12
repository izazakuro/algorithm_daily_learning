package main

import (
	"fmt"
)

const mod = int64(1e9)

func main() {

	var n, k int
	fmt.Scan(&n, &k)

	dp := make([]int64, n+1)

	if n < k {
		fmt.Println(1)
		return
	}

	r, l := 0, 0

	var sum int64
	sum = 0
	for r <= n {
		if r < k {
			dp[r] = 1
			sum = (sum + dp[r]) % mod
			r++
		} else {
			for r-l != k {
				sum = (sum - dp[l] + mod) % mod
				l++
			}
			dp[r] = sum
			sum = (sum + dp[r]) % mod
			r++
		}
	}

	fmt.Println(dp[n] % mod)

	return
}
