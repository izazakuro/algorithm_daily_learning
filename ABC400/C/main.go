package main

import (
	"fmt"
)

func main() {
	var n int64
	fmt.Scan(&n)

	const inf int64 = 1e9
	var l, r, m int64
	var ans int64 = 0

	// First binary search: 2 * b^2 <= n
	l, r = 0, inf
	for l+1 < r {
		m = (l + r) / 2
		if m*m*2 <= n {
			l = m
		} else {
			r = m
		}
	}
	ans += l

	// Second binary search: 4 * b^2 <= n
	l, r = 0, inf
	for l+1 < r {
		m = (l + r) / 2
		if m*m*4 <= n {
			l = m
		} else {
			r = m
		}
	}
	ans += l

	fmt.Println(ans)
}
