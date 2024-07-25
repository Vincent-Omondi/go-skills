package main

import "fmt"

func FifthAndSkip(str string) string {
	if str == "" {
		return "\n"
	}

	// Trim leading spaces
	for len(str) > 0 && str[0] == ' ' {
		str = str[1:]
	}

	// Trim trailing spaces
	for len(str) > 0 && str[len(str)-1] == ' ' {
		str = str[:len(str)-1]
	}

	// Replace multiple internal spaces with a single space
	var trimmed string
	inSpace := false
	for _, ch := range str {
		if ch == ' ' {
			if !inSpace {
				trimmed += ""
				inSpace = true
			}
		} else {
			trimmed += string(ch)
			inSpace = false
		}
	}

	if len(trimmed) < 5 {
		return "Invalid Input$\n"
	}

	var result string
	for i := 0; i < len(trimmed); i += 5 {
		endIndex := i + 5
		if endIndex > len(trimmed) {
			endIndex = len(trimmed)
		}
		substr := trimmed[i:endIndex]
		if len(substr) >= 6 {
			substr = substr[:5] + substr[6:]
		}
		result += substr
		if i+5 < len(trimmed) {
			result += " "
		}
	}

	return result + "$\n"
}

func main() {
	fmt.Print(FifthAndSkip("abcdefghijklmnopqrstuwxyz"))
	fmt.Print(FifthAndSkip("This  is  a  short  sentence"))
	fmt.Print(FifthAndSkip("1234"))
}
