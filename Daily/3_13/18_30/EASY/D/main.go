package main

import (
	"fmt"
)

func main() {
	var n int

	fmt.Scan(&n)

	m := make(map[int]int, n+1)

	var num int

	for i := 0; i < 4*n-1; i++ {
		fmt.Scan(&num)
		m[num]++
	}

	for i := 1; i <= n; i++ {
		if m[i] != 4 {
			fmt.Println(i)
			return
		}
	}
}
