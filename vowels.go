package main

import (
	"fmt"
	"slices"
)

func main() {
	var a rune
	fmt.Scanf("%c", &a)

	vowels := []rune{'a', 'o', 'e', 'i', 'u', 'y'}

	fmt.Println(slices.Contains(vowels, a))

}
