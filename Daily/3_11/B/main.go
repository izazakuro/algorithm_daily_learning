package main

import "fmt"

func main() {

	var s string

	fmt.Scan(&s)

	var result string

	for i := 0; i < len(s); i++ {
		if s[i] != 'a' && s[i] != 'i' && s[i] != 'u' && s[i] != 'e' && s[i] != 'o' {
			result += string(s[i])
		}
	}

	fmt.Println(result)
}
