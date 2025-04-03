package main

// 56. 合并区间

import "sort"

func merge(intervals [][]int) [][]int {
	n := len(intervals)

	var res [][]int

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res = append(res, intervals[0])

	for i := 1; i < n; i++ {
		if intervals[i][0] <= res[len(res)-1][1] {
			res[len(res)-1][1] = max(intervals[i][1], res[len(res)-1][1])
		} else {
			res = append(res, intervals[i])
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

// 贪心策略：
// 1. 排序
// 2. 判断是否重合，重合就对结果的右区间进行修改
// 注：判断直接用结果中的区间进行判断
