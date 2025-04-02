package main

// 280. 摆动排序

import "fmt"

func wiggleSort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		if (i%2 == 0 && nums[i] > nums[i+1]) || (i%2 == 1 && nums[i] < nums[i+1]) {
			temp := nums[i+1]
			nums[i+1] = nums[i]
			nums[i] = temp
		}
	}
	fmt.Println(nums)
}

// 贪心策略：
// 1. 遍历数组，判断当前元素和下一个元素的大小关系
// 2. 如果当前元素是偶数下标，且当前元素大于下一个元素，则交换两个元素
// 3. 如果当前元素是奇数下标，且当前元素小于下一个元素，则交换两个元素
// 4. 最终返回的数组就是摆动排序后的结果
// 5. 时间复杂度：O(n)，空间复杂度：O(1)
