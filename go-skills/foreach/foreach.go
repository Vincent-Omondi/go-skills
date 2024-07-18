package main

import "github.com/01-edu/z01"

func main() {
	a := []int{1, -2, 3, 4, 5, 6}
	ForEach(PrintNbr, a)
}

func ForEach(f func(int), a []int) {
	for _, num := range a {
		f(num)
	}
	z01.PrintRune('\n')
}

func PrintNbr(n int) {
	Isnegative := false
	if n < 0 {
		Isnegative = true
		n = -n
	}
	s := ""
	for ; n > 0; n /= 10 {
		s = string(rune(n%10)+'0') + s
	}

	if Isnegative {
		s = "-" + s
	}

	for _, c := range s {
		z01.PrintRune(c)
	}
}
