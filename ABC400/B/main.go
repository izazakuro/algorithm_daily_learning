package main

import (
	"fmt"
	"math"
)

func main() {
	var n, m int64
	fmt.Scan(&n, &m)

	var res int64 = 0
	for i := int64(0); i <= m; i++ {
		term := sekiSafe(n, i)
		if term == -1 || res > 1e9-term {
			fmt.Println("inf")
			return
		}
		res += term
	}

	fmt.Println(res)
}

func sekiSafe(a int64, index int64) int64 {
	res := int64(1)
	for i := int64(0); i < index; i++ {
		if willOverflow(res, a) {
			return -1
		}
		res *= a
	}
	return res
}

func willOverflow(a, b int64) bool {
	if a == 0 || b == 0 {
		return false
	}
	return a > math.MaxInt64/b
}
