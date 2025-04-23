package main

// 1423. 可获得最大点数

func maxScore(cardPoints []int, k int) int {
	n := len(cardPoints)
	l, r := 0, 0
	m := n - k
	sum := 0
	sum_all := 0

	min_val := 1000000000

	for r < n {
		sum_all += cardPoints[r]
		sum += cardPoints[r]

		if r-l+1 > m {
			sum -= cardPoints[l]
			l++
		}

		if r-l+1 == m {
			min_val = min(sum, min_val)
		}

		r++
	}

	return sum_all - min_val

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 思路：因为只能取两边的牌，所以中间的部分一定是一个定长的连续数组， 因此，需要维护一个 n-k 长度的滑动窗口
//      之后取窗口中总和的最小值，用数组总和 - 最小值 = 最大值
