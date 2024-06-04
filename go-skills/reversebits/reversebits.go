package main

import "fmt"

// ReverseBits reverses the bits of a byte
func ReverseBits(oct byte) byte {
	var result byte = 0
	for i := 0; i < 8; i++ {
		bit := (oct >> i) & 1      // Extract the i-th bit from oct
		result |= (bit << (7 - i)) // Set the bit in the reversed position
	}
	return result
}

func main() {
	var oct byte = 0x26 // 0010 0110 in binary
	reversed := ReverseBits(oct)
	fmt.Printf("%08b\n", reversed) // Should print: 0110 0100
}


// package main

// import (
// 	"github.com/01-edu/z01"
// )

// // ReverseBits reverses the bits of a byte
// func ReverseBits(oct byte) byte {
// 	var result byte = 0
// 	for i := 0; i < 8; i++ {
// 		bit := (oct >> i) & 1        // Extract the i-th bit from oct
// 		result |= (bit << (7 - i))   // Set the bit in the reversed position
// 	}
// 	return result
// }

// // PrintByteAsBinary prints a byte in binary format using z01.PrintRune
// func PrintByteAsBinary(b byte) {
// 	for i := 7; i >= 0; i-- {
// 		bit := (b >> i) & 1
// 		if bit == 1 {
// 			z01.PrintRune('1')
// 		} else {
// 			z01.PrintRune('0')
// 		}
// 	}
// 	z01.PrintRune('\n')
// }

// func main() {
// 	var oct byte = 0x26  // 0010 0110 in binary
// 	reversed := ReverseBits(oct)
// 	PrintByteAsBinary(reversed)  // Should print: 0110 0100
// }
