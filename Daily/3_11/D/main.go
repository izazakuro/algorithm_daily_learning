package main

import "fmt"

func main() {

	var n, m int

	fmt.Scan(&n, &m)

	S := make([][]byte, n+1)

	T := make([][]byte, m+1)

	for i := 1; i <= n; i++ {
		S[i] = make([]byte, n+1)
	}
	for i := 1; i <= m; i++ {
		T[i] = make([]byte, m+1)

	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			fmt.Scan(&S[i][j])
		}
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= m; j++ {
			fmt.Scan(&T[i][j])
		}
	}

	for a := 1; a <= n-m+1; a++ {
		for b := 1; b <= n-m+1; b++ {
			for i := 1; i <= m; i++ {
				for j := 1; j <= m; j++ {
					if S[a+i-1][b+j-1] == T[i][j] {

						fmt.Printf("%d %d", a, b)
						return

					}
				}
			}
		}

	}

}
