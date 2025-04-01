package main

// 455. 分发饼干

import "sort"

func findContentChildren(g []int, s []int) int {

	sort.Ints(g)
	sort.Ints(s)

	res := 0

	index := len(s) - 1

	for i := len(g) - 1; i >= 0; i-- {
		if index >= 0 && s[index] >= g[i] {
			res++
			index--
		}
	}

	return res

}

// 贪心策略：用大饼干喂大孩子，不浪费大饼干
