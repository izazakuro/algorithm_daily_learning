package main

// 斐波那契数列
func fib(n int) int {
	arr := make([]int, 31)
	arr[0] = 0
	arr[1] = 1

	for i := 2; i <= 30; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}

	return arr[n]
}
