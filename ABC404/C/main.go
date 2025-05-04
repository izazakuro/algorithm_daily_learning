package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func NewGraph(N, M int) [][]int {
	G := make([][]int, N+1)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	for i := 1; i <= M; i++ {
		scanner.Scan()
		a, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		b, _ := strconv.Atoi(scanner.Text())
		G[a] = append(G[a], b)
		G[b] = append(G[b], a)
	}

	return G
}

func isCycleGraph(n int, g [][]int) bool {
	for i := 1; i <= n; i++ {
		if len(g[i]) != 2 {
			return false
		}
	}

	visited := make([]bool, n+1)
	var dfs func(int)
	dfs = func(v int) {
		visited[v] = true
		for _, u := range g[v] {
			if !visited[u] {
				dfs(u)
			}
		}
	}
	dfs(1)

	for i := 1; i <= n; i++ {
		if !visited[i] {
			return false
		}
	}

	edgeCount := 0
	for i := 1; i <= n; i++ {
		edgeCount += len(g[i])
	}
	if edgeCount/2 != n {
		return false
	}

	return true
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	g := NewGraph(n, m)

	if isCycleGraph(n, g) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
