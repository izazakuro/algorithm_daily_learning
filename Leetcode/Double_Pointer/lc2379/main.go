package main

// 2379. 得到k个黑块的最少涂色次数

func minimumRecolors(blocks string, k int) int {
	n := len(blocks)
	l, r := 0, 0

	result := n

	count := 0

	for r < n {
		if blocks[r] == 'W' {
			count++
		}

		if r-l+1 > k {
			if blocks[l] == 'W' {
				count--
			}
			l++
		}

		if r-l+1 == k {
			result = min(result, count)
		}

		r++
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
