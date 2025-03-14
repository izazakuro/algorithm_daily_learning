package main

import "fmt"

func binary_search(l, r, a int, arr []int) int {
	for l <= r {
		mid := (l + r) / 2
		if a < arr[mid] {
			r = mid - 1
		}
		if a == arr[mid] {
			return mid
		}
		if a > arr[mid] {
			l = mid + 1
		}
	}
	return -1
}

func main() {

	var n, a int
	fmt.Scan(&n, &a)

	arr := make([]int, n+1)

	for i := 1; i <= n; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Println(binary_search(1, n, a, arr))

}
