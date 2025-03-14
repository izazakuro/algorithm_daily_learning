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

func dfs(g [][]int, pos int, visited []bool) []bool {
	visited[pos] = true
	for i := 0; i < len(g[pos]); i++ {
		next := g[pos][i]
		if visited[next] == false {
			visited = dfs(g, next, visited)
		}
	}
	return visited

}

func main() {

	var n, m int

	fmt.Scan(&n, &m)

	G := NewGraph(n, m)

	visited := make([]bool, n+1)

	visited = dfs(G, 1, visited)

	for i := 1; i <= n; i++ {
		if visited[i] == false {
			fmt.Println("The graph is not connected.")
			return
		}
	}

	fmt.Println("The graph is connected.")

}
