package main

func tribonacci(n int) int {
	arr := make([]int, 38)
	arr[0] = 0
	arr[1] = 1
	arr[2] = 1

	for i := 3; i <= 37; i++ {
		arr[i] = arr[i-3] + arr[i-2] + arr[i-1]
	}

	return arr[n]
}
