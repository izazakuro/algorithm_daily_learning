package main

// 1343. 大小为k且平局值大于阈值的子数组

func numOfSubarrays(arr []int, k int, threshold int) int {
	n := len(arr)
	l, r := 0, 0
	sum := 0
	result := 0

	for r < n {

		sum += arr[r]

		if r-l+1 > k {
			sum -= arr[l]
			l++
		}

		if float64(sum/k) >= float64(threshold) && r-l+1 == k {
			result++
		}

		r++
	}

	return result
}
