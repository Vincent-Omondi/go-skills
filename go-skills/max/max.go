package main

import "fmt"

func Max(a []int) int {
	max := a[0]
	for i := 0; i < len(a); i++ {
		if a[i] > max {
			max = a[i]
		}
	}
	return max
}

func main() {
	a := []int{2223, 123, 1, 1111, 5225, 93}
	max := Max(a)
	fmt.Println(max)
}
