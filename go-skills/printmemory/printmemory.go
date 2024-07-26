package main

import (
	"fmt"
)

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}

func PrintMemory(arr [10]byte) {
	// Displaying the memory in hexadecimal
	for i := 0; i < len(arr); i += 4 {
		for j := 0; j < 4 && i+j < len(arr); j++ {
			fmt.Printf("%02x", arr[i+j])
		}
		fmt.Println()
	}

	// Displaying the ASCII characters, replacing non-printable characters with a dot
	for i := 0; i < len(arr); i++ {
		if arr[i] >= 32 && arr[i] <= 126 {
			fmt.Printf("%c", arr[i])
		} else {
			fmt.Print(".")
		}
	}
	fmt.Println()
}
