package main

// 118. 杨辉三角

func generate(numRows int) [][]int {
    dp := make([][]int, numRows)
    
    for i := 0 ; i < numRows ; i ++ {
        dp[i] = make([]int,0)
    }

    dp[0] = append(dp[0], 1)

    for i := 1; i < numRows ; i ++ {
        for j := 0 ; j <= i ; j ++ {
            if j != 0 && j != i {
                dp[i] = append(dp[i],dp[i-1][j-1] + dp[i-1][j])
            } else {
                dp[i] = append(dp[i], 1)
            }
        }
    }
    return dp
}