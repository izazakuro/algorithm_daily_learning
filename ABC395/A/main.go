package main

import "fmt"

func main() {

	var n int
	fmt.Scan(&n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	lastPos := make(map[int]int)
	result := -1
	for i := 0; i < n; i++ {
		val := arr[i]
		if pos, ok := lastPos[val]; ok {
			length := i - pos + 1
			if result == -1 || length < result {
				result = length
			}
		}
		lastPos[val] = i
	}
	fmt.Println(result)
}
