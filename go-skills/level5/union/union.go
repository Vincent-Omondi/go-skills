package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		return
	}

	union := make(map[rune]bool)

	for _, c := range os.Args[1] + os.Args[2] {
		if !union[c] {
			union[c] = true
			z01.PrintRune(c)
		}
	}
	z01.PrintRune('\n')
}
