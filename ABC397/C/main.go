package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	arr := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Scan(&arr[i])
	}

	front := make([]int, n+1)
	m1 := make(map[int]bool)
	count := 0

	for i := 1; i <= n; i++ {
		if !m1[arr[i]] {
			m1[arr[i]] = true
			count++
		}
		front[i] = count
	}

	back := make([]int, n+2)
	m2 := make(map[int]bool)
	count = 0

	for i := n; i >= 1; i-- {
		if !m2[arr[i]] {
			m2[arr[i]] = true
			count++
		}
		back[i] = count
	}

	result := 0
	for i := 1; i < n; i++ {
		result = max(result, front[i]+back[i+1])
	}

	fmt.Println(result)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
