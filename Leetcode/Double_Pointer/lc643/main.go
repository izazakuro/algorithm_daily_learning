package main

// 643. 子数组最大平均数

func findMaxAverage(nums []int, k int) float64 {
	n := len(nums)
	if n < k {
		sum := 0.0
		for _, val := range nums {
			sum += float64(val)
		}
		return sum / float64(n)
	}
	l, r := 0, 0

	result := float64(-1e9)
	sum := 0.0

	for r < n {

		sum += float64(nums[r])

		if r-l+1 > k {
			sum -= float64(nums[l])
			l++
		}
		if r-l+1 == k {
			result = max(result, sum/float64(k))
		}
		r++
	}
	return result
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}
