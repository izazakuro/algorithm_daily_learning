package main

import (
	"fmt"
	"strings"
)

func main() {

	var s string

	fmt.Scan(&s)

	upper := strings.ToUpper(s)

	fmt.Println(upper)

}
