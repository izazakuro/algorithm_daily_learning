package main

import "fmt"

func main() {

	var n, m int

	fmt.Scan(&n, &m)

	arr := make([][]int, n+1)
	for i := range arr {
		arr[i] = make([]int, m+1)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			fmt.Scan(&arr[i][j])
		}
	}

	for i := 1; i <= m; i++ {
		for j := 2; j <= n; j++ {
			arr[i][j] += arr[i][j-1]
		}
	}

	for j := 1; j <= m; j++ {
		for i := 2; i <= n; i++ {
			arr[i][j] += arr[i-1][j]
		}
	}

	var q int
	fmt.Scan(&q)
	var x1, y1, x2, y2 int
	for i := 0; i < q; i++ {
		fmt.Scan(&x1, &y1, &x2, &y2)
		fmt.Println(arr[x2][y2] + arr[x1-1][y1-1] - arr[x1-1][y2] - arr[x2][y1-1])
	}
}
