package main

import (
	"os"

	"github.com/01-edu/z01"
)

// Converts a string to an integer
func atoi(s string) int {
	i := 0
	for _, n := range s {
		if n < '0' || n > '9' {
			return -1
		}
		i = i*10 + int(n-'0')
	}
	return i
}

// Converts an integer to a hexadecimal string
func itoa_hex(n int) string {
	if n == 0 {
		return "0"
	}

	hexChars := "0123456789abcdef"
	s := ""
	for ; n > 0; n /= 16 {
		s = string(hexChars[n%16]) + s
	}
	return s
}

// Prints a string using z01.PrintRune
func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func main() {
	if len(os.Args) != 2 {
		return
	}

	num := atoi(os.Args[1])
	if num < 0 {
		printStr("ERROR")
		return
	}

	hex := itoa_hex(num)
	printStr(hex)
}
