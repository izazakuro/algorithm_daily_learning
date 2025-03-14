package common

import (
	"bufio"
	"os"
	"sort"
	"strconv"
)

func binary_search(l, r, a int, arr []int) int {
	for l <= r {
		mid := (l + r) / 2
		if a < arr[mid] {
			r = mid - 1
		}
		if a == arr[mid] {
			return mid
		}
		if a > arr[mid] {
			l = mid + 1
		}
	}
	return -1
}

func lower_bound(arr []int, x int) int {
	index := sort.Search(len(arr), func(i int) bool {
		return arr[i] >= x
	})
	return index
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func absInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// Graph

type Edge struct {
	to int
	w  int
}

func ReadGraph() ([][]Edge, int, int) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	M, _ := strconv.Atoi(scanner.Text())

	G := make([][]Edge, N)

	for i := 0; i < M; i++ {
		scanner.Scan()
		u, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		v, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		w, _ := strconv.Atoi(scanner.Text())

		u--
		v--

		G[u] = append(G[u], Edge{to: v, w: w})
		G[v] = append(G[v], Edge{to: u, w: w})
	}

	return G, N, M
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
