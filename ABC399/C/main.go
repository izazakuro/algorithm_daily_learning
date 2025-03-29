package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	G := NewGraph(n, m)
	count := 0

	visited := make([]bool, n+1)

	for i := 1; i <= n; i++ {
		if !visited[i] {
			dfs(i, -1, visited, G, &count)
		}
	}

	fmt.Println(count / 2)

}

func dfs(now int, parent int, visited []bool, g [][]int, count *int) {
	visited[now] = true
	for _, v := range g[now] {
		if !visited[v] {
			dfs(v, now, visited, g, count)
		} else if v != parent {
			*count++
		}
	}
}

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
