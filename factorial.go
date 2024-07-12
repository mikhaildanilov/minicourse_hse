package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	var result = 1
	for i := 1; i <= a; i++ {
		result *= i
	}
	fmt.Println(result)
}
