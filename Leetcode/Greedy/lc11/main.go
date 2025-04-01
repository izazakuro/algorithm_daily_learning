package main

// 盛最多水的容器
func maxArea(height []int) int {
    l := 0
    r := len(height) - 1

    area := 0
    max_val := 0

    for l != r {
        area = min(height[l], height[r]) * (r-l)
        max_val = max(max_val, area)
        if height[l] < height[r] {
            l ++
        } else {
            r --
        }
    }

    return max_val
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

// 贪心算法：每次移动最小的一边
// 如果取到的值大于最大值，则更新