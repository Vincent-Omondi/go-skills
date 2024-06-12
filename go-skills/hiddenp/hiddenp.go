package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		return
	}

	s1, s2 := os.Args[1], os.Args[2]

	j := 0

	for _, r := range s1 {
		found := false
		for j < len(s2) && !found {
			if rune(s2[j]) == r {
				found = true
			}
			j++
		}
		if !found {
			z01.PrintRune('0')
			z01.PrintRune('\n')
			return
		}
	}
	z01.PrintRune('1')
	z01.PrintRune('\n')
}
