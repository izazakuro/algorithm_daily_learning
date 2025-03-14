package main

import (
	"fmt"
)

func cal(x float32) float32 {

	return x*x*x + x
}

func main() {

	var n int

	fmt.Scan(&n)

	var l float32 = 0.0
	var r float32 = 100.0
	var mid float32

	for i := 0; i < 20; i++ {
		mid = (l + r) / 2.0
		sum := cal(float32(mid))
		if sum > float32(n) {
			r = mid
		}
		if sum < float32(n) {
			l = mid
		}
	}

	fmt.Printf("%.5f", mid)
}
