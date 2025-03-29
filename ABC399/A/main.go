package main

import "fmt"

func main() {
	var n int
	var s, t string

	fmt.Scan(&n)

	fmt.Scan(&s, &t)

	result := 0

	for i := 0; i < n; i++ {
		if s[i] != t[i] {
			result++
		}
	}

	fmt.Println(result)

}
