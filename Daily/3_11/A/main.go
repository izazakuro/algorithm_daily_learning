package main

import (
	"fmt"
)

func main() {

	var n int
	fmt.Scan(&n)

	arr := make([]int, n+1)

	for i := 1; i <= n; i++ {
		fmt.Scan(&arr[i])
	}

	result := -1

	temp := arr[1]

	for i := 2; i <= n; i++ {
		if arr[i] > temp {
			result = i
			temp = arr[i]
			break
		}
	}
	fmt.Println(result)

}
