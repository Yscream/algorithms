package main

import (
	"fmt"
)

func main() {
	s := []string{"h", "e", "l", "l", "o"}
	reverseString(s)
}

func reverseString(s []string) {
	var start int = 0
	var end int = len(s) - 1

	for start < end {
		s[start], s[end] = s[end], s[start]
		start++
		end--
		fmt.Println(s)
	}
}
