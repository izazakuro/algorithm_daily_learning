package main

import (
	"fmt"
	"math"
)

func main() {

	var x1, y1, x2, y2 int64

	fmt.Scan(&x1, &y1, &x2, &y2)

	seki := (x2-x1)*(x2-x1) + (y2-y1)*(y2-y1)

	if seki <= 20 {
		if seki <= 20 {
			floatRoot := math.Sqrt(float64(seki))
			intRoot := int64(math.Round(floatRoot))

			if intRoot*intRoot == seki {
				fmt.Println("Yes")
			} else {
				fmt.Println("No")
			}
		} else {
			fmt.Println("No")
		}
	}

}
