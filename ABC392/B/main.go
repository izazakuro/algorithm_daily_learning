package main

import "fmt"

func main() {

	var n, m int

	fmt.Scan(&n, &m)

	h := make(map[int]bool, n+1)

	for i := 0; i < m; i++ {
		var a int

		fmt.Scan(&a)

		h[a] = true
	}

	result := 0

	for i := 1; i <= n; i++ {
		if h[i] == false {
			result++
		}
	}

	fmt.Println(result)

	for i := 1; i <= n; i++ {
		if !h[i] {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()

}
