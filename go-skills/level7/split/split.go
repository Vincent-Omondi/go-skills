package main

import "fmt"

func Split(s, sep string) []string {
	word := ""
	result := []string{}
	sepLen := len(sep)
	i := 0

	for i < len(s) {
		if i+sepLen <= len(s) && s[i:i+sepLen] == sep {
			result = append(result, word)
			word = ""
			i += sepLen
		} else {
			word += string(s[i])
			i++
		}
	}
	result = append(result, word)

	return result
}

func main() {
	s := "HelloHAhowHAareHAyou?"
	fmt.Printf("%#v\n", Split(s, "HA"))
}
