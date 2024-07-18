package main

import "fmt"

func main() {
	fmt.Println(Lcm(2, 7))
	fmt.Println(Lcm(0, 4))
}

func Lcm(first, second int) int {
	if first == 0 || second == 0 {
		return 0
	}

	gcd := gcd(first, second)
	return (first / gcd) * second
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
