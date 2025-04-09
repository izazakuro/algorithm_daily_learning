package main

// 213. 打家劫舍2

func rob(nums []int) int {

	n := len(nums)
	if n == 1 {
		return nums[0]
	} else if n < 3 {
		return max(nums[0], nums[1])
	}
	dp1 := make([]int, len(nums))
	dp2 := make([]int, len(nums))

	dp2[1] = nums[1]
	dp2[2] = max(nums[1], nums[2])
	dp1[0] = nums[0]
	dp1[1] = max(nums[0], nums[1])

	for i := 2; i < len(nums)-1; i++ {
		dp1[i] = max(dp1[i-2]+nums[i], dp1[i-1])
	}

	for i := 3; i < len(nums); i++ {
		dp2[i] = max(dp2[i-2]+nums[i], dp2[i-1])
	}

	return max(dp1[len(nums)-2], dp2[len(nums)-1])

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// dp含义：0～i个房间能偷取的最大值为dp[i]
// 环形 = 不能同时偷第一个房间和最后一个房间，分别对从第一个房间到倒数第二个房间
//       以及第二个房间到最后一个房间进行dp求解，最后取两个情况的最大值
