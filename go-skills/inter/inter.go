package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) < 3 {
		return
	}

	s1, s2 := os.Args[1], os.Args[2]
	seens1, seens2 := make(map[rune]bool), make(map[rune]bool)

	for _, r := range s2 {
		seens2[r] = true
	}

	for _, c := range s1 {
		if seens2[c] && !seens1[c] {
			seens1[c] = true
			z01.PrintRune(c)
		}
	}
	z01.PrintRune('\n')
}
