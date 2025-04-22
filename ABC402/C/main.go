package main

import (
	"fmt"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	hash := make([]map[int]bool, m)

	reverse := make(map[int][]int)

	cleared := make([]bool, m)

	for i := 0; i < m; i++ {
		var k int
		fmt.Scan(&k)
		hash[i] = make(map[int]bool)
		for j := 0; j < k; j++ {
			var num int
			fmt.Scan(&num)
			hash[i][num] = true
			reverse[num] = append(reverse[num], i)
		}
	}

	result := 0

	for i := 0; i < n; i++ {
		var b int
		fmt.Scan(&b)

		for _, j := range reverse[b] {
			if hash[j][b] {
				delete(hash[j], b)
				if !cleared[j] && len(hash[j]) == 0 {

					result++

					cleared[j] = true
				}
			}
		}
		fmt.Println(result)

	}

}
