package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) <= 1 {
		z01.PrintRune('0')
		z01.PrintRune('\n')
		return
	}

	count := len(os.Args[1:])
	z01.PrintRune(rune(count) + '0')
	z01.PrintRune('\n')
}
