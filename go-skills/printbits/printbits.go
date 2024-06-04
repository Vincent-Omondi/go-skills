package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	num := atoi(os.Args[1])
	if num > 255 {
		printBinary("00000000")
		return
	}

	binary := ""
	for i := 7; i >= 0; i-- {
		if num&(1<<i) != 0 {
			binary += "1"
		} else {
			binary += "0"
		}
	}
	printBinary(binary)
}

func atoi(s string) int {
	i := 0
	for _, n := range s {
		if n < '0' || n > '9' {
			return 0
		}
		i = i*10 + int(n-'0')
	}
	return i
}

func printBinary(binary string) {
	for _, c := range binary {
		z01.PrintRune(c)
	}
}
