package main

import (
	"fmt"
)

type node struct {
	x int
	y int
}

func main() {

	var n, r, c int
	var s string

	fmt.Scan(&n, &r, &c)
	fmt.Scan(&s)

	result := ""

	for i := 0; i < n; i++ {
		switch s[i] {
		case 'N':
			r += 1
		case 'W':
			c += 1
		case 'S':
			r -= 1
		case 'E':
			c -= 1
		}
		if r == 0 && c == 0 {
			result = result + "1"
		} else {
			result = result + "0"
		}
	}

	fmt.Println(result)

}
