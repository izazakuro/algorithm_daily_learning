package main

import "fmt"

func main() {
	var s string

	fmt.Scan(&s)

	result := 0

	for i := 0; i < len(s); i++ {
		if s[i] == 'i' && i != len(s)-1 && s[i+1] != 'o' {
			result++
		}
		if i > 0 && s[i] == 'o' && s[i-1] != 'i' {
			result++
		}
		if i == len(s)-1 && s[len(s)-1] == 'i' {
			result++
		}
		if i == 0 && s[i] == 'o' {
			result++
		}
	}
	fmt.Println(result)
}
