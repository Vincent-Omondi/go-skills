package main

import "fmt"

func ReduceInt(a []int, f func(int, int) int) {
	acc := a[0]

	for _, v := range a[1:] {
		acc = f(acc, v)
	}
	fmt.Println(acc)
}

func main() {
	mul := func(acc int, cur int) int {
		return acc * cur
	}
	sum := func(acc int, cur int) int {
		return acc + cur
	}
	div := func(acc int, cur int) int {
		return acc / cur
	}
	as := []int{500, 2,}
	ReduceInt(as, mul)
	ReduceInt(as, sum)
	ReduceInt(as, div)
}
