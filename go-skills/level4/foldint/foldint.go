package main

import (
	"fmt"

	"github.com/01-edu/z01"
)

func FoldInt(f func(int, int) int, a []int, n int) {
	ac := n
	for _, v := range a {
		ac = f(ac, v)
	}

	num := ac
	if num == 0 {
		z01.PrintRune('0')
	}

	digits := ""
	for ; num > 0; num /= 10 {
		digits = string(rune(num%10)+'0') + digits
	}
	for _, v := range digits {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}

func Add(ac int, v int) int { return ac + v }
func Sub(ac int, v int) int { return ac - v }
func Mul(ac int, v int) int { return ac * v }
func Div(ac int, v int) int { return ac / v }

func main() {
	table := []int{1, 2, 3}
	ac := 93
	FoldInt(Add, table, ac)
	FoldInt(Mul, table, ac)
	FoldInt(Sub, table, ac)
	fmt.Println()

	table = []int{0}
	FoldInt(Add, table, ac)
	FoldInt(Mul, table, ac)
	FoldInt(Sub, table, ac)
}


