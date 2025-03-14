package main

import "fmt"

func main() {
	var s string

	fmt.Scan(&s)

	if len(s)%2 != 0 {
		fmt.Println("No")
		return
	}

	l, r := 0, 1
	for r < len(s) {
		if (s[l] == '(' && s[r] == ')') ||
			(s[l] == '[' && s[r] == ']') ||
			(s[l] == '<' && s[r] == '>') {
			fmt.Println("Yes")
			return
		}
		r++
		l++
	}
	fmt.Println("No")
}
