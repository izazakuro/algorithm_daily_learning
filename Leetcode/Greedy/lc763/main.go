package main

// 763. 划分字母区间

func partitionLabels(s string) []int {
	m := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		m[s[i]] = i
	}

	var res []int

	l, r := 0, 0
	for i := 0; i < len(s); i++ {
		r = max(r, m[s[i]])
		if i == r {
			res = append(res, r-l+1)
			l = i + 1
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 贪心策略： 计算每个字母出现的最远距离
// 对字符串进行遍历，中间取各字母最远距离的最大值，
// 遍历到对应下标时，表示已经遍历到对应区间最大值
