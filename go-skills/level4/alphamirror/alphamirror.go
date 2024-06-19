package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) < 2 {
		return
	}
	arg := os.Args[1]

	result := ""
	for _, r := range arg {
		if 'a' <= r && r <= 'z' {
			result += string('a' - r + 'z')
		} else if 'A' <= r && r <= 'Z' {
			result += string('A' - r + 'Z')
		} else {
			result += string(r)
		}
	}
	for _, v := range result {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}
