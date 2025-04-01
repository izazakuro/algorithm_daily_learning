package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInts(reader *bufio.Reader) []int {
	line, _ := reader.ReadString('\n')
	fields := strings.Fields(line)
	nums := make([]int, len(fields))
	for i, s := range fields {
		nums[i], _ = strconv.Atoi(s)
	}
	return nums
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	line, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(line))

	p := readInts(reader)
	p = append([]int{0}, p...)

	q := readInts(reader)
	q = append([]int{0}, q...)

	m := make([]int, n+1)
	for i := 1; i <= n; i++ {
		m[q[i]] = q[p[i]]
	}

	writer := bufio.NewWriter(os.Stdout)
	for i := 1; i <= n; i++ {
		fmt.Fprintf(writer, "%d ", m[i])
	}
	writer.Flush()
}
