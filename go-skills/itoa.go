package main

import "fmt"

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}

	Isnegative := false

	if n < 0 {
		Isnegative = true
		n = -n
	}

	s := ""
	for ; n > 0; n /= 10 {
		s = string(rune(n%10)+'0') + s
	}

	if Isnegative {
		s = "-" + s
	}
	return s
}

func main() {
	fmt.Println(Itoa(12345)) // Expected output: 12345
	fmt.Println(Itoa(-1234)) // Expected output: -1234
	fmt.Println(Itoa(0))     // Expected output: 987654321
}
