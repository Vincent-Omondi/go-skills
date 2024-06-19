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

	var rot string
	for _, r := range arg {
		if 'a' <= r && r <= 'z' {
			rot += string((r-'a'+13)%26 + 'a')
		} else if 'A' <= r && r <= 'Z' {
			rot += string((r-'A'+13)%26 + 'A')
		} else {
			rot += string(r)
		}
	}
	for _, char := range rot {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
