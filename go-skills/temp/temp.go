package main

import (
	"fmt"
)

func main() {
	// fmt.Print(FifthAndSkip("abcdefghijklmnopqrstuwxyz"))
	fmt.Print(FifthAndSkip("This is a short sentence"))
	// fmt.Print(FifthAndSkip("1234"))
}

func FifthAndSkip(str string) string {
	words := ""
	for _, r := range str {
		if r == ' ' {
			continue
		}
		words += string(r)
	}

	res := ""
	for i := 0; i < len(words); i += 5 {
		end := i + 5
		if end > len(words) {
			end = len(words)
		}
		strslice := words[i:end]

		if len(strslice) >= 6 {
			strslice = strslice[:5] + strslice[6:]
		}

		res += strslice
	}
	return res + "\n"
}
