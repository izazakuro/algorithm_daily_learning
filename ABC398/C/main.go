package main

import "fmt"

func main() {

	var n int

	fmt.Scan(&n)

	arr := make([]int32, n)
	m := make(map[int32]int)

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
		m[arr[i]]++
	}

	var temp int32
	temp = -1

	result := 0

	for i := 0; i < n; i++ {
		if m[arr[i]] == 1 {
			if arr[i] > temp {
				result = i + 1
				temp = arr[i]
			}
		}
	}

	if temp == -1 {
		result = -1
	}
	fmt.Println(result)

}
