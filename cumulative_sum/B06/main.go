package main

import "fmt"

func main() {

	var n int

	fmt.Scan(&n)

	arr := make([]int, n+1)

	for i := 1; i <= n; i++ {
		fmt.Scan(&arr[i])
	}

	lose := make([]int, n+1)
	win := make([]int, n+1)

	if arr[1] == 0 {
		lose[1] = 1
	} else {
		win[1] = 1
	}

	for i := 2; i <= n; i++ {
		if arr[i] == 0 {
			lose[i] += lose[i-1] + 1
			win[i] += win[i-1]
		} else {
			win[i] += win[i-1] + 1
			lose[i] += lose[i-1]
		}
	}

	var q int

	fmt.Scan(&q)

	for i := 0; i < q; i++ {
		var l, r int
		fmt.Scan(&l, &r)
		if win[r]-win[l-1] > lose[r]-lose[l-1] {
			fmt.Println("win")
		} else if win[r]-win[l-1] == lose[r]-lose[l-1] {
			fmt.Println("draw")
		} else {
			fmt.Println("lose")
		}

	}

}
