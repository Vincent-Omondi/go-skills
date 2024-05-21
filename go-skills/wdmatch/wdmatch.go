package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		return
	}

	s1, s2 := os.Args[1], os.Args[2]

	pos := 0

	for _, r := range s1 {
		found := false
		for pos < len(s2) && !found {
			if rune(s2[pos]) == r {
				found = true
			}
			pos++
		}
		if !found {
			return
		}
	}

	fmt.Println(s1)
}
