package main

import "fmt"

func main() {

	var T, N int

	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		fmt.Scan(&N)
		result := 0
		var A int32
		for i := 0; i < N; i++ {
			fmt.Scan(&A)
			if A%2 != 0 {
				result++
			}
		}
		fmt.Println(result)
	}
}
