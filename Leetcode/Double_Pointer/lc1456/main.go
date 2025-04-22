package main

// 1456. 定长字母中元音的最大数目

func maxVowels(s string, k int) int {
	n := len(s)
	l, r := 0, 0
	m := make(map[byte]bool)
	m['a'] = true
	m['e'] = true
	m['i'] = true
	m['o'] = true
	m['u'] = true
	result := 0
	count := 0

	for r < n {
		c := s[r]
		if m[c] {
			count++
		}
		if r-l+1 > k {
			c1 := s[l]
			if m[c1] {
				count--
			}
			l++
		}
		result = max(result, count)
		r++

	}

	return result

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
