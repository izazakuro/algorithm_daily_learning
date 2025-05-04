package main

import "fmt"

func main() {

	var s string
	m := make(map[byte]bool)

	fmt.Scan(&s)

	for i := 0; i < len(s); i++ {
		c := s[i]
		m[c] = true
	}

	for i := 'a'; i <= 'z'; i++ {
		if !m[byte(i)] {
			var c byte
			c = byte(i)
			fmt.Printf("%c\n", c)
			return
		}
	}
}
