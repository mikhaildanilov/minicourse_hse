package main

import (
	"fmt"
)

func main() {
	var a string
	fmt.Scan(&a)
	var result string
	for i := len(a) - 1; i >= 0; i-- {
		result += string(a[i])
	}
	fmt.Println(result)
}
