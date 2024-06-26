package main

import "github.com/01-edu/z01"

func LastRune(s string) rune {
	var last rune
	for _, r := range s {
		last = r
	}
	return last
}

func main() {
	z01.PrintRune(LastRune("Hello!"))
	z01.PrintRune(LastRune("Salut!"))
	z01.PrintRune(LastRune("Ola!"))
	z01.PrintRune(LastRune(""))
	z01.PrintRune('\n')
}
