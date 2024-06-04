package main

import (
	"fmt"
	"os"
)

func main() {
	// Check the number of arguments
	if len(os.Args) != 3 {
		fmt.Println()
		return
	}

	// Get the input strings
	str1 := os.Args[1]
	str2 := os.Args[2]

	// Initialize a map to track seen characters
	seen := make(map[rune]bool)
	var result []rune

	// Function to add unique characters to the result
	addUniqueChars := func(str string) {
		for _, char := range str {
			if !seen[char] {
				seen[char] = true
				result = append(result, char)
			}
		}
	}

	// Add characters from both strings
	addUniqueChars(str1)
	addUniqueChars(str2)

	// Print the result
	fmt.Println(string(result))
}
