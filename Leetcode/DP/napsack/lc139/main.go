package main

import "fmt"

// 139. 单词拆分

func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)

	dp[0] = true

	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			temp := s[j:i]
			for _, val := range wordDict {
				if val == temp && dp[j] == true {
					dp[i] = true
					break
				}
			}
			fmt.Println(dp)
		}
	}

	return dp[len(s)]
}

// dp含义：0～i是否存在于字典中，存在则为true
// 递推：dp[i]中， 若dp[j]为true，且dp[j:i]也存在于字典中，则dp[i]为true
