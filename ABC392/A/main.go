package main

import "fmt"

func main() {

	arr := make([]int, 3)

	for i := 0; i < 3; i++ {
		fmt.Scan(&arr[i])
	}

	if arr[0]*arr[1] == arr[2] {
		fmt.Println("Yes")
		return
	}
	if arr[0]*arr[2] == arr[1] {
		fmt.Println("Yes")
		return
	}
	if arr[1]*arr[2] == arr[0] {
		fmt.Println("Yes")
		return
	}

	fmt.Println("No")

	return

}
