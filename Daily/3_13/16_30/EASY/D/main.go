package main

import (
	"fmt"
)

func main() {
	var m int

	fmt.Scan(&m)

	arr := make([]int, m+1)

	sum := 0

	for i := 1; i <= m; i++ {
		fmt.Scan(&arr[i])
		sum += arr[i]
	}

	num := sum/2 + 1

	index := 1
	a, b := 1, 1

	for num > 0 {
		if num-arr[index] > 0 {
			num -= arr[index]
		} else if num-arr[index] == 0 {
			a = index
			b = arr[index]
			break
		} else if num-arr[index] < 0 {
			a = index
			b = num
			break
		}
		index++
	}

	fmt.Printf("%d %d", a, b)

}
