package main

// 分发糖果
func candy(ratings []int) int {
    n := len(ratings)
    candy := make([]int, n)

    for i := 0 ; i < n ; i ++ {
        candy[i] = 1
    }

    for i := 1 ; i < n ; i ++ {
        if ratings[i] > ratings[i-1] {
            candy[i] = candy[i-1] + 1
        }
    }

    for i := n - 2 ; i >= 0 ; i -- {
        if ratings[i] > ratings[i+1]{
            candy[i] = max(candy[i+1] + 1 , candy[i])
        }
    }
    sum := 0

    for _, val := range candy {
        sum += val
    }

    return sum
}

func max(a,b int) int {
    if a > b {
        return a
    }
    return b
}

// 贪心策略：从前到后遍历一遍，找出右比左大的情况，右+1
// 再从后到前遍历一遍，取max(左比右大 左+1，右比左大)
// 看两边的情况都先看一边，再从另一边遍历