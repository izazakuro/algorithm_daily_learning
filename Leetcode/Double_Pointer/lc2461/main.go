package main

// 2461. 长度为k子数组中的最大和

func maximumSubarraySum(nums []int, k int) int64 {
	n := len(nums)
	l, r := 0, 0
	m := make(map[int]int)
	var result int
	result = 0
	sum := 0

	for r < n {
		m[nums[r]]++
		sum += nums[r]

		if r-l+1 > k {
			m[nums[l]]--
			sum -= nums[l]
			if m[nums[l]] == 0 {
				delete(m, nums[l])
			}
			l++
		}

		if r-l+1 == k && len(m) == k {
			result = max(sum, result)
		}
		r++
	}

	return int64(result)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
