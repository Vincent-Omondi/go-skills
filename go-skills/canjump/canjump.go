package main

import (
	"fmt"
)

func main() {
	input1 := []uint{2, 3, 1, 1, 4}
	fmt.Println(CanJump(input1))

	input2 := []uint{3, 2, 1, 0, 4}
	fmt.Println(CanJump(input2))

	input3 := []uint{0}
	fmt.Println(CanJump(input3))

	input4 := []uint{3, 2, 1, 2, 4, 6}
	fmt.Println(CanJump(input4))
}

func CanJump(slice []uint) bool {
	i := 0
	for i < len(slice) {
		if i == len(slice)-1 {
			return true
		}

		if int(slice[i]) == 0 {
			break
		}
		i += int(slice[i])
	}
	return false
}
