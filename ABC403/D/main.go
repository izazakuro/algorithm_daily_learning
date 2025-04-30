package main

import (
	"fmt"
)

func main() {
	var n, d int
	fmt.Scan(&n, &d)

	arr := make([]int, n)
	count := make(map[int]int)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
		count[arr[i]]++
	}

	res := 0

	for _, v := range arr {
		if count[v] == 0 {
			continue
		}
		if count[v+d] > 0 {
			if count[v]+count[v+2*d] < count[v+d] {
				res += count[v]
				res += count[v+2*d]
				count[v+2*d] = 0
				count[v] = 0
			} else {
				res += count[v+d]
				count[v+d] = 0
			}
		}
	}

	fmt.Println(res)
}
