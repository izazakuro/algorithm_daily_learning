package main

import "fmt"

func main() {
	var n int

	fmt.Scan(&n)

	arr := make([]string, n+1)

	for i := 1; i <= n; i++ {
		fmt.Scan(&arr[i])
	}

	result := ""

	for i := 1; i <= 50; i++ {
		for j := 1; j <= n; j++ {
			if len(arr[j]) == i {
				result += arr[j]
			}
		}
	}

	fmt.Println(result)
}
