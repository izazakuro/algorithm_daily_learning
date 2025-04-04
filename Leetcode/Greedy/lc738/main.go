package main

import "strconv"

func monotoneIncreasingDigits(n int) int {
	s := []byte(strconv.Itoa(n))
	pos := len(s)

	for i := len(s) - 1; i > 0; i-- {
		if s[i] < s[i-1] {
			s[i-1]--
			pos = i
		}
	}

	for i := pos; i < len(s); i++ {
		s[i] = '9'
	}

	res, _ := strconv.Atoi(string(s))
	return res
}

// 贪心策略：从后向前遍历，找到处于最前方的小于i-1的位置，将其之后全部变为9
