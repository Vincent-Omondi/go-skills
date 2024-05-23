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

	digits := []rune{}
	for count > 0 {
		digit := (rune(count%10) + '0')
		digits = append(digits, digit)
		count /= 10
	}

	for i := len(digits) - 1; i >= 0; i-- {
		z01.PrintRune(digits[i])
	}
	z01.PrintRune('\n')
}
