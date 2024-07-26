package main

import (
	"fmt"
)

func main() {
	fmt.Print(FifthAndSkip("abcdefghijklmnopqrstuwxyz"))
	fmt.Print(FifthAndSkip("This is a short sentence"))
	fmt.Print(FifthAndSkip("1234"))
}

func FifthAndSkip(str string) string {
	words := ""
	for _, r := range str {
		if r == ' ' {
			continue
		}
		words += string(r)
	}

	if len(words) == 0 {
		return "Invalid Input\n"
	}

	res := ""
	count := 0
	for i := 0; i < len(words); i++ {

		if count == 5 {
			count = 0
			continue
		}
		res += string(words[i])
		count++

		if count == 5 && i+1 < len(words) {
			res += " "
		}
	}
	return res + "\n"
}
