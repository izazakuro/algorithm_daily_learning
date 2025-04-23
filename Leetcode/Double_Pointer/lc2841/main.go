package main

// 2841. 几乎唯一子数组的最大和

func maxSum(nums []int, m int, k int) int64 {
	n := len(nums)
	l, r := 0, 0
	result := 0
	sum := 0

	m1 := make(map[int]int)

	for r < n {
		m1[nums[r]]++
		sum += nums[r]

		if r-l+1 > k {
			sum -= nums[l]
			m1[nums[l]]--
			if m1[nums[l]] == 0 {
				delete(m1, nums[l])
			}
			l++
		}

		if r-l+1 == k && len(m1) >= m {
			result = max(result, sum)
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
