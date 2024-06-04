package main

import (
	"os"

	"github.com/01-edu/z01"
)

func Piglatin(word string) string {
	vowels := "aeiou"

	if containRune(vowels, rune(word[0])) {
		return word + "ay"
	}

	for i, v := range word {
		if containRune(vowels, v) {
			return word[i:] + word[:i] + "ay"
		}
	}
	return "No vowel"
}

func containRune(s string, r rune) bool {
	for _, c := range s {
		if c == r {
			return true
		}
	}
	return false
}

func main() {
	if len(os.Args) != 2 {
		return
	}
	s := os.Args[1]
	word := Piglatin(s)

	for _, v := range word {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}
