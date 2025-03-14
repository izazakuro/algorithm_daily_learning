package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func nextPermutation(a []int) bool {
	i := len(a) - 2
	for i >= 0 && a[i] >= a[i+1] {
		i--
	}
	if i < 0 {
		return false
	}

	j := len(a) - 1
	for a[j] <= a[i] {
		j--
	}

	a[i], a[j] = a[j], a[i]

	left, right := i+1, len(a)-1
	for left < right {
		a[left], a[right] = a[right], a[left]
		left++
		right--
	}
	return true
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

func main() {

	var N, M int

	fmt.Scan(&N, &M)

	g1 := NewGraph(N, M)
	g2 := NewGraph(N, M)

	p := make([]int, N+1)
	for i := 1; i <= N; i++ {
		p[i] = i
	}

	for {
		ok := true
	Outer:
		for i := 1; i <= N; i++ {
			for j := 1; j <= N; j++ {
				if g1[i][j] != g2[p[i]][p[j]] {
					ok = false
					break Outer
				}
			}
		}
		if ok {
			fmt.Println("Yes")
			return
		}
		if !nextPermutation(p) {
			break
		}
	}

	fmt.Println("No")

}
