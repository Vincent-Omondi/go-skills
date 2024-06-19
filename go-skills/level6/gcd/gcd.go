package main

import (
	"os"

	"github.com/01-edu/z01"
)

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
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

func main() {
	if len(os.Args) != 3 {
		return
	}

	num1, num2 := atoi(os.Args[1]), atoi(os.Args[2])

	res := itoa(gcd(num1, num2))

	for _,r := range res {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func itoa(n int) string {
	s := ""
	for ; n > 0; n /= 10 {
		s = string(byte(n%10+'0')) + s
	}
	return s
}
