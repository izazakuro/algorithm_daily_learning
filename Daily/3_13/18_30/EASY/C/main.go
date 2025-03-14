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

func main() {

	var n, m int

	fmt.Scan(&n, &m)

	G := NewGraph(n, m)

	count := 0

	for i := 1; i <= n; i++ {
		for _, neighbor := range G[i] {
			for _, next := range G[neighbor] {
				if next == i {
					continue
				}
				for _, last := range G[next] {
					if last == i {
						count++
					}
				}
			}
		}
	}

	fmt.Println(count / 6)

}
