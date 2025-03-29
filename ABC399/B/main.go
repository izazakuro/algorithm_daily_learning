package main

import (
	"fmt"
	"sort"
)

type p struct {
	num   int
	index int
	rank  int
}

func main() {

	var n int

	fmt.Scan(&n)

	arr := make([]p, n)

	var num int

	for i := 0; i < n; i++ {
		fmt.Scan(&num)
		arr[i].num = num
		arr[i].index = i + 1
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i].num > arr[j].num
	})

	x := arr[0].num
	r := 1

	for i := 0; i < n; i++ {
		if arr[i].num != x {
			r = i + 1
			x = arr[i].num
		}
		arr[i].rank = r
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i].index < arr[j].index
	})

	for i := 0; i < n; i++ {
		fmt.Println(arr[i].rank)
	}

}
