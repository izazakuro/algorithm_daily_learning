package main

import "fmt"

func main() {

	var n int

	fmt.Scan(&n)

	for i := 0; i <= n+1; i++ {
		for j := 0; j <= n+1; j++ {
			for k := 0; k <= n+1; k++ {
				if i+j+k <= n {
					fmt.Println(i, j, k)
				}
			}
		}
	}
}
