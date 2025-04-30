package main

// 300. 最长递增子序列

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))

	res := 1

	dp[0] = 1

	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[j]+1, dp[i])
			}
		}
		res = max(res, dp[i])
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// dp含义：0～i最长的单调递增子序列长度
// 用双层循环遍历，dp[i] = max(dp[j](j<i), dp[i])
