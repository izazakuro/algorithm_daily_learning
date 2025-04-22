package main

import (
	"bufio"
	"fmt"
	"os"
)

const MAXN = 10
const MAXX = 105

var (
	N, X int
	S    [MAXN]int
	C    [MAXN]int
	P    [MAXN]float64

	memo [MAXX][1 << MAXN]float64
	vis  [MAXX][1 << MAXN]bool
)

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func dfs(money int, used int) float64 {
	if vis[money][used] {
		return memo[money][used]
	}
	vis[money][used] = true

	res := 0.0

	for i := 0; i < N; i++ {
		if (used>>i)&1 == 1 || money < C[i] {
			continue
		}

		nextMoney := money - C[i]

		success := P[i] * (float64(S[i]) + dfs(nextMoney, used|(1<<i)))

		fail := (1 - P[i]) * dfs(nextMoney, used)

		res = max(res, success+fail)
	}

	memo[money][used] = res
	return res
}

func main() {
	fmt.Scan(&N, &X)
	for i := 0; i < N; i++ {
		fmt.Scan(&S[i], &C[i])
		var p int
		fmt.Scan(&p)
		P[i] = float64(p) / 100.0
	}

	ans := dfs(X, 0)

	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	fmt.Fprintf(writer, "%.10f\n", ans)
}
