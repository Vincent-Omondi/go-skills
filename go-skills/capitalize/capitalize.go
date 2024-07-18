package main

import "fmt"

func main() {
	fmt.Println(Capitalize("Hello! How are you? How+are+things+4you?"))
}

func Capitalize(s string) string {
	var result []rune
	wdstart := true

	for _, c := range s {
		if c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' || c >= '0' && c <= '9' {
			if wdstart {
				result = append(result, toupper(c))
				wdstart = false
			} else {
				result = append(result, tolower(c))
			}
		} else {
			result = append(result, c)
			wdstart = true
		}
	}
	return string(result)
}

func toupper(r rune) rune {
	if r >= 'a' && r <= 'z' {
		return r - 'a' + 'A'
	}
	return r
}

func tolower(r rune) rune {
	if r >= 'A' && r <= 'Z' {
		return r - 'A' + 'a'
	}
	return r
}
