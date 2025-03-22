package main

import "fmt"

func main() {

	var n int

	fmt.Scan(&n)

	if n%2 == 0 {
		for i := 0; i < n/2-1; i++ {
			fmt.Printf("-")
		}
		fmt.Printf("==")
		for i := 0; i < n/2-1; i++ {
			fmt.Printf("-")

		}
	} else {
		for i := 0; i < n/2; i++ {
			fmt.Printf("-")
		}
		fmt.Printf("=")
		for i := 0; i < n/2; i++ {
			fmt.Printf("-")

		}
	}

}
