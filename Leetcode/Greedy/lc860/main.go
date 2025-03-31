package main

func lemonadeChange(bills []int) bool {
	m := make(map[int]int)

	for _, val := range bills {
		if val == 5 {
			m[val]++
		} else if val == 10 {
			m[5]--
			if m[5] < 0 {
				return false
			}
			m[val]++
		} else {
			if m[10] != 0 {
				m[10]--
				m[5]--
				if m[5] < 0 {
					return false
				}
			} else if m[5] >= 3 {
				m[5] -= 3
			} else {
				return false

			}
		}
	}
	return true
}

// 贪心策略：先使用10找零，如果val=10只能用5找零，20:3张5/1张10 1张5
