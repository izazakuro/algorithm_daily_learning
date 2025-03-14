package main

import "fmt"

func main() {

	arr := make([][]int, 9)

	a, b, c := 0, 0, 0

	for i := 0; i < 9; i++ {
		arr[i] = make([]int, 9)
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Scan(&arr[i][j])
		}
	}

	for i := 0; i < 9; i++ {
		a, b = 0, 0
		for j := 0; j < 9; j++ {
			a += arr[i][j]
			b += arr[j][i]
		}
		if a != 45 || b != 45 {
			fmt.Println("No")
			return
		}
	}

	for a := 0; a < 3; a++ {
		for b := 0; b < 3; b++ {
			c = 0
			for i := a * 3; i < (a+1)*3; i++ {
				for j := b * 3; j < (b+1)*3; j++ {
					c += arr[i][j]
				}
			}
			if c != 45 {

				fmt.Println("No")
				return

			}
		}
	}

	fmt.Println("Yes")

}
