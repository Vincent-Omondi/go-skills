package main

import (
	"fmt"
	"strings"
)

func IsPalindrome(s string) bool{
	s = strings.ToLower(s)
	s1 := ""
	for _, v := range s{
		if v >= 'a' && v <= 'z' {
			s1 += string(v)
		}

	}
	res := s1
	s2 := ""

	for i:= len(res) -1; i >= 0; i--{
		s2 += string(res[i])
	}

	return res == s2
}



func main(){
	fmt.Println(IsPalindrome("A man, a plan, a canal, Panama!")) // true
	fmt.Println(IsPalindrome("Hello, World!"))                  // false
	fmt.Println(IsPalindrome("Madam"))                           // true
	fmt.Println(IsPalindrome("No lemon, no melon"))             // true
}