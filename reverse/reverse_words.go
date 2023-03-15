package main

import "strings"

func main() {
	s := "Let's take LeetCode contest"
	reverseWords(s)
}

func reverseWords(s string) string {
	strs := strings.Split(s, " ")

	for i := range strs {
		revw := reverseWord(strs[i])
		strs[i] = revw
	}

	return strings.Join(strs, " ")
}

func reverseWord(s string) string {
	b := []byte(s)

	var start int = 0
	var end int = len(b) - 1

	for start < end {
		b[start], b[end] = b[end], b[start]
		start++
		end--
	}
	return string(b)
}
