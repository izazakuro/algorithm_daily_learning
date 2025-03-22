package main

import "fmt"

func main() {

	var s string

	fmt.Scan(&s)

	result := 0

	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			for k := i + 2; k < len(s); k++ {
				if j-i == k-j && s[i] == 'A' && s[j] == 'B' && s[k] == 'C' {
					result++
				}
			}
		}
	}

	fmt.Println(result)

}
