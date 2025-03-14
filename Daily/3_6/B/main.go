package main

import "fmt"

func main() {

	s := "Takahashi"

	var n int

	fmt.Scan(&n)

	var s1 string
	result := 0

	for i := 0; i < n; i++ {
		fmt.Scan(&s1)
		if s1 == s {
			result++
		}
	}

	fmt.Println(result)

}
