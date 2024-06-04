package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("ERROR: cannot convert to roman digit")
		return
	}

	num, err := strconv.Atoi(os.Args[1])
	if err != nil || num <= 0 || num >= 4000 {
		fmt.Println("ERROR: cannot convert to roman digit")
		return
	}

	romanMap := []struct {
		Value  int
		Symbol string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	result := ""
	calculation := ""

	for _, entry := range romanMap {
		for num >= entry.Value {
			result += entry.Symbol
			if calculation != "" {
				calculation += "+"
			}
			switch entry.Symbol {
			case "M", "D", "C", "L", "X", "V", "I":
				calculation += entry.Symbol
			case "CM":
				calculation += "(M-C)"
			case "CD":
				calculation += "(D-C)"
			case "XC":
				calculation += "(C-X)"
			case "XL":
				calculation += "(L-X)"
			case "IX":
				calculation += "(X-I)"
			case "IV":
				calculation += "(V-I)"
			}
			num -= entry.Value
		}
	}

	fmt.Println(calculation)
	fmt.Println(result)
}
