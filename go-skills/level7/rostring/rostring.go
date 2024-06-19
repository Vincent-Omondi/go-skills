package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println()
		return
	}

	s := os.Args[1]
	words := splitWords(s)
	if len(words) == 0 {
		fmt.Println()
		return
	}
	rotate := append(words[1:], words[0])

	for i := 0; i <= len(rotate)-1; i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(rotate[i])
	}
	fmt.Println()
}

func splitWords(s string) []string {
	word := ""
	words := []string{}
	inWord := false
	for _, v := range s {
		if v != ' ' {
			word += string(v)
			inWord = true
		} else if word != "" {
			words = append(words, word)
			word = ""
			inWord = false
		}
	}
	if inWord {
		words = append(words, word)
	}
	return words
}
