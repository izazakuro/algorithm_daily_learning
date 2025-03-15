package main

import (
	"fmt"
	"sort"
)

func main() {

	var n int

	fmt.Scan(&n)

	arr := make([]int, n)
	arr2 := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
		arr2[i] = arr[i]
	}

	sort.Ints(arr)

	for i := 0; i < n; i++ {
		if arr2[i] == arr[n-2] {
			fmt.Println(i + 1)
			return
		}
	}
}
