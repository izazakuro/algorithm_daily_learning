package main

import "fmt"

func main() {
	var n, k, a int

	fmt.Scan(&n, &k, &a)

	result := (k + a - 1) % n

	if result == 0 {
		result = n
	}

	fmt.Println(result)
}
