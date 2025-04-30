package main

import (
	"fmt"
)

func main() {
	var n, m, q int

	fmt.Scan(&n, &m, &q)

	user := make([]map[int]bool, n+1)

	for i := 1; i <= n; i++ {
		user[i] = make(map[int]bool)
	}

	for i := 0; i < q; i++ {
		var query, x, y int
		fmt.Scan(&query, &x)

		switch query {
		case 1:
			fmt.Scan(&y)
			user[x][y] = true
		case 2:
			user[x][200001] = true
		case 3:
			fmt.Scan(&y)
			if user[x][y] == true || user[x][200001] == true {
				fmt.Println("Yes")
			} else {
				fmt.Println("No")
			}
		}

	}

	return

}
