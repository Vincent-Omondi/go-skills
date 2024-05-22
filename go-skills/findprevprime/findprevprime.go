package main

import "fmt"

func FindPrevPrime(n int) int {
	for i := n; i > 0; i-- {
		if IsPrime(i) {
			return i
		}
	}
	return 0
}

func IsPrime(n int) bool {
	if n < 0 || n == 1 {
		return false
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(FindPrevPrime(25)) // Output: 5
	fmt.Println(FindPrevPrime(4))  // Output: 3
	fmt.Println(FindPrevPrime(-2))
	fmt.Println(FindPrevPrime(3))
	fmt.Println(FindPrevPrime(5))
	fmt.Println(FindPrevPrime(0))
	fmt.Println(FindPrevPrime(11))
	fmt.Println(FindPrevPrime(29))
}
