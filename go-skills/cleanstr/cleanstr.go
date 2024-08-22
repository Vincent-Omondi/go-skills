package main

import (
	"fmt"
	"os"
)

func main(){
	if len(os.Args) != 2{
		fmt.Println()
		return 
	}
	s:= os.Args[1]

	for range s{
		if s[0] == ' ' {
			s = s[1:]
		}

		if s[len(s)-1] == ' '{
			s = s[:len(s)-1]
		}
	}

	result := s

	for i := 0; i < len(result)-1; i++{
		if result[i] == 32 && result[i+1] == 32{
			result = result[:i] + result[i+1:]
			i--
		}
	}
	fmt.Println(result)
}