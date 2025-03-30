package main

// 加油站
func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	sum_g := 0
	sum_c := 0
	cur := 0
	start := 0
	for i := 0; i < n; i++ {
		sum_g += gas[i]
		sum_c += cost[i]
	}
	if sum_g < sum_c {
		return -1
	}
	for i := 0; i < n; i++ {
		cur += (gas[i] - cost[i])
		if cur < 0 {
			start = i + 1
			cur = 0
		}
	}

	return start
}

// 贪心策略：如果累计小于0，则前方到目前下标位置都不可能，从i+1重新开始遍历
