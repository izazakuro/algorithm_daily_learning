package main

import "sort"

// 406. 根据身高重建队列

func reconstructQueue(people [][]int) [][]int {
	n := len(people)
	var res [][]int

	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}
		return people[i][0] > people[j][0]
	})

	for i := 0; i < n; i++ {
		pos := people[i][1]
		res = insert(res, pos, people[i])
	}

	return res
}

func insert(matrix [][]int, index int, row []int) [][]int {
	matrix = append(matrix[:index], append([][]int{row}, matrix[index:]...)...)
	return matrix
}

// 贪心策略；先按照身高降序排列，保证每个元算前面的所有元素都比自己大
// 之后按照排位顺序向对应的位置进行插入即可，因为前面的元素都比自己大，
// 所以插入进去之后，前面只有比自己大的元素，而因为元素是从后面插入进来的
// 只会是比当前位置小的元素，所以不影响当前位置元素的结果
