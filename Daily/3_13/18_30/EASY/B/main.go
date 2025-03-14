package main

import "fmt"

func function(x int) int {
	return x*x + 2*x + 3
}

func main() {

	var t int

	fmt.Scan(&t)

	ft := function(t)
	f1 := ft + t
	f2 := function(f1)
	f3 := function(ft)
	result := function(f2 + f3)
	fmt.Println(result)

}
