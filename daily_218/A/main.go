package main

import "fmt"

func main() {

	var n int64

	fmt.Scan(&n)

	if 2 <= n && n <= 4 {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}

}
