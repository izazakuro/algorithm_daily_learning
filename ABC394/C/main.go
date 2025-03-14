package main

import "fmt"

func dfs(s []rune, index int) []rune {

	if index != 0 && s[index] == 'A' && s[index-1] == 'W' {
		s[index] = 'C'
		s[index-1] = 'A'
		dfs(s, index-1)
	}
	return s

}

func main() {

	var s string

	fmt.Scan(&s)

	runes := []rune(s)

	index := 1

	for index < len(runes) {
		if runes[index] == 'A' {
			runes = dfs(runes, index)
		}
		index++
	}

	fmt.Println(string(runes))

}
