package main

import "sort"

// 435. 无重叠区间

func eraseOverlapIntervals(intervals [][]int) int {

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := 0

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= intervals[i-1][1] {
			continue
		} else {
			res++
			intervals[i][1] = min(intervals[i-1][1], intervals[i][1])
		}
	}

	return res

}

// 贪心策略：
// 1. 排序
// 2. 如果区间不重叠，什么都不做
// 3. 如果重叠，更新右边界为两个区间右边界中的最小值，以便判断后续的区间是否重叠