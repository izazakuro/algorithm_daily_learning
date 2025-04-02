package main

import "sort"

// 452. 用最少数量的箭引爆气球

func findMinArrowShots(points [][]int) int {
	n := len(points)
	res := 1

	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	for i := 1; i < n; i++ {
		if points[i][0] > points[i-1][1] {
			res++
		} else {
			points[i][1] = min(points[i-1][1], points[i][1])
		}
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 贪心策略：射气球重叠最多的地方
// 1. 排序
// 2. 如果本气球跟上一个气球不重叠，即本气球的左边界大于上一个气球的右边界，则箭++
// 3. 如果重叠，更新右边界为两个气球中的最小值，以便判断后续的气球是否重叠
