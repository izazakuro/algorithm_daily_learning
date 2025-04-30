package main

// 34. 在排序数组中查找第一个和最后一个位置

func searchRange(nums []int, target int) []int {
    s := lowerBoudry(nums, target)
    if s == len(nums) || nums[s] != target {
        return []int{-1,-1}
    }

    e := lowerBoudry(nums, target + 1) - 1
    return []int{s,e}
}

func lowerBoudry(nums []int, target int) int {
    l,r := 0,len(nums) - 1

    for l <= r {
        mid := l +(r - l)/ 2

        if nums[mid] >= target {
            r = mid - 1
        } else {
            l = mid + 1
        }
    }    

    return l
}