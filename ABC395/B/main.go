package main

import (
	"fmt"
)

func main() {

	var n, q int

	fmt.Scan(&n, &q)

	m := make(map[int]int, n)

	for i := 1; i <= n; i++ {
		m[i] = i
	}

	var op, a, b int

	for i := 0; i < q; i++ {
		fmt.Scan(&op)
		switch op {
		case 1:
			fmt.Scan(&a, &b)
			m[a] = b
			continue
		case 2:
			fmt.Scan(&a, &b)
			for j := 0; j < len(m); j++ {
				if m[j] == a {
					m[j] = b
				} else if m[j] == b {
					m[j] = a
				}
			}
			continue
		case 3:
			fmt.Scan(&a)
			fmt.Println(m[a])
		}
	}

}
