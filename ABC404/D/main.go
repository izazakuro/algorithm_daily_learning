package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	n, m        int
	cost        []int
	animalInZoo [][]bool
	minCost     = int(1e18)
	memo        = map[string]int{}
)

func dfs(seen []int, currentCost int) {
	if currentCost >= minCost {
		return
	}

	done := true
	for _, cnt := range seen {
		if cnt < 2 {
			done = false
			break
		}
	}
	if done {
		if currentCost < minCost {
			minCost = currentCost
		}
		return
	}

	key := fmt.Sprint(seen)
	if prev, ok := memo[key]; ok && prev <= currentCost {
		return
	}
	memo[key] = currentCost

	for i := 0; i < n; i++ {
		newSeen := make([]int, m)
		copy(newSeen, seen)

		for j := 0; j < m; j++ {
			if animalInZoo[i][j] {
				if newSeen[j] < 2 {
					newSeen[j]++
				}
			}
		}
		dfs(newSeen, currentCost+cost[i])
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ = strconv.Atoi(scanner.Text())
	scanner.Scan()
	m, _ = strconv.Atoi(scanner.Text())

	cost = make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		cost[i], _ = strconv.Atoi(scanner.Text())
	}

	animalInZoo = make([][]bool, n)
	for i := 0; i < n; i++ {
		animalInZoo[i] = make([]bool, m)
	}

	for i := 0; i < m; i++ {
		scanner.Scan()
		k, _ := strconv.Atoi(scanner.Text())
		for j := 0; j < k; j++ {
			scanner.Scan()
			zooID, _ := strconv.Atoi(scanner.Text())
			zooID--
			animalInZoo[zooID][i] = true
		}
	}

	seen := make([]int, m)
	dfs(seen, 0)
	fmt.Println(minCost)
}
