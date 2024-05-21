package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) <= 1 {
		return
	}

	args := os.Args[1:]

	firstparam := args[0]

	for _, v := range firstparam {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}
