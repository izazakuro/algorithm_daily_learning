package main

import (
	"fmt"
	"sort"
)

func lower_bound(arr []int, x int) int {
	index := sort.Search(len(arr), func(i int) bool {
		return arr[i] >= x
	})
	return index
}

func main() {
	var n int

	fmt.Scan(&n)

	arr := make([]int, n+1)

	for i := 1; i <= n; i++ {
		fmt.Scan(&arr[i])
	}

	var q int

	fmt.Scan(&q)

	sort.Ints(arr)

	for i := 0; i < q; i++ {
		var x int
		fmt.Scan(&x)
		fmt.Println(lower_bound(arr, x) - 1)
	}

}
