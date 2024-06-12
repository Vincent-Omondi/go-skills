package main

import (
	"fmt"
	"os"
)

func Revwstr(s string) string {
	word := ""
	result := ""
	for _, v := range s {
		if v != ' ' {
			word += string(v)
		} else {
			result = " " + word + result
			word = ""
		}
	}
	return word + result
}

func main() {
	if len(os.Args) != 2 {
		return
	}
	s := os.Args[1]

	if len(s) == 0 {
		return
	}

	fmt.Println(Revwstr(s))
}

// below is string reverse
// if len(s) == 0 {
// 	return ""
// }
// return s[len(s)-1:] + Revwstr(s[:len(s)-1])
