package main

import "fmt"

func main() {

	var n int

	fmt.Scan(&n)

	res := 400 / n

	if 400%n != 0 {
		fmt.Println(-1)
	} else {
		fmt.Println(res)
	}

	return
}
