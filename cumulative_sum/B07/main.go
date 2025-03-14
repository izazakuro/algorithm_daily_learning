package main

import "fmt"

func main() {

	var t, n int

	fmt.Scan(&t, &n)

	arr := make([]int, t+2)

	for i := 1; i <= n; i++ {
		var l, r int
		fmt.Scan(&l, &r)
		arr[l] += 1
		arr[r] -= 1
	}

	for i := 1; i <= t; i++ {
		arr[i] += arr[i-1]
	}

	for i := 0; i < t; i++ {
		fmt.Println(arr[i])
	}
}
