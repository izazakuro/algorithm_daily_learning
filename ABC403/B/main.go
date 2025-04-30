package main

import "fmt"

func main() {
	var t, u string
	fmt.Scan(&t, &u)

	for i := 0; i <= len(t)-len(u); i++ {
		match := true
		for j := 0; j < len(u); j++ {
			if t[i+j] != u[j] && t[i+j] != '?' {
				match = false
				break
			}
		}
		if match {
			fmt.Println("Yes")
			return
		}
	}

	fmt.Println("No")
}
