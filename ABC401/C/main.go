package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)
	buf := make([]byte, 0, 1024*1024)
	in.Buffer(buf, 1024*1024)

	readInt := func() int {
		in.Scan()
		n := 0
		for _, b := range in.Bytes() {
			n = n*10 + int(b-'0')
		}
		return n
	}

	N := readInt()
	M := readInt()

	badCount := make([]int, M)
	foodToDish := make([][]int, N+1)

	for i := 0; i < M; i++ {
		k := readInt()
		for j := 0; j < k; j++ {
			a := readInt()
			foodToDish[a] = append(foodToDish[a], i)
			badCount[i]++
		}
	}

	isCleared := make([]bool, M)

	B := make([]int, N)
	for i := 0; i < N; i++ {
		B[i] = readInt()
	}

	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	result := 0

	for i := 0; i < N; i++ {
		b := B[i]
		for _, dish := range foodToDish[b] {
			badCount[dish]--
			if !isCleared[dish] && badCount[dish] == 0 {
				result++
				isCleared[dish] = true
			}
		}
		fmt.Fprintln(writer, result)
	}
}
