package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	n, _ := strconv.Atoi(os.Args[1])
	if n <= 1 {
		return
	}

	primeFactors := getPrimeFactors(n)

	for i, factor := range primeFactors {
		if i > 0 {
			z01.PrintRune('*')
		}

		s := Itoa(factor)
		for _, r := range s {
			z01.PrintRune(r)
		}
	}
	z01.PrintRune('\n')
}

func getPrimeFactors(n int) []int {
	factors := []int{}
	for i := 2; i*i <= n; i++ {
		for ; n%i == 0; n /= i {
			factors = append(factors, i)
		}
	}
	if n > 1 {
		factors = append(factors, n)
	}
	return factors
}

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}

	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}

	var digits []rune
	for n > 0 {
		digit := n % 10
		digits = append([]rune{rune('0' + digit)}, digits...)
		n /= 10
	}
	return string(digits)
}
