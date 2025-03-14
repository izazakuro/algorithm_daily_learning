package main

import "fmt"

func main() {

	var n, q int

	fmt.Scan(&n, &q)

	arr := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Scan(&arr[i])
	}

	for i := 2; i <= n; i++ {
		arr[i] += arr[i-1]
	}

	var l, r int

	for i := 0; i < q; i++ {
		fmt.Scan(&l, &r)
		fmt.Println(arr[r] - arr[l-1])
	}

}
