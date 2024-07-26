package main

import "github.com/01-edu/z01"

func main() {
	PrintComb()
}

func PrintComb() {
	for i := '0'; i <= '7'; i++ {
		for j := i + 1; j <= '8'; j++ {
			for h := j + 1; h <= '9'; h++ {
				z01.PrintRune(i)
				z01.PrintRune(j)
				z01.PrintRune(h)
				if i == '7' && j == '8' && h == '9' {
					break
				} else {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}

		}
	}
}