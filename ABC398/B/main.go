package main

import "fmt"

func main() {

	m := make(map[int]int)

	var a int

	for i := 0; i < 7; i++ {
		fmt.Scan(&a)
		m[a]++
	}

	for i := 1; i <= 13; i++ {
		for j := 1; j <= 13; j++ {
			if i != j && m[i] >= 3 && m[j] >= 2 {
				fmt.Println("Yes")
				return
			}
		}
	}

	fmt.Println("No")

}
