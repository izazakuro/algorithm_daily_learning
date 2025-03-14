package main

import "fmt"

func main() {

	s := ""

	var num int
	for i := 0; i < 26; i++ {
		fmt.Scan(&num)
		ch := rune(num + 96)
		s += string(ch)
	}

	fmt.Println(s)

}
