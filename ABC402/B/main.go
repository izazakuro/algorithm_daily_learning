package main

import "fmt"

func main() {

	var q int

	fmt.Scan(&q)
	queue := make([]int, 0)

	for i := 0; i < q; i++ {
		var n, x int
		fmt.Scan(&n)
		switch n {
		case 1:
			fmt.Scan(&x)
			queue = append(queue, x)
		case 2:
			if len(queue) == 0 {
				continue
			}
			fmt.Println(queue[0])
			queue = queue[1:]
		}
	}

	return

}
