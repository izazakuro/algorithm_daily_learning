package main

import (
	"fmt"
)

type Edge struct {
	to int
	w  rune
}

func main() {
	var n int

	fmt.Scan(&n)

	G := make([][]Edge, n+1)

	for i := 0; i <= n; i++ {
		G[i] = make([]Edge, n+1)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			var w rune
			fmt.Scan(&w)
			G[i] = append(G[i], Edge{to: j, w: w})
		}
	}

}
