package main

import "fmt"

func main() {

	var n int

	fmt.Scan(&n)

	if n >= 200 && n <= 299 {
		fmt.Println("Success")
	} else {
		fmt.Println("Failure")
	}

	return
}
