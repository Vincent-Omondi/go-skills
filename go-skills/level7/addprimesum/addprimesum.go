package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	n := atoi(os.Args[1])
	sum := 0

	for i := n; i > 0; i-- {
		if IsPrime(i) {
			sum += i
		}
	}
	s := ""
	for ; sum > 0; sum /= 10 {
		s = string(byte(sum%10)+'0') + s
	}
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func IsPrime(x int) bool {
	if x < 2 {
		return false
	}

	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
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
