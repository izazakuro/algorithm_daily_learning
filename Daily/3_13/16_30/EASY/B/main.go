package main

import "fmt"

func main() {

	var a, b, c, d, e, f, x int

	fmt.Scan(&a, &b, &c, &d, &e, &f, &x)
	tmp := x
	t := 0
	ao := 0
	for x > 0 {
		if x-a > 0 {
			t += a * b
			x -= a

		} else if x-a <= 0 {
			t += x * b
			x -= a

		}
		x -= c
	}

	x = tmp

	for x > 0 {
		if x-d > 0 {
			ao += d * e
			x -= d

		} else if x-d <= 0 {
			ao += x * e
			x -= d

		}
		x -= f
	}

	if t > ao {
		fmt.Println("Takahashi")
	} else if t == ao {
		fmt.Println("Draw")
	} else {
		fmt.Println("Aoki")
	}

}
