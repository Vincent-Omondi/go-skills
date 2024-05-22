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
	for i := 1; i <= 9; i++ {
		str := itoa(i*num) + "\n"
		for _, char := range []byte(itoa(i) + " x " + itoa(num) + " = " + str) {
			z01.PrintRune(rune(char))
		}
	}
}

func atoi(s string) int {
	i := 0
	for _, c := range s {
		if c < '0' || c > '9' {
			return i
		}
		i = i*10 + int(c-'0')
	}
	return i
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	s := ""
	for ; n > 0; n /= 10 {
		s = string(byte('0'+n%10)) + s
	}
	return s
}
