package main

import "fmt"

func main() {
	var n int
	var s string

	fmt.Scan(&n)
	fmt.Scan(&s)

	for i := 1; i <= n-1; i++ {
		result := 0
		l, r := 0, i
		for r < len(s) {
			if s[l] != s[r] {
				result++
				l++
				r++
			} else {
				break
			}
		}
		fmt.Println(result)
	}
}
