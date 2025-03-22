package main

import "fmt"

func main() {

	var s1, s2 string

	fmt.Scan(&s1, &s2)

	if s1 == "sick" && s2 != "sick" {
		fmt.Println(2)
	} else if s1 != "sick" && s2 == "sick" {
		fmt.Println(3)
	} else if s1 == "sick" && s2 == "sick" {
		fmt.Println(1)
	} else if s1 == "fine" && s2 == "fine" {
		fmt.Println(4)
	}

}
