package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	n := atoi(os.Args[1])

	if n > 0 && (n&(n-1)) == 0 {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}

func atoi(s string) int {
	i := 0
	for _, n := range s {
		if n < '0' || n > '9' {
			return 0
		}
		i = i*10 + int(n-'0')
	}
	return i
}
