package main

// 股票买卖的最佳时机2
func maxProfit(prices []int) int {
    ans := 0
    for i := 1 ; i < len(prices) ; i ++ {
        ans += max(0, prices[i] - prices[i-1])
    }
    return ans
}

func max(a,b int) int {
    if a > b {
        return a
    }

    return b
}

// 贪心：只要涨了就卖 跟0取最大值