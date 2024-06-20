package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	arg := os.Args[1]

	if arg == " " {
		return
	}

	for len(arg) > 0 && arg[0] == ' ' {
		arg = arg[1:]
	}

	for len(arg) > 0 && arg[len(arg)-1] == ' ' {
		arg = arg[:len(arg)-1]
	}

	var word string

	for _, char := range arg {
		if char != ' ' {
			word += string(char)
		} else {
			word = ""
		}
	}
	for _, r := range word {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
