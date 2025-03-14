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

	var N, M int

	fmt.Scan(&N, &M)

	G := NewGraph(N, M)

	for i := 1; i <= N; i++ {
		fmt.Printf("%d: {", i)
		for j := 0; j < len(G[i]); j++ {
			if j >= 1 {
				fmt.Printf(", ")
			}
			fmt.Printf("%d", G[i][j])
		}
		fmt.Println("}")
	}

}
