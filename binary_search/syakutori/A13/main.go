package main

import (
	"fmt"
)

func main() {

	var n, k int

	fmt.Scan(&n, &k)

	arr := make([]int, n+1)
	r := make([]int, n+1)

	for i := 1; i <= n; i++ {
		fmt.Scan(&arr[i])
	}

	for i := 1; i <= n-1; i++ {
		if i == 1 {
			r[i] = 1
		} else {
			r[i] = r[i-1]
		}

		for r[i] < n && arr[r[i]+1]-arr[i] <= k {
			r[i]++
		}
	}

	result := 0
	for i := 0; i <= n-1; i++ {
		result += r[i] - i
	}

	fmt.Println(result)

}
