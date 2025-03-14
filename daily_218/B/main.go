package main

import "fmt"

func main() {

	var n int

	fmt.Scan(&n)

	t, a := 0, 0

	for i := 0; i < n; i++ {
		var x int
		fmt.Scan(&x)
		t += x
		fmt.Scan(&x)
		a += x
	}

	if t > a {
		fmt.Println("Takahashi")
	} else if t == a {
		fmt.Println("Draw")
	} else {
		fmt.Println("Aoki")

	}

}
