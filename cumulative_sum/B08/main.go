package main

import "fmt"

var arr = [1501][1501]int{}

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var x, y int
		fmt.Scan(&x, &y)
		arr[x][y]++
	}

	for i := 1; i < len(arr); i++ {
		for j := 1; j < len(arr[i]); j++ {
			arr[i][j] += arr[i][j-1]
		}
	}

	for j := 1; j < len(arr[0]); j++ {
		for i := 2; i < len(arr); i++ {
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
