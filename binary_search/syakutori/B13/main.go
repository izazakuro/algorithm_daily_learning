package main

import "fmt"

func main() {

	var n, k int

	fmt.Scan(&n, &k)

	arr := make([]int, n+1)
	r := make([]int, n+1)

	for i := 1; i <= n; i++ {
		fmt.Scan(&arr[i])
	}

	for i := 2; i <= n; i++ {
		arr[i] += arr[i-1]
	}

	for i := 1; i <= n; i++ {
		if i == 1 {
			r[i] = 0
		} else {
			r[i] = r[i-1]
		}

		for r[i] < n && arr[r[i]+1]-arr[i-1] <= k {
			r[i] += 1
		}
	}

	result := 0
	for i := 1; i <= n; i++ {
		result += r[i] - i + 1
	}

	fmt.Println(result)
}
