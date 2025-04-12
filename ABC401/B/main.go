package main

import "fmt"

func main() {

	var s string

	var n int

	fmt.Scan(&n)

	isLogin := false

	res := 0

	for i := 0; i < n; i++ {
		fmt.Scan(&s)
		if !isLogin && s == "login" {
			isLogin = true
			continue
		}
		if isLogin && s == "logout" {
			isLogin = false
			continue
		}
		if !isLogin && s == "private" {
			res++
		} else if isLogin && s == "login" {
			continue
		} else if !isLogin && s == "logout" {
			continue
		}
	}

	fmt.Println(res)

	return
}
