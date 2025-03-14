package main

import "fmt"

func main() {

	var d, n int

	fmt.Scan(&d, &n)

	arr := make([]int, d+2)

	for i := 1; i <= n; i++ {
		var l, r int
		fmt.Scan(&l, &r)
		arr[l] += 1
		arr[r+1] -= 1
	}

	for i := 2; i <= d; i++ {
		arr[i] += arr[i-1]
	}

	for i := 1; i <= d; i++ {
		fmt.Println(arr[i])
	}

}
