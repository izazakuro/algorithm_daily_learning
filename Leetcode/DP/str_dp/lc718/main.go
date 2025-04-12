package main

// 718. 最长重复子数组

func findLength(nums1 []int, nums2 []int) int {

	n := len(nums1)
	m := len(nums2)

	dp := make([][]int, n+1)

	res := 0

	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			}
			res = max(res, dp[i][j])
		}
	}

	return res

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// dp含义，数组1的0～i-1 和 数组2的0～j-1中的最长公共子数组的长度
// 因为第一行和第一列没有实际意义，因此，为i-1和j-1
// 递推： if nums1[i-1] == nums2[j-1] ; dp[i][j] = dp[i-1][j-1]+1
