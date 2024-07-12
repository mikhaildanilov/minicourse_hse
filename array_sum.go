package main

import (
	"fmt"
)

func sum(arr []int) (int) {
	var result = 0
	for _, a := range arr {
		result += a
	}
	return result
}

func main() {
	fmt.Println(sum([]int{1, 2, 3, 4, 5, 6, 7, 8}))
}
