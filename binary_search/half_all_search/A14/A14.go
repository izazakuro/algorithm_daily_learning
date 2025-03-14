package main

import (
	"fmt"
	"sort"
)

func main() {
	var N, K int
	fmt.Scan(&N, &K)

	A := make([]int, N+1)
	B := make([]int, N+1)
	C := make([]int, N+1)
	D := make([]int, N+1)

	for i := 1; i <= N; i++ {
		fmt.Scan(&A[i])
	}
	for i := 1; i <= N; i++ {
		fmt.Scan(&B[i])
	}
	for i := 1; i <= N; i++ {
		fmt.Scan(&C[i])
	}
	for i := 1; i <= N; i++ {
		fmt.Scan(&D[i])
	}

	P := make([]int, N*N+1)
	Q := make([]int, N*N+1)

	idx := 1
	for x := 1; x <= N; x++ {
		for y := 1; y <= N; y++ {
			P[idx] = A[x] + B[y]
			idx++
		}
	}

	idx = 1
	for z := 1; z <= N; z++ {
		for w := 1; w <= N; w++ {
			Q[idx] = C[z] + D[w]
			idx++
		}
	}

	sort.Slice(Q[1:], func(i, j int) bool {
		return Q[1:][i] < Q[1:][j]
	})

	for i := 1; i <= N*N; i++ {
		need := K - P[i]
		pos := sort.Search(N*N, func(mid int) bool {
			// 返回一个pos下标，根据true和false确定查找区间
			return Q[mid+1] >= need
		})

		// 终止条件
		if pos < N*N && Q[pos+1] == need {
			fmt.Println("Yes")
			return
		}
	}

	fmt.Println("No")
}

// 1. 确保在区间内
// 2. 单调递增或递减