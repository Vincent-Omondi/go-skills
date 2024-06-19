package main

import (
	"os"
)

func atoi(s string) int {
	sign := 1
	if len(s) > 0 && s[0] == '-' {
		if s[0] == '-' {
			sign = -1
		}
		s = s[1:]
	}

	i := 0
	for _, c := range s {
		if c < '0' || c > '9' {
			return 0
		}
		i = i*10 + int(c-'0')
	}
	return i * sign
}

func itoa(n int) string {
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
		s = string(byte(n%10+'0')) + s
	}

	if Isnegative {
		s = "-" + s
	}

	return s
}

func main() {
	args := os.Args[1:]
	if len(args) != 3 {
		return
	}
	a, b, op := atoi(args[0]), atoi(args[2]), args[1]

	// Check for potential overflow before calculation
	if (b > 0 && a > (1<<63-1)-b) || (b < 0 && a < (-1<<63)-b) {
		return
	}

	var result int
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			os.Stdout.WriteString("No division by 0\n")
			return
		}
		result = a / b
	case "%":
		if b == 0 {
			os.Stdout.WriteString("No modulo by 0\n")
			return
		}
		result = a % b
	default:
		return
	}

	// Only print if the result is within the valid range
	if result >= -1<<63 && result <= 1<<63-1 {
		os.Stdout.WriteString(itoa(result) + "\n")
	}
}
