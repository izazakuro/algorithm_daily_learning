package main

import "fmt"

func main() {

	var n int
	fmt.Scan(&n)

	p := make([]int, n+1)
	q := make([]int, n+1)

	for i := 1; i <= n; i++ {
		fmt.Scan(&p[i])
	}
	for i := 1; i <= n; i++ {
		fmt.Scan(&q[i])
	}

	m := make([]int, n+1)

	for i := 1; i <= n; i++ {
		m[q[i]] = q[p[i]]
	}

	for i := 1; i <= n; i++ {
		fmt.Printf("%d ", m[i])
	}
}
