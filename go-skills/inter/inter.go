package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		return
	}

	s1, s2 := os.Args[1], os.Args[2]
	seen := make(map[rune]bool)
	result := ""

	for _, r := range s1 {
		if !seen[r] && containRune(s2, r) {
			seen[r] = true
			result += string(r)
		}
	}
	fmt.Println(result)
}

func containRune(s string, r rune) bool {
	for _, c := range s {
		if c == r {
			return true
		}
	}
	return false
}
