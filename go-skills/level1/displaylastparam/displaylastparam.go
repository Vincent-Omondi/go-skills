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

	lastparam := args[len(args)-1]

	for _, v := range lastparam {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}
