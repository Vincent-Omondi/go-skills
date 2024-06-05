package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	for _, c := range os.Args[1] {
		var repeatCount int
		if c >= 'a' && c <= 'z' {
			repeatCount = int(c - 'a' + 1)
		} else if c >= 'A' && c <= 'Z' {
			repeatCount = int(c - 'A' + 1)
		} else {
			repeatCount = 1
		}

		for i := 0; i < repeatCount; i++ {
			z01.PrintRune(c)
		}
	}
	z01.PrintRune('\n')
}
