package main

import "fmt"

func main() {

	var n, m int

	fmt.Scan(&n, &m)

	count := 0

	var h int

	for i := 0; i < n; i++ {
		fmt.Scan(&h)
		m -= h
		if m < 0 {
			fmt.Println(count)
			return
		} else if m == 0 {
			fmt.Println(count + 1)
			return
		} else {
			count++
		}
	}
	fmt.Println(count)

}
