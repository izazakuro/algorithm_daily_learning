package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	arr := make([]int, n+1)

	for i := 1; i <= n; i++ {
		fmt.Scan(&arr[i])
	}

	res := 0

	for i := 1; i <= n; i += 2 {
		res += arr[i]
	}

	fmt.Println(res)

	return
}
