package main

import "fmt"

func check(x, n, k int, arr []int) bool {
	var sum int
	for i := 1; i <= n; i++ {
		sum += x / arr[i]
	}
	if sum >= k {
		return true
	}
	return false
}

func main() {

	var n, k int

	fmt.Scan(&n, &k)

	arr := make([]int, n+1)

	for i := 1; i <= n; i++ {
		fmt.Scan(&arr[i])
	}

	l := 1
	r := 1000000000

	for l < r {
		mid := (l + r) / 2
		ok := check(mid, n, k, arr)
		if !ok {
			l = mid + 1
		}
		if ok {
			r = mid
		}
	}

	fmt.Println(l)

}
