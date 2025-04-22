package main

// 2090. 半径为k的子数组平均值

func getAverages(nums []int, k int) []int {
	n := len(nums)
	l, r := 0, 0
	avg := make([]int, 0)
	sum := 0

	for r < n {
		sum += nums[r]

		if r-l > k {
			sum -= nums[l]
			l++
		}
		if (r-l == k) && (r+k < n) {
			temp := 0
			for i := r + 1; i <= r+k; i++ {
				temp += nums[i]
			}
			avg = append(avg, (temp+sum)/(2*k+1))
		} else {
			avg = append(avg, -1)
		}

		r++
	}

	return avg
}
