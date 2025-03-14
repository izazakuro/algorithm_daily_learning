package main

import "fmt"

func main() {

	var n int

	fmt.Scan(&n)

	arr := make([][]int, n+1)

	for i := 0; i <= n; i++ {
		arr[i] = make([]int, n+1)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Scan(&arr[i][j])
		}
	}
	element := 1
	for i := 1; i <= n; i++ {
		if element >= i {
			element = arr[element][i]
		} else {
			element = arr[i][element]
		}
	}

	fmt.Println(element)

}
