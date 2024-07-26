package main

import (
	"fmt"
)

func main() {
	fmt.Print(NotDecimal("0.1"))
	fmt.Print(NotDecimal("174.2"))
	fmt.Print(NotDecimal("0.1255"))
	fmt.Print(NotDecimal("1.20525856"))
	fmt.Print(NotDecimal("-0.0f00d00"))
	fmt.Print(NotDecimal(""))
	fmt.Print(NotDecimal("-19.525856"))
	fmt.Print(NotDecimal("1952"))
}

func NotDecimal(dec string) string {
	if len(dec) == 0 {
		return "\n"
	}

	for _, r := range dec {
		if (r < '0' || r > '9') && r != '.' && r != '-' {
			return dec + "\n"
		}
	}

	dotIndex := -1

	for i, r := range dec {
		if r == '.' {
			dotIndex = i
			break
		}
	}

	if dotIndex == -1 || (dotIndex == len(dec)-2 && dec[dotIndex+1] == '0') {
		return dec + "\n"
	}
	return dec[:dotIndex] + dec[dotIndex+1:] + "\n"

}
