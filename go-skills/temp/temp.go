package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 || containsHelpFlag(os.Args[1:]) {
		fmt.Println("options: abcdefghijklmnopqrstuvwxyz")
		return
	}

	var options int

	for _, arg := range os.Args[1:] {
		if len(arg) < 2 || arg[0] != '-' {
			fmt.Println("Invalid Option")
			return
		}

		for _, ch := range arg[1:] {
			if ch < 'a' || ch > 'z' {
				fmt.Println("Invalid Option")
				return
			}
			options |= 1 << (ch - 'a')
		}
	}

	fmt.Printf("%08b %08b %08b %08b\n",
		(options>>24)&0xFF,
		(options>>16)&0xFF,
		(options>>8)&0xFF,
		options&0xFF)
}

func containsHelpFlag(args []string) bool {
	for _, arg := range args {
		if len(arg) > 1 && arg[1] == 'h' {
			return true
		}
	}
	return false
}
