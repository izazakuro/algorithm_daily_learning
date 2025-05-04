package main

import "fmt"

func main() {

	var n int

	fmt.Scan(&n)

	arr := make([][]byte, n)
	arr2 := make([][]byte, n)

	for i := 0; i < n; i++ {
		arr[i] = make([]byte, n)
		arr2[i] = make([]byte, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Scan(&arr[i][j])
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Scan(&arr2[i][j])
		}
	}

	for i := 0 ; i < 3 ; i ++ {
		
	}


}

func rotate90Clockwise(matrix [][]byte) {
	n := len(matrix)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			matrix[i][j], matrix[i][n-1-j] = matrix[i][n-1-j], matrix[i][j]
		}
	}
}
