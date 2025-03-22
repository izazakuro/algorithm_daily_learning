package main

import "fmt"

func main() {

	var x float32

	fmt.Scan(&x)

	if x >= 38.0 {
		fmt.Println(1)
	} else if (x >= 37.5) && (x < 38.0) {
		fmt.Println(2)
	} else if x < 37.5 {
		fmt.Println(3)
	}
}
