package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func NewGraph(N, M int) ([][]int, int) {

	m := make(map[int]bool, N+1)

	num := 0
	temp := 0

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
		if a == b && !m[a] {
			num++
			m[a] = true
			continue
		}
		if a > b {
			temp = a
			a = b
			b = temp
		}
		if m[a] == true && m[b] == true {
			num++
		} else {
			m[a] = true
			m[b] = true
		}

	}

	return G, num

}

func main() {

	var n, m int

	fmt.Scan(&n, &m)

	_, result := NewGraph(n, m)

	fmt.Println(result)

}
