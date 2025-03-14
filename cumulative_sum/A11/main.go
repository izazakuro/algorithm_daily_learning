package main

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	var n int
	fmt.Scan(&n)

	arr := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Scan(&arr[i])
	}

	P := make([]int, n+1)
	P[1] = arr[1]
	for i := 2; i <= n; i++ {
		P[i] = max(P[i-1], arr[i])
	}

	Q := make([]int, n+2)
	Q[n] = arr[n]
	for i := n - 1; i >= 1; i-- {
		Q[i] = max(Q[i+1], arr[i])
	}

	var d int
	fmt.Scan(&d)

	for i := 0; i < d; i++ {
		var l, r int
		fmt.Scan(&l, &r)
		fmt.Println(max(P[l-1], Q[r+1]))
	}
}
