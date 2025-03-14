package main

import "fmt"

func main() {

	var n, k int

	fmt.Scan(&n, &k)

	A := make([]int, n+1)

	for i := 1; i <= n; i++ {
		fmt.Scan(A[i])
	}

	B := make([]int, 2*n+1)

}
