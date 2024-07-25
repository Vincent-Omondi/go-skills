package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 4 {
		return
	}

	s1, s2, s3 := os.Args[1], os.Args[2], os.Args[3]

	res := ""

	for _, c := range s1 {
		if string(c) == s2 {
			res += s3
			// string(c) = s3
		} else {
			res += string(c)
		}
	}
	fmt.Println(res)
}
